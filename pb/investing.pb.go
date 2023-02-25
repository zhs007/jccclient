// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: investing.proto

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

// InvestingMode - investing mode
type InvestingMode int32

const (
	// INVESTINGMODE_ASSETS - get assets
	InvestingMode_INVESTINGMODE_ASSETS InvestingMode = 0
	// INVESTINGMODE_ASSET - get asset
	InvestingMode_INVESTINGMODE_ASSET InvestingMode = 1
	// INVESTINGMODE_HD - get asset historical data
	InvestingMode_INVESTINGMODE_HD InvestingMode = 2
)

// Enum value maps for InvestingMode.
var (
	InvestingMode_name = map[int32]string{
		0: "INVESTINGMODE_ASSETS",
		1: "INVESTINGMODE_ASSET",
		2: "INVESTINGMODE_HD",
	}
	InvestingMode_value = map[string]int32{
		"INVESTINGMODE_ASSETS": 0,
		"INVESTINGMODE_ASSET":  1,
		"INVESTINGMODE_HD":     2,
	}
)

func (x InvestingMode) Enum() *InvestingMode {
	p := new(InvestingMode)
	*p = x
	return p
}

func (x InvestingMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvestingMode) Descriptor() protoreflect.EnumDescriptor {
	return file_investing_proto_enumTypes[0].Descriptor()
}

func (InvestingMode) Type() protoreflect.EnumType {
	return &file_investing_proto_enumTypes[0]
}

func (x InvestingMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvestingMode.Descriptor instead.
func (InvestingMode) EnumDescriptor() ([]byte, []int) {
	return file_investing_proto_rawDescGZIP(), []int{0}
}

type InvestingHistoricalData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts     int64 `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Open   int64 `protobuf:"varint,2,opt,name=open,proto3" json:"open,omitempty"`
	Close  int64 `protobuf:"varint,3,opt,name=close,proto3" json:"close,omitempty"`
	High   int64 `protobuf:"varint,4,opt,name=high,proto3" json:"high,omitempty"`
	Low    int64 `protobuf:"varint,5,opt,name=low,proto3" json:"low,omitempty"`
	Volume int64 `protobuf:"varint,6,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *InvestingHistoricalData) Reset() {
	*x = InvestingHistoricalData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_investing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvestingHistoricalData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvestingHistoricalData) ProtoMessage() {}

func (x *InvestingHistoricalData) ProtoReflect() protoreflect.Message {
	mi := &file_investing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvestingHistoricalData.ProtoReflect.Descriptor instead.
func (*InvestingHistoricalData) Descriptor() ([]byte, []int) {
	return file_investing_proto_rawDescGZIP(), []int{0}
}

func (x *InvestingHistoricalData) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *InvestingHistoricalData) GetOpen() int64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *InvestingHistoricalData) GetClose() int64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *InvestingHistoricalData) GetHigh() int64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *InvestingHistoricalData) GetLow() int64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *InvestingHistoricalData) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

// InvestingAsset - investing asset
type InvestingAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url      string                     `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Data     []*InvestingHistoricalData `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	NameCode string                     `protobuf:"bytes,4,opt,name=nameCode,proto3" json:"nameCode,omitempty"`
}

func (x *InvestingAsset) Reset() {
	*x = InvestingAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_investing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvestingAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvestingAsset) ProtoMessage() {}

func (x *InvestingAsset) ProtoReflect() protoreflect.Message {
	mi := &file_investing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvestingAsset.ProtoReflect.Descriptor instead.
func (*InvestingAsset) Descriptor() ([]byte, []int) {
	return file_investing_proto_rawDescGZIP(), []int{1}
}

func (x *InvestingAsset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InvestingAsset) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *InvestingAsset) GetData() []*InvestingHistoricalData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *InvestingAsset) GetNameCode() string {
	if x != nil {
		return x.NameCode
	}
	return ""
}

// InvestingAssets - investing assets
type InvestingAssets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assets []*InvestingAsset `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
}

func (x *InvestingAssets) Reset() {
	*x = InvestingAssets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_investing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvestingAssets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvestingAssets) ProtoMessage() {}

func (x *InvestingAssets) ProtoReflect() protoreflect.Message {
	mi := &file_investing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvestingAssets.ProtoReflect.Descriptor instead.
func (*InvestingAssets) Descriptor() ([]byte, []int) {
	return file_investing_proto_rawDescGZIP(), []int{2}
}

func (x *InvestingAssets) GetAssets() []*InvestingAsset {
	if x != nil {
		return x.Assets
	}
	return nil
}

// RequestInvesting - request investing
type RequestInvesting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode      InvestingMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.InvestingMode" json:"mode,omitempty"`
	Url       string        `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	StartData string        `protobuf:"bytes,3,opt,name=startData,proto3" json:"startData,omitempty"`
	EndData   string        `protobuf:"bytes,4,opt,name=endData,proto3" json:"endData,omitempty"`
}

func (x *RequestInvesting) Reset() {
	*x = RequestInvesting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_investing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestInvesting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestInvesting) ProtoMessage() {}

func (x *RequestInvesting) ProtoReflect() protoreflect.Message {
	mi := &file_investing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestInvesting.ProtoReflect.Descriptor instead.
func (*RequestInvesting) Descriptor() ([]byte, []int) {
	return file_investing_proto_rawDescGZIP(), []int{3}
}

func (x *RequestInvesting) GetMode() InvestingMode {
	if x != nil {
		return x.Mode
	}
	return InvestingMode_INVESTINGMODE_ASSETS
}

func (x *RequestInvesting) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RequestInvesting) GetStartData() string {
	if x != nil {
		return x.StartData
	}
	return ""
}

func (x *RequestInvesting) GetEndData() string {
	if x != nil {
		return x.EndData
	}
	return ""
}

// ReplyInvesting - reply investing
type ReplyInvesting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode InvestingMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.InvestingMode" json:"mode,omitempty"`
	// Types that are assignable to Reply:
	//	*ReplyInvesting_Asset
	//	*ReplyInvesting_Assets
	Reply isReplyInvesting_Reply `protobuf_oneof:"reply"`
}

