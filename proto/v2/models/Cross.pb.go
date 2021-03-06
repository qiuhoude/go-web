// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Cross.proto

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

// 服务器注册
type ServerRegistRq struct {
	ServerId             *int32   `protobuf:"varint,1,req,name=serverId" json:"serverId,omitempty"`
	Index                *int32   `protobuf:"varint,2,req,name=index" json:"index,omitempty"`
	Total                *int32   `protobuf:"varint,3,req,name=total" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerRegistRq) Reset()         { *m = ServerRegistRq{} }
func (m *ServerRegistRq) String() string { return proto.CompactTextString(m) }
func (*ServerRegistRq) ProtoMessage()    {}
func (*ServerRegistRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e60e98cf7de90f, []int{0}
}
func (m *ServerRegistRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerRegistRq.Unmarshal(m, b)
}
func (m *ServerRegistRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerRegistRq.Marshal(b, m, deterministic)
}
func (m *ServerRegistRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerRegistRq.Merge(m, src)
}
func (m *ServerRegistRq) XXX_Size() int {
	return xxx_messageInfo_ServerRegistRq.Size(m)
}
func (m *ServerRegistRq) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerRegistRq.DiscardUnknown(m)
}

var xxx_messageInfo_ServerRegistRq proto.InternalMessageInfo

func (m *ServerRegistRq) GetServerId() int32 {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return 0
}

func (m *ServerRegistRq) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *ServerRegistRq) GetTotal() int32 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

var E_ServerRegistRq_Ext = &proto.ExtensionDesc{
	ExtendedType:  (*Base)(nil),
	ExtensionType: (*ServerRegistRq)(nil),
	Field:         10001,
	Name:          "ServerRegistRq.ext",
	Tag:           "bytes,10001,opt,name=ext",
	Filename:      "Cross.proto",
}

// 游戏服 -> 跨服 发送的心跳
type HeartRq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartRq) Reset()         { *m = HeartRq{} }
func (m *HeartRq) String() string { return proto.CompactTextString(m) }
func (*HeartRq) ProtoMessage()    {}
func (*HeartRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e60e98cf7de90f, []int{1}
}
func (m *HeartRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartRq.Unmarshal(m, b)
}
func (m *HeartRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartRq.Marshal(b, m, deterministic)
}
func (m *HeartRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartRq.Merge(m, src)
}
func (m *HeartRq) XXX_Size() int {
	return xxx_messageInfo_HeartRq.Size(m)
}
func (m *HeartRq) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartRq.DiscardUnknown(m)
}

var xxx_messageInfo_HeartRq proto.InternalMessageInfo

var E_HeartRq_Ext = &proto.ExtensionDesc{
	ExtendedType:  (*Base)(nil),
	ExtensionType: (*HeartRq)(nil),
	Field:         10003,
	Name:          "HeartRq.ext",
	Tag:           "bytes,10003,opt,name=ext",
	Filename:      "Cross.proto",
}

type HeartRs struct {
	Time                 *int64   `protobuf:"varint,1,req,name=time" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartRs) Reset()         { *m = HeartRs{} }
func (m *HeartRs) String() string { return proto.CompactTextString(m) }
func (*HeartRs) ProtoMessage()    {}
func (*HeartRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e60e98cf7de90f, []int{2}
}
func (m *HeartRs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartRs.Unmarshal(m, b)
}
func (m *HeartRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartRs.Marshal(b, m, deterministic)
}
func (m *HeartRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartRs.Merge(m, src)
}
func (m *HeartRs) XXX_Size() int {
	return xxx_messageInfo_HeartRs.Size(m)
}
func (m *HeartRs) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartRs.DiscardUnknown(m)
}

var xxx_messageInfo_HeartRs proto.InternalMessageInfo

func (m *HeartRs) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

var E_HeartRs_Ext = &proto.ExtensionDesc{
	ExtendedType:  (*Base)(nil),
	ExtensionType: (*HeartRs)(nil),
	Field:         10004,
	Name:          "HeartRs.ext",
	Tag:           "bytes,10004,opt,name=ext",
	Filename:      "Cross.proto",
}

