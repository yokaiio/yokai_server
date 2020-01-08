// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pubsub/pubsub.proto

package pubsub

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	account "github.com/yokaiio/yokai_server/proto/account"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

/////////////////////////////////////////////////
// pub/sub
/////////////////////////////////////////////////
type PubStartGate struct {
	Info                 *account.LiteAccount `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PubStartGate) Reset()         { *m = PubStartGate{} }
func (m *PubStartGate) String() string { return proto.CompactTextString(m) }
func (*PubStartGate) ProtoMessage()    {}
func (*PubStartGate) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce310d0bb9f289ed, []int{0}
}

func (m *PubStartGate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubStartGate.Unmarshal(m, b)
}
func (m *PubStartGate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubStartGate.Marshal(b, m, deterministic)
}
func (m *PubStartGate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubStartGate.Merge(m, src)
}
func (m *PubStartGate) XXX_Size() int {
	return xxx_messageInfo_PubStartGate.Size(m)
}
func (m *PubStartGate) XXX_DiscardUnknown() {
	xxx_messageInfo_PubStartGate.DiscardUnknown(m)
}

var xxx_messageInfo_PubStartGate proto.InternalMessageInfo

func (m *PubStartGate) GetInfo() *account.LiteAccount {
	if m != nil {
		return m.Info
	}
	return nil
}

type PubGateResult struct {
	Info                 *account.LiteAccount `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Win                  bool                 `protobuf:"varint,2,opt,name=win,proto3" json:"win,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PubGateResult) Reset()         { *m = PubGateResult{} }
func (m *PubGateResult) String() string { return proto.CompactTextString(m) }
func (*PubGateResult) ProtoMessage()    {}
func (*PubGateResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce310d0bb9f289ed, []int{1}
}

func (m *PubGateResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubGateResult.Unmarshal(m, b)
}
func (m *PubGateResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubGateResult.Marshal(b, m, deterministic)
}
func (m *PubGateResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubGateResult.Merge(m, src)
}
func (m *PubGateResult) XXX_Size() int {
	return xxx_messageInfo_PubGateResult.Size(m)
}
func (m *PubGateResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PubGateResult.DiscardUnknown(m)
}

var xxx_messageInfo_PubGateResult proto.InternalMessageInfo

func (m *PubGateResult) GetInfo() *account.LiteAccount {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *PubGateResult) GetWin() bool {
	if m != nil {
		return m.Win
	}
	return false
}

type PubExpirePlayer struct {
	PlayerId             int64    `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	GameId               int32    `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubExpirePlayer) Reset()         { *m = PubExpirePlayer{} }
func (m *PubExpirePlayer) String() string { return proto.CompactTextString(m) }
func (*PubExpirePlayer) ProtoMessage()    {}
func (*PubExpirePlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce310d0bb9f289ed, []int{2}
}

func (m *PubExpirePlayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubExpirePlayer.Unmarshal(m, b)
}
func (m *PubExpirePlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubExpirePlayer.Marshal(b, m, deterministic)
}
func (m *PubExpirePlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubExpirePlayer.Merge(m, src)
}
func (m *PubExpirePlayer) XXX_Size() int {
	return xxx_messageInfo_PubExpirePlayer.Size(m)
}
func (m *PubExpirePlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_PubExpirePlayer.DiscardUnknown(m)
}

var xxx_messageInfo_PubExpirePlayer proto.InternalMessageInfo

func (m *PubExpirePlayer) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *PubExpirePlayer) GetGameId() int32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

type PubExpireLitePlayer struct {
	PlayerId             int64    `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	GameId               int32    `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubExpireLitePlayer) Reset()         { *m = PubExpireLitePlayer{} }
func (m *PubExpireLitePlayer) String() string { return proto.CompactTextString(m) }
func (*PubExpireLitePlayer) ProtoMessage()    {}
func (*PubExpireLitePlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce310d0bb9f289ed, []int{3}
}

func (m *PubExpireLitePlayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubExpireLitePlayer.Unmarshal(m, b)
}
func (m *PubExpireLitePlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubExpireLitePlayer.Marshal(b, m, deterministic)
}
func (m *PubExpireLitePlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubExpireLitePlayer.Merge(m, src)
}
func (m *PubExpireLitePlayer) XXX_Size() int {
	return xxx_messageInfo_PubExpireLitePlayer.Size(m)
}
func (m *PubExpireLitePlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_PubExpireLitePlayer.DiscardUnknown(m)
}

var xxx_messageInfo_PubExpireLitePlayer proto.InternalMessageInfo

func (m *PubExpireLitePlayer) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *PubExpireLitePlayer) GetGameId() int32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func init() {
	proto.RegisterType((*PubStartGate)(nil), "yokai_pubsub.PubStartGate")
	proto.RegisterType((*PubGateResult)(nil), "yokai_pubsub.PubGateResult")
	proto.RegisterType((*PubExpirePlayer)(nil), "yokai_pubsub.PubExpirePlayer")
	proto.RegisterType((*PubExpireLitePlayer)(nil), "yokai_pubsub.PubExpireLitePlayer")
}

func init() { proto.RegisterFile("pubsub/pubsub.proto", fileDescriptor_ce310d0bb9f289ed) }

var fileDescriptor_ce310d0bb9f289ed = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x59, 0xab, 0xb5, 0x8e, 0x15, 0x25, 0x45, 0x2c, 0xf5, 0x52, 0xf6, 0xd4, 0x83, 0x64,
	0x41, 0xef, 0x82, 0x82, 0x94, 0xa2, 0x87, 0x35, 0xde, 0xbc, 0x2c, 0xc9, 0x6e, 0xac, 0xc1, 0x76,
	0x13, 0xb2, 0x33, 0x6a, 0xff, 0xbd, 0x24, 0xa9, 0xde, 0xa5, 0xa7, 0xf7, 0x66, 0x32, 0xef, 0x0b,
	0xc9, 0xc0, 0xc8, 0x91, 0xea, 0x48, 0x15, 0x49, 0xb8, 0xf3, 0x16, 0x2d, 0x1b, 0x6e, 0xec, 0x87,
	0x34, 0x55, 0xea, 0x4d, 0xce, 0x65, 0x5d, 0x5b, 0x6a, 0xb1, 0xd8, 0x6a, 0x1a, 0xca, 0x6f, 0x61,
	0x58, 0x92, 0x7a, 0x41, 0xe9, 0x71, 0x2e, 0x51, 0x33, 0x0e, 0xfb, 0xa6, 0x7d, 0xb3, 0xe3, 0x6c,
	0x9a, 0xcd, 0x8e, 0xaf, 0x27, 0x3c, 0x31, 0x7e, 0x33, 0x4f, 0x06, 0xf5, 0x5d, 0xf2, 0x22, 0xce,
	0xe5, 0xcf, 0x70, 0x52, 0x92, 0x0a, 0x51, 0xa1, 0x3b, 0x5a, 0xe1, 0x7f, 0x01, 0xec, 0x0c, 0x7a,
	0x5f, 0xa6, 0x1d, 0xef, 0x4d, 0xb3, 0xd9, 0x40, 0x04, 0x9b, 0xcf, 0xe1, 0xb4, 0x24, 0xf5, 0xf0,
	0xed, 0x8c, 0xd7, 0xe5, 0x4a, 0x6e, 0xb4, 0x67, 0x97, 0x70, 0xe4, 0xa2, 0xab, 0x4c, 0x13, 0xc9,
	0x3d, 0x31, 0x48, 0x8d, 0x45, 0xc3, 0x2e, 0xe0, 0x70, 0x29, 0xd7, 0x3a, 0x1c, 0x05, 0xca, 0x81,
	0xe8, 0x87, 0x72, 0xd1, 0xe4, 0x8f, 0x30, 0xfa, 0x03, 0x85, 0x8b, 0x77, 0x81, 0xdd, 0xf3, 0xd7,
	0xab, 0xa5, 0xc1, 0x77, 0x52, 0xbc, 0xb6, 0xeb, 0x22, 0xbe, 0xca, 0xd8, 0xa4, 0x55, 0xa7, 0xfd,
	0xa7, 0xf6, 0x45, 0xfc, 0xd1, 0xed, 0x0e, 0x54, 0x3f, 0x56, 0x37, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x5a, 0x9c, 0x90, 0x71, 0x9b, 0x01, 0x00, 0x00,
}
