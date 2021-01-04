// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: game/rune.proto

package game

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

////////////////////////////////////////////////
// Rune
type RuneAtt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttType  int32 `protobuf:"varint,1,opt,name=AttType,proto3" json:"AttType,omitempty"`
	AttValue int64 `protobuf:"varint,2,opt,name=AttValue,proto3" json:"AttValue,omitempty"`
}

func (x *RuneAtt) Reset() {
	*x = RuneAtt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rune_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuneAtt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuneAtt) ProtoMessage() {}

func (x *RuneAtt) ProtoReflect() protoreflect.Message {
	mi := &file_game_rune_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuneAtt.ProtoReflect.Descriptor instead.
func (*RuneAtt) Descriptor() ([]byte, []int) {
	return file_game_rune_proto_rawDescGZIP(), []int{0}
}

func (x *RuneAtt) GetAttType() int32 {
	if x != nil {
		return x.AttType
	}
	return 0
}

func (x *RuneAtt) GetAttValue() int64 {
	if x != nil {
		return x.AttValue
	}
	return 0
}

type Rune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64      `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	TypeId     int32      `protobuf:"varint,2,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
	EquipObjId int64      `protobuf:"varint,3,opt,name=EquipObjId,proto3" json:"EquipObjId,omitempty"` // 装备者objid
	Atts       []*RuneAtt `protobuf:"bytes,4,rep,name=Atts,proto3" json:"Atts,omitempty"`              // 1条主属性+5条副属性
}

func (x *Rune) Reset() {
	*x = Rune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rune_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rune) ProtoMessage() {}

func (x *Rune) ProtoReflect() protoreflect.Message {
	mi := &file_game_rune_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rune.ProtoReflect.Descriptor instead.
func (*Rune) Descriptor() ([]byte, []int) {
	return file_game_rune_proto_rawDescGZIP(), []int{1}
}

func (x *Rune) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Rune) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *Rune) GetEquipObjId() int64 {
	if x != nil {
		return x.EquipObjId
	}
	return 0
}

func (x *Rune) GetAtts() []*RuneAtt {
	if x != nil {
		return x.Atts
	}
	return nil
}

type C2M_AddRune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId int32 `protobuf:"varint,1,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
}

func (x *C2M_AddRune) Reset() {
	*x = C2M_AddRune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rune_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_AddRune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_AddRune) ProtoMessage() {}

func (x *C2M_AddRune) ProtoReflect() protoreflect.Message {
	mi := &file_game_rune_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_AddRune.ProtoReflect.Descriptor instead.
func (*C2M_AddRune) Descriptor() ([]byte, []int) {
	return file_game_rune_proto_rawDescGZIP(), []int{2}
}

func (x *C2M_AddRune) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

type C2M_DelRune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *C2M_DelRune) Reset() {
	*x = C2M_DelRune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rune_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_DelRune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_DelRune) ProtoMessage() {}

func (x *C2M_DelRune) ProtoReflect() protoreflect.Message {
	mi := &file_game_rune_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_DelRune.ProtoReflect.Descriptor instead.
func (*C2M_DelRune) Descriptor() ([]byte, []int) {
	return file_game_rune_proto_rawDescGZIP(), []int{3}
}

func (x *C2M_DelRune) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type C2M_QueryRunes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2M_QueryRunes) Reset() {
	*x = C2M_QueryRunes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rune_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_QueryRunes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_QueryRunes) ProtoMessage() {}

func (x *C2M_QueryRunes) ProtoReflect() protoreflect.Message {
	mi := &file_game_rune_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_QueryRunes.ProtoReflect.Descriptor instead.
func (*C2M_QueryRunes) Descriptor() ([]byte, []int) {
	return file_game_rune_proto_rawDescGZIP(), []int{4}
}

type M2C_RuneList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runes []*Rune `protobuf:"bytes,1,rep,name=runes,proto3" json:"runes,omitempty"`
}

func (x *M2C_RuneList) Reset() {
	*x = M2C_RuneList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rune_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_RuneList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_RuneList) ProtoMessage() {}

func (x *M2C_RuneList) ProtoReflect() protoreflect.Message {
	mi := &file_game_rune_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_RuneList.ProtoReflect.Descriptor instead.
func (*M2C_RuneList) Descriptor() ([]byte, []int) {
	return file_game_rune_proto_rawDescGZIP(), []int{5}
}

func (x *M2C_RuneList) GetRunes() []*Rune {
	if x != nil {
		return x.Runes
	}
	return nil
}

type M2C_RuneAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rune *Rune `protobuf:"bytes,1,opt,name=rune,proto3" json:"rune,omitempty"`
}

func (x *M2C_RuneAdd) Reset() {
	*x = M2C_RuneAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rune_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_RuneAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_RuneAdd) ProtoMessage() {}

func (x *M2C_RuneAdd) ProtoReflect() protoreflect.Message {
	mi := &file_game_rune_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_RuneAdd.ProtoReflect.Descriptor instead.
func (*M2C_RuneAdd) Descriptor() ([]byte, []int) {
	return file_game_rune_proto_rawDescGZIP(), []int{6}
}

func (x *M2C_RuneAdd) GetRune() *Rune {
	if x != nil {
		return x.Rune
	}
	return nil
}

type M2C_RuneUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rune *Rune `protobuf:"bytes,1,opt,name=rune,proto3" json:"rune,omitempty"`
}

func (x *M2C_RuneUpdate) Reset() {
	*x = M2C_RuneUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rune_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_RuneUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_RuneUpdate) ProtoMessage() {}

func (x *M2C_RuneUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_game_rune_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_RuneUpdate.ProtoReflect.Descriptor instead.
func (*M2C_RuneUpdate) Descriptor() ([]byte, []int) {
	return file_game_rune_proto_rawDescGZIP(), []int{7}
}

func (x *M2C_RuneUpdate) GetRune() *Rune {
	if x != nil {
		return x.Rune
	}
	return nil
}

type M2C_DelRune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuneId int64 `protobuf:"varint,1,opt,name=RuneId,proto3" json:"RuneId,omitempty"`
}

func (x *M2C_DelRune) Reset() {
	*x = M2C_DelRune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rune_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_DelRune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_DelRune) ProtoMessage() {}

func (x *M2C_DelRune) ProtoReflect() protoreflect.Message {
	mi := &file_game_rune_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_DelRune.ProtoReflect.Descriptor instead.
func (*M2C_DelRune) Descriptor() ([]byte, []int) {
	return file_game_rune_proto_rawDescGZIP(), []int{8}
}

func (x *M2C_DelRune) GetRuneId() int64 {
	if x != nil {
		return x.RuneId
	}
	return 0
}

type C2M_PutonRune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	RuneId int64 `protobuf:"varint,2,opt,name=RuneId,proto3" json:"RuneId,omitempty"`
}

func (x *C2M_PutonRune) Reset() {
	*x = C2M_PutonRune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rune_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_PutonRune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_PutonRune) ProtoMessage() {}

func (x *C2M_PutonRune) ProtoReflect() protoreflect.Message {
	mi := &file_game_rune_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_PutonRune.ProtoReflect.Descriptor instead.
func (*C2M_PutonRune) Descriptor() ([]byte, []int) {
	return file_game_rune_proto_rawDescGZIP(), []int{9}
}

func (x *C2M_PutonRune) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *C2M_PutonRune) GetRuneId() int64 {
	if x != nil {
		return x.RuneId
	}
	return 0
}

type C2M_TakeoffRune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	Pos    int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (x *C2M_TakeoffRune) Reset() {
	*x = C2M_TakeoffRune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rune_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_TakeoffRune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_TakeoffRune) ProtoMessage() {}

func (x *C2M_TakeoffRune) ProtoReflect() protoreflect.Message {
	mi := &file_game_rune_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_TakeoffRune.ProtoReflect.Descriptor instead.
func (*C2M_TakeoffRune) Descriptor() ([]byte, []int) {
	return file_game_rune_proto_rawDescGZIP(), []int{10}
}

func (x *C2M_TakeoffRune) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *C2M_TakeoffRune) GetPos() int32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

var File_game_rune_proto protoreflect.FileDescriptor

var file_game_rune_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x72, 0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x65, 0x41,
	0x74, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x74, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x41, 0x74, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x41, 0x74, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x41, 0x74, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x71, 0x0a, 0x04, 0x52, 0x75, 0x6e, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x71, 0x75, 0x69,
	0x70, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x45, 0x71,
	0x75, 0x69, 0x70, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x41, 0x74, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x75,
	0x6e, 0x65, 0x41, 0x74, 0x74, 0x52, 0x04, 0x41, 0x74, 0x74, 0x73, 0x22, 0x25, 0x0a, 0x0b, 0x43,
	0x32, 0x4d, 0x5f, 0x41, 0x64, 0x64, 0x52, 0x75, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x22, 0x1d, 0x0a, 0x0b, 0x43, 0x32, 0x4d, 0x5f, 0x44, 0x65, 0x6c, 0x52, 0x75, 0x6e,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x64, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x32, 0x4d, 0x5f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x75,
	0x6e, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x0c, 0x4d, 0x32, 0x43, 0x5f, 0x52, 0x75, 0x6e, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x65, 0x52, 0x05,
	0x72, 0x75, 0x6e, 0x65, 0x73, 0x22, 0x2d, 0x0a, 0x0b, 0x4d, 0x32, 0x43, 0x5f, 0x52, 0x75, 0x6e,
	0x65, 0x41, 0x64, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x65, 0x52, 0x04,
	0x72, 0x75, 0x6e, 0x65, 0x22, 0x30, 0x0a, 0x0e, 0x4d, 0x32, 0x43, 0x5f, 0x52, 0x75, 0x6e, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x65,
	0x52, 0x04, 0x72, 0x75, 0x6e, 0x65, 0x22, 0x25, 0x0a, 0x0b, 0x4d, 0x32, 0x43, 0x5f, 0x44, 0x65,
	0x6c, 0x52, 0x75, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x75, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x3f, 0x0a,
	0x0d, 0x43, 0x32, 0x4d, 0x5f, 0x50, 0x75, 0x74, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x75, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x3b,
	0x0a, 0x0f, 0x43, 0x32, 0x4d, 0x5f, 0x54, 0x61, 0x6b, 0x65, 0x6f, 0x66, 0x66, 0x52, 0x75, 0x6e,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x6f, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x6f, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x65,
	0x64, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_rune_proto_rawDescOnce sync.Once
	file_game_rune_proto_rawDescData = file_game_rune_proto_rawDesc
)

func file_game_rune_proto_rawDescGZIP() []byte {
	file_game_rune_proto_rawDescOnce.Do(func() {
		file_game_rune_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_rune_proto_rawDescData)
	})
	return file_game_rune_proto_rawDescData
}

var file_game_rune_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_game_rune_proto_goTypes = []interface{}{
	(*RuneAtt)(nil),         // 0: game.RuneAtt
	(*Rune)(nil),            // 1: game.Rune
	(*C2M_AddRune)(nil),     // 2: game.C2M_AddRune
	(*C2M_DelRune)(nil),     // 3: game.C2M_DelRune
	(*C2M_QueryRunes)(nil),  // 4: game.C2M_QueryRunes
	(*M2C_RuneList)(nil),    // 5: game.M2C_RuneList
	(*M2C_RuneAdd)(nil),     // 6: game.M2C_RuneAdd
	(*M2C_RuneUpdate)(nil),  // 7: game.M2C_RuneUpdate
	(*M2C_DelRune)(nil),     // 8: game.M2C_DelRune
	(*C2M_PutonRune)(nil),   // 9: game.C2M_PutonRune
	(*C2M_TakeoffRune)(nil), // 10: game.C2M_TakeoffRune
}
var file_game_rune_proto_depIdxs = []int32{
	0, // 0: game.Rune.Atts:type_name -> game.RuneAtt
	1, // 1: game.M2C_RuneList.runes:type_name -> game.Rune
	1, // 2: game.M2C_RuneAdd.rune:type_name -> game.Rune
	1, // 3: game.M2C_RuneUpdate.rune:type_name -> game.Rune
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_game_rune_proto_init() }
func file_game_rune_proto_init() {
	if File_game_rune_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_rune_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuneAtt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_rune_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rune); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_rune_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_AddRune); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_rune_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_DelRune); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_rune_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_QueryRunes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_rune_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_RuneList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_rune_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_RuneAdd); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_rune_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_RuneUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_rune_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_DelRune); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_rune_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_PutonRune); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_rune_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_TakeoffRune); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_rune_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_rune_proto_goTypes,
		DependencyIndexes: file_game_rune_proto_depIdxs,
		MessageInfos:      file_game_rune_proto_msgTypes,
	}.Build()
	File_game_rune_proto = out.File
	file_game_rune_proto_rawDesc = nil
	file_game_rune_proto_goTypes = nil
	file_game_rune_proto_depIdxs = nil
}
