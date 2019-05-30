package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	port := ":5000"
	//在Go语言中通过ResolveTCPAddr获取一个TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	CheckErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr) // 监听某个端口
	CheckErr(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}

}

// 处理客户端的请求
func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:]) // 相等于每次最多读 512字节 , n表示读到的字节数
		if err != nil {
			return
		}
		remoteAddr := conn.RemoteAddr()
		fmt.Println("Receive from client", remoteAddr.String(), string(buf[0:n]))
		_, err2 := conn.Write([]byte("hi client"))
		if err2 != nil {
			return
		}
	}
}

func CheckErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
