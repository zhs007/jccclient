// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: publictransit.proto

package pb

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

// PublicTransitMode - PublicTransit mode
type PublicTransitMode int32

const (
	// PTM_TOKYOMETRO_SUBWAYS - subways
	PublicTransitMode_PTM_TOKYOMETRO_SUBWAYS PublicTransitMode = 0
	// PTM_TOKYOMETRO_LINE - line
	PublicTransitMode_PTM_TOKYOMETRO_LINE PublicTransitMode = 1
	// PTM_KOSTUMETROTOKYO_SUBWAYS - subways
	PublicTransitMode_PTM_KOSTUMETROTOKYO_SUBWAYS PublicTransitMode = 2
	// PTM_JRAILPASS_SUBWAYS - subways
	PublicTransitMode_PTM_JRAILPASS_SUBWAYS PublicTransitMode = 3
)

// Enum value maps for PublicTransitMode.
var (
	PublicTransitMode_name = map[int32]string{
		0: "PTM_TOKYOMETRO_SUBWAYS",
		1: "PTM_TOKYOMETRO_LINE",
		2: "PTM_KOSTUMETROTOKYO_SUBWAYS",
		3: "PTM_JRAILPASS_SUBWAYS",
	}
	PublicTransitMode_value = map[string]int32{
		"PTM_TOKYOMETRO_SUBWAYS":      0,
		"PTM_TOKYOMETRO_LINE":         1,
		"PTM_KOSTUMETROTOKYO_SUBWAYS": 2,
		"PTM_JRAILPASS_SUBWAYS":       3,
	}
)

func (x PublicTransitMode) Enum() *PublicTransitMode {
	p := new(PublicTransitMode)
	*p = x
	return p
}

func (x PublicTransitMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PublicTransitMode) Descriptor() protoreflect.EnumDescriptor {
	return file_publictransit_proto_enumTypes[0].Descriptor()
}

func (PublicTransitMode) Type() protoreflect.EnumType {
	return &file_publictransit_proto_enumTypes[0]
}

func (x PublicTransitMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PublicTransitMode.Descriptor instead.
func (PublicTransitMode) EnumDescriptor() ([]byte, []int) {
	return file_publictransit_proto_rawDescGZIP(), []int{0}
}

// PublicTransitLine - PublicTransit line
type PublicTransitLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url      string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Company  string   `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Country  string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	City     string   `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	Stations []string `protobuf:"bytes,6,rep,name=stations,proto3" json:"stations,omitempty"`
}

func (x *PublicTransitLine) Reset() {
	*x = PublicTransitLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publictransit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicTransitLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicTransitLine) ProtoMessage() {}

func (x *PublicTransitLine) ProtoReflect() protoreflect.Message {
	mi := &file_publictransit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicTransitLine.ProtoReflect.Descriptor instead.
func (*PublicTransitLine) Descriptor() ([]byte, []int) {
	return file_publictransit_proto_rawDescGZIP(), []int{0}
}

func (x *PublicTransitLine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PublicTransitLine) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PublicTransitLine) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *PublicTransitLine) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *PublicTransitLine) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *PublicTransitLine) GetStations() []string {
	if x != nil {
		return x.Stations
	}
	return nil
}

// PublicTransitLines - PublicTransit lines
type PublicTransitLines struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lines []*PublicTransitLine `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
}

func (x *PublicTransitLines) Reset() {
	*x = PublicTransitLines{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publictransit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicTransitLines) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicTransitLines) ProtoMessage() {}

func (x *PublicTransitLines) ProtoReflect() protoreflect.Message {
	mi := &file_publictransit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicTransitLines.ProtoReflect.Descriptor instead.
func (*PublicTransitLines) Descriptor() ([]byte, []int) {
	return file_publictransit_proto_rawDescGZIP(), []int{1}
}

func (x *PublicTransitLines) GetLines() []*PublicTransitLine {
	if x != nil {
		return x.Lines
	}
	return nil
}

// RequestPublicTransit - request PublicTransit
type RequestPublicTransit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode PublicTransitMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.PublicTransitMode" json:"mode,omitempty"`
	Url  string            `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *RequestPublicTransit) Reset() {
	*x = RequestPublicTransit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publictransit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPublicTransit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPublicTransit) ProtoMessage() {}

func (x *RequestPublicTransit) ProtoReflect() protoreflect.Message {
	mi := &file_publictransit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPublicTransit.ProtoReflect.Descriptor instead.
func (*RequestPublicTransit) Descriptor() ([]byte, []int) {
	return file_publictransit_proto_rawDescGZIP(), []int{2}
}

func (x *RequestPublicTransit) GetMode() PublicTransitMode {
	if x != nil {
		return x.Mode
	}
	return PublicTransitMode_PTM_TOKYOMETRO_SUBWAYS
}

func (x *RequestPublicTransit) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// ReplyPublicTransit - reply PublicTransit
type ReplyPublicTransit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode PublicTransitMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.PublicTransitMode" json:"mode,omitempty"`
	// Types that are assignable to Reply:
	//	*ReplyPublicTransit_Lines
	Reply isReplyPublicTransit_Reply `protobuf_oneof:"reply"`
}

func (x *ReplyPublicTransit) Reset() {
	*x = ReplyPublicTransit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publictransit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyPublicTransit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyPublicTransit) ProtoMessage() {}

