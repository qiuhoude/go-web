package pcodec

import (
	"encoding/binary"
	"io"
)

// 基础包
type Package struct {
	Length int32  // 长度大小
	PbMsg  []byte // pb协议信息
}

func NewPack(data *[]byte) *Package {
	return &Package{
		Length: int32(len(*data)),
		PbMsg:  *data,
	}
}

// 打包
func (p *Package) Pack(writer io.Writer) error {
	var err error
	err = binary.Write(writer, binary.BigEndian, &p.Length)
	err = binary.Write(writer, binary.BigEndian, &p.PbMsg)
	return err
}

//
func (p *Package) Unpack(reader io.Reader) error {
	var err error
	err = binary.Read(reader, binary.BigEndian, &p.Length)
	p.PbMsg = make([]byte, p.Length) // 创建大小
	err = binary.Read(reader, binary.BigEndian, &p.PbMsg)
	return err
}
