// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: global/item.proto

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

type C2S_DelItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *C2S_DelItem) Reset() {
	*x = C2S_DelItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_DelItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_DelItem) ProtoMessage() {}

func (x *C2S_DelItem) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_DelItem.ProtoReflect.Descriptor instead.
func (*C2S_DelItem) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{0}
}

func (x *C2S_DelItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type C2S_UseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId int64 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
}

func (x *C2S_UseItem) Reset() {
	*x = C2S_UseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_UseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_UseItem) ProtoMessage() {}

func (x *C2S_UseItem) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_UseItem.ProtoReflect.Descriptor instead.
func (*C2S_UseItem) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{1}
}

func (x *C2S_UseItem) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

type C2S_QueryItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2S_QueryItems) Reset() {
	*x = C2S_QueryItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_QueryItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_QueryItems) ProtoMessage() {}

func (x *C2S_QueryItems) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_QueryItems.ProtoReflect.Descriptor instead.
func (*C2S_QueryItems) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{2}
}

type S2C_ItemList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *S2C_ItemList) Reset() {
	*x = S2C_ItemList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_ItemList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_ItemList) ProtoMessage() {}

func (x *S2C_ItemList) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_ItemList.ProtoReflect.Descriptor instead.
func (*S2C_ItemList) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{3}
}

func (x *S2C_ItemList) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type S2C_ItemAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *S2C_ItemAdd) Reset() {
	*x = S2C_ItemAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_ItemAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_ItemAdd) ProtoMessage() {}

func (x *S2C_ItemAdd) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_ItemAdd.ProtoReflect.Descriptor instead.
func (*S2C_ItemAdd) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{4}
}

func (x *S2C_ItemAdd) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type S2C_ItemUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *S2C_ItemUpdate) Reset() {
	*x = S2C_ItemUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_ItemUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_ItemUpdate) ProtoMessage() {}

func (x *S2C_ItemUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_ItemUpdate.ProtoReflect.Descriptor instead.
func (*S2C_ItemUpdate) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{5}
}

func (x *S2C_ItemUpdate) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type S2C_DelItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId int64 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
}

func (x *S2C_DelItem) Reset() {
	*x = S2C_DelItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_DelItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_DelItem) ProtoMessage() {}

func (x *S2C_DelItem) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_DelItem.ProtoReflect.Descriptor instead.
func (*S2C_DelItem) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{6}
}

func (x *S2C_DelItem) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

/////////////////////////////////////////////
// 装备
type C2S_PutonEquip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId  int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	EquipId int64 `protobuf:"varint,2,opt,name=EquipId,proto3" json:"EquipId,omitempty"`
}

func (x *C2S_PutonEquip) Reset() {
	*x = C2S_PutonEquip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_PutonEquip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_PutonEquip) ProtoMessage() {}

func (x *C2S_PutonEquip) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_PutonEquip.ProtoReflect.Descriptor instead.
func (*C2S_PutonEquip) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{7}
}

func (x *C2S_PutonEquip) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *C2S_PutonEquip) GetEquipId() int64 {
	if x != nil {
		return x.EquipId
	}
	return 0
}

type C2S_TakeoffEquip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	Pos    int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"` // 0-3
}

func (x *C2S_TakeoffEquip) Reset() {
	*x = C2S_TakeoffEquip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_TakeoffEquip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_TakeoffEquip) ProtoMessage() {}

func (x *C2S_TakeoffEquip) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_TakeoffEquip.ProtoReflect.Descriptor instead.
func (*C2S_TakeoffEquip) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{8}
}

func (x *C2S_TakeoffEquip) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *C2S_TakeoffEquip) GetPos() int32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

// 升级
type C2S_EquipLevelup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId     int64   `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`                // 升级的装备id
	StuffItems []int64 `protobuf:"varint,2,rep,packed,name=StuffItems,proto3" json:"StuffItems,omitempty"` // 吞噬的装备id列表
	ExpItems   []int64 `protobuf:"varint,3,rep,packed,name=ExpItems,proto3" json:"ExpItems,omitempty"`     // 经验道具
}

func (x *C2S_EquipLevelup) Reset() {
	*x = C2S_EquipLevelup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_EquipLevelup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_EquipLevelup) ProtoMessage() {}

func (x *C2S_EquipLevelup) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_EquipLevelup.ProtoReflect.Descriptor instead.
func (*C2S_EquipLevelup) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{9}
}

func (x *C2S_EquipLevelup) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *C2S_EquipLevelup) GetStuffItems() []int64 {
	if x != nil {
		return x.StuffItems
	}
	return nil
}

func (x *C2S_EquipLevelup) GetExpItems() []int64 {
	if x != nil {
		return x.ExpItems
	}
	return nil
}

// 突破
type C2S_EquipPromote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId int64 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"` // 突破的装备id
}

func (x *C2S_EquipPromote) Reset() {
	*x = C2S_EquipPromote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_EquipPromote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_EquipPromote) ProtoMessage() {}

func (x *C2S_EquipPromote) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_EquipPromote.ProtoReflect.Descriptor instead.
func (*C2S_EquipPromote) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{10}
}

func (x *C2S_EquipPromote) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

// 升星
type C2S_EquipStarup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId int64 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"` // 升星的装备id
}

func (x *C2S_EquipStarup) Reset() {
	*x = C2S_EquipStarup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_EquipStarup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_EquipStarup) ProtoMessage() {}

func (x *C2S_EquipStarup) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_EquipStarup.ProtoReflect.Descriptor instead.
func (*C2S_EquipStarup) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{11}
}

func (x *C2S_EquipStarup) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

// 装备信息更新
type S2C_EquipUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EquipId   int64      `protobuf:"varint,1,opt,name=EquipId,proto3" json:"EquipId,omitempty"`
	EquipData *EquipData `protobuf:"bytes,2,opt,name=EquipData,proto3" json:"EquipData,omitempty"`
}

func (x *S2C_EquipUpdate) Reset() {
	*x = S2C_EquipUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_EquipUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_EquipUpdate) ProtoMessage() {}

func (x *S2C_EquipUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_EquipUpdate.ProtoReflect.Descriptor instead.
func (*S2C_EquipUpdate) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{12}
}

func (x *S2C_EquipUpdate) GetEquipId() int64 {
	if x != nil {
		return x.EquipId
	}
	return 0
}

func (x *S2C_EquipUpdate) GetEquipData() *EquipData {
	if x != nil {
		return x.EquipData
	}
	return nil
}

/////////////////////////////////////////////
// 晶石
type S2C_CrystalUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrystalId   int64        `protobuf:"varint,1,opt,name=CrystalId,proto3" json:"CrystalId,omitempty"`    // 晶石id
	CrystalData *CrystalData `protobuf:"bytes,2,opt,name=CrystalData,proto3" json:"CrystalData,omitempty"` // 晶石数据
}

func (x *S2C_CrystalUpdate) Reset() {
	*x = S2C_CrystalUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_CrystalUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_CrystalUpdate) ProtoMessage() {}

func (x *S2C_CrystalUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_CrystalUpdate.ProtoReflect.Descriptor instead.
func (*S2C_CrystalUpdate) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{13}
}

func (x *S2C_CrystalUpdate) GetCrystalId() int64 {
	if x != nil {
		return x.CrystalId
	}
	return 0
}

func (x *S2C_CrystalUpdate) GetCrystalData() *CrystalData {
	if x != nil {
		return x.CrystalData
	}
	return nil
}

type C2S_PutonCrystal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId    int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	CrystalId int64 `protobuf:"varint,2,opt,name=CrystalId,proto3" json:"CrystalId,omitempty"`
}

func (x *C2S_PutonCrystal) Reset() {
	*x = C2S_PutonCrystal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_PutonCrystal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_PutonCrystal) ProtoMessage() {}

func (x *C2S_PutonCrystal) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_PutonCrystal.ProtoReflect.Descriptor instead.
func (*C2S_PutonCrystal) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{14}
}

func (x *C2S_PutonCrystal) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *C2S_PutonCrystal) GetCrystalId() int64 {
	if x != nil {
		return x.CrystalId
	}
	return 0
}

type C2S_TakeoffCrystal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	Pos    int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (x *C2S_TakeoffCrystal) Reset() {
	*x = C2S_TakeoffCrystal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_TakeoffCrystal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_TakeoffCrystal) ProtoMessage() {}

func (x *C2S_TakeoffCrystal) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_TakeoffCrystal.ProtoReflect.Descriptor instead.
func (*C2S_TakeoffCrystal) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{15}
}

func (x *C2S_TakeoffCrystal) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *C2S_TakeoffCrystal) GetPos() int32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

// 升级
type C2S_CrystalLevelup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId     int64   `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`                // 升级的晶石id
	StuffItems []int64 `protobuf:"varint,2,rep,packed,name=StuffItems,proto3" json:"StuffItems,omitempty"` // 吞噬的晶石id列表
	ExpItems   []int64 `protobuf:"varint,3,rep,packed,name=ExpItems,proto3" json:"ExpItems,omitempty"`     // 经验道具
}

func (x *C2S_CrystalLevelup) Reset() {
	*x = C2S_CrystalLevelup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_CrystalLevelup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_CrystalLevelup) ProtoMessage() {}

func (x *C2S_CrystalLevelup) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_CrystalLevelup.ProtoReflect.Descriptor instead.
func (*C2S_CrystalLevelup) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{16}
}

func (x *C2S_CrystalLevelup) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *C2S_CrystalLevelup) GetStuffItems() []int64 {
	if x != nil {
		return x.StuffItems
	}
	return nil
}

func (x *C2S_CrystalLevelup) GetExpItems() []int64 {
	if x != nil {
		return x.ExpItems
	}
	return nil
}

// 测试att
type S2C_CrystalAttUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrystalId int64   `protobuf:"varint,1,opt,name=CrystalId,proto3" json:"CrystalId,omitempty"`
	AttValue  []int32 `protobuf:"varint,2,rep,packed,name=AttValue,proto3" json:"AttValue,omitempty"`
}

func (x *S2C_CrystalAttUpdate) Reset() {
	*x = S2C_CrystalAttUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_item_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_CrystalAttUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_CrystalAttUpdate) ProtoMessage() {}

func (x *S2C_CrystalAttUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_global_item_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_CrystalAttUpdate.ProtoReflect.Descriptor instead.
func (*S2C_CrystalAttUpdate) Descriptor() ([]byte, []int) {
	return file_global_item_proto_rawDescGZIP(), []int{17}
}

func (x *S2C_CrystalAttUpdate) GetCrystalId() int64 {
	if x != nil {
		return x.CrystalId
	}
	return 0
}

func (x *S2C_CrystalAttUpdate) GetAttValue() []int32 {
	if x != nil {
		return x.AttValue
	}
	return nil
}

var File_global_item_proto protoreflect.FileDescriptor

var file_global_item_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x1a, 0x13, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1d, 0x0a, 0x0b, 0x43, 0x32, 0x53, 0x5f, 0x44, 0x65, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22,
	0x25, 0x0a, 0x0b, 0x43, 0x32, 0x53, 0x5f, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16,
	0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x32, 0x53, 0x5f, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x32, 0x0a, 0x0c, 0x53, 0x32, 0x43, 0x5f,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2f, 0x0a, 0x0b,
	0x53, 0x32, 0x43, 0x5f, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x64, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x32, 0x0a,
	0x0e, 0x53, 0x32, 0x43, 0x5f, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x20, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x32, 0x43, 0x5f, 0x44, 0x65, 0x6c, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0e, 0x43, 0x32, 0x53, 0x5f,
	0x50, 0x75, 0x74, 0x6f, 0x6e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65,
	0x72, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x72, 0x6f,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x71, 0x75, 0x69, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x71, 0x75, 0x69, 0x70, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x10,
	0x43, 0x32, 0x53, 0x5f, 0x54, 0x61, 0x6b, 0x65, 0x6f, 0x66, 0x66, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x6f, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x6f, 0x73, 0x22, 0x66, 0x0a, 0x10, 0x43, 0x32,
	0x53, 0x5f, 0x45, 0x71, 0x75, 0x69, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x75, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x75, 0x66, 0x66, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x53, 0x74, 0x75, 0x66,
	0x66, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x78, 0x70, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x45, 0x78, 0x70, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x2a, 0x0a, 0x10, 0x43, 0x32, 0x53, 0x5f, 0x45, 0x71, 0x75, 0x69, 0x70, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x29,
	0x0a, 0x0f, 0x43, 0x32, 0x53, 0x5f, 0x45, 0x71, 0x75, 0x69, 0x70, 0x53, 0x74, 0x61, 0x72, 0x75,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x5c, 0x0a, 0x0f, 0x53, 0x32, 0x43,
	0x5f, 0x45, 0x71, 0x75, 0x69, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x45, 0x71, 0x75, 0x69, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x09, 0x45, 0x71, 0x75, 0x69, 0x70, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x45, 0x71,
	0x75, 0x69, 0x70, 0x44, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x11, 0x53, 0x32, 0x43, 0x5f, 0x43,
	0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0b, 0x43, 0x72,
	0x79, 0x73, 0x74, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x48, 0x0a, 0x10, 0x43, 0x32, 0x53, 0x5f, 0x50, 0x75, 0x74, 0x6f, 0x6e, 0x43, 0x72,
	0x79, 0x73, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x12, 0x43,
	0x32, 0x53, 0x5f, 0x54, 0x61, 0x6b, 0x65, 0x6f, 0x66, 0x66, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x6f, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x6f, 0x73, 0x22, 0x68, 0x0a, 0x12, 0x43,
	0x32, 0x53, 0x5f, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x75,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x75,
	0x66, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x53,
	0x74, 0x75, 0x66, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x78, 0x70,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x45, 0x78, 0x70,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x50, 0x0a, 0x14, 0x53, 0x32, 0x43, 0x5f, 0x43, 0x72, 0x79,
	0x73, 0x74, 0x61, 0x6c, 0x41, 0x74, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x41,
	0x74, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x41,
	0x74, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x34, 0x5a, 0x29, 0x62, 0x69, 0x74, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x70, 0x6c, 0x75, 0x73,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0xaa, 0x02, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_global_item_proto_rawDescOnce sync.Once
	file_global_item_proto_rawDescData = file_global_item_proto_rawDesc
)

