//*********************************************************************************************************************
//  Description: interface between router agents and SDK service manager
//
//  Copyright (c) 2018 Nokia
//*********************************************************************************************************************

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.7.1
// source: networkinstance_service.proto

package protos

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

/// Represents network instance type.
type NetworkInstanceData_NetInstType int32

const (
	NetworkInstanceData_DEFAULT NetworkInstanceData_NetInstType = 0 // Default network instance type
	NetworkInstanceData_L3VRF   NetworkInstanceData_NetInstType = 1 // L3VRF network instance type
)

// Enum value maps for NetworkInstanceData_NetInstType.
var (
	NetworkInstanceData_NetInstType_name = map[int32]string{
		0: "DEFAULT",
		1: "L3VRF",
	}
	NetworkInstanceData_NetInstType_value = map[string]int32{
		"DEFAULT": 0,
		"L3VRF":   1,
	}
)

func (x NetworkInstanceData_NetInstType) Enum() *NetworkInstanceData_NetInstType {
	p := new(NetworkInstanceData_NetInstType)
	*p = x
	return p
}

func (x NetworkInstanceData_NetInstType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkInstanceData_NetInstType) Descriptor() protoreflect.EnumDescriptor {
	return file_networkinstance_service_proto_enumTypes[0].Descriptor()
}

func (NetworkInstanceData_NetInstType) Type() protoreflect.EnumType {
	return &file_networkinstance_service_proto_enumTypes[0]
}

func (x NetworkInstanceData_NetInstType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkInstanceData_NetInstType.Descriptor instead.
func (NetworkInstanceData_NetInstType) EnumDescriptor() ([]byte, []int) {
	return file_networkinstance_service_proto_rawDescGZIP(), []int{2, 0}
}

//*
// Represents network instance subscription request.
type NetworkInstanceSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NetworkInstanceSubscriptionRequest) Reset() {
	*x = NetworkInstanceSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networkinstance_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInstanceSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInstanceSubscriptionRequest) ProtoMessage() {}

func (x *NetworkInstanceSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_networkinstance_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInstanceSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*NetworkInstanceSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_networkinstance_service_proto_rawDescGZIP(), []int{0}
}

//*
// Represents network instance key.
type NetworkInstanceKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstName string `protobuf:"bytes,1,opt,name=inst_name,json=instName,proto3" json:"inst_name,omitempty"` // Network instance name
}

func (x *NetworkInstanceKey) Reset() {
	*x = NetworkInstanceKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networkinstance_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInstanceKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInstanceKey) ProtoMessage() {}

func (x *NetworkInstanceKey) ProtoReflect() protoreflect.Message {
	mi := &file_networkinstance_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInstanceKey.ProtoReflect.Descriptor instead.
func (*NetworkInstanceKey) Descriptor() ([]byte, []int) {
	return file_networkinstance_service_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkInstanceKey) GetInstName() string {
	if x != nil {
		return x.InstName
	}
	return ""
}

//*
// Represents network instance data.
type NetworkInstanceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetInstId    uint32                          `protobuf:"varint,1,opt,name=net_inst_id,json=netInstId,proto3" json:"net_inst_id,omitempty"`                                             // Network instance identifier
	BaseName     string                          `protobuf:"bytes,2,opt,name=base_name,json=baseName,proto3" json:"base_name,omitempty"`                                                   // Base name
	LoopbackAddr *IpAddrPrefLenPb                `protobuf:"bytes,3,opt,name=loopback_addr,json=loopbackAddr,proto3" json:"loopback_addr,omitempty"`                                       // Loopback address of network instance
	OperIsUp     bool                            `protobuf:"varint,4,opt,name=oper_is_up,json=operIsUp,proto3" json:"oper_is_up,omitempty"`                                                // Operation status
	RouterId     string                          `protobuf:"bytes,5,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`                                                   // Router identifier
	InstType     NetworkInstanceData_NetInstType `protobuf:"varint,6,opt,name=inst_type,json=instType,proto3,enum=srlinux.sdk.NetworkInstanceData_NetInstType" json:"inst_type,omitempty"` // Network instance type
}

func (x *NetworkInstanceData) Reset() {
	*x = NetworkInstanceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networkinstance_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInstanceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInstanceData) ProtoMessage() {}

func (x *NetworkInstanceData) ProtoReflect() protoreflect.Message {
	mi := &file_networkinstance_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInstanceData.ProtoReflect.Descriptor instead.
func (*NetworkInstanceData) Descriptor() ([]byte, []int) {
	return file_networkinstance_service_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkInstanceData) GetNetInstId() uint32 {
	if x != nil {
		return x.NetInstId
	}
	return 0
}

func (x *NetworkInstanceData) GetBaseName() string {
	if x != nil {
		return x.BaseName
	}
	return ""
}

func (x *NetworkInstanceData) GetLoopbackAddr() *IpAddrPrefLenPb {
	if x != nil {
		return x.LoopbackAddr
	}
	return nil
}

func (x *NetworkInstanceData) GetOperIsUp() bool {
	if x != nil {
		return x.OperIsUp
	}
	return false
}

func (x *NetworkInstanceData) GetRouterId() string {
	if x != nil {
		return x.RouterId
	}
	return ""
}

func (x *NetworkInstanceData) GetInstType() NetworkInstanceData_NetInstType {
	if x != nil {
		return x.InstType
	}
	return NetworkInstanceData_DEFAULT
}

//*
// Represents network instance notification.
type NetworkInstanceNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op   SdkMgrOperation      `protobuf:"varint,1,opt,name=op,proto3,enum=srlinux.sdk.SdkMgrOperation" json:"op,omitempty"` // Operation such as create, delete, or update
	Key  *NetworkInstanceKey  `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`                                 // Network key
	Data *NetworkInstanceData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`                               // Network data
}

func (x *NetworkInstanceNotification) Reset() {
	*x = NetworkInstanceNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networkinstance_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInstanceNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInstanceNotification) ProtoMessage() {}

func (x *NetworkInstanceNotification) ProtoReflect() protoreflect.Message {
	mi := &file_networkinstance_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInstanceNotification.ProtoReflect.Descriptor instead.
func (*NetworkInstanceNotification) Descriptor() ([]byte, []int) {
	return file_networkinstance_service_proto_rawDescGZIP(), []int{3}
}

func (x *NetworkInstanceNotification) GetOp() SdkMgrOperation {
	if x != nil {
		return x.Op
	}
	return SdkMgrOperation_Create
}

func (x *NetworkInstanceNotification) GetKey() *NetworkInstanceKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *NetworkInstanceNotification) GetData() *NetworkInstanceData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_networkinstance_service_proto protoreflect.FileDescriptor

var file_networkinstance_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x1a, 0x10, 0x73, 0x64,
	0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24,
	0x0a, 0x22, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x12, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6e, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xc2, 0x02, 0x0a, 0x13, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1e, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0d,
	0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x50, 0x72, 0x65, 0x66, 0x4c, 0x65, 0x6e, 0x50,
	0x62, 0x52, 0x0c, 0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x1c, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x73, 0x5f, 0x75, 0x70, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x49, 0x73, 0x55, 0x70, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x09, 0x69, 0x6e,
	0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e,
	0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x4e, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x73,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x25, 0x0a, 0x0b, 0x4e, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x33, 0x56, 0x52, 0x46, 0x10, 0x01, 0x22, 0xb4, 0x01, 0x0a,
	0x1b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x02,
	0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e,
	0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x64, 0x6b, 0x4d, 0x67, 0x72, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x31, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75,
	0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x72,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x1e, 0x5a, 0x1c, 0x6e, 0x6f, 0x6b, 0x69, 0x61, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_networkinstance_service_proto_rawDescOnce sync.Once
	file_networkinstance_service_proto_rawDescData = file_networkinstance_service_proto_rawDesc
)

func file_networkinstance_service_proto_rawDescGZIP() []byte {
	file_networkinstance_service_proto_rawDescOnce.Do(func() {
		file_networkinstance_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_networkinstance_service_proto_rawDescData)
	})
	return file_networkinstance_service_proto_rawDescData
}

var file_networkinstance_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_networkinstance_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_networkinstance_service_proto_goTypes = []interface{}{
	(NetworkInstanceData_NetInstType)(0),       // 0: srlinux.sdk.NetworkInstanceData.NetInstType
	(*NetworkInstanceSubscriptionRequest)(nil), // 1: srlinux.sdk.NetworkInstanceSubscriptionRequest
	(*NetworkInstanceKey)(nil),                 // 2: srlinux.sdk.NetworkInstanceKey
	(*NetworkInstanceData)(nil),                // 3: srlinux.sdk.NetworkInstanceData
	(*NetworkInstanceNotification)(nil),        // 4: srlinux.sdk.NetworkInstanceNotification
	(*IpAddrPrefLenPb)(nil),                    // 5: srlinux.sdk.IpAddrPrefLenPb
	(SdkMgrOperation)(0),                       // 6: srlinux.sdk.SdkMgrOperation
}
var file_networkinstance_service_proto_depIdxs = []int32{
	5, // 0: srlinux.sdk.NetworkInstanceData.loopback_addr:type_name -> srlinux.sdk.IpAddrPrefLenPb
	0, // 1: srlinux.sdk.NetworkInstanceData.inst_type:type_name -> srlinux.sdk.NetworkInstanceData.NetInstType
	6, // 2: srlinux.sdk.NetworkInstanceNotification.op:type_name -> srlinux.sdk.SdkMgrOperation
	2, // 3: srlinux.sdk.NetworkInstanceNotification.key:type_name -> srlinux.sdk.NetworkInstanceKey
	3, // 4: srlinux.sdk.NetworkInstanceNotification.data:type_name -> srlinux.sdk.NetworkInstanceData
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_networkinstance_service_proto_init() }
func file_networkinstance_service_proto_init() {
	if File_networkinstance_service_proto != nil {
		return
	}
	file_sdk_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_networkinstance_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInstanceSubscriptionRequest); i {
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
		file_networkinstance_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInstanceKey); i {
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
		file_networkinstance_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInstanceData); i {
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
		file_networkinstance_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInstanceNotification); i {
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
			RawDescriptor: file_networkinstance_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_networkinstance_service_proto_goTypes,
		DependencyIndexes: file_networkinstance_service_proto_depIdxs,
		EnumInfos:         file_networkinstance_service_proto_enumTypes,
		MessageInfos:      file_networkinstance_service_proto_msgTypes,
	}.Build()
	File_networkinstance_service_proto = out.File
	file_networkinstance_service_proto_rawDesc = nil
	file_networkinstance_service_proto_goTypes = nil
	file_networkinstance_service_proto_depIdxs = nil
}
