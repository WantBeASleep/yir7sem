// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: uzisplitted.proto

package uzisplitted

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UziSplitted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UziId   string   `protobuf:"bytes,100,opt,name=uzi_id,json=uziId,proto3" json:"uzi_id,omitempty"`
	PagesId []string `protobuf:"bytes,200,rep,name=pages_id,json=pagesId,proto3" json:"pages_id,omitempty"`
}

func (x *UziSplitted) Reset() {
	*x = UziSplitted{}
	mi := &file_uzisplitted_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UziSplitted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UziSplitted) ProtoMessage() {}

func (x *UziSplitted) ProtoReflect() protoreflect.Message {
	mi := &file_uzisplitted_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UziSplitted.ProtoReflect.Descriptor instead.
func (*UziSplitted) Descriptor() ([]byte, []int) {
	return file_uzisplitted_proto_rawDescGZIP(), []int{0}
}

func (x *UziSplitted) GetUziId() string {
	if x != nil {
		return x.UziId
	}
	return ""
}

func (x *UziSplitted) GetPagesId() []string {
	if x != nil {
		return x.PagesId
	}
	return nil
}

var File_uzisplitted_proto protoreflect.FileDescriptor

var file_uzisplitted_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x7a, 0x69, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0b, 0x75, 0x7a, 0x69, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x74,
	0x65, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x75, 0x7a, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x75, 0x7a, 0x69, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x49, 0x64, 0x42, 0x2f, 0x5a, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x2f, 0x75, 0x7a, 0x69, 0x73, 0x70,
	0x6c, 0x69, 0x74, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uzisplitted_proto_rawDescOnce sync.Once
	file_uzisplitted_proto_rawDescData = file_uzisplitted_proto_rawDesc
)

func file_uzisplitted_proto_rawDescGZIP() []byte {
	file_uzisplitted_proto_rawDescOnce.Do(func() {
		file_uzisplitted_proto_rawDescData = protoimpl.X.CompressGZIP(file_uzisplitted_proto_rawDescData)
	})
	return file_uzisplitted_proto_rawDescData
}

var file_uzisplitted_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_uzisplitted_proto_goTypes = []any{
	(*UziSplitted)(nil), // 0: uziSplitted
}
var file_uzisplitted_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_uzisplitted_proto_init() }
func file_uzisplitted_proto_init() {
	if File_uzisplitted_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_uzisplitted_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_uzisplitted_proto_goTypes,
		DependencyIndexes: file_uzisplitted_proto_depIdxs,
		MessageInfos:      file_uzisplitted_proto_msgTypes,
	}.Build()
	File_uzisplitted_proto = out.File
	file_uzisplitted_proto_rawDesc = nil
	file_uzisplitted_proto_goTypes = nil
	file_uzisplitted_proto_depIdxs = nil
}
