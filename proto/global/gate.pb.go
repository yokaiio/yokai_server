// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: gate.proto

package global

import (
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

type CmdType int32

const (
	CmdType_NEW    CmdType = 0 // new connect
	CmdType_RECONN CmdType = 1 // re connect
)

// Enum value maps for CmdType.
var (
	CmdType_name = map[int32]string{
		0: "NEW",
		1: "RECONN",
	}
	CmdType_value = map[string]int32{
		"NEW":    0,
		"RECONN": 1,
	}
)

func (x CmdType) Enum() *CmdType {
	p := new(CmdType)
	*p = x
	return p
}

func (x CmdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CmdType) Descriptor() protoreflect.EnumDescriptor {
	return file_gate_proto_enumTypes[0].Descriptor()
}

func (CmdType) Type() protoreflect.EnumType {
	return &file_gate_proto_enumTypes[0]
}

func (x CmdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CmdType.Descriptor instead.
func (CmdType) EnumDescriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{0}
}

type SrcType int32

const (
	SrcType_CLIENT SrcType = 0 // client direct
	SrcType_GATE   SrcType = 1 // gate forward
)

// Enum value maps for SrcType.
var (
	SrcType_name = map[int32]string{
		0: "CLIENT",
		1: "GATE",
	}
	SrcType_value = map[string]int32{
		"CLIENT": 0,
		"GATE":   1,
	}
)

func (x SrcType) Enum() *SrcType {
	p := new(SrcType)
	*p = x
	return p
}

func (x SrcType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SrcType) Descriptor() protoreflect.EnumDescriptor {
	return file_gate_proto_enumTypes[1].Descriptor()
}

func (SrcType) Type() protoreflect.EnumType {
	return &file_gate_proto_enumTypes[1]
}

func (x SrcType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SrcType.Descriptor instead.
func (SrcType) EnumDescriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{1}
}

type ErrorCode int32

const (
	ErrorCode_Success             ErrorCode = 0
	ErrorCode_Unauthorized        ErrorCode = 401
	ErrorCode_Forbidden           ErrorCode = 403
	ErrorCode_NotFound            ErrorCode = 404
	ErrorCode_InternalServerError ErrorCode = 500
	ErrorCode_ServiceUnavailable  ErrorCode = 503
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:   "Success",
		401: "Unauthorized",
		403: "Forbidden",
		404: "NotFound",
		500: "InternalServerError",
		503: "ServiceUnavailable",
	}
	ErrorCode_value = map[string]int32{
		"Success":             0,
		"Unauthorized":        401,
		"Forbidden":           403,
		"NotFound":            404,
		"InternalServerError": 500,
		"ServiceUnavailable":  503,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_gate_proto_enumTypes[2].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_gate_proto_enumTypes[2]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{2}
}

type MapFieldEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MapFieldEntry) Reset() {
	*x = MapFieldEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapFieldEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapFieldEntry) ProtoMessage() {}

func (x *MapFieldEntry) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapFieldEntry.ProtoReflect.Descriptor instead.
func (*MapFieldEntry) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{0}
}

func (x *MapFieldEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MapFieldEntry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Handshake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd          CmdType          `protobuf:"varint,1,opt,name=Cmd,proto3,enum=proto.CmdType" json:"Cmd,omitempty"` //
	Src          SrcType          `protobuf:"varint,2,opt,name=Src,proto3,enum=proto.SrcType" json:"Src,omitempty"` //
	ServiceName  string           `protobuf:"bytes,3,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`     // Service name（new connection OR stateless service）
	ServerId     string           `protobuf:"bytes,4,opt,name=ServerId,proto3" json:"ServerId,omitempty"`           // Used when reconnecting stateful services
	ServerIP     string           `protobuf:"bytes,5,opt,name=ServerIP,proto3" json:"ServerIP,omitempty"`           // Backend ip
	ServerPort   int32            `protobuf:"varint,6,opt,name=ServerPort,proto3" json:"ServerPort,omitempty"`      // Backend port
	ClientAddr   string           `protobuf:"bytes,7,opt,name=ClientAddr,proto3" json:"ClientAddr,omitempty"`       // The address of the client when the gate forward
	UserID       string           `protobuf:"bytes,8,opt,name=UserID,proto3" json:"UserID,omitempty"`               // Userid used for actor
	ClientVer    string           `protobuf:"bytes,9,opt,name=ClientVer,proto3" json:"ClientVer,omitempty"`         // client version
	ClientResVer string           `protobuf:"bytes,10,opt,name=ClientResVer,proto3" json:"ClientResVer,omitempty"`  // client resource version
	Meta         []*MapFieldEntry `protobuf:"bytes,11,rep,name=Meta,proto3" json:"Meta,omitempty"`                  // custom meta
}

func (x *Handshake) Reset() {
	*x = Handshake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handshake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handshake) ProtoMessage() {}

func (x *Handshake) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handshake.ProtoReflect.Descriptor instead.
func (*Handshake) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{1}
}

func (x *Handshake) GetCmd() CmdType {
	if x != nil {
		return x.Cmd
	}
	return CmdType_NEW
}

func (x *Handshake) GetSrc() SrcType {
	if x != nil {
		return x.Src
	}
	return SrcType_CLIENT
}

func (x *Handshake) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Handshake) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Handshake) GetServerIP() string {
	if x != nil {
		return x.ServerIP
	}
	return ""
}

func (x *Handshake) GetServerPort() int32 {
	if x != nil {
		return x.ServerPort
	}
	return 0
}

func (x *Handshake) GetClientAddr() string {
	if x != nil {
		return x.ClientAddr
	}
	return ""
}

func (x *Handshake) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Handshake) GetClientVer() string {
	if x != nil {
		return x.ClientVer
	}
	return ""
}

func (x *Handshake) GetClientResVer() string {
	if x != nil {
		return x.ClientResVer
	}
	return ""
}

func (x *Handshake) GetMeta() []*MapFieldEntry {
	if x != nil {
		return x.Meta
	}
	return nil
}

type HandshakeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ErrorCode `protobuf:"varint,1,opt,name=Code,proto3,enum=proto.ErrorCode" json:"Code,omitempty"` //
	Desc string    `protobuf:"bytes,2,opt,name=Desc,proto3" json:"Desc,omitempty"`                       // success or err desc
}

func (x *HandshakeResp) Reset() {
	*x = HandshakeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeResp) ProtoMessage() {}

func (x *HandshakeResp) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeResp.ProtoReflect.Descriptor instead.
func (*HandshakeResp) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{2}
}

func (x *HandshakeResp) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_Success
}

func (x *HandshakeResp) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

var File_gate_proto protoreflect.FileDescriptor

var file_gate_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0d, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xed, 0x02, 0x0a,
	0x09, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x43, 0x6d,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x20, 0x0a, 0x03,
	0x53, 0x72, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x72, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x53, 0x72, 0x63, 0x12, 0x20,
	0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x50, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x50, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x12, 0x22,
	0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x56,
	0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x49, 0x0a, 0x0d,
	0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x44, 0x65, 0x73, 0x63, 0x2a, 0x1e, 0x0a, 0x07, 0x43, 0x6d, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x57, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x45, 0x43, 0x4f, 0x4e, 0x4e, 0x10, 0x01, 0x2a, 0x1f, 0x0a, 0x07, 0x53, 0x72, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x47, 0x41, 0x54, 0x45, 0x10, 0x01, 0x2a, 0x7d, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x10, 0x91, 0x03, 0x12, 0x0e, 0x0a, 0x09, 0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64,
	0x65, 0x6e, 0x10, 0x93, 0x03, 0x12, 0x0d, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e,
	0x64, 0x10, 0x94, 0x03, 0x12, 0x18, 0x0a, 0x13, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0xf4, 0x03, 0x12, 0x17,
	0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x10, 0xf7, 0x03, 0x42, 0x39, 0x5a, 0x2f, 0x65, 0x2e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x6d, 0x6d, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f,
	0x2f, 0x62, 0x6c, 0x61, 0x64, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0xaa, 0x02, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gate_proto_rawDescOnce sync.Once
	file_gate_proto_rawDescData = file_gate_proto_rawDesc
)

func file_gate_proto_rawDescGZIP() []byte {
	file_gate_proto_rawDescOnce.Do(func() {
		file_gate_proto_rawDescData = protoimpl.X.CompressGZIP(file_gate_proto_rawDescData)
	})
	return file_gate_proto_rawDescData
}

var file_gate_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_gate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gate_proto_goTypes = []interface{}{
	(CmdType)(0),          // 0: proto.CmdType
	(SrcType)(0),          // 1: proto.SrcType
	(ErrorCode)(0),        // 2: proto.ErrorCode
	(*MapFieldEntry)(nil), // 3: proto.MapFieldEntry
	(*Handshake)(nil),     // 4: proto.Handshake
	(*HandshakeResp)(nil), // 5: proto.HandshakeResp
}
var file_gate_proto_depIdxs = []int32{
	0, // 0: proto.Handshake.Cmd:type_name -> proto.CmdType
	1, // 1: proto.Handshake.Src:type_name -> proto.SrcType
	3, // 2: proto.Handshake.Meta:type_name -> proto.MapFieldEntry
	2, // 3: proto.HandshakeResp.Code:type_name -> proto.ErrorCode
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_gate_proto_init() }
func file_gate_proto_init() {
	if File_gate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapFieldEntry); i {
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
		file_gate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handshake); i {
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
		file_gate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeResp); i {
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
			RawDescriptor: file_gate_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gate_proto_goTypes,
		DependencyIndexes: file_gate_proto_depIdxs,
		EnumInfos:         file_gate_proto_enumTypes,
		MessageInfos:      file_gate_proto_msgTypes,
	}.Build()
	File_gate_proto = out.File
	file_gate_proto_rawDesc = nil
	file_gate_proto_goTypes = nil
	file_gate_proto_depIdxs = nil
}