// 登陆、退出跨服
type CrossLoginRq struct {
	OpType               *int32         `protobuf:"varint,1,opt,name=opType" json:"opType,omitempty"`
	Player               *CrossPlayerPb `protobuf:"bytes,2,opt,name=player" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CrossLoginRq) Reset()         { *m = CrossLoginRq{} }
func (m *CrossLoginRq) String() string { return proto.CompactTextString(m) }
func (*CrossLoginRq) ProtoMessage()    {}
func (*CrossLoginRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e60e98cf7de90f, []int{3}
}
func (m *CrossLoginRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrossLoginRq.Unmarshal(m, b)
}
func (m *CrossLoginRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrossLoginRq.Marshal(b, m, deterministic)
}
func (m *CrossLoginRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossLoginRq.Merge(m, src)
}
func (m *CrossLoginRq) XXX_Size() int {
	return xxx_messageInfo_CrossLoginRq.Size(m)
}
func (m *CrossLoginRq) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossLoginRq.DiscardUnknown(m)
}

var xxx_messageInfo_CrossLoginRq proto.InternalMessageInfo

func (m *CrossLoginRq) GetOpType() int32 {
	if m != nil && m.OpType != nil {
		return *m.OpType
	}
	return 0
}

func (m *CrossLoginRq) GetPlayer() *CrossPlayerPb {
	if m != nil {
		return m.Player
	}
	return nil
}

var E_CrossLoginRq_Ext = &proto.ExtensionDesc{
	ExtendedType:  (*Base)(nil),
	ExtensionType: (*CrossLoginRq)(nil),
	Field:         10005,
	Name:          "CrossLoginRq.ext",
	Tag:           "bytes,10005,opt,name=ext",
	Filename:      "Cross.proto",
}

type CrossLoginRs struct {
	OpType               *int32        `protobuf:"varint,1,opt,name=opType" json:"opType,omitempty"`
	Hero                 []*FortHeroPb `protobuf:"bytes,2,rep,name=hero" json:"hero,omitempty"`
	CurKillNum           *int32        `protobuf:"varint,3,opt,name=curKillNum" json:"curKillNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CrossLoginRs) Reset()         { *m = CrossLoginRs{} }
func (m *CrossLoginRs) String() string { return proto.CompactTextString(m) }
func (*CrossLoginRs) ProtoMessage()    {}
func (*CrossLoginRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e60e98cf7de90f, []int{4}
}
func (m *CrossLoginRs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrossLoginRs.Unmarshal(m, b)
}
func (m *CrossLoginRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrossLoginRs.Marshal(b, m, deterministic)
}
func (m *CrossLoginRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossLoginRs.Merge(m, src)
}
func (m *CrossLoginRs) XXX_Size() int {
	return xxx_messageInfo_CrossLoginRs.Size(m)
}
func (m *CrossLoginRs) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossLoginRs.DiscardUnknown(m)
}

var xxx_messageInfo_CrossLoginRs proto.InternalMessageInfo

func (m *CrossLoginRs) GetOpType() int32 {
	if m != nil && m.OpType != nil {
		return *m.OpType
	}
	return 0
}

func (m *CrossLoginRs) GetHero() []*FortHeroPb {
	if m != nil {
		return m.Hero
	}
	return nil
}

func (m *CrossLoginRs) GetCurKillNum() int32 {
	if m != nil && m.CurKillNum != nil {
		return *m.CurKillNum
	}
	return 0
}

var E_CrossLoginRs_Ext = &proto.ExtensionDesc{
	ExtendedType:  (*Base)(nil),
	ExtensionType: (*CrossLoginRs)(nil),
	Field:         10006,
	Name:          "CrossLoginRs.ext",
	Tag:           "bytes,10006,opt,name=ext",
	Filename:      "Cross.proto",
}

// 发送跨服聊天
type SendCrossChatRq struct {
	Chat                 *Chat    `protobuf:"bytes,1,req,name=chat" json:"chat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendCrossChatRq) Reset()         { *m = SendCrossChatRq{} }
func (m *SendCrossChatRq) String() string { return proto.CompactTextString(m) }
func (*SendCrossChatRq) ProtoMessage()    {}
func (*SendCrossChatRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e60e98cf7de90f, []int{5}
}
func (m *SendCrossChatRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendCrossChatRq.Unmarshal(m, b)
}
func (m *SendCrossChatRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendCrossChatRq.Marshal(b, m, deterministic)
}
func (m *SendCrossChatRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendCrossChatRq.Merge(m, src)
}
func (m *SendCrossChatRq) XXX_Size() int {
	return xxx_messageInfo_SendCrossChatRq.Size(m)
}
func (m *SendCrossChatRq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendCrossChatRq.DiscardUnknown(m)
}

var xxx_messageInfo_SendCrossChatRq proto.InternalMessageInfo

func (m *SendCrossChatRq) GetChat() *Chat {
	if m != nil {
		return m.Chat
	}
	return nil
}

var E_SendCrossChatRq_Ext = &proto.ExtensionDesc{
	ExtendedType:  (*Base)(nil),
	ExtensionType: (*SendCrossChatRq)(nil),
	Field:         10007,
	Name:          "SendCrossChatRq.ext",
	Tag:           "bytes,10007,opt,name=ext",
	Filename:      "Cross.proto",
}

// 选择将领加入或
type ChoiceHeroRq struct {
	Heros                []*CrossHeroPb `protobuf:"bytes,1,rep,name=heros" json:"heros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ChoiceHeroRq) Reset()         { *m = ChoiceHeroRq{} }
func (m *ChoiceHeroRq) String() string { return proto.CompactTextString(m) }
func (*ChoiceHeroRq) ProtoMessage()    {}
func (*ChoiceHeroRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e60e98cf7de90f, []int{6}
}
func (m *ChoiceHeroRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChoiceHeroRq.Unmarshal(m, b)
}
func (m *ChoiceHeroRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChoiceHeroRq.Marshal(b, m, deterministic)
}
func (m *ChoiceHeroRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChoiceHeroRq.Merge(m, src)
}
func (m *ChoiceHeroRq) XXX_Size() int {
	return xxx_messageInfo_ChoiceHeroRq.Size(m)
}
func (m *ChoiceHeroRq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChoiceHeroRq.DiscardUnknown(m)
}

var xxx_messageInfo_ChoiceHeroRq proto.InternalMessageInfo

func (m *ChoiceHeroRq) GetHeros() []*CrossHeroPb {
	if m != nil {
		return m.Heros
	}
	return nil
}

var E_ChoiceHeroRq_Ext = &proto.ExtensionDesc{
	ExtendedType:  (*Base)(nil),
	ExtensionType: (*ChoiceHeroRq)(nil),
	Field:         10009,
	Name:          "ChoiceHeroRq.ext",
	Tag:           "bytes,10009,opt,name=ext",
	Filename:      "Cross.proto",
}

// 扣除资源返回值 游戏服->跨服
type CrossAwardOpRq struct {
	TaskId               *int64        `protobuf:"varint,1,opt,name=taskId" json:"taskId,omitempty"`
	ReqAwards            *CrossAwardPb `protobuf:"bytes,2,opt,name=reqAwards" json:"reqAwards,omitempty"`
	Success              *bool         `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CrossAwardOpRq) Reset()         { *m = CrossAwardOpRq{} }
func (m *CrossAwardOpRq) String() string { return proto.CompactTextString(m) }
func (*CrossAwardOpRq) ProtoMessage()    {}
func (*CrossAwardOpRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e60e98cf7de90f, []int{7}
}
func (m *CrossAwardOpRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrossAwardOpRq.Unmarshal(m, b)
}
func (m *CrossAwardOpRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrossAwardOpRq.Marshal(b, m, deterministic)
}
func (m *CrossAwardOpRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossAwardOpRq.Merge(m, src)
}
func (m *CrossAwardOpRq) XXX_Size() int {
	return xxx_messageInfo_CrossAwardOpRq.Size(m)
}
func (m *CrossAwardOpRq) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossAwardOpRq.DiscardUnknown(m)
}

var xxx_messageInfo_CrossAwardOpRq proto.InternalMessageInfo

func (m *CrossAwardOpRq) GetTaskId() int64 {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return 0
}

func (m *CrossAwardOpRq) GetReqAwards() *CrossAwardPb {
	if m != nil {
		return m.ReqAwards
	}
	return nil
}

func (m *CrossAwardOpRq) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

var E_CrossAwardOpRq_Ext = &proto.ExtensionDesc{
	ExtendedType:  (*Base)(nil),
	ExtensionType: (*CrossAwardOpRq)(nil),
	Field:         10010,
	Name:          "CrossAwardOpRq.ext",
	Tag:           "bytes,10010,opt,name=ext",
	Filename:      "Cross.proto",
}

// 异步资源扣除资源  跨服->游戏服
type CrossAwardOpRs struct {
	TaskId               *int64        `protobuf:"varint,1,opt,name=taskId" json:"taskId,omitempty"`
	ReqAwards            *CrossAwardPb `protobuf:"bytes,2,opt,name=reqAwards" json:"reqAwards,omitempty"`
	RollBack             *bool         `protobuf:"varint,3,opt,name=rollBack" json:"rollBack,omitempty"`
	Cmd                  *int32        `protobuf:"varint,4,opt,name=cmd" json:"cmd,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CrossAwardOpRs) Reset()         { *m = CrossAwardOpRs{} }
func (m *CrossAwardOpRs) String() string { return proto.CompactTextString(m) }
func (*CrossAwardOpRs) ProtoMessage()    {}
func (*CrossAwardOpRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e60e98cf7de90f, []int{8}
}
func (m *CrossAwardOpRs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrossAwardOpRs.Unmarshal(m, b)
}
func (m *CrossAwardOpRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrossAwardOpRs.Marshal(b, m, deterministic)
}
func (m *CrossAwardOpRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossAwardOpRs.Merge(m, src)
}
func (m *CrossAwardOpRs) XXX_Size() int {
	return xxx_messageInfo_CrossAwardOpRs.Size(m)
}
func (m *CrossAwardOpRs) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossAwardOpRs.DiscardUnknown(m)
}

var xxx_messageInfo_CrossAwardOpRs proto.InternalMessageInfo

func (m *CrossAwardOpRs) GetTaskId() int64 {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return 0
}

func (m *CrossAwardOpRs) GetReqAwards() *CrossAwardPb {
	if m != nil {
		return m.ReqAwards
	}
	return nil
}

func (m *CrossAwardOpRs) GetRollBack() bool {
	if m != nil && m.RollBack != nil {
		return *m.RollBack
	}
	return false
}

func (m *CrossAwardOpRs) GetCmd() int32 {
	if m != nil && m.Cmd != nil {
		return *m.Cmd
	}
	return 0
}

var E_CrossAwardOpRs_Ext = &proto.ExtensionDesc{
	ExtendedType:  (*Base)(nil),
	ExtensionType: (*CrossAwardOpRs)(nil),
	Field:         10011,
	Name:          "CrossAwardOpRs.ext",
	Tag:           "bytes,10011,opt,name=ext",
	Filename:      "Cross.proto",
}

// 将领复活操作
type CrossHeroReviveRq struct {
	Hero                 *CrossHeroPb `protobuf:"bytes,1,opt,name=hero" json:"hero,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CrossHeroReviveRq) Reset()         { *m = CrossHeroReviveRq{} }
func (m *CrossHeroReviveRq) String() string { return proto.CompactTextString(m) }
func (*CrossHeroReviveRq) ProtoMessage()    {}
func (*CrossHeroReviveRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e60e98cf7de90f, []int{9}
}
func (m *CrossHeroReviveRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrossHeroReviveRq.Unmarshal(m, b)
}
func (m *CrossHeroReviveRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrossHeroReviveRq.Marshal(b, m, deterministic)
}
func (m *CrossHeroReviveRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossHeroReviveRq.Merge(m, src)
}
func (m *CrossHeroReviveRq) XXX_Size() int {
	return xxx_messageInfo_CrossHeroReviveRq.Size(m)
}
func (m *CrossHeroReviveRq) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossHeroReviveRq.DiscardUnknown(m)
}

var xxx_messageInfo_CrossHeroReviveRq proto.InternalMessageInfo

func (m *CrossHeroReviveRq) GetHero() *CrossHeroPb {
	if m != nil {
		return m.Hero
	}
	return nil
}

var E_CrossHeroReviveRq_Ext = &proto.ExtensionDesc{
	ExtendedType:  (*Base)(nil),
	ExtensionType: (*CrossHeroReviveRq)(nil),
	Field:         10013,
	Name:          "CrossHeroReviveRq.ext",
	Tag:           "bytes,10013,opt,name=ext",
	Filename:      "Cross.proto",
}

// 将领的属性同步
type CrossHeroSyncRq struct {
	Hero                 []*CrossHeroPb `protobuf:"bytes,1,rep,name=hero" json:"hero,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CrossHeroSyncRq) Reset()         { *m = CrossHeroSyncRq{} }
func (m *CrossHeroSyncRq) String() string { return proto.CompactTextString(m) }
func (*CrossHeroSyncRq) ProtoMessage()    {}
func (*CrossHeroSyncRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e60e98cf7de90f, []int{10}
}
func (m *CrossHeroSyncRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrossHeroSyncRq.Unmarshal(m, b)
}
func (m *CrossHeroSyncRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrossHeroSyncRq.Marshal(b, m, deterministic)
}
func (m *CrossHeroSyncRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossHeroSyncRq.Merge(m, src)
}
func (m *CrossHeroSyncRq) XXX_Size() int {
	return xxx_messageInfo_CrossHeroSyncRq.Size(m)
}
func (m *CrossHeroSyncRq) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossHeroSyncRq.DiscardUnknown(m)
}

var xxx_messageInfo_CrossHeroSyncRq proto.InternalMessageInfo

func (m *CrossHeroSyncRq) GetHero() []*CrossHeroPb {
	if m != nil {
		return m.Hero
	}
	return nil
}

var E_CrossHeroSyncRq_Ext = &proto.ExtensionDesc{
	ExtendedType:  (*Base)(nil),
	ExtensionType: (*CrossHeroSyncRq)(nil),
	Field:         10015,
	Name:          "CrossHeroSyncRq.ext",
	Tag:           "bytes,10015,opt,name=ext",
	Filename:      "Cross.proto",
}

// 跨服开始结束推送给游戏服
type CrossStartFinishRs struct {
	IsStart              *bool    `protobuf:"varint,1,req,name=isStart" json:"isStart,omitempty"`
	WinCamp              *int32   `protobuf:"varint,2,opt,name=winCamp" json:"winCamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CrossStartFinishRs) Reset()         { *m = CrossStartFinishRs{} }
func (m *CrossStartFinishRs) String() string { return proto.CompactTextString(m) }
func (*CrossStartFinishRs) ProtoMessage()    {}
func (*CrossStartFinishRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e60e98cf7de90f, []int{11}
}
func (m *CrossStartFinishRs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrossStartFinishRs.Unmarshal(m, b)
}
func (m *CrossStartFinishRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrossStartFinishRs.Marshal(b, m, deterministic)
}
func (m *CrossStartFinishRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossStartFinishRs.Merge(m, src)
}
func (m *CrossStartFinishRs) XXX_Size() int {
	return xxx_messageInfo_CrossStartFinishRs.Size(m)
}
func (m *CrossStartFinishRs) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossStartFinishRs.DiscardUnknown(m)
}

var xxx_messageInfo_CrossStartFinishRs proto.InternalMessageInfo

func (m *CrossStartFinishRs) GetIsStart() bool {
	if m != nil && m.IsStart != nil {
		return *m.IsStart
	}
	return false
}

func (m *CrossStartFinishRs) GetWinCamp() int32 {
	if m != nil && m.WinCamp != nil {
		return *m.WinCamp
	}
	return 0
}

var E_CrossStartFinishRs_Ext = &proto.ExtensionDesc{
	ExtendedType:  (*Base)(nil),
	ExtensionType: (*CrossStartFinishRs)(nil),
	Field:         10016,
	Name:          "CrossStartFinishRs.ext",
	Tag:           "bytes,10016,opt,name=ext",
	Filename:      "Cross.proto",
}

// GM命令
type GmDoSomeRq struct {
	Strs                 []string `protobuf:"bytes,1,rep,name=strs" json:"strs,omitempty"`
	RoleId               *int64   `protobuf:"varint,2,opt,name=roleId" json:"roleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GmDoSomeRq) Reset()         { *m = GmDoSomeRq{} }
func (m *GmDoSomeRq) String() string { return proto.CompactTextString(m) }
func (*GmDoSomeRq) ProtoMessage()    {}
func (*GmDoSomeRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e60e98cf7de90f, []int{12}
}
func (m *GmDoSomeRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GmDoSomeRq.Unmarshal(m, b)
}
func (m *GmDoSomeRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GmDoSomeRq.Marshal(b, m, deterministic)
}
func (m *GmDoSomeRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GmDoSomeRq.Merge(m, src)
}
func (m *GmDoSomeRq) XXX_Size() int {
	return xxx_messageInfo_GmDoSomeRq.Size(m)
}
func (m *GmDoSomeRq) XXX_DiscardUnknown() {
	xxx_messageInfo_GmDoSomeRq.DiscardUnknown(m)
}

var xxx_messageInfo_GmDoSomeRq proto.InternalMessageInfo

func (m *GmDoSomeRq) GetStrs() []string {
	if m != nil {
		return m.Strs
	}
	return nil
}

func (m *GmDoSomeRq) GetRoleId() int64 {
	if m != nil && m.RoleId != nil {
		return *m.RoleId
	}
	return 0
}

var E_GmDoSomeRq_Ext = &proto.ExtensionDesc{
	ExtendedType:  (*Base)(nil),
	ExtensionType: (*GmDoSomeRq)(nil),
	Field:         10017,
	Name:          "GmDoSomeRq.ext",
	Tag:           "bytes,10017,opt,name=ext",
	Filename:      "Cross.proto",
}

func init() {
	proto.RegisterExtension(E_ServerRegistRq_Ext)
	proto.RegisterType((*ServerRegistRq)(nil), "ServerRegistRq")
	proto.RegisterExtension(E_HeartRq_Ext)
	proto.RegisterType((*HeartRq)(nil), "HeartRq")
	proto.RegisterExtension(E_HeartRs_Ext)
	proto.RegisterType((*HeartRs)(nil), "HeartRs")
	proto.RegisterExtension(E_CrossLoginRq_Ext)
	proto.RegisterType((*CrossLoginRq)(nil), "CrossLoginRq")
	proto.RegisterExtension(E_CrossLoginRs_Ext)
	proto.RegisterType((*CrossLoginRs)(nil), "CrossLoginRs")
	proto.RegisterExtension(E_SendCrossChatRq_Ext)
	proto.RegisterType((*SendCrossChatRq)(nil), "SendCrossChatRq")
	proto.RegisterExtension(E_ChoiceHeroRq_Ext)
	proto.RegisterType((*ChoiceHeroRq)(nil), "ChoiceHeroRq")
	proto.RegisterExtension(E_CrossAwardOpRq_Ext)
	proto.RegisterType((*CrossAwardOpRq)(nil), "CrossAwardOpRq")
	proto.RegisterExtension(E_CrossAwardOpRs_Ext)
	proto.RegisterType((*CrossAwardOpRs)(nil), "CrossAwardOpRs")
	proto.RegisterExtension(E_CrossHeroReviveRq_Ext)
	proto.RegisterType((*CrossHeroReviveRq)(nil), "CrossHeroReviveRq")
	proto.RegisterExtension(E_CrossHeroSyncRq_Ext)
	proto.RegisterType((*CrossHeroSyncRq)(nil), "CrossHeroSyncRq")
	proto.RegisterExtension(E_CrossStartFinishRs_Ext)
	proto.RegisterType((*CrossStartFinishRs)(nil), "CrossStartFinishRs")
	proto.RegisterExtension(E_GmDoSomeRq_Ext)
	proto.RegisterType((*GmDoSomeRq)(nil), "GmDoSomeRq")
}

func init() { proto.RegisterFile("Cross.proto", fileDescriptor_91e60e98cf7de90f) }

var fileDescriptor_91e60e98cf7de90f = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x5d, 0x6f, 0x12, 0x4f,
	0x14, 0xc6, 0x03, 0x0b, 0x2d, 0x3d, 0x50, 0xe8, 0x7f, 0x7a, 0xb3, 0xff, 0x1a, 0x0d, 0x99, 0x68,
	0x24, 0x26, 0x0e, 0xc9, 0x5e, 0x78, 0x2f, 0x98, 0xda, 0xc6, 0xda, 0xe2, 0xd2, 0x0b, 0xe3, 0x95,
	0xfb, 0x32, 0xb2, 0x93, 0xee, 0xee, 0xb0, 0x33, 0x03, 0x14, 0xbf, 0x80, 0xd7, 0xc6, 0xb7, 0xa8,
	0x31, 0xea, 0x37, 0x35, 0x3b, 0x0b, 0xb5, 0xb3, 0xf4, 0xca, 0x3b, 0xe6, 0xcc, 0x3c, 0xbf, 0xe7,
	0xf0, 0x9c, 0x03, 0xd0, 0x1c, 0x0a, 0x2e, 0x25, 0x99, 0x0a, 0xae, 0xf8, 0x01, 0x0c, 0x3c, 0x49,
	0x57, 0x9f, 0x5b, 0x43, 0x9e, 0x24, 0x3c, 0x5d, 0x9d, 0x3a, 0x63, 0x2a, 0x98, 0x17, 0xb3, 0xb7,
	0xab, 0x6b, 0xfc, 0x06, 0xda, 0x63, 0x2a, 0xe6, 0x54, 0xb8, 0x74, 0xc2, 0xa4, 0x72, 0x33, 0xb4,
	0x07, 0x0d, 0xa9, 0x2b, 0xc7, 0xa1, 0x5d, 0xe9, 0x56, 0x7b, 0x75, 0xb4, 0x0b, 0x75, 0x96, 0x86,
	0xf4, 0xd2, 0xae, 0xae, 0x8f, 0x8a, 0x2b, 0x2f, 0xb6, 0xad, 0xfc, 0xe8, 0xdc, 0x05, 0x8b, 0x5e,
	0x2a, 0x54, 0x27, 0xb9, 0xa9, 0xfd, 0xfe, 0xb4, 0x5b, 0xe9, 0x35, 0x9d, 0x0e, 0x31, 0xa9, 0xb8,
	0x07, 0xdb, 0x47, 0xd4, 0x13, 0xca, 0xcd, 0x9c, 0xdb, 0x86, 0xe0, 0x43, 0x21, 0x68, 0x90, 0xd5,
	0x35, 0x7e, 0xb4, 0x7e, 0x29, 0x51, 0x0b, 0x6a, 0x8a, 0x25, 0x54, 0xb7, 0x61, 0x95, 0x74, 0x1f,
	0x4d, 0x9d, 0xc4, 0x3e, 0xb4, 0x74, 0x06, 0x27, 0x7c, 0xc2, 0x52, 0x37, 0x43, 0x6d, 0xd8, 0xe2,
	0xd3, 0xf3, 0xe5, 0x34, 0x97, 0x57, 0x7a, 0x75, 0x74, 0x07, 0xb6, 0xa6, 0xb1, 0xb7, 0xa4, 0xc2,
	0xae, 0x6a, 0x65, 0x9b, 0xe8, 0xe7, 0x23, 0x5d, 0x1b, 0xf9, 0x0e, 0x36, 0xf0, 0x9f, 0x0a, 0xfc,
	0x2e, 0xb9, 0xce, 0xc4, 0x99, 0xe1, 0x21, 0x37, 0x3c, 0xfe, 0x87, 0x5a, 0x44, 0x05, 0xb7, 0xab,
	0x5d, 0xab, 0xd7, 0x74, 0x9a, 0xe4, 0x90, 0x0b, 0x75, 0x44, 0x05, 0x1f, 0xf9, 0x08, 0x01, 0x04,
	0x33, 0xf1, 0x8c, 0xc5, 0xf1, 0xe9, 0x2c, 0xb1, 0xad, 0xfc, 0x79, 0xc9, 0xf2, 0xf3, 0x0d, 0x96,
	0x12, 0x3f, 0x87, 0xce, 0x98, 0xa6, 0xa1, 0xae, 0x0d, 0x23, 0x2f, 0x9f, 0xd0, 0x3e, 0xd4, 0x82,
	0xc8, 0x53, 0x3a, 0x96, 0xa6, 0x53, 0x27, 0x79, 0xd9, 0xb9, 0x67, 0xb0, 0xbe, 0x14, 0xac, 0x3d,
	0x52, 0xd2, 0xe2, 0x33, 0x68, 0x0d, 0x23, 0xce, 0x02, 0x9a, 0xb7, 0xe5, 0x66, 0xe8, 0x16, 0xd4,
	0xf3, 0x8e, 0xa5, 0x5d, 0xd1, 0x2d, 0xb7, 0x0a, 0xf3, 0xa2, 0xe7, 0x52, 0x7f, 0x5f, 0xaf, 0xfa,
	0xbb, 0x06, 0xc0, 0x4b, 0x68, 0x6b, 0xc9, 0xe3, 0x85, 0x27, 0xc2, 0xb3, 0x69, 0x11, 0xbc, 0xf2,
	0xe4, 0x85, 0x5e, 0x9f, 0x4a, 0xcf, 0x42, 0x5d, 0xd8, 0x11, 0x34, 0xd3, 0xf7, 0x72, 0x95, 0xfd,
	0xea, 0x3b, 0xea, 0xda, 0xc8, 0x47, 0x1d, 0xd8, 0x96, 0xb3, 0x20, 0xa0, 0x52, 0xea, 0x60, 0x1a,
	0xa5, 0x9d, 0xfa, 0xb6, 0xde, 0x29, 0xd3, 0x08, 0xbf, 0xab, 0x94, 0xbc, 0xe5, 0x3f, 0x78, 0xef,
	0x41, 0x43, 0xf0, 0x38, 0x1e, 0x78, 0xc1, 0x45, 0x61, 0x8e, 0x9a, 0x60, 0x05, 0x49, 0x68, 0xd7,
	0xf4, 0x88, 0xcc, 0x4e, 0xbe, 0xdf, 0xd8, 0x89, 0xc4, 0x2f, 0xe1, 0xbf, 0xab, 0xdc, 0x5c, 0x3a,
	0x67, 0x73, 0xea, 0x66, 0xe8, 0x60, 0xb5, 0x0c, 0x15, 0xad, 0x31, 0x93, 0xbd, 0x6f, 0x60, 0x7f,
	0x14, 0x58, 0x44, 0x36, 0x20, 0xf8, 0x1c, 0x3a, 0x57, 0xc5, 0xf1, 0x32, 0x0d, 0x0c, 0xee, 0xe6,
	0xc4, 0xcc, 0x2d, 0xf8, 0xb9, 0xde, 0x82, 0x12, 0x02, 0xbf, 0x06, 0xa4, 0x4b, 0x63, 0xe5, 0x09,
	0x75, 0xc8, 0x52, 0x26, 0x23, 0x57, 0xe6, 0x63, 0x60, 0x45, 0x49, 0xaf, 0x56, 0x23, 0x2f, 0x2c,
	0x58, 0x3a, 0xf4, 0x92, 0xa9, 0xce, 0xae, 0xee, 0xf4, 0x0c, 0xfc, 0xaf, 0x02, 0xbf, 0x4f, 0x36,
	0x59, 0xf8, 0x04, 0xe0, 0x69, 0xf2, 0x84, 0x8f, 0x79, 0x92, 0x47, 0xd1, 0x82, 0x9a, 0x54, 0xa2,
	0x58, 0xb2, 0x9d, 0x7c, 0x48, 0x82, 0xc7, 0xf4, 0x38, 0xd4, 0x54, 0xcb, 0xe9, 0x1a, 0xd4, 0xdf,
	0x05, 0xb5, 0x49, 0xfe, 0xea, 0x07, 0x2f, 0xe0, 0x20, 0xe0, 0x09, 0x89, 0x66, 0x69, 0x28, 0x68,
	0x18, 0xd0, 0x54, 0x91, 0x89, 0x97, 0x50, 0x92, 0x2c, 0xc8, 0xd4, 0x1f, 0x6c, 0x17, 0x3f, 0x64,
	0xff, 0xd5, 0x83, 0x09, 0x53, 0xd1, 0xcc, 0x27, 0x01, 0x4f, 0xfa, 0x19, 0x9b, 0x45, 0x7c, 0x16,
	0xd2, 0xfe, 0x84, 0x3f, 0x5c, 0x50, 0xbf, 0xaf, 0xff, 0xee, 0xfa, 0x73, 0xa7, 0x9f, 0xf0, 0x90,
	0xc6, 0xf2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xcd, 0xad, 0x9b, 0x31, 0x05, 0x00, 0x00,
}
