package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/qiuhoude/go-web/proto/v2/models"
	"net"
	"os"
	"time"
)

const (
	maxSize    = 1048576
	lengthSize = 4
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

	writer := bufio.NewWriter(conn)
	// 开个协程进行读数据
	go recvMsg(&conn)

	sendbegin(writer)
	time.AfterFunc(3*time.Second, func() {
		sendlogin(writer)
	})

	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		sendChatMsg(writer, sender.Text())
		if sender.Text() == "stop" {
			return
		}
	}
}

func sendChatMsg(writer *bufio.Writer, msg string) {
	basePb := &models.Base{
		Cmd: proto.Int32(models.E_SendChatRq_Ext.Field),
	}
	dataPb := &models.SendChatRq{
		Channel: proto.Int32(1),
		Content: []string{msg},
	}
	proto.SetExtension(basePb, models.E_SendChatRq_Ext, dataPb)
	bdata, _ := proto.Marshal(basePb)
	sendMsg(writer, bdata)
}

func sendlogin(writer *bufio.Writer) {
	basePb := &models.Base{
		Cmd: proto.Int32(models.E_RoleLoginRq_Ext.Field),
	}
	dataPb := &models.RoleLoginRq{}
	proto.SetExtension(basePb, models.E_RoleLoginRq_Ext, dataPb)
	bdata, _ := proto.Marshal(basePb)
	sendMsg(writer, bdata)
}

func sendbegin(writer *bufio.Writer) {

	basePb := &models.Base{
		Cmd: proto.Int32(models.E_BeginGameRq_Ext.Field),
	}
	beginRqPb := &models.BeginGameRq{
		ServerId:   proto.Int32(5),
		KeyId:      proto.Int64(53248),
		Token:      proto.String(`eb00937b79dc45a18494a1e357334ec7`),
		DeviceNo:   proto.String(`00000000-2625-0b64-7b72-55e30033c587`),
		CurVersion: proto.String(`1.0.0`),
	}
	proto.SetExtension(basePb, models.E_BeginGameRq_Ext, beginRqPb)
	beginData, _ := proto.Marshal(basePb)
	sendMsg(writer, beginData)

}

func sendMsg(writer *bufio.Writer, data []byte) error {
	var err error
	err = binary.Write(writer, binary.BigEndian, int32(len(data)))
	err = binary.Write(writer, binary.BigEndian, data)
	writer.Flush()
	return err
}

func recvMsg(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	for {
		(*conn).SetReadDeadline(time.Now().Add(3 * time.Minute))
		lengthByte, err := reader.Peek(lengthSize)
		if err != nil {
			break
		}
		lengthBuff := bytes.NewBuffer(lengthByte)
		var length int32
		err = binary.Read(lengthBuff, binary.BigEndian, &length)
		if err != nil {
			break
		}
		if length > maxSize || length < 0 { //说明是错误数据调过1字节
			reader.Discard(1)
			continue
		}
		reader.Discard(lengthSize)
		packBytes := make([]byte, length)
		basePb := &models.Base{}
		reader.Read(packBytes)
		proto.Unmarshal(packBytes, basePb)
		fmt.Println(basePb)
		(*conn).SetReadDeadline(time.Time{})
	}
	fmt.Println("读协程死掉了")
}
