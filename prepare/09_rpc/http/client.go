package main

import (
	"fmt"
	"github.com/qiuhoude/go-web/prepare/09_rpc/moudle"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	ret := 0
	//调用远程方法
	//注意第三个参数是指针类型
	err = client.Call("Rect.Area", moudle.Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)
	err = client.Call("Rect.Perimeter", moudle.Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)

	client.Go()
}
