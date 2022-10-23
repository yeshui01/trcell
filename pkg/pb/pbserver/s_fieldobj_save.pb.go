// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: s_fieldobj_save.proto

package pbserver

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

// 装备数据
type DBEquipOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EquipID int64 `protobuf:"varint,1,opt,name=EquipID,proto3" json:"EquipID,omitempty"`
}

func (x *DBEquipOne) Reset() {
	*x = DBEquipOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_fieldobj_save_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBEquipOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBEquipOne) ProtoMessage() {}

func (x *DBEquipOne) ProtoReflect() protoreflect.Message {
	mi := &file_s_fieldobj_save_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBEquipOne.ProtoReflect.Descriptor instead.
func (*DBEquipOne) Descriptor() ([]byte, []int) {
	return file_s_fieldobj_save_proto_rawDescGZIP(), []int{0}
}

func (x *DBEquipOne) GetEquipID() int64 {
	if x != nil {
		return x.EquipID
	}
	return 0
}

type DBPlayerEquip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EquipList map[int64]*DBEquipOne `protobuf:"bytes,1,rep,name=EquipList,proto3" json:"EquipList,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DBPlayerEquip) Reset() {
	*x = DBPlayerEquip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_fieldobj_save_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBPlayerEquip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBPlayerEquip) ProtoMessage() {}

func (x *DBPlayerEquip) ProtoReflect() protoreflect.Message {
	mi := &file_s_fieldobj_save_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBPlayerEquip.ProtoReflect.Descriptor instead.
func (*DBPlayerEquip) Descriptor() ([]byte, []int) {
	return file_s_fieldobj_save_proto_rawDescGZIP(), []int{1}
}

func (x *DBPlayerEquip) GetEquipList() map[int64]*DBEquipOne {
	if x != nil {
		return x.EquipList
	}
	return nil
}

// 角色额外数据
type DBRoleScopeUid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idx     int64           `protobuf:"varint,1,opt,name=Idx,proto3" json:"Idx,omitempty"`
	UidPool map[int64]int64 `protobuf:"bytes,2,rep,name=UidPool,proto3" json:"UidPool,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *DBRoleScopeUid) Reset() {
	*x = DBRoleScopeUid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_fieldobj_save_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBRoleScopeUid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBRoleScopeUid) ProtoMessage() {}

func (x *DBRoleScopeUid) ProtoReflect() protoreflect.Message {
	mi := &file_s_fieldobj_save_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBRoleScopeUid.ProtoReflect.Descriptor instead.
func (*DBRoleScopeUid) Descriptor() ([]byte, []int) {
	return file_s_fieldobj_save_proto_rawDescGZIP(), []int{2}
}

func (x *DBRoleScopeUid) GetIdx() int64 {
	if x != nil {
		return x.Idx
	}
	return 0
}

func (x *DBRoleScopeUid) GetUidPool() map[int64]int64 {
	if x != nil {
		return x.UidPool
	}
	return nil
}

type DBRoleExtraData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScopeUid *DBRoleScopeUid `protobuf:"bytes,1,opt,name=ScopeUid,proto3" json:"ScopeUid,omitempty"`
}

func (x *DBRoleExtraData) Reset() {
	*x = DBRoleExtraData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_fieldobj_save_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBRoleExtraData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBRoleExtraData) ProtoMessage() {}

func (x *DBRoleExtraData) ProtoReflect() protoreflect.Message {
	mi := &file_s_fieldobj_save_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBRoleExtraData.ProtoReflect.Descriptor instead.
func (*DBRoleExtraData) Descriptor() ([]byte, []int) {
	return file_s_fieldobj_save_proto_rawDescGZIP(), []int{3}
}

func (x *DBRoleExtraData) GetScopeUid() *DBRoleScopeUid {
	if x != nil {
		return x.ScopeUid
	}
	return nil
}

// 货币数据
type DBRoleCoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinList map[int32]int64 `protobuf:"bytes,1,rep,name=CoinList,proto3" json:"CoinList,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 货币列表 key:货币id value:货币数量
}

func (x *DBRoleCoin) Reset() {
	*x = DBRoleCoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_fieldobj_save_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBRoleCoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBRoleCoin) ProtoMessage() {}

func (x *DBRoleCoin) ProtoReflect() protoreflect.Message {
	mi := &file_s_fieldobj_save_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBRoleCoin.ProtoReflect.Descriptor instead.
func (*DBRoleCoin) Descriptor() ([]byte, []int) {
	return file_s_fieldobj_save_proto_rawDescGZIP(), []int{4}
}

func (x *DBRoleCoin) GetCoinList() map[int32]int64 {
	if x != nil {
		return x.CoinList
	}
	return nil
}

// 背包数据
type DBBagPos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosID     int32 `protobuf:"varint,1,opt,name=PosID,proto3" json:"PosID,omitempty"`
	GoodsType int32 `protobuf:"varint,2,opt,name=GoodsType,proto3" json:"GoodsType,omitempty"`
	ExcelID   int32 `protobuf:"varint,3,opt,name=ExcelID,proto3" json:"ExcelID,omitempty"`
	Num       int64 `protobuf:"varint,4,opt,name=Num,proto3" json:"Num,omitempty"`
	AttachID  int64 `protobuf:"varint,5,opt,name=AttachID,proto3" json:"AttachID,omitempty"`
}

func (x *DBBagPos) Reset() {
	*x = DBBagPos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_fieldobj_save_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBBagPos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBBagPos) ProtoMessage() {}

func (x *DBBagPos) ProtoReflect() protoreflect.Message {
	mi := &file_s_fieldobj_save_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBBagPos.ProtoReflect.Descriptor instead.
func (*DBBagPos) Descriptor() ([]byte, []int) {
	return file_s_fieldobj_save_proto_rawDescGZIP(), []int{5}
}

func (x *DBBagPos) GetPosID() int32 {
	if x != nil {
		return x.PosID
	}
	return 0
}

func (x *DBBagPos) GetGoodsType() int32 {
	if x != nil {
		return x.GoodsType
	}
	return 0
}

func (x *DBBagPos) GetExcelID() int32 {
	if x != nil {
		return x.ExcelID
	}
	return 0
}

func (x *DBBagPos) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *DBBagPos) GetAttachID() int64 {
	if x != nil {
		return x.AttachID
	}
	return 0
}

type DBGoodsBag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BagType int32               `protobuf:"varint,1,opt,name=BagType,proto3" json:"BagType,omitempty"`                                                                                         // 背包类型
	PosIdx  int32               `protobuf:"varint,2,opt,name=PosIdx,proto3" json:"PosIdx,omitempty"`                                                                                           // 格子生成索引
	PosList map[int32]*DBBagPos `protobuf:"bytes,3,rep,name=PosList,proto3" json:"PosList,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 背包格子列表
}

func (x *DBGoodsBag) Reset() {
	*x = DBGoodsBag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_fieldobj_save_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBGoodsBag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBGoodsBag) ProtoMessage() {}

func (x *DBGoodsBag) ProtoReflect() protoreflect.Message {
	mi := &file_s_fieldobj_save_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBGoodsBag.ProtoReflect.Descriptor instead.
func (*DBGoodsBag) Descriptor() ([]byte, []int) {
	return file_s_fieldobj_save_proto_rawDescGZIP(), []int{6}
}

func (x *DBGoodsBag) GetBagType() int32 {
	if x != nil {
		return x.BagType
	}
	return 0
}

func (x *DBGoodsBag) GetPosIdx() int32 {
	if x != nil {
		return x.PosIdx
	}
	return 0
}

func (x *DBGoodsBag) GetPosList() map[int32]*DBBagPos {
	if x != nil {
		return x.PosList
	}
	return nil
}

type DBBagList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BagList []*DBGoodsBag `protobuf:"bytes,1,rep,name=BagList,proto3" json:"BagList,omitempty"` // 背包列表
}

func (x *DBBagList) Reset() {
	*x = DBBagList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_fieldobj_save_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBBagList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBBagList) ProtoMessage() {}

func (x *DBBagList) ProtoReflect() protoreflect.Message {
	mi := &file_s_fieldobj_save_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBBagList.ProtoReflect.Descriptor instead.
func (*DBBagList) Descriptor() ([]byte, []int) {
	return file_s_fieldobj_save_proto_rawDescGZIP(), []int{7}
}

func (x *DBBagList) GetBagList() []*DBGoodsBag {
	if x != nil {
		return x.BagList
	}
	return nil
}

var File_s_fieldobj_save_proto protoreflect.FileDescriptor

var file_s_fieldobj_save_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6f, 0x62, 0x6a, 0x5f, 0x73, 0x61, 0x76,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x22, 0x26, 0x0a, 0x0a, 0x44, 0x42, 0x45, 0x71, 0x75, 0x69, 0x70, 0x4f, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x45, 0x71, 0x75, 0x69, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x45, 0x71, 0x75, 0x69, 0x70, 0x49, 0x44, 0x22, 0xa9, 0x01, 0x0a, 0x0d, 0x44, 0x42,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x12, 0x44, 0x0a, 0x09, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x70, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x42, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x45, 0x71, 0x75, 0x69, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x1a, 0x52, 0x0a, 0x0e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x44, 0x42, 0x45, 0x71, 0x75, 0x69, 0x70, 0x4f, 0x6e, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9f, 0x01, 0x0a, 0x0e, 0x44, 0x42, 0x52, 0x6f, 0x6c, 0x65,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x55, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x64, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x49, 0x64, 0x78, 0x12, 0x3f, 0x0a, 0x07, 0x55, 0x69,
	0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x42, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x55, 0x69, 0x64, 0x2e, 0x55, 0x69, 0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x55, 0x69, 0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x1a, 0x3a, 0x0a, 0x0c, 0x55,
	0x69, 0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x47, 0x0a, 0x0f, 0x44, 0x42, 0x52, 0x6f, 0x6c,
	0x65, 0x45, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x08, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x42, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x55, 0x69, 0x64, 0x52, 0x08, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x55, 0x69, 0x64,
	0x22, 0x89, 0x01, 0x0a, 0x0a, 0x44, 0x42, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x12,
	0x3e, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x70, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x42, 0x52,
	0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x1a,
	0x3b, 0x0a, 0x0d, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x86, 0x01, 0x0a,
	0x08, 0x44, 0x42, 0x42, 0x61, 0x67, 0x50, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x6f, 0x73,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x50, 0x6f, 0x73, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x45, 0x78, 0x63, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x49, 0x44, 0x22, 0xcb, 0x01, 0x0a, 0x0a, 0x44, 0x42, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x42, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x42, 0x61, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x50, 0x6f, 0x73, 0x49, 0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x50, 0x6f, 0x73, 0x49, 0x64, 0x78, 0x12, 0x3b, 0x0a, 0x07, 0x50, 0x6f, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x44, 0x42, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x61, 0x67, 0x2e, 0x50, 0x6f,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x50, 0x6f, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x1a, 0x4e, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x44, 0x42, 0x42, 0x61, 0x67, 0x50, 0x6f, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x3b, 0x0a, 0x09, 0x44, 0x42, 0x42, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x07, 0x42, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x42, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x42, 0x61, 0x67, 0x52, 0x07, 0x42, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_s_fieldobj_save_proto_rawDescOnce sync.Once
	file_s_fieldobj_save_proto_rawDescData = file_s_fieldobj_save_proto_rawDesc
)

