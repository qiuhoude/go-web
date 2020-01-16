package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/qiuhoude/go-web/proto/v2/models"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	maxSize    = 1048576
	lengthSize = 4
)

var Version = "No Version Provided ..."

func main() {
	info := getLoginInfo("ax1", "000")
	serverList(info.GetKeyId())
	//tcp(1, *info.KeyId, *info.Token)
}

func serverList(accountKey int64) {
	basePb := &models.Base{
		Cmd: proto.Int32(models.E_ServerListRq_Ext.Field),
	}
	dataPb := &models.ServerListRq{
		AccountKey: proto.Int64(accountKey),
	}
	_ = proto.SetExtension(basePb, models.E_DoLoginRq_Ext, dataPb)
	pbData, _ := proto.Marshal(basePb)
	url := "http://192.168.1.151:9200/honor_account/account/account.do"
	resp, err := http.Post(url, "application/octet-stream", bytes.NewReader(encode(pbData)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resBasePb := &models.Base{}
		_ = proto.Unmarshal(decode(body), resBasePb)
		fmt.Println(resBasePb)
		extension, _ := proto.GetExtension(resBasePb, models.E_ServerListRs_Ext)
		res := extension.(*models.ServerListRs)
		fmt.Printf("%v\n", res)
	}

}

func encode(data []byte) []byte {
	buf := new(bytes.Buffer)
	length := uint16(len(data))
	_ = binary.Write(buf, binary.BigEndian, length)
	_ = binary.Write(buf, binary.BigEndian, data)
	return buf.Bytes()
}

func decode(data []byte) []byte {
	buf := bytes.NewBuffer(data)
	var length uint16
	_ = binary.Read(buf, binary.BigEndian, &length)
	return buf.Bytes()
}

func getLoginInfo(account, pwd string) *models.DoLoginRs {

	basePb := &models.Base{
		Cmd: proto.Int32(models.E_DoLoginRq_Ext.Field),
	}
	dataPb := &models.DoLoginRq{
		Sid:         proto.String(fmt.Sprintf("%s_%s", account, pwd)),
		BaseVersion: proto.String("1.4.0"),
		Version:     proto.String("1.0.0"),
		DeviceNo:    proto.String("00000000-2625-0b64-7b72-55e30033c587"),
		Plat:        proto.String("self"),
	}
	err := proto.SetExtension(basePb, models.E_DoLoginRq_Ext, dataPb)

	pbData, _ := proto.Marshal(basePb)
	url := "http://192.168.1.151:9200/honor_account/account/account.do"
	resp, err := http.Post(url, "application/octet-stream", bytes.NewReader(encode(pbData)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resBasePb := &models.Base{}
		_ = proto.Unmarshal(decode(body), resBasePb)
		fmt.Println(resBasePb)
		extension, _ := proto.GetExtension(resBasePb, models.E_DoLoginRs_Ext)
		res := extension.(*models.DoLoginRs)
		return res
	}
	log.Fatal("未获取到Token")
	return nil
}

func tcp(serverId int32, keyId int64, token string) {
	fmt.Println("Client Version is:", Version)
	//go run -ldflags "-X main.Version=1.6.6" client.go 编译版本
	strIP := "192.168.1.151:9201"
	var conn net.Conn
	var err error
	addr, err := net.ResolveTCPAddr("tcp", strIP)

	//连接服务器
	for conn, err = net.DialTCP("tcp", nil, addr); err != nil; conn, err = net.Dial("tcp", strIP) {
		fmt.Println("connect", strIP, "fail")
		time.Sleep(time.Second)
		fmt.Println("reconnect...")
	}
	fmt.Println("connect", strIP, "success")
	defer conn.Close()

	//writer := bufio.NewWriter(conn)
	// 开个协程进行读数据
	fc := warpConn(conn)
	go recvMsg2(fc)

	_ = fc.WriteFrame(begin(serverId, keyId, token))
	time.AfterFunc(3*time.Second, func() {
		_ = fc.WriteFrame(login())
	})

	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		_ = fc.WriteFrame(chatMsg(sender.Text()))
		if sender.Text() == "stop" {
			return
		}
	}
}

func chatMsg(msg string) []byte {
	basePb := &models.Base{
		Cmd: proto.Int32(models.E_SendChatRq_Ext.Field),
	}
	dataPb := &models.SendChatRq{
		Channel: proto.Int32(1),
		Content: []string{msg},
	}
	_ = proto.SetExtension(basePb, models.E_SendChatRq_Ext, dataPb)
	bdata, _ := proto.Marshal(basePb)
	return bdata
}

func login() []byte {
	basePb := &models.Base{
		Cmd: proto.Int32(models.E_RoleLoginRq_Ext.Field),
	}
	dataPb := &models.RoleLoginRq{}
	_ = proto.SetExtension(basePb, models.E_RoleLoginRq_Ext, dataPb)
	bdata, _ := proto.Marshal(basePb)
	return bdata
}

func begin(serverId int32, keyId int64, token string) []byte {
	basePb := &models.Base{
		Cmd: proto.Int32(models.E_BeginGameRq_Ext.Field),
	}
	beginRqPb := &models.BeginGameRq{
		ServerId:   proto.Int32(serverId),
		KeyId:      proto.Int64(keyId),
		Token:      proto.String(token),
		DeviceNo:   proto.String(`00000000-2625-0b64-7b72-55e30033c587`),
		CurVersion: proto.String(`1.0.0`),
	}
	_ = proto.SetExtension(basePb, models.E_BeginGameRq_Ext, beginRqPb)
	beginData, _ := proto.Marshal(basePb)
	return beginData
}

func sendMsg(writer *bufio.Writer, data []byte) error {
	var err error
	err = binary.Write(writer, binary.BigEndian, int32(len(data)))
	err = binary.Write(writer, binary.BigEndian, data)
	writer.Flush()
	return err
}

func recvMsg(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		conn.SetReadDeadline(time.Now().Add(3 * time.Minute))
		lengthByte, err := reader.Peek(lengthSize)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("远程链接：%s已经关闭！\n", conn.RemoteAddr().String())
			}
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
		conn.SetReadDeadline(time.Time{})
	}
	fmt.Println("读协程死掉了")
}
