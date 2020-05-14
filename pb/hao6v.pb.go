// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: hao6v.proto

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

type Hao6VResType int32

const (
	Hao6VResType_H6VRT_UNKNOW   Hao6VResType = 0
	Hao6VResType_H6VRT_ED2K     Hao6VResType = 1
	Hao6VResType_H6VRT_MAGNET   Hao6VResType = 2
	Hao6VResType_H6VRT_LIVE     Hao6VResType = 3
	Hao6VResType_H6VRT_BAIGUPAN Hao6VResType = 4
	Hao6VResType_H6VRT_XUNLEI   Hao6VResType = 5
)

// Enum value maps for Hao6VResType.
var (
	Hao6VResType_name = map[int32]string{
		0: "H6VRT_UNKNOW",
		1: "H6VRT_ED2K",
		2: "H6VRT_MAGNET",
		3: "H6VRT_LIVE",
		4: "H6VRT_BAIGUPAN",
		5: "H6VRT_XUNLEI",
	}
	Hao6VResType_value = map[string]int32{
		"H6VRT_UNKNOW":   0,
		"H6VRT_ED2K":     1,
		"H6VRT_MAGNET":   2,
		"H6VRT_LIVE":     3,
		"H6VRT_BAIGUPAN": 4,
		"H6VRT_XUNLEI":   5,
	}
)

func (x Hao6VResType) Enum() *Hao6VResType {
	p := new(Hao6VResType)
	*p = x
	return p
}

func (x Hao6VResType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Hao6VResType) Descriptor() protoreflect.EnumDescriptor {
	return file_hao6v_proto_enumTypes[0].Descriptor()
}

func (Hao6VResType) Type() protoreflect.EnumType {
	return &file_hao6v_proto_enumTypes[0]
}

func (x Hao6VResType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Hao6VResType.Descriptor instead.
func (Hao6VResType) EnumDescriptor() ([]byte, []int) {
	return file_hao6v_proto_rawDescGZIP(), []int{0}
}

// Hao6vMode - hao6v mode
type Hao6VMode int32

const (
	// H6VM_NEWPAGE - new page
	Hao6VMode_H6VM_NEWPAGE Hao6VMode = 0
	// H6VM_RESPAGE - res page
	Hao6VMode_H6VM_RESPAGE Hao6VMode = 1
)

// Enum value maps for Hao6VMode.
var (
	Hao6VMode_name = map[int32]string{
		0: "H6VM_NEWPAGE",
		1: "H6VM_RESPAGE",
	}
	Hao6VMode_value = map[string]int32{
		"H6VM_NEWPAGE": 0,
		"H6VM_RESPAGE": 1,
	}
)

func (x Hao6VMode) Enum() *Hao6VMode {
	p := new(Hao6VMode)
	*p = x
	return p
}

func (x Hao6VMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Hao6VMode) Descriptor() protoreflect.EnumDescriptor {
	return file_hao6v_proto_enumTypes[1].Descriptor()
}

func (Hao6VMode) Type() protoreflect.EnumType {
	return &file_hao6v_proto_enumTypes[1]
}

func (x Hao6VMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Hao6VMode.Descriptor instead.
func (Hao6VMode) EnumDescriptor() ([]byte, []int) {
	return file_hao6v_proto_rawDescGZIP(), []int{1}
}

// Hao6v Resource infomation
type Hao6VResInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Hao6VResType `protobuf:"varint,1,opt,name=type,proto3,enum=jarviscrawlercore.Hao6VResType" json:"type,omitempty"`
	Url  string       `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Name string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Code string       `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Hao6VResInfo) Reset() {
	*x = Hao6VResInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hao6v_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hao6VResInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hao6VResInfo) ProtoMessage() {}

func (x *Hao6VResInfo) ProtoReflect() protoreflect.Message {
	mi := &file_hao6v_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hao6VResInfo.ProtoReflect.Descriptor instead.
func (*Hao6VResInfo) Descriptor() ([]byte, []int) {
	return file_hao6v_proto_rawDescGZIP(), []int{0}
}

func (x *Hao6VResInfo) GetType() Hao6VResType {
	if x != nil {
		return x.Type
	}
	return Hao6VResType_H6VRT_UNKNOW
}

func (x *Hao6VResInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Hao6VResInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hao6VResInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

// Hao6v Node
type Hao6VNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fullname
	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	// resource id
	Resid string `protobuf:"bytes,2,opt,name=resid,proto3" json:"resid,omitempty"`
	// lst
	Lst []*Hao6VResInfo `protobuf:"bytes,3,rep,name=lst,proto3" json:"lst,omitempty"`
	// 片名
	Title []string `protobuf:"bytes,4,rep,name=title,proto3" json:"title,omitempty"`
	// 导演
	Director []string `protobuf:"bytes,5,rep,name=director,proto3" json:"director,omitempty"`
	// url
	Url string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	// cover
	Cover string `protobuf:"bytes,7,opt,name=cover,proto3" json:"cover,omitempty"`
	// 导演
	Fulldirector string `protobuf:"bytes,8,opt,name=fulldirector,proto3" json:"fulldirector,omitempty"`
}

func (x *Hao6VNode) Reset() {
	*x = Hao6VNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hao6v_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hao6VNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hao6VNode) ProtoMessage() {}

func (x *Hao6VNode) ProtoReflect() protoreflect.Message {
	mi := &file_hao6v_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hao6VNode.ProtoReflect.Descriptor instead.
func (*Hao6VNode) Descriptor() ([]byte, []int) {
	return file_hao6v_proto_rawDescGZIP(), []int{1}
}

func (x *Hao6VNode) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *Hao6VNode) GetResid() string {
	if x != nil {
		return x.Resid
	}
	return ""
}

func (x *Hao6VNode) GetLst() []*Hao6VResInfo {
	if x != nil {
		return x.Lst
	}
	return nil
}

func (x *Hao6VNode) GetTitle() []string {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *Hao6VNode) GetDirector() []string {
	if x != nil {
		return x.Director
	}
	return nil
}

func (x *Hao6VNode) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Hao6VNode) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *Hao6VNode) GetFulldirector() string {
	if x != nil {
		return x.Fulldirector
	}
	return ""
}

