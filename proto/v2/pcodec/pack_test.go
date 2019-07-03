package pcodec

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"log"
	"testing"
	"time"
)

func TestPack(t *testing.T) {
	msg := []byte("现在时间是:" + time.Now().Format("2006-01-02 15:04:05"))
	pack := &Package{
		Length: int32(len(msg)),
		PbMsg:  msg,
	}

	buf := new(bytes.Buffer)
	// 写入四次，模拟TCP粘包效果
	pack.Pack(buf)
	buf.Write([]byte("213213")) // 模拟个错误数据包
	pack.Pack(buf)
	pack.Pack(buf)
	pack.Pack(buf)

	// 使用scanner的方式
	/*scanner := bufio.NewScanner(buf)
	scanner.Split(ScanPBMsg)
	for scanner.Scan() {
		scannedPack := new(Package)
		scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
		log.Printf("%d,%s", scannedPack.Length, scannedPack.PbMsg)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("无效数据包")
	}*/
	const (
		maxSize = int32(140000)
	)
	reader := bufio.NewReader(buf)
	for reader.Size() > 0 {
		lengthByte, err := reader.Peek(4)
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
		packBytes := make([]byte, int32(length+4))
		reader.Read(packBytes)
		pack := new(Package)
		pack.Unpack(bytes.NewReader(packBytes))
		log.Printf("%d,%s", pack.Length, pack.PbMsg)
	}

}
