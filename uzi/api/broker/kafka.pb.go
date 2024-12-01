// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: kafka.proto

package broker

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

type UziUpload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UziUpload) Reset() {
	*x = UziUpload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UziUpload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UziUpload) ProtoMessage() {}

func (x *UziUpload) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_kafka_proto_rawDescGZIP(), []int{0}
}

func (x *UziUpload) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int64 `protobuf:"varint,100,opt,name=x,proto3" json:"x,omitempty"`
	Y int64 `protobuf:"varint,200,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{1}
}

func (x *Point) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Tirads struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tirads_23 float64 `protobuf:"fixed64,1,opt,name=tirads_23,json=tirads23,proto3" json:"tirads_23,omitempty"`
	Tirads_4  float64 `protobuf:"fixed64,2,opt,name=tirads_4,json=tirads4,proto3" json:"tirads_4,omitempty"`
	Tirads_5  float64 `protobuf:"fixed64,3,opt,name=tirads_5,json=tirads5,proto3" json:"tirads_5,omitempty"`
}

func (x *Tirads) Reset() {
	*x = Tirads{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tirads) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tirads) ProtoMessage() {}

func (x *Tirads) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tirads.ProtoReflect.Descriptor instead.
func (*Tirads) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{2}
}

func (x *Tirads) GetTirads_23() float64 {
	if x != nil {
		return x.Tirads_23
	}
	return 0
}

func (x *Tirads) GetTirads_4() float64 {
	if x != nil {
		return x.Tirads_4
	}
	return 0
}

func (x *Tirads) GetTirads_5() float64 {
	if x != nil {
		return x.Tirads_5
	}
	return 0
}

type KafkaSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,100,opt,name=id,proto3" json:"id,omitempty"`
	FormationId string   `protobuf:"bytes,200,opt,name=formation_id,json=formationId,proto3" json:"formation_id,omitempty"`
	ImageId     string   `protobuf:"bytes,300,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	Contor      []*Point `protobuf:"bytes,400,rep,name=contor,proto3" json:"contor,omitempty"`
	Tirads      *Tirads  `protobuf:"bytes,500,opt,name=tirads,proto3" json:"tirads,omitempty"`
}

func (x *KafkaSegment) Reset() {
	*x = KafkaSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaSegment) ProtoMessage() {}

func (x *KafkaSegment) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaSegment.ProtoReflect.Descriptor instead.
func (*KafkaSegment) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{3}
}

func (x *KafkaSegment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *KafkaSegment) GetFormationId() string {
	if x != nil {
		return x.FormationId
	}
	return ""
}

func (x *KafkaSegment) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *KafkaSegment) GetContor() []*Point {
	if x != nil {
		return x.Contor
	}
	return nil
}

func (x *KafkaSegment) GetTirads() *Tirads {
	if x != nil {
		return x.Tirads
	}
	return nil
}

type KafkaFormation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,100,opt,name=id,proto3" json:"id,omitempty"`
	Tirads *Tirads `protobuf:"bytes,200,opt,name=tirads,proto3" json:"tirads,omitempty"`
	Ai     bool    `protobuf:"varint,300,opt,name=ai,proto3" json:"ai,omitempty"`
}

func (x *KafkaFormation) Reset() {
	*x = KafkaFormation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaFormation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaFormation) ProtoMessage() {}

func (x *KafkaFormation) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaFormation.ProtoReflect.Descriptor instead.
func (*KafkaFormation) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{4}
}

func (x *KafkaFormation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *KafkaFormation) GetTirads() *Tirads {
	if x != nil {
		return x.Tirads
	}
	return nil
}

func (x *KafkaFormation) GetAi() bool {
	if x != nil {
		return x.Ai
	}
	return false
}

type UziProcessed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Formations []*KafkaFormation `protobuf:"bytes,100,rep,name=formations,proto3" json:"formations,omitempty"`
	Segments   []*KafkaSegment   `protobuf:"bytes,200,rep,name=segments,proto3" json:"segments,omitempty"`
}

func (x *UziProcessed) Reset() {
	*x = UziProcessed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UziProcessed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UziProcessed) ProtoMessage() {}

func (x *UziProcessed) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UziProcessed.ProtoReflect.Descriptor instead.
func (*UziProcessed) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{5}
}

func (x *UziProcessed) GetFormations() []*KafkaFormation {
	if x != nil {
		return x.Formations
	}
	return nil
}

func (x *UziProcessed) GetSegments() []*KafkaSegment {
	if x != nil {
		return x.Segments
	}
	return nil
}

type UziSplitted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UziId   string   `protobuf:"bytes,100,opt,name=uzi_id,json=uziId,proto3" json:"uzi_id,omitempty"`
	PagesId []string `protobuf:"bytes,200,rep,name=pages_id,json=pagesId,proto3" json:"pages_id,omitempty"`
}

func (x *UziSplitted) Reset() {
	*x = UziSplitted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UziSplitted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UziSplitted) ProtoMessage() {}

func (x *UziSplitted) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_kafka_proto_rawDescGZIP(), []int{6}
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

var File_kafka_proto protoreflect.FileDescriptor

