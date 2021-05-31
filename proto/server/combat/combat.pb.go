// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: server/combat/combat.proto

package combat

import (
	global "github.com/east-eden/server/proto/global"
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

// 关卡战斗
type StageCombatRq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId          int32                `protobuf:"varint,1,opt,name=StageId,proto3" json:"StageId,omitempty"`                  // 关卡id
	AttackId         int64                `protobuf:"varint,2,opt,name=AttackId,proto3" json:"AttackId,omitempty"`                // 进攻方id -- 玩家id
	AttackEntityList []*global.EntityInfo `protobuf:"bytes,3,rep,name=AttackEntityList,proto3" json:"AttackEntityList,omitempty"` // 进攻方英雄信息
}

func (x *StageCombatRq) Reset() {
	*x = StageCombatRq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_combat_combat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StageCombatRq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StageCombatRq) ProtoMessage() {}

func (x *StageCombatRq) ProtoReflect() protoreflect.Message {
	mi := &file_server_combat_combat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StageCombatRq.ProtoReflect.Descriptor instead.
func (*StageCombatRq) Descriptor() ([]byte, []int) {
	return file_server_combat_combat_proto_rawDescGZIP(), []int{0}
}

func (x *StageCombatRq) GetStageId() int32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *StageCombatRq) GetAttackId() int64 {
	if x != nil {
		return x.AttackId
	}
	return 0
}

func (x *StageCombatRq) GetAttackEntityList() []*global.EntityInfo {
	if x != nil {
		return x.AttackEntityList
	}
	return nil
}

type StageCombatRs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Win       bool   `protobuf:"varint,1,opt,name=Win,proto3" json:"Win,omitempty"`                    // 战斗结果
	Objective []bool `protobuf:"varint,2,rep,packed,name=Objective,proto3" json:"Objective,omitempty"` // 关卡条件达成
}

func (x *StageCombatRs) Reset() {
	*x = StageCombatRs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_combat_combat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StageCombatRs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StageCombatRs) ProtoMessage() {}

func (x *StageCombatRs) ProtoReflect() protoreflect.Message {
	mi := &file_server_combat_combat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StageCombatRs.ProtoReflect.Descriptor instead.
func (*StageCombatRs) Descriptor() ([]byte, []int) {
	return file_server_combat_combat_proto_rawDescGZIP(), []int{1}
}

func (x *StageCombatRs) GetWin() bool {
	if x != nil {
		return x.Win
	}
	return false
}

func (x *StageCombatRs) GetObjective() []bool {
	if x != nil {
		return x.Objective
	}
	return nil
}

var File_server_combat_combat_proto protoreflect.FileDescriptor

var file_server_combat_combat_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2f,
	0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f,
	0x6d, 0x62, 0x61, 0x74, 0x1a, 0x13, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d,
	0x62, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x0d, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x49,
	0x64, 0x12, 0x3d, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x3f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x57, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03,
	0x57, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x08, 0x52, 0x09, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x32, 0x4e, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x62, 0x61,
	0x74, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x71, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x62, 0x61,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x73, 0x22,
	0x00, 0x42, 0x32, 0x5a, 0x30, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x70, 0x6c, 0x75, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63,
	0x6f, 0x6d, 0x62, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_combat_combat_proto_rawDescOnce sync.Once
	file_server_combat_combat_proto_rawDescData = file_server_combat_combat_proto_rawDesc
)

func file_server_combat_combat_proto_rawDescGZIP() []byte {
	file_server_combat_combat_proto_rawDescOnce.Do(func() {
		file_server_combat_combat_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_combat_combat_proto_rawDescData)
	})
	return file_server_combat_combat_proto_rawDescData
}

var file_server_combat_combat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_combat_combat_proto_goTypes = []interface{}{
	(*StageCombatRq)(nil),     // 0: combat.StageCombatRq
	(*StageCombatRs)(nil),     // 1: combat.StageCombatRs
	(*global.EntityInfo)(nil), // 2: proto.EntityInfo
}
var file_server_combat_combat_proto_depIdxs = []int32{
	2, // 0: combat.StageCombatRq.AttackEntityList:type_name -> proto.EntityInfo
	0, // 1: combat.CombatService.StageCombat:input_type -> combat.StageCombatRq
	1, // 2: combat.CombatService.StageCombat:output_type -> combat.StageCombatRs
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_server_combat_combat_proto_init() }
func file_server_combat_combat_proto_init() {
	if File_server_combat_combat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_combat_combat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StageCombatRq); i {
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
		file_server_combat_combat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StageCombatRs); i {
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
			RawDescriptor: file_server_combat_combat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_combat_combat_proto_goTypes,
		DependencyIndexes: file_server_combat_combat_proto_depIdxs,
		MessageInfos:      file_server_combat_combat_proto_msgTypes,
	}.Build()
	File_server_combat_combat_proto = out.File
	file_server_combat_combat_proto_rawDesc = nil
	file_server_combat_combat_proto_goTypes = nil
	file_server_combat_combat_proto_depIdxs = nil
}
