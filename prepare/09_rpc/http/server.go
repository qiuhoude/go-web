package main

import (
	"github.com/qiuhoude/go-web/prepare/09_rpc/moudle"
	"log"
	"net/http"
	"net/rpc"
)

func main() {

	rect := moudle.Rect{}
	//注册一个rect服务
	rpc.Register(&rect)
	//把服务处理绑定到http协议上
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