// Hao6vNewPage - hao6v new page
type Hao6VNewPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lst []*Hao6VNode `protobuf:"bytes,1,rep,name=lst,proto3" json:"lst,omitempty"`
}

func (x *Hao6VNewPage) Reset() {
	*x = Hao6VNewPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hao6v_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hao6VNewPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hao6VNewPage) ProtoMessage() {}

func (x *Hao6VNewPage) ProtoReflect() protoreflect.Message {
	mi := &file_hao6v_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hao6VNewPage.ProtoReflect.Descriptor instead.
func (*Hao6VNewPage) Descriptor() ([]byte, []int) {
	return file_hao6v_proto_rawDescGZIP(), []int{2}
}

func (x *Hao6VNewPage) GetLst() []*Hao6VNode {
	if x != nil {
		return x.Lst
	}
	return nil
}

// RequestHao6v - request hao6v
type RequestHao6V struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode Hao6VMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.Hao6VMode" json:"mode,omitempty"`
	Url  string    `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *RequestHao6V) Reset() {
	*x = RequestHao6V{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hao6v_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHao6V) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHao6V) ProtoMessage() {}

func (x *RequestHao6V) ProtoReflect() protoreflect.Message {
	mi := &file_hao6v_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHao6V.ProtoReflect.Descriptor instead.
func (*RequestHao6V) Descriptor() ([]byte, []int) {
	return file_hao6v_proto_rawDescGZIP(), []int{3}
}

func (x *RequestHao6V) GetMode() Hao6VMode {
	if x != nil {
		return x.Mode
	}
	return Hao6VMode_H6VM_NEWPAGE
}

func (x *RequestHao6V) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// ReplyHao6v - reply hao6v
type ReplyHao6V struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode Hao6VMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.Hao6VMode" json:"mode,omitempty"`
	// Types that are assignable to Reply:
	//	*ReplyHao6V_Newpage
	//	*ReplyHao6V_Res
	Reply isReplyHao6V_Reply `protobuf_oneof:"reply"`
}