func file_global_item_proto_rawDescGZIP() []byte {
	file_global_item_proto_rawDescOnce.Do(func() {
		file_global_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_global_item_proto_rawDescData)
	})
	return file_global_item_proto_rawDescData
}

var file_global_item_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_global_item_proto_goTypes = []interface{}{
	(*C2S_DelItem)(nil),          // 0: global.C2S_DelItem
	(*C2S_UseItem)(nil),          // 1: global.C2S_UseItem
	(*C2S_QueryItems)(nil),       // 2: global.C2S_QueryItems
	(*S2C_ItemList)(nil),         // 3: global.S2C_ItemList
	(*S2C_ItemAdd)(nil),          // 4: global.S2C_ItemAdd
	(*S2C_ItemUpdate)(nil),       // 5: global.S2C_ItemUpdate
	(*S2C_DelItem)(nil),          // 6: global.S2C_DelItem
	(*C2S_PutonEquip)(nil),       // 7: global.C2S_PutonEquip
	(*C2S_TakeoffEquip)(nil),     // 8: global.C2S_TakeoffEquip
	(*C2S_EquipLevelup)(nil),     // 9: global.C2S_EquipLevelup
	(*C2S_EquipPromote)(nil),     // 10: global.C2S_EquipPromote
	(*C2S_EquipStarup)(nil),      // 11: global.C2S_EquipStarup
	(*S2C_EquipUpdate)(nil),      // 12: global.S2C_EquipUpdate
	(*S2C_CrystalUpdate)(nil),    // 13: global.S2C_CrystalUpdate
	(*C2S_PutonCrystal)(nil),     // 14: global.C2S_PutonCrystal
	(*C2S_TakeoffCrystal)(nil),   // 15: global.C2S_TakeoffCrystal
	(*C2S_CrystalLevelup)(nil),   // 16: global.C2S_CrystalLevelup
	(*S2C_CrystalAttUpdate)(nil), // 17: global.S2C_CrystalAttUpdate
	(*Item)(nil),                 // 18: global.Item
	(*EquipData)(nil),            // 19: global.EquipData
	(*CrystalData)(nil),          // 20: global.CrystalData
}
var file_global_item_proto_depIdxs = []int32{
	18, // 0: global.S2C_ItemList.items:type_name -> global.Item
	18, // 1: global.S2C_ItemAdd.item:type_name -> global.Item
	18, // 2: global.S2C_ItemUpdate.item:type_name -> global.Item
	19, // 3: global.S2C_EquipUpdate.EquipData:type_name -> global.EquipData
	20, // 4: global.S2C_CrystalUpdate.CrystalData:type_name -> global.CrystalData
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_global_item_proto_init() }
func file_global_item_proto_init() {
	if File_global_item_proto != nil {
		return
	}
	file_global_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_global_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_DelItem); i {
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
		file_global_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_UseItem); i {
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
		file_global_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_QueryItems); i {
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
		file_global_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_ItemList); i {
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
		file_global_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_ItemAdd); i {
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
		file_global_item_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_ItemUpdate); i {
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
		file_global_item_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_DelItem); i {
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
		file_global_item_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_PutonEquip); i {
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
		file_global_item_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_TakeoffEquip); i {
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
		file_global_item_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_EquipLevelup); i {
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
		file_global_item_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_EquipPromote); i {
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
		file_global_item_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_EquipStarup); i {
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
		file_global_item_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_EquipUpdate); i {
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
		file_global_item_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_CrystalUpdate); i {
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
		file_global_item_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_PutonCrystal); i {
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
		file_global_item_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_TakeoffCrystal); i {
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
		file_global_item_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_CrystalLevelup); i {
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
		file_global_item_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_CrystalAttUpdate); i {
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
			RawDescriptor: file_global_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_global_item_proto_goTypes,
		DependencyIndexes: file_global_item_proto_depIdxs,
		MessageInfos:      file_global_item_proto_msgTypes,
	}.Build()
	File_global_item_proto = out.File
	file_global_item_proto_rawDesc = nil
	file_global_item_proto_goTypes = nil
	file_global_item_proto_depIdxs = nil
}
