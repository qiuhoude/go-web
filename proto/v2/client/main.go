package main

import (
	"bufio"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/qiuhoude/go-web/proto/v2/models"
	"net"
	"os"
	"time"
)

const (
	maxSize  = int32(1048576)
	bodySize = int32(4)
)

func main() {
	strIP := "127.0.0.1:9201"
	var conn net.Conn
	var err error

	//连接服务器
	for conn, err = net.Dial("tcp", strIP); err != nil; conn, err = net.Dial("tcp", strIP) {
		fmt.Println("connect", strIP, "fail")
		time.Sleep(time.Second)
		fmt.Println("reconnect...")
	}
	fmt.Println("connect", strIP, "success")
	defer conn.Close()

	// 开个协程进行读数据
	go func(conn net.Conn) {
		//reader := bufio.NewReader(conn)
		//for {
		//	reader.Discard()
		//	reader.Reset()
		//
		//}

	}(conn)

	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		stSend := &models.BeginGameRq{}

		//protobuf编码
		pData, err := proto.Marshal(stSend)
		if err != nil {
			panic(err)
		}
		//发送
		conn.Write(pData)
		if sender.Text() == "stop" {
			return
		}
	}
}
