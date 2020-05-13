// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: oabt.proto

package jccclient

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

type OABTResType int32

const (
	OABTResType_OABTRT_ED2K   OABTResType = 0
	OABTResType_OABTRT_MAGNET OABTResType = 1
)

// Enum value maps for OABTResType.
var (
	OABTResType_name = map[int32]string{
		0: "OABTRT_ED2K",
		1: "OABTRT_MAGNET",
	}
	OABTResType_value = map[string]int32{
		"OABTRT_ED2K":   0,
		"OABTRT_MAGNET": 1,
	}
)

func (x OABTResType) Enum() *OABTResType {
	p := new(OABTResType)
	*p = x
	return p
}

func (x OABTResType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OABTResType) Descriptor() protoreflect.EnumDescriptor {
	return file_oabt_proto_enumTypes[0].Descriptor()
}

func (OABTResType) Type() protoreflect.EnumType {
	return &file_oabt_proto_enumTypes[0]
}

func (x OABTResType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OABTResType.Descriptor instead.
func (OABTResType) EnumDescriptor() ([]byte, []int) {
	return file_oabt_proto_rawDescGZIP(), []int{0}
}

// OABTMode - oabt mode
type OABTMode int32

const (
	// OABTM_PAGE - page
	OABTMode_OABTM_PAGE OABTMode = 0
)

// Enum value maps for OABTMode.
var (
	OABTMode_name = map[int32]string{
		0: "OABTM_PAGE",
	}
	OABTMode_value = map[string]int32{
		"OABTM_PAGE": 0,
	}
)

func (x OABTMode) Enum() *OABTMode {
	p := new(OABTMode)
	*p = x
	return p
}

func (x OABTMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OABTMode) Descriptor() protoreflect.EnumDescriptor {
	return file_oabt_proto_enumTypes[1].Descriptor()
}

func (OABTMode) Type() protoreflect.EnumType {
	return &file_oabt_proto_enumTypes[1]
}

func (x OABTMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OABTMode.Descriptor instead.
func (OABTMode) EnumDescriptor() ([]byte, []int) {
	return file_oabt_proto_rawDescGZIP(), []int{1}
}

// OABT Resource infomation
type OABTResInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type OABTResType `protobuf:"varint,1,opt,name=type,proto3,enum=jarviscrawlercore.OABTResType" json:"type,omitempty"`
	Url  string      `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *OABTResInfo) Reset() {
	*x = OABTResInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oabt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OABTResInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OABTResInfo) ProtoMessage() {}

func (x *OABTResInfo) ProtoReflect() protoreflect.Message {
	mi := &file_oabt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OABTResInfo.ProtoReflect.Descriptor instead.
func (*OABTResInfo) Descriptor() ([]byte, []int) {
	return file_oabt_proto_rawDescGZIP(), []int{0}
}

func (x *OABTResInfo) GetType() OABTResType {
	if x != nil {
		return x.Type
	}
	return OABTResType_OABTRT_ED2K
}

func (x *OABTResInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// OABT Node
type OABTNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fullname
	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	// resource id
	Resid string `protobuf:"bytes,2,opt,name=resid,proto3" json:"resid,omitempty"`
	// category
	Cat int32 `protobuf:"varint,3,opt,name=cat,proto3" json:"cat,omitempty"`
	// lst
	Lst []*OABTResInfo `protobuf:"bytes,4,rep,name=lst,proto3" json:"lst,omitempty"`
}

func (x *OABTNode) Reset() {
	*x = OABTNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oabt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OABTNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OABTNode) ProtoMessage() {}

func (x *OABTNode) ProtoReflect() protoreflect.Message {
	mi := &file_oabt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OABTNode.ProtoReflect.Descriptor instead.
func (*OABTNode) Descriptor() ([]byte, []int) {
	return file_oabt_proto_rawDescGZIP(), []int{1}
}

func (x *OABTNode) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *OABTNode) GetResid() string {
	if x != nil {
		return x.Resid
	}
	return ""
}

func (x *OABTNode) GetCat() int32 {
	if x != nil {
		return x.Cat
	}
	return 0
}

func (x *OABTNode) GetLst() []*OABTResInfo {
	if x != nil {
		return x.Lst
	}
	return nil
}

type OABTPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lst []*OABTNode `protobuf:"bytes,1,rep,name=lst,proto3" json:"lst,omitempty"`
}

func (x *OABTPage) Reset() {
	*x = OABTPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oabt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OABTPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OABTPage) ProtoMessage() {}

func (x *OABTPage) ProtoReflect() protoreflect.Message {
	mi := &file_oabt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OABTPage.ProtoReflect.Descriptor instead.
func (*OABTPage) Descriptor() ([]byte, []int) {
	return file_oabt_proto_rawDescGZIP(), []int{2}
}

func (x *OABTPage) GetLst() []*OABTNode {
	if x != nil {
		return x.Lst
	}
	return nil
}

// RequestOABT - request oabt
type RequestOABT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode      OABTMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.OABTMode" json:"mode,omitempty"`
	PageIndex int32    `protobuf:"varint,2,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
}

func (x *RequestOABT) Reset() {
	*x = RequestOABT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oabt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestOABT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestOABT) ProtoMessage() {}

func (x *RequestOABT) ProtoReflect() protoreflect.Message {
	mi := &file_oabt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestOABT.ProtoReflect.Descriptor instead.
func (*RequestOABT) Descriptor() ([]byte, []int) {
	return file_oabt_proto_rawDescGZIP(), []int{3}
}

func (x *RequestOABT) GetMode() OABTMode {
	if x != nil {
		return x.Mode
	}
	return OABTMode_OABTM_PAGE
}

func (x *RequestOABT) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

// ReplyOABT - reply oabt
type ReplyOABT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode OABTMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.OABTMode" json:"mode,omitempty"`
	// Types that are assignable to Reply:
	//	*ReplyOABT_Page
	Reply isReplyOABT_Reply `protobuf_oneof:"reply"`
}

func (x *ReplyOABT) Reset() {
	*x = ReplyOABT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oabt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyOABT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyOABT) ProtoMessage() {}

func (x *ReplyOABT) ProtoReflect() protoreflect.Message {
	mi := &file_oabt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyOABT.ProtoReflect.Descriptor instead.
func (*ReplyOABT) Descriptor() ([]byte, []int) {
	return file_oabt_proto_rawDescGZIP(), []int{4}
}

func (x *ReplyOABT) GetMode() OABTMode {
	if x != nil {
		return x.Mode
	}
	return OABTMode_OABTM_PAGE
}

func (m *ReplyOABT) GetReply() isReplyOABT_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (x *ReplyOABT) GetPage() *OABTPage {
	if x, ok := x.GetReply().(*ReplyOABT_Page); ok {
		return x.Page
	}
	return nil
}

type isReplyOABT_Reply interface {
	isReplyOABT_Reply()
}

type ReplyOABT_Page struct {
	Page *OABTPage `protobuf:"bytes,100,opt,name=page,proto3,oneof"`
}

func (*ReplyOABT_Page) isReplyOABT_Reply() {}

var File_oabt_proto protoreflect.FileDescriptor

var file_oabt_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6f, 0x61, 0x62, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6a, 0x61,
	0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x22,
	0x53, 0x0a, 0x0b, 0x4f, 0x41, 0x42, 0x54, 0x52, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6a,
	0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x4f, 0x41, 0x42, 0x54, 0x52, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x4f, 0x41, 0x42, 0x54, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x65, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65,
	0x73, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x63, 0x61, 0x74, 0x12, 0x30, 0x0a, 0x03, 0x6c, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c,
	0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x41, 0x42, 0x54, 0x52, 0x65, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x03, 0x6c, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x08, 0x4f, 0x41, 0x42, 0x54, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x6c, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x41, 0x42, 0x54, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x03, 0x6c,
	0x73, 0x74, 0x22, 0x5c, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x41, 0x42,
	0x54, 0x12, 0x2f, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x41, 0x42, 0x54, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x22, 0x78, 0x0a, 0x09, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4f, 0x41, 0x42, 0x54, 0x12, 0x2f, 0x0a,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6a, 0x61,
	0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x4f, 0x41, 0x42, 0x54, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x31,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a,
	0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x4f, 0x41, 0x42, 0x54, 0x50, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x31, 0x0a, 0x0b, 0x4f, 0x41,
	0x42, 0x54, 0x52, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x41, 0x42,
	0x54, 0x52, 0x54, 0x5f, 0x45, 0x44, 0x32, 0x4b, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x41,
	0x42, 0x54, 0x52, 0x54, 0x5f, 0x4d, 0x41, 0x47, 0x4e, 0x45, 0x54, 0x10, 0x01, 0x2a, 0x1a, 0x0a,
	0x08, 0x4f, 0x41, 0x42, 0x54, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x41, 0x42,
	0x54, 0x4d, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x10, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2f, 0x7a, 0x68, 0x73, 0x30, 0x30, 0x37, 0x2f, 0x6a, 0x63, 0x63, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oabt_proto_rawDescOnce sync.Once
	file_oabt_proto_rawDescData = file_oabt_proto_rawDesc
)

func file_oabt_proto_rawDescGZIP() []byte {
	file_oabt_proto_rawDescOnce.Do(func() {
		file_oabt_proto_rawDescData = protoimpl.X.CompressGZIP(file_oabt_proto_rawDescData)
	})
	return file_oabt_proto_rawDescData
}

var file_oabt_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_oabt_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_oabt_proto_goTypes = []interface{}{
	(OABTResType)(0),    // 0: jarviscrawlercore.OABTResType
	(OABTMode)(0),       // 1: jarviscrawlercore.OABTMode
	(*OABTResInfo)(nil), // 2: jarviscrawlercore.OABTResInfo
	(*OABTNode)(nil),    // 3: jarviscrawlercore.OABTNode
	(*OABTPage)(nil),    // 4: jarviscrawlercore.OABTPage
	(*RequestOABT)(nil), // 5: jarviscrawlercore.RequestOABT
	(*ReplyOABT)(nil),   // 6: jarviscrawlercore.ReplyOABT
}
var file_oabt_proto_depIdxs = []int32{
	0, // 0: jarviscrawlercore.OABTResInfo.type:type_name -> jarviscrawlercore.OABTResType
	2, // 1: jarviscrawlercore.OABTNode.lst:type_name -> jarviscrawlercore.OABTResInfo
	3, // 2: jarviscrawlercore.OABTPage.lst:type_name -> jarviscrawlercore.OABTNode
	1, // 3: jarviscrawlercore.RequestOABT.mode:type_name -> jarviscrawlercore.OABTMode
	1, // 4: jarviscrawlercore.ReplyOABT.mode:type_name -> jarviscrawlercore.OABTMode
	4, // 5: jarviscrawlercore.ReplyOABT.page:type_name -> jarviscrawlercore.OABTPage
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_oabt_proto_init() }
func file_oabt_proto_init() {
	if File_oabt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oabt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OABTResInfo); i {
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
		file_oabt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OABTNode); i {
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
		file_oabt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OABTPage); i {
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
		file_oabt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestOABT); i {
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
		file_oabt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyOABT); i {
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
	file_oabt_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ReplyOABT_Page)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oabt_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oabt_proto_goTypes,
		DependencyIndexes: file_oabt_proto_depIdxs,
		EnumInfos:         file_oabt_proto_enumTypes,
		MessageInfos:      file_oabt_proto_msgTypes,
	}.Build()
	File_oabt_proto = out.File
	file_oabt_proto_rawDesc = nil
	file_oabt_proto_goTypes = nil
	file_oabt_proto_depIdxs = nil
}
