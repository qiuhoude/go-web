package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"time"
)

/*
#!/usr/bin/env bash
case `uname -s` in
    Linux*)     sslConfig=/etc/ssl/openssl.cnf;;
    Darwin*)    sslConfig=/System/Library/OpenSSL/openssl.cnf;;
esac
openssl req \
    -newkey rsa:2048 \
    -x509 \
    -nodes \
    -keyout server.key \
    -new \
    -out server.pem \
    -subj /CN=localhost \
    -reqexts SAN \
    -extensions SAN \
    -config <(cat $sslConfig \
        <(printf '[SAN]\nsubjectAltName=DNS:localhost')) \
    -sha256 \
    -days 3650
签名证书脚本

1. 生成 512位的私钥
openssl genrsa -out server.key 512

2. 身份证申请（CSR）文件生成 scr(certificate signing request) 客户端必须可以通过域名 locahost 访问到这个 HTTPS 服务
openssl req -nodes -key server.key -subj '/CN=localhost' -out server.csr

3. 签署身份证, 自签名(self-sign) 用自己的私钥 签署自己的CSR
openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -out server.crt

k8s中 apiServer 签署
会用一个CA的私钥来签署一个CSR,

*/
/*
 http代理
*/

func main() {
	pemPath := "server.crt" // 证书路径
	keyPath := "server.key" // 私钥路径
	proto := "https"        // 协议 http 或https

	server := &http.Server{
		Addr: ":443",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodConnect {
				//支持https websocket deng ... tcp
				handleTunneling(w, r)
			} else {
				//直接http代理
				handleHTTP(w, r)
			}
		}),
		// 关闭http2
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	if proto == "http" {
		log.Fatal(server.ListenAndServe())
	} else {
		log.Fatal(server.ListenAndServeTLS(pemPath, keyPath))
	}
}

func handleTunneling(w http.ResponseWriter, r *http.Request) {
	//设置超时防止大量超时导致服务器资源不大量占用
	dest_conn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	//类型转换，Hijacker 接口容许接管连接.在此之后调用者有责任管理这个链接
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	//接管连接
	client_conn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	reqDump, err := httputil.DumpRequest(r, true)
	fmt.Printf("req-> %q", reqDump)

	go transfer(dest_conn, client_conn)
	go transfer(client_conn, dest_conn)
}

//转发连接的数据
func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

func handleHTTP(w http.ResponseWriter, req *http.Request) {
	//roudtrip 传递发送的请求返回响应的结果
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	//把目标服务器的响应header复制
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

//复制响应头
func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
