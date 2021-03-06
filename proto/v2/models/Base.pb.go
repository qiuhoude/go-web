// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Base.proto

package models

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Base struct {
	Cmd                          *int32   `protobuf:"varint,1,req,name=cmd" json:"cmd,omitempty"`
	Code                         *int32   `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Param                        *int64   `protobuf:"varint,3,opt,name=param" json:"param,omitempty"`
	LordId                       *int64   `protobuf:"varint,4,opt,name=lordId" json:"lordId,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *Base) Reset()         { *m = Base{} }
func (m *Base) String() string { return proto.CompactTextString(m) }
func (*Base) ProtoMessage()    {}
func (*Base) Descriptor() ([]byte, []int) {
	return fileDescriptor_70eaf8cd7d110228, []int{0}
}

var extRange_Base = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*Base) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Base
}

func (m *Base) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Base.Unmarshal(m, b)
}
func (m *Base) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Base.Marshal(b, m, deterministic)
}
func (m *Base) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Base.Merge(m, src)
}
func (m *Base) XXX_Size() int {
	return xxx_messageInfo_Base.Size(m)
}
func (m *Base) XXX_DiscardUnknown() {
	xxx_messageInfo_Base.DiscardUnknown(m)
}

var xxx_messageInfo_Base proto.InternalMessageInfo

func (m *Base) GetCmd() int32 {
	if m != nil && m.Cmd != nil {
		return *m.Cmd
	}
	return 0
}

func (m *Base) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *Base) GetParam() int64 {
	if m != nil && m.Param != nil {
		return *m.Param
	}
	return 0
}

func (m *Base) GetLordId() int64 {
	if m != nil && m.LordId != nil {
		return *m.LordId
	}
	return 0
}

func init() {
	proto.RegisterType((*Base)(nil), "Base")
}

func init() { proto.RegisterFile("Base.proto", fileDescriptor_70eaf8cd7d110228) }

var fileDescriptor_70eaf8cd7d110228 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x1c, 0xce, 0x3d, 0x6b, 0x80, 0x30,
	0x10, 0x80, 0x61, 0xfc, 0x44, 0xae, 0x1f, 0x94, 0x4c, 0xa1, 0x93, 0x74, 0x12, 0xa1, 0x09, 0xf4,
	0x27, 0xb8, 0xb5, 0x93, 0x73, 0xb7, 0x24, 0x77, 0xa8, 0x60, 0x3c, 0x1b, 0x4d, 0x5d, 0xf3, 0xd3,
	0x8b, 0x1d, 0x9f, 0xe9, 0x7d, 0x01, 0x06, 0x73, 0x90, 0xda, 0x03, 0x9f, 0xfc, 0xf6, 0x05, 0xe5,
	0x2d, 0xf1, 0x00, 0x85, 0xf3, 0x28, 0xb3, 0x36, 0xef, 0x2a, 0xf1, 0x08, 0xa5, 0x63, 0x24, 0x99,
	0xb7, 0x59, 0x57, 0x89, 0x27, 0xa8, 0x76, 0x13, 0x8c, 0x97, 0x45, 0x9b, 0x75, 0x85, 0x78, 0x86,
	0x7a, 0xe5, 0x80, 0x9f, 0x28, 0xcb, 0xdb, 0x7d, 0xd3, 0xe0, 0x4b, 0x4a, 0x29, 0xe5, 0xc3, 0x08,
	0xaf, 0x8e, 0xbd, 0x9a, 0xe3, 0x86, 0x81, 0xd0, 0xd1, 0x76, 0xaa, 0xc9, 0x78, 0x52, 0xfe, 0x52,
	0xbb, 0x1d, 0xea, 0xbb, 0x33, 0xda, 0xef, 0x7e, 0x5a, 0xce, 0x39, 0x5a, 0xe5, 0xd8, 0xeb, 0x9f,
	0x25, 0xce, 0x1c, 0x91, 0xf4, 0xc4, 0xef, 0x17, 0x59, 0xfd, 0xff, 0xa4, 0x7f, 0x3f, 0xb4, 0x67,
	0xa4, 0xf5, 0xf8, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x85, 0xc5, 0xa8, 0x7d, 0xaa, 0x00, 0x00, 0x00,
}
