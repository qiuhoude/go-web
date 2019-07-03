package pcodec

import (
	"bytes"
	"encoding/binary"
)

func ScanPBMsg(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if len(data) < 4 {

	}
	var length int32
	binary.Read(bytes.NewReader(data[:4]), binary.BigEndian, &length)
	// 过长进行丢弃

	if int(length+4) <= len(data) {
		return int(length + 4), data[:int(length)+4], nil
	}
	return 0, nil, nil
}
