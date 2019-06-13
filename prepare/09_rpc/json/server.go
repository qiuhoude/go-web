package main

import (
	"fmt"
	"github.com/qiuhoude/go-web/prepare/09_rpc/moudle"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	rect := new(moudle.Rect)
	//注册一个rect服务
	rpc.Register(rect)

	addr, err := net.ResolveTCPAddr("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s \n", "start connection")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go func() {
			fmt.Fprintf(os.Stdout, "%s", "new client in coming\n")
			jsonrpc.ServeConn(conn)
		}()

	}

}