func (x *ReplyInvesting) Reset() {
	*x = ReplyInvesting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_investing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyInvesting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyInvesting) ProtoMessage() {}

func (x *ReplyInvesting) ProtoReflect() protoreflect.Message {
	mi := &file_investing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyInvesting.ProtoReflect.Descriptor instead.
func (*ReplyInvesting) Descriptor() ([]byte, []int) {
	return file_investing_proto_rawDescGZIP(), []int{4}
}

func (x *ReplyInvesting) GetMode() InvestingMode {
	if x != nil {
		return x.Mode
	}
	return InvestingMode_INVESTINGMODE_ASSETS
}

func (m *ReplyInvesting) GetReply() isReplyInvesting_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (x *ReplyInvesting) GetAsset() *InvestingAsset {
	if x, ok := x.GetReply().(*ReplyInvesting_Asset); ok {
		return x.Asset
	}
	return nil
}

func (x *ReplyInvesting) GetAssets() *InvestingAssets {
	if x, ok := x.GetReply().(*ReplyInvesting_Assets); ok {
		return x.Assets
	}
	return nil
}

type isReplyInvesting_Reply interface {
	isReplyInvesting_Reply()
}

type ReplyInvesting_Asset struct {
	Asset *InvestingAsset `protobuf:"bytes,100,opt,name=asset,proto3,oneof"`
}

type ReplyInvesting_Assets struct {
	Assets *InvestingAssets `protobuf:"bytes,101,opt,name=assets,proto3,oneof"`
}

func (*ReplyInvesting_Asset) isReplyInvesting_Reply() {}

func (*ReplyInvesting_Assets) isReplyInvesting_Reply() {}

var File_investing_proto protoreflect.FileDescriptor

var file_investing_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72,
	0x63, 0x6f, 0x72, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x17, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x6f, 0x70, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69,
	0x67, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6c, 0x6f, 0x77,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x76,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x3e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x4c, 0x0a,
	0x0f, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x12, 0x39, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x10,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x34, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x22, 0xc8, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x34, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65,
	0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69,
	0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x76,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x73, 0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x05, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61,
	0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x48, 0x00, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x58, 0x0a, 0x0d, 0x49,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x14,
	0x49, 0x4e, 0x56, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x53,
	0x53, 0x45, 0x54, 0x53, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x45, 0x53, 0x54,
	0x49, 0x4e, 0x47, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x01, 0x12,
	0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x4d, 0x4f, 0x44, 0x45,
	0x5f, 0x48, 0x44, 0x10, 0x02, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x73, 0x30, 0x30, 0x37, 0x2f, 0x6a, 0x63, 0x63, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_investing_proto_rawDescOnce sync.Once
	file_investing_proto_rawDescData = file_investing_proto_rawDesc
)

func file_investing_proto_rawDescGZIP() []byte {
	file_investing_proto_rawDescOnce.Do(func() {
		file_investing_proto_rawDescData = protoimpl.X.CompressGZIP(file_investing_proto_rawDescData)
	})
	return file_investing_proto_rawDescData
}

var file_investing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_investing_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_investing_proto_goTypes = []interface{}{
	(InvestingMode)(0),              // 0: jarviscrawlercore.InvestingMode
	(*InvestingHistoricalData)(nil), // 1: jarviscrawlercore.InvestingHistoricalData
	(*InvestingAsset)(nil),          // 2: jarviscrawlercore.InvestingAsset
	(*InvestingAssets)(nil),         // 3: jarviscrawlercore.InvestingAssets
	(*RequestInvesting)(nil),        // 4: jarviscrawlercore.RequestInvesting
	(*ReplyInvesting)(nil),          // 5: jarviscrawlercore.ReplyInvesting
}
var file_investing_proto_depIdxs = []int32{
	1, // 0: jarviscrawlercore.InvestingAsset.data:type_name -> jarviscrawlercore.InvestingHistoricalData
	2, // 1: jarviscrawlercore.InvestingAssets.assets:type_name -> jarviscrawlercore.InvestingAsset
	0, // 2: jarviscrawlercore.RequestInvesting.mode:type_name -> jarviscrawlercore.InvestingMode
	0, // 3: jarviscrawlercore.ReplyInvesting.mode:type_name -> jarviscrawlercore.InvestingMode
	2, // 4: jarviscrawlercore.ReplyInvesting.asset:type_name -> jarviscrawlercore.InvestingAsset
	3, // 5: jarviscrawlercore.ReplyInvesting.assets:type_name -> jarviscrawlercore.InvestingAssets
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_investing_proto_init() }
func file_investing_proto_init() {
	if File_investing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_investing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvestingHistoricalData); i {
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
		file_investing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvestingAsset); i {
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
		file_investing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvestingAssets); i {
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
		file_investing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestInvesting); i {
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
		file_investing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyInvesting); i {
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
	file_investing_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ReplyInvesting_Asset)(nil),
		(*ReplyInvesting_Assets)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_investing_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_investing_proto_goTypes,
		DependencyIndexes: file_investing_proto_depIdxs,
		EnumInfos:         file_investing_proto_enumTypes,
		MessageInfos:      file_investing_proto_msgTypes,
	}.Build()
	File_investing_proto = out.File
	file_investing_proto_rawDesc = nil
	file_investing_proto_goTypes = nil
	file_investing_proto_depIdxs = nil
}