// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: s_common.proto

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

// 表数据
type NetNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneID    int32 `protobuf:"varint,1,opt,name=ZoneID,proto3" json:"ZoneID,omitempty"`
	NodeType  int32 `protobuf:"varint,2,opt,name=NodeType,proto3" json:"NodeType,omitempty"`
	NodeIndex int32 `protobuf:"varint,3,opt,name=NodeIndex,proto3" json:"NodeIndex,omitempty"`
}

func (x *NetNode) Reset() {
	*x = NetNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetNode) ProtoMessage() {}

func (x *NetNode) ProtoReflect() protoreflect.Message {
	mi := &file_s_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetNode.ProtoReflect.Descriptor instead.
func (*NetNode) Descriptor() ([]byte, []int) {
	return file_s_common_proto_rawDescGZIP(), []int{0}
}

func (x *NetNode) GetZoneID() int32 {
	if x != nil {
		return x.ZoneID
	}
	return 0
}

func (x *NetNode) GetNodeType() int32 {
	if x != nil {
		return x.NodeType
	}
	return 0
}

func (x *NetNode) GetNodeIndex() int32 {
	if x != nil {
		return x.NodeIndex
	}
	return 0
}

// 服务器节点信息
type ServerNodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneID    int32 `protobuf:"varint,1,opt,name=ZoneID,proto3" json:"ZoneID,omitempty"`
	NodeType  int32 `protobuf:"varint,2,opt,name=NodeType,proto3" json:"NodeType,omitempty"`
	NodeIndex int32 `protobuf:"varint,3,opt,name=NodeIndex,proto3" json:"NodeIndex,omitempty"`
}

func (x *ServerNodeInfo) Reset() {
	*x = ServerNodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerNodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerNodeInfo) ProtoMessage() {}

func (x *ServerNodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_s_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerNodeInfo.ProtoReflect.Descriptor instead.
func (*ServerNodeInfo) Descriptor() ([]byte, []int) {
	return file_s_common_proto_rawDescGZIP(), []int{1}
}

func (x *ServerNodeInfo) GetZoneID() int32 {
	if x != nil {
		return x.ZoneID
	}
	return 0
}

func (x *ServerNodeInfo) GetNodeType() int32 {
	if x != nil {
		return x.NodeType
	}
	return 0
}

func (x *ServerNodeInfo) GetNodeIndex() int32 {
	if x != nil {
		return x.NodeIndex
	}
	return 0
}

// 服务器用玩家信息
type LogicPlayerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleID   int64  `protobuf:"varint,1,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
}

func (x *LogicPlayerData) Reset() {
	*x = LogicPlayerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogicPlayerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogicPlayerData) ProtoMessage() {}

func (x *LogicPlayerData) ProtoReflect() protoreflect.Message {
	mi := &file_s_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogicPlayerData.ProtoReflect.Descriptor instead.
func (*LogicPlayerData) Descriptor() ([]byte, []int) {
	return file_s_common_proto_rawDescGZIP(), []int{2}
}

func (x *LogicPlayerData) GetRoleID() int64 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

func (x *LogicPlayerData) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

var File_s_common_proto protoreflect.FileDescriptor

var file_s_common_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x5b, 0x0a, 0x07, 0x4e, 0x65,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x62, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e,
	0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x45, 0x0a, 0x0f, 0x4c,
	0x6f, 0x67, 0x69, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_s_common_proto_rawDescOnce sync.Once
	file_s_common_proto_rawDescData = file_s_common_proto_rawDesc
)

func file_s_common_proto_rawDescGZIP() []byte {
	file_s_common_proto_rawDescOnce.Do(func() {
		file_s_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_s_common_proto_rawDescData)
	})
	return file_s_common_proto_rawDescData
}

var file_s_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_s_common_proto_goTypes = []interface{}{
	(*NetNode)(nil),         // 0: pbserver.NetNode
	(*ServerNodeInfo)(nil),  // 1: pbserver.ServerNodeInfo
	(*LogicPlayerData)(nil), // 2: pbserver.LogicPlayerData
}
var file_s_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_s_common_proto_init() }
func file_s_common_proto_init() {
	if File_s_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_s_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetNode); i {
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
		file_s_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerNodeInfo); i {
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
		file_s_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogicPlayerData); i {
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
			RawDescriptor: file_s_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_s_common_proto_goTypes,
		DependencyIndexes: file_s_common_proto_depIdxs,
		MessageInfos:      file_s_common_proto_msgTypes,
	}.Build()
	File_s_common_proto = out.File
	file_s_common_proto_rawDesc = nil
	file_s_common_proto_goTypes = nil
	file_s_common_proto_depIdxs = nil
}