func (x *ReplyPublicTransit) ProtoReflect() protoreflect.Message {
	mi := &file_publictransit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyPublicTransit.ProtoReflect.Descriptor instead.
func (*ReplyPublicTransit) Descriptor() ([]byte, []int) {
	return file_publictransit_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyPublicTransit) GetMode() PublicTransitMode {
	if x != nil {
		return x.Mode
	}
	return PublicTransitMode_PTM_TOKYOMETRO_SUBWAYS
}

func (m *ReplyPublicTransit) GetReply() isReplyPublicTransit_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (x *ReplyPublicTransit) GetLines() *PublicTransitLines {
	if x, ok := x.GetReply().(*ReplyPublicTransit_Lines); ok {
		return x.Lines
	}
	return nil
}

type isReplyPublicTransit_Reply interface {
	isReplyPublicTransit_Reply()
}

type ReplyPublicTransit_Lines struct {
	Lines *PublicTransitLines `protobuf:"bytes,100,opt,name=lines,proto3,oneof"`
}

func (*ReplyPublicTransit_Lines) isReplyPublicTransit_Reply() {}

var File_publictransit_proto protoreflect.FileDescriptor

var file_publictransit_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61,
	0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x11, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x50, 0x0a, 0x12, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x3a,
	0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x4c,
	0x69, 0x6e, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x62, 0x0a, 0x14, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x24, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x96,
	0x01, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77,
	0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12,
	0x3d, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x4c, 0x69, 0x6e, 0x65, 0x73, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x42, 0x07,
	0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x84, 0x01, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x16, 0x50, 0x54, 0x4d, 0x5f, 0x54, 0x4f, 0x4b, 0x59, 0x4f, 0x4d, 0x45, 0x54, 0x52, 0x4f, 0x5f,
	0x53, 0x55, 0x42, 0x57, 0x41, 0x59, 0x53, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x54, 0x4d,
	0x5f, 0x54, 0x4f, 0x4b, 0x59, 0x4f, 0x4d, 0x45, 0x54, 0x52, 0x4f, 0x5f, 0x4c, 0x49, 0x4e, 0x45,
	0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x54, 0x4d, 0x5f, 0x4b, 0x4f, 0x53, 0x54, 0x55, 0x4d,
	0x45, 0x54, 0x52, 0x4f, 0x54, 0x4f, 0x4b, 0x59, 0x4f, 0x5f, 0x53, 0x55, 0x42, 0x57, 0x41, 0x59,
	0x53, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x54, 0x4d, 0x5f, 0x4a, 0x52, 0x41, 0x49, 0x4c,
	0x50, 0x41, 0x53, 0x53, 0x5f, 0x53, 0x55, 0x42, 0x57, 0x41, 0x59, 0x53, 0x10, 0x03, 0x42, 0x20,
	0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x73,
	0x30, 0x30, 0x37, 0x2f, 0x6a, 0x63, 0x63, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_publictransit_proto_rawDescOnce sync.Once
	file_publictransit_proto_rawDescData = file_publictransit_proto_rawDesc
)

func file_publictransit_proto_rawDescGZIP() []byte {
	file_publictransit_proto_rawDescOnce.Do(func() {
		file_publictransit_proto_rawDescData = protoimpl.X.CompressGZIP(file_publictransit_proto_rawDescData)
	})
	return file_publictransit_proto_rawDescData
}

var file_publictransit_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_publictransit_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_publictransit_proto_goTypes = []interface{}{
	(PublicTransitMode)(0),       // 0: jarviscrawlercore.PublicTransitMode
	(*PublicTransitLine)(nil),    // 1: jarviscrawlercore.PublicTransitLine
	(*PublicTransitLines)(nil),   // 2: jarviscrawlercore.PublicTransitLines
	(*RequestPublicTransit)(nil), // 3: jarviscrawlercore.RequestPublicTransit
	(*ReplyPublicTransit)(nil),   // 4: jarviscrawlercore.ReplyPublicTransit
}
var file_publictransit_proto_depIdxs = []int32{
	1, // 0: jarviscrawlercore.PublicTransitLines.lines:type_name -> jarviscrawlercore.PublicTransitLine
	0, // 1: jarviscrawlercore.RequestPublicTransit.mode:type_name -> jarviscrawlercore.PublicTransitMode
	0, // 2: jarviscrawlercore.ReplyPublicTransit.mode:type_name -> jarviscrawlercore.PublicTransitMode
	2, // 3: jarviscrawlercore.ReplyPublicTransit.lines:type_name -> jarviscrawlercore.PublicTransitLines
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_publictransit_proto_init() }
func file_publictransit_proto_init() {
	if File_publictransit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publictransit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicTransitLine); i {
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
		file_publictransit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicTransitLines); i {
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
		file_publictransit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPublicTransit); i {
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
		file_publictransit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyPublicTransit); i {
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
	file_publictransit_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ReplyPublicTransit_Lines)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_publictransit_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_publictransit_proto_goTypes,
		DependencyIndexes: file_publictransit_proto_depIdxs,
		EnumInfos:         file_publictransit_proto_enumTypes,
		MessageInfos:      file_publictransit_proto_msgTypes,
	}.Build()
	File_publictransit_proto = out.File
	file_publictransit_proto_rawDesc = nil
	file_publictransit_proto_goTypes = nil
	file_publictransit_proto_depIdxs = nil
}