func (x *ReplyHao6V) Reset() {
	*x = ReplyHao6V{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hao6v_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyHao6V) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyHao6V) ProtoMessage() {}

func (x *ReplyHao6V) ProtoReflect() protoreflect.Message {
	mi := &file_hao6v_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyHao6V.ProtoReflect.Descriptor instead.
func (*ReplyHao6V) Descriptor() ([]byte, []int) {
	return file_hao6v_proto_rawDescGZIP(), []int{4}
}

func (x *ReplyHao6V) GetMode() Hao6VMode {
	if x != nil {
		return x.Mode
	}
	return Hao6VMode_H6VM_NEWPAGE
}

func (m *ReplyHao6V) GetReply() isReplyHao6V_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (x *ReplyHao6V) GetNewpage() *Hao6VNewPage {
	if x, ok := x.GetReply().(*ReplyHao6V_Newpage); ok {
		return x.Newpage
	}
	return nil
}

func (x *ReplyHao6V) GetRes() *Hao6VNode {
	if x, ok := x.GetReply().(*ReplyHao6V_Res); ok {
		return x.Res
	}
	return nil
}

type isReplyHao6V_Reply interface {
	isReplyHao6V_Reply()
}

type ReplyHao6V_Newpage struct {
	Newpage *Hao6VNewPage `protobuf:"bytes,100,opt,name=newpage,proto3,oneof"`
}

type ReplyHao6V_Res struct {
	Res *Hao6VNode `protobuf:"bytes,101,opt,name=res,proto3,oneof"`
}

func (*ReplyHao6V_Newpage) isReplyHao6V_Reply() {}

func (*ReplyHao6V_Res) isReplyHao6V_Reply() {}

var File_hao6v_proto protoreflect.FileDescriptor

var file_hao6v_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x61, 0x6f, 0x36, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6a,
	0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65,
	0x22, 0x7d, 0x0a, 0x0c, 0x48, 0x61, 0x6f, 0x36, 0x76, 0x52, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x48, 0x61, 0x6f, 0x36, 0x76, 0x52, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0xee, 0x01, 0x0a, 0x09, 0x48, 0x61, 0x6f, 0x36, 0x76, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x73,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x73, 0x69, 0x64, 0x12,
	0x31, 0x0a, 0x03, 0x6c, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6a,
	0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x48, 0x61, 0x6f, 0x36, 0x76, 0x52, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x6c,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c,
	0x66, 0x75, 0x6c, 0x6c, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6c, 0x6c, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x22, 0x3e, 0x0a, 0x0c, 0x48, 0x61, 0x6f, 0x36, 0x76, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x2e, 0x0a, 0x03, 0x6c, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x48, 0x61, 0x6f, 0x36, 0x76, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x03, 0x6c, 0x73, 0x74,
	0x22, 0x52, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x61, 0x6f, 0x36, 0x76,
	0x12, 0x30, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x48, 0x61, 0x6f, 0x36, 0x76, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0xb6, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x61,
	0x6f, 0x36, 0x76, 0x12, 0x30, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65,
	0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x48, 0x61, 0x6f, 0x36, 0x76, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x48, 0x61, 0x6f, 0x36, 0x76,
	0x4e, 0x65, 0x77, 0x50, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x48, 0x61, 0x6f, 0x36, 0x76, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52,
	0x03, 0x72, 0x65, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x78, 0x0a,
	0x0c, 0x48, 0x61, 0x6f, 0x36, 0x76, 0x52, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x0c, 0x48, 0x36, 0x56, 0x52, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x48, 0x36, 0x56, 0x52, 0x54, 0x5f, 0x45, 0x44, 0x32, 0x4b, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x48, 0x36, 0x56, 0x52, 0x54, 0x5f, 0x4d, 0x41, 0x47, 0x4e, 0x45, 0x54, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x36, 0x56, 0x52, 0x54, 0x5f, 0x4c, 0x49, 0x56, 0x45, 0x10,
	0x03, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x36, 0x56, 0x52, 0x54, 0x5f, 0x42, 0x41, 0x49, 0x47, 0x55,
	0x50, 0x41, 0x4e, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x36, 0x56, 0x52, 0x54, 0x5f, 0x58,
	0x55, 0x4e, 0x4c, 0x45, 0x49, 0x10, 0x05, 0x2a, 0x2f, 0x0a, 0x09, 0x48, 0x61, 0x6f, 0x36, 0x76,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x36, 0x56, 0x4d, 0x5f, 0x4e, 0x45, 0x57,
	0x50, 0x41, 0x47, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x36, 0x56, 0x4d, 0x5f, 0x52,
	0x45, 0x53, 0x50, 0x41, 0x47, 0x45, 0x10, 0x01, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x73, 0x30, 0x30, 0x37, 0x2f, 0x6a, 0x63,
	0x63, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_hao6v_proto_rawDescOnce sync.Once
	file_hao6v_proto_rawDescData = file_hao6v_proto_rawDesc
)

func file_hao6v_proto_rawDescGZIP() []byte {
	file_hao6v_proto_rawDescOnce.Do(func() {
		file_hao6v_proto_rawDescData = protoimpl.X.CompressGZIP(file_hao6v_proto_rawDescData)
	})
	return file_hao6v_proto_rawDescData
}

var file_hao6v_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_hao6v_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_hao6v_proto_goTypes = []interface{}{
	(Hao6VResType)(0),    // 0: jarviscrawlercore.Hao6vResType
	(Hao6VMode)(0),       // 1: jarviscrawlercore.Hao6vMode
	(*Hao6VResInfo)(nil), // 2: jarviscrawlercore.Hao6vResInfo
	(*Hao6VNode)(nil),    // 3: jarviscrawlercore.Hao6vNode
	(*Hao6VNewPage)(nil), // 4: jarviscrawlercore.Hao6vNewPage
	(*RequestHao6V)(nil), // 5: jarviscrawlercore.RequestHao6v
	(*ReplyHao6V)(nil),   // 6: jarviscrawlercore.ReplyHao6v
}
var file_hao6v_proto_depIdxs = []int32{
	0, // 0: jarviscrawlercore.Hao6vResInfo.type:type_name -> jarviscrawlercore.Hao6vResType
	2, // 1: jarviscrawlercore.Hao6vNode.lst:type_name -> jarviscrawlercore.Hao6vResInfo
	3, // 2: jarviscrawlercore.Hao6vNewPage.lst:type_name -> jarviscrawlercore.Hao6vNode
	1, // 3: jarviscrawlercore.RequestHao6v.mode:type_name -> jarviscrawlercore.Hao6vMode
	1, // 4: jarviscrawlercore.ReplyHao6v.mode:type_name -> jarviscrawlercore.Hao6vMode
	4, // 5: jarviscrawlercore.ReplyHao6v.newpage:type_name -> jarviscrawlercore.Hao6vNewPage
	3, // 6: jarviscrawlercore.ReplyHao6v.res:type_name -> jarviscrawlercore.Hao6vNode
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_hao6v_proto_init() }
func file_hao6v_proto_init() {
	if File_hao6v_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hao6v_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hao6VResInfo); i {
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
		file_hao6v_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hao6VNode); i {
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
		file_hao6v_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hao6VNewPage); i {
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
		file_hao6v_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHao6V); i {
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
		file_hao6v_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyHao6V); i {
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
	file_hao6v_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ReplyHao6V_Newpage)(nil),
		(*ReplyHao6V_Res)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hao6v_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hao6v_proto_goTypes,
		DependencyIndexes: file_hao6v_proto_depIdxs,
		EnumInfos:         file_hao6v_proto_enumTypes,
		MessageInfos:      file_hao6v_proto_msgTypes,
	}.Build()
	File_hao6v_proto = out.File
	file_hao6v_proto_rawDesc = nil
	file_hao6v_proto_goTypes = nil
	file_hao6v_proto_depIdxs = nil
}