func file_s_fieldobj_save_proto_rawDescGZIP() []byte {
	file_s_fieldobj_save_proto_rawDescOnce.Do(func() {
		file_s_fieldobj_save_proto_rawDescData = protoimpl.X.CompressGZIP(file_s_fieldobj_save_proto_rawDescData)
	})
	return file_s_fieldobj_save_proto_rawDescData
}

var file_s_fieldobj_save_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_s_fieldobj_save_proto_goTypes = []interface{}{
	(*DBEquipOne)(nil),      // 0: pbserver.DBEquipOne
	(*DBPlayerEquip)(nil),   // 1: pbserver.DBPlayerEquip
	(*DBRoleScopeUid)(nil),  // 2: pbserver.DBRoleScopeUid
	(*DBRoleExtraData)(nil), // 3: pbserver.DBRoleExtraData
	(*DBRoleCoin)(nil),      // 4: pbserver.DBRoleCoin
	(*DBBagPos)(nil),        // 5: pbserver.DBBagPos
	(*DBGoodsBag)(nil),      // 6: pbserver.DBGoodsBag
	(*DBBagList)(nil),       // 7: pbserver.DBBagList
	nil,                     // 8: pbserver.DBPlayerEquip.EquipListEntry
	nil,                     // 9: pbserver.DBRoleScopeUid.UidPoolEntry
	nil,                     // 10: pbserver.DBRoleCoin.CoinListEntry
	nil,                     // 11: pbserver.DBGoodsBag.PosListEntry
}
var file_s_fieldobj_save_proto_depIdxs = []int32{
	8,  // 0: pbserver.DBPlayerEquip.EquipList:type_name -> pbserver.DBPlayerEquip.EquipListEntry
	9,  // 1: pbserver.DBRoleScopeUid.UidPool:type_name -> pbserver.DBRoleScopeUid.UidPoolEntry
	2,  // 2: pbserver.DBRoleExtraData.ScopeUid:type_name -> pbserver.DBRoleScopeUid
	10, // 3: pbserver.DBRoleCoin.CoinList:type_name -> pbserver.DBRoleCoin.CoinListEntry
	11, // 4: pbserver.DBGoodsBag.PosList:type_name -> pbserver.DBGoodsBag.PosListEntry
	6,  // 5: pbserver.DBBagList.BagList:type_name -> pbserver.DBGoodsBag
	0,  // 6: pbserver.DBPlayerEquip.EquipListEntry.value:type_name -> pbserver.DBEquipOne
	5,  // 7: pbserver.DBGoodsBag.PosListEntry.value:type_name -> pbserver.DBBagPos
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_s_fieldobj_save_proto_init() }
func file_s_fieldobj_save_proto_init() {
	if File_s_fieldobj_save_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_s_fieldobj_save_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBEquipOne); i {
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
		file_s_fieldobj_save_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBPlayerEquip); i {
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
		file_s_fieldobj_save_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBRoleScopeUid); i {
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
		file_s_fieldobj_save_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBRoleExtraData); i {
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
		file_s_fieldobj_save_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBRoleCoin); i {
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
		file_s_fieldobj_save_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBBagPos); i {
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
		file_s_fieldobj_save_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBGoodsBag); i {
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
		file_s_fieldobj_save_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBBagList); i {
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
			RawDescriptor: file_s_fieldobj_save_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_s_fieldobj_save_proto_goTypes,
		DependencyIndexes: file_s_fieldobj_save_proto_depIdxs,
		MessageInfos:      file_s_fieldobj_save_proto_msgTypes,
	}.Build()
	File_s_fieldobj_save_proto = out.File
	file_s_fieldobj_save_proto_rawDesc = nil
	file_s_fieldobj_save_proto_goTypes = nil
	file_s_fieldobj_save_proto_depIdxs = nil
}
