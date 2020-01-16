package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	u, err := url.Parse("https://localhost:443")
	if err != nil {
		panic(err)
	}
	tr := &http.Transport{
		Proxy: http.ProxyURL(u),
		// Disable HTTP/2.
		TLSNextProto:    make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 不验证证书

	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", dump)
}
