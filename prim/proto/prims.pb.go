// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: prims.proto

package proto

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

type PrimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PrimRequest) Reset() {
	*x = PrimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prims_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimRequest) ProtoMessage() {}

func (x *PrimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prims_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimRequest.ProtoReflect.Descriptor instead.
func (*PrimRequest) Descriptor() ([]byte, []int) {
	return file_prims_proto_rawDescGZIP(), []int{0}
}

func (x *PrimRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type PrimResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PrimResponse) Reset() {
	*x = PrimResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prims_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimResponse) ProtoMessage() {}

func (x *PrimResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prims_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimResponse.ProtoReflect.Descriptor instead.
func (*PrimResponse) Descriptor() ([]byte, []int) {
	return file_prims_proto_rawDescGZIP(), []int{1}
}

func (x *PrimResponse) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_prims_proto protoreflect.FileDescriptor

var file_prims_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70,
	0x72, 0x69, 0x6d, 0x22, 0x25, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x0c, 0x50, 0x72,
	0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x32, 0x43, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x34, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x11,
	0x2e, 0x70, 0x72, 0x69, 0x6d, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x58, 0x6e, 0x55, 0x73, 0x32, 0x35, 0x2f, 0x67, 0x52,
	0x50, 0x43, 0x50, 0x72, 0x69, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prims_proto_rawDescOnce sync.Once
	file_prims_proto_rawDescData = file_prims_proto_rawDesc
)

func file_prims_proto_rawDescGZIP() []byte {
	file_prims_proto_rawDescOnce.Do(func() {
		file_prims_proto_rawDescData = protoimpl.X.CompressGZIP(file_prims_proto_rawDescData)
	})
	return file_prims_proto_rawDescData
}

var file_prims_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_prims_proto_goTypes = []interface{}{
	(*PrimRequest)(nil),  // 0: prim.PrimRequest
	(*PrimResponse)(nil), // 1: prim.PrimResponse
}
var file_prims_proto_depIdxs = []int32{
	0, // 0: prim.PrimService.Primaries:input_type -> prim.PrimRequest
	1, // 1: prim.PrimService.Primaries:output_type -> prim.PrimResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_prims_proto_init() }
func file_prims_proto_init() {
	if File_prims_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prims_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimRequest); i {
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
		file_prims_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimResponse); i {
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
			RawDescriptor: file_prims_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prims_proto_goTypes,
		DependencyIndexes: file_prims_proto_depIdxs,
		MessageInfos:      file_prims_proto_msgTypes,
	}.Build()
	File_prims_proto = out.File
	file_prims_proto_rawDesc = nil
	file_prims_proto_goTypes = nil
	file_prims_proto_depIdxs = nil
}