package main

import (
	"encoding/binary"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/qiuhoude/go-web/proto/v2/models"
	"github.com/smallnest/goframe"
	"io"
	"log"
	"net"
	"time"
)

func warpConn(conn net.Conn) goframe.FrameConn {
	encoderConfig := goframe.EncoderConfig{
		ByteOrder:                       binary.BigEndian,
		LengthFieldLength:               4,
		LengthAdjustment:                0,
		LengthIncludesLengthFieldLength: false,
	}

	decoderConfig := goframe.DecoderConfig{
		ByteOrder:           binary.BigEndian,
		LengthFieldOffset:   0,
		LengthFieldLength:   4,
		LengthAdjustment:    0,
		InitialBytesToStrip: 4,
	}
	fc := goframe.NewLengthFieldBasedFrameConn(encoderConfig, decoderConfig, conn)
	return fc

}
func recvMsg2(fc goframe.FrameConn) {
	conn := fc.Conn()
	for {
		_ = conn.SetReadDeadline(time.Now().Add(3 * time.Minute))
		frameData, err := fc.ReadFrame()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("远程链接：%s已经关闭！\n", conn.RemoteAddr().String())
			}
			break
		}

		basePb := &models.Base{}
		err = proto.Unmarshal(frameData, basePb)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("-> %v\n", basePb)
	}

}