var file_kafka_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x79,
	0x69, 0x72, 0x2e, 0x75, 0x7a, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x22, 0x1b, 0x0a, 0x09,
	0x75, 0x7a, 0x69, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x05, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x78,
	0x12, 0x0d, 0x0a, 0x01, 0x79, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x79, 0x22,
	0x5b, 0x0a, 0x06, 0x54, 0x69, 0x72, 0x61, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x72,
	0x61, 0x64, 0x73, 0x5f, 0x32, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x74, 0x69,
	0x72, 0x61, 0x64, 0x73, 0x32, 0x33, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x69, 0x72, 0x61, 0x64, 0x73,
	0x5f, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x74, 0x69, 0x72, 0x61, 0x64, 0x73,
	0x34, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x69, 0x72, 0x61, 0x64, 0x73, 0x5f, 0x35, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x07, 0x74, 0x69, 0x72, 0x61, 0x64, 0x73, 0x35, 0x22, 0xbd, 0x01, 0x0a,
	0x0c, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0xc8, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xac, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x2d, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x74, 0x6f, 0x72, 0x18, 0x90, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x79, 0x69, 0x72, 0x2e, 0x75, 0x7a, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x74, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x06,
	0x74, 0x69, 0x72, 0x61, 0x64, 0x73, 0x18, 0xf4, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x79, 0x69, 0x72, 0x2e, 0x75, 0x7a, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x54, 0x69,
	0x72, 0x61, 0x64, 0x73, 0x52, 0x06, 0x74, 0x69, 0x72, 0x61, 0x64, 0x73, 0x22, 0x61, 0x0a, 0x0e,
	0x4b, 0x61, 0x66, 0x6b, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e,
	0x0a, 0x06, 0x74, 0x69, 0x72, 0x61, 0x64, 0x73, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x79, 0x69, 0x72, 0x2e, 0x75, 0x7a, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e,
	0x54, 0x69, 0x72, 0x61, 0x64, 0x73, 0x52, 0x06, 0x74, 0x69, 0x72, 0x61, 0x64, 0x73, 0x12, 0x0f,
	0x0a, 0x02, 0x61, 0x69, 0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x61, 0x69, 0x22,
	0x87, 0x01, 0x0a, 0x0c, 0x75, 0x7a, 0x69, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x12, 0x3d, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x64,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x79, 0x69, 0x72, 0x2e, 0x75, 0x7a, 0x69, 0x2e, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x38, 0x0a, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0xc8, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x79, 0x69, 0x72, 0x2e, 0x75, 0x7a, 0x69, 0x2e, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x40, 0x0a, 0x0b, 0x75, 0x7a, 0x69,
	0x53, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x75, 0x7a, 0x69, 0x5f,
	0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x7a, 0x69, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0xc8, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x73, 0x49, 0x64, 0x42, 0x1b, 0x5a, 0x19, 0x79,
	0x69, 0x72, 0x2f, 0x75, 0x7a, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x3b, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kafka_proto_rawDescOnce sync.Once
	file_kafka_proto_rawDescData = file_kafka_proto_rawDesc
)

func file_kafka_proto_rawDescGZIP() []byte {
	file_kafka_proto_rawDescOnce.Do(func() {
		file_kafka_proto_rawDescData = protoimpl.X.CompressGZIP(file_kafka_proto_rawDescData)
	})
	return file_kafka_proto_rawDescData
}

var file_kafka_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_kafka_proto_goTypes = []any{
	(*UziUpload)(nil),      // 0: yir.uzi.kafka.uziUpload
	(*Point)(nil),          // 1: yir.uzi.kafka.Point
	(*Tirads)(nil),         // 2: yir.uzi.kafka.Tirads
	(*KafkaSegment)(nil),   // 3: yir.uzi.kafka.KafkaSegment
	(*KafkaFormation)(nil), // 4: yir.uzi.kafka.KafkaFormation
	(*UziProcessed)(nil),   // 5: yir.uzi.kafka.uziProcessed
	(*UziSplitted)(nil),    // 6: yir.uzi.kafka.uziSplitted
}
var file_kafka_proto_depIdxs = []int32{
	1, // 0: yir.uzi.kafka.KafkaSegment.contor:type_name -> yir.uzi.kafka.Point
	2, // 1: yir.uzi.kafka.KafkaSegment.tirads:type_name -> yir.uzi.kafka.Tirads
	2, // 2: yir.uzi.kafka.KafkaFormation.tirads:type_name -> yir.uzi.kafka.Tirads
	4, // 3: yir.uzi.kafka.uziProcessed.formations:type_name -> yir.uzi.kafka.KafkaFormation
	3, // 4: yir.uzi.kafka.uziProcessed.segments:type_name -> yir.uzi.kafka.KafkaSegment
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_kafka_proto_init() }
func file_kafka_proto_init() {
	if File_kafka_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kafka_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UziUpload); i {
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
		file_kafka_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Point); i {
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
		file_kafka_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Tirads); i {
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
		file_kafka_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*KafkaSegment); i {
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
		file_kafka_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*KafkaFormation); i {
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
		file_kafka_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UziProcessed); i {
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
		file_kafka_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UziSplitted); i {
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
			RawDescriptor: file_kafka_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kafka_proto_goTypes,
		DependencyIndexes: file_kafka_proto_depIdxs,
		MessageInfos:      file_kafka_proto_msgTypes,
	}.Build()
	File_kafka_proto = out.File
	file_kafka_proto_rawDesc = nil
	file_kafka_proto_goTypes = nil
	file_kafka_proto_depIdxs = nil
}
