// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: uziupload.proto

package uziupload

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

type UziUpload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UziId string `protobuf:"bytes,100,opt,name=uzi_id,json=uziId,proto3" json:"uzi_id,omitempty"`
}

func (x *UziUpload) Reset() {
	*x = UziUpload{}
	mi := &file_uziupload_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UziUpload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UziUpload) ProtoMessage() {}

func (x *UziUpload) ProtoReflect() protoreflect.Message {
	mi := &file_uziupload_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UziUpload.ProtoReflect.Descriptor instead.
func (*UziUpload) Descriptor() ([]byte, []int) {
	return file_uziupload_proto_rawDescGZIP(), []int{0}
}

func (x *UziUpload) GetUziId() string {
	if x != nil {
		return x.UziId
	}
	return ""
}

var File_uziupload_proto protoreflect.FileDescriptor

var file_uziupload_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x7a, 0x69, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x22, 0x0a, 0x09, 0x55, 0x7a, 0x69, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x75, 0x7a, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x75, 0x7a, 0x69, 0x49, 0x64, 0x42, 0x2d, 0x5a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x2f, 0x75, 0x7a, 0x69, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uziupload_proto_rawDescOnce sync.Once
	file_uziupload_proto_rawDescData = file_uziupload_proto_rawDesc
)

func file_uziupload_proto_rawDescGZIP() []byte {
	file_uziupload_proto_rawDescOnce.Do(func() {
		file_uziupload_proto_rawDescData = protoimpl.X.CompressGZIP(file_uziupload_proto_rawDescData)
	})
	return file_uziupload_proto_rawDescData
}

var file_uziupload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_uziupload_proto_goTypes = []any{
	(*UziUpload)(nil), // 0: UziUpload
}
var file_uziupload_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_uziupload_proto_init() }
func file_uziupload_proto_init() {
	if File_uziupload_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_uziupload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_uziupload_proto_goTypes,
		DependencyIndexes: file_uziupload_proto_depIdxs,
		MessageInfos:      file_uziupload_proto_msgTypes,
	}.Build()
	File_uziupload_proto = out.File
	file_uziupload_proto_rawDesc = nil
	file_uziupload_proto_goTypes = nil
	file_uziupload_proto_depIdxs = nil
}
