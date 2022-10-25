// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: poc-grpc/pocgrpc.proto

package __

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

// Broadcast request with a message
type BroadcastReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BroadcastReq) Reset() {
	*x = BroadcastReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poc_grpc_pocgrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastReq) ProtoMessage() {}

func (x *BroadcastReq) ProtoReflect() protoreflect.Message {
	mi := &file_poc_grpc_pocgrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastReq.ProtoReflect.Descriptor instead.
func (*BroadcastReq) Descriptor() ([]byte, []int) {
	return file_poc_grpc_pocgrpc_proto_rawDescGZIP(), []int{0}
}

func (x *BroadcastReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Broadcast response with a message
type BroadcastRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BroadcastRes) Reset() {
	*x = BroadcastRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poc_grpc_pocgrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRes) ProtoMessage() {}

func (x *BroadcastRes) ProtoReflect() protoreflect.Message {
	mi := &file_poc_grpc_pocgrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRes.ProtoReflect.Descriptor instead.
func (*BroadcastRes) Descriptor() ([]byte, []int) {
	return file_poc_grpc_pocgrpc_proto_rawDescGZIP(), []int{1}
}

func (x *BroadcastRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_poc_grpc_pocgrpc_proto protoreflect.FileDescriptor

var file_poc_grpc_pocgrpc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6f, 0x63, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6f, 0x63, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6f, 0x63, 0x67, 0x72, 0x70,
	0x63, 0x22, 0x28, 0x0a, 0x0c, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x48, 0x0a, 0x0b, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x12, 0x15, 0x2e, 0x70, 0x6f, 0x63, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x6f, 0x63, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x42,
	0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_poc_grpc_pocgrpc_proto_rawDescOnce sync.Once
	file_poc_grpc_pocgrpc_proto_rawDescData = file_poc_grpc_pocgrpc_proto_rawDesc
)

func file_poc_grpc_pocgrpc_proto_rawDescGZIP() []byte {
	file_poc_grpc_pocgrpc_proto_rawDescOnce.Do(func() {
		file_poc_grpc_pocgrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_poc_grpc_pocgrpc_proto_rawDescData)
	})
	return file_poc_grpc_pocgrpc_proto_rawDescData
}

var file_poc_grpc_pocgrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_poc_grpc_pocgrpc_proto_goTypes = []interface{}{
	(*BroadcastReq)(nil), // 0: pocgrpc.BroadcastReq
	(*BroadcastRes)(nil), // 1: pocgrpc.BroadcastRes
}
var file_poc_grpc_pocgrpc_proto_depIdxs = []int32{
	0, // 0: pocgrpc.Broadcaster.Broadcast:input_type -> pocgrpc.BroadcastReq
	1, // 1: pocgrpc.Broadcaster.Broadcast:output_type -> pocgrpc.BroadcastRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_poc_grpc_pocgrpc_proto_init() }
func file_poc_grpc_pocgrpc_proto_init() {
	if File_poc_grpc_pocgrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_poc_grpc_pocgrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastReq); i {
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
		file_poc_grpc_pocgrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastRes); i {
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
			RawDescriptor: file_poc_grpc_pocgrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_poc_grpc_pocgrpc_proto_goTypes,
		DependencyIndexes: file_poc_grpc_pocgrpc_proto_depIdxs,
		MessageInfos:      file_poc_grpc_pocgrpc_proto_msgTypes,
	}.Build()
	File_poc_grpc_pocgrpc_proto = out.File
	file_poc_grpc_pocgrpc_proto_rawDesc = nil
	file_poc_grpc_pocgrpc_proto_goTypes = nil
	file_poc_grpc_pocgrpc_proto_depIdxs = nil
}