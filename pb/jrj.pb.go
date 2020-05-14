// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: jrj.proto

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

// JRJMode - jrj mode
type JRJMode int32

const (
	// JRJM_FUNDS - funds
	JRJMode_JRJM_FUNDS JRJMode = 0
	// JRJM_FUND - fund
	JRJMode_JRJM_FUND JRJMode = 1
	// JRJM_FUNDVALUE - fund value
	JRJMode_JRJM_FUNDVALUE JRJMode = 2
	// JRJM_FUNDMANAGER - fund manager
	JRJMode_JRJM_FUNDMANAGER JRJMode = 3
)

// Enum value maps for JRJMode.
var (
	JRJMode_name = map[int32]string{
		0: "JRJM_FUNDS",
		1: "JRJM_FUND",
		2: "JRJM_FUNDVALUE",
		3: "JRJM_FUNDMANAGER",
	}
	JRJMode_value = map[string]int32{
		"JRJM_FUNDS":       0,
		"JRJM_FUND":        1,
		"JRJM_FUNDVALUE":   2,
		"JRJM_FUNDMANAGER": 3,
	}
)

func (x JRJMode) Enum() *JRJMode {
	p := new(JRJMode)
	*p = x
	return p
}

func (x JRJMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JRJMode) Descriptor() protoreflect.EnumDescriptor {
	return file_jrj_proto_enumTypes[0].Descriptor()
}

func (JRJMode) Type() protoreflect.EnumType {
	return &file_jrj_proto_enumTypes[0]
}

func (x JRJMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JRJMode.Descriptor instead.
func (JRJMode) EnumDescriptor() ([]byte, []int) {
	return file_jrj_proto_rawDescGZIP(), []int{0}
}

// JRJFundSize - jrj fund size
type JRJFundSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size float32 `protobuf:"fixed32,1,opt,name=size,proto3" json:"size,omitempty"`
	Time int64   `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *JRJFundSize) Reset() {
	*x = JRJFundSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jrj_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JRJFundSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JRJFundSize) ProtoMessage() {}

func (x *JRJFundSize) ProtoReflect() protoreflect.Message {
	mi := &file_jrj_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JRJFundSize.ProtoReflect.Descriptor instead.
func (*JRJFundSize) Descriptor() ([]byte, []int) {
	return file_jrj_proto_rawDescGZIP(), []int{0}
}

func (x *JRJFundSize) GetSize() float32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *JRJFundSize) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

// JRJFundManager - jrj fund manager
type JRJFundManager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StartTime int64  `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime   int64  `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	BirthYear int64  `protobuf:"varint,4,opt,name=birthYear,proto3" json:"birthYear,omitempty"`
	Sex       bool   `protobuf:"varint,5,opt,name=sex,proto3" json:"sex,omitempty"`
	Education string `protobuf:"bytes,6,opt,name=education,proto3" json:"education,omitempty"`
	Country   string `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty"`
	Resume    string `protobuf:"bytes,8,opt,name=resume,proto3" json:"resume,omitempty"`
}

func (x *JRJFundManager) Reset() {
	*x = JRJFundManager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jrj_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JRJFundManager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JRJFundManager) ProtoMessage() {}

func (x *JRJFundManager) ProtoReflect() protoreflect.Message {
	mi := &file_jrj_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JRJFundManager.ProtoReflect.Descriptor instead.
func (*JRJFundManager) Descriptor() ([]byte, []int) {
	return file_jrj_proto_rawDescGZIP(), []int{1}
}

func (x *JRJFundManager) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JRJFundManager) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *JRJFundManager) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *JRJFundManager) GetBirthYear() int64 {
	if x != nil {
		return x.BirthYear
	}
	return 0
}

func (x *JRJFundManager) GetSex() bool {
	if x != nil {
		return x.Sex
	}
	return false
}

func (x *JRJFundManager) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *JRJFundManager) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *JRJFundManager) GetResume() string {
	if x != nil {
		return x.Resume
	}
	return ""
}

type JRJFundDayValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date       string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Value      int32  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	TotalValue int32  `protobuf:"varint,3,opt,name=totalValue,proto3" json:"totalValue,omitempty"`
}

func (x *JRJFundDayValue) Reset() {
	*x = JRJFundDayValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jrj_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JRJFundDayValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JRJFundDayValue) ProtoMessage() {}

func (x *JRJFundDayValue) ProtoReflect() protoreflect.Message {
	mi := &file_jrj_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JRJFundDayValue.ProtoReflect.Descriptor instead.
func (*JRJFundDayValue) Descriptor() ([]byte, []int) {
	return file_jrj_proto_rawDescGZIP(), []int{2}
}

func (x *JRJFundDayValue) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *JRJFundDayValue) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *JRJFundDayValue) GetTotalValue() int32 {
	if x != nil {
		return x.TotalValue
	}
	return 0
}

// JRJFundValue - jrj fund value
type JRJFundValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// Deprecated: Do not use.
	Value []float32 `protobuf:"fixed32,2,rep,packed,name=value,proto3" json:"value,omitempty"`
	// Deprecated: Do not use.
	TotalValue []float32 `protobuf:"fixed32,3,rep,packed,name=totalValue,proto3" json:"totalValue,omitempty"`
	// Deprecated: Do not use.
	IValue []int32 `protobuf:"varint,4,rep,packed,name=iValue,proto3" json:"iValue,omitempty"`
	// Deprecated: Do not use.
	ITotalValue []int32 `protobuf:"varint,5,rep,packed,name=iTotalValue,proto3" json:"iTotalValue,omitempty"`
	// Deprecated: Do not use.
	Date   []string           `protobuf:"bytes,6,rep,name=date,proto3" json:"date,omitempty"`
	Values []*JRJFundDayValue `protobuf:"bytes,7,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *JRJFundValue) Reset() {
	*x = JRJFundValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jrj_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JRJFundValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JRJFundValue) ProtoMessage() {}

func (x *JRJFundValue) ProtoReflect() protoreflect.Message {
	mi := &file_jrj_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JRJFundValue.ProtoReflect.Descriptor instead.
func (*JRJFundValue) Descriptor() ([]byte, []int) {
	return file_jrj_proto_rawDescGZIP(), []int{3}
}

func (x *JRJFundValue) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

// Deprecated: Do not use.
func (x *JRJFundValue) GetValue() []float32 {
	if x != nil {
		return x.Value
	}
	return nil
}

// Deprecated: Do not use.
func (x *JRJFundValue) GetTotalValue() []float32 {
	if x != nil {
		return x.TotalValue
	}
	return nil
}

// Deprecated: Do not use.
func (x *JRJFundValue) GetIValue() []int32 {
	if x != nil {
		return x.IValue
	}
	return nil
}

// Deprecated: Do not use.
func (x *JRJFundValue) GetITotalValue() []int32 {
	if x != nil {
		return x.ITotalValue
	}
	return nil
}

// Deprecated: Do not use.
func (x *JRJFundValue) GetDate() []string {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *JRJFundValue) GetValues() []*JRJFundDayValue {
	if x != nil {
		return x.Values
	}
	return nil
}

// JRJFund - jrj fund
type JRJFund struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       string            `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name       string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tags       []string          `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	CreateTime int64             `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Size       []*JRJFundSize    `protobuf:"bytes,5,rep,name=size,proto3" json:"size,omitempty"`
	Company    string            `protobuf:"bytes,6,opt,name=company,proto3" json:"company,omitempty"`
	Managers   []*JRJFundManager `protobuf:"bytes,7,rep,name=managers,proto3" json:"managers,omitempty"`
}

func (x *JRJFund) Reset() {
	*x = JRJFund{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jrj_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JRJFund) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JRJFund) ProtoMessage() {}

func (x *JRJFund) ProtoReflect() protoreflect.Message {
	mi := &file_jrj_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JRJFund.ProtoReflect.Descriptor instead.
func (*JRJFund) Descriptor() ([]byte, []int) {
	return file_jrj_proto_rawDescGZIP(), []int{4}
}

func (x *JRJFund) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *JRJFund) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JRJFund) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *JRJFund) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *JRJFund) GetSize() []*JRJFundSize {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *JRJFund) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *JRJFund) GetManagers() []*JRJFundManager {
	if x != nil {
		return x.Managers
	}
	return nil
}

// JRJFunds - jrj funds
type JRJFunds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes []string `protobuf:"bytes,1,rep,name=codes,proto3" json:"codes,omitempty"`
}

func (x *JRJFunds) Reset() {
	*x = JRJFunds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jrj_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JRJFunds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JRJFunds) ProtoMessage() {}

func (x *JRJFunds) ProtoReflect() protoreflect.Message {
	mi := &file_jrj_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JRJFunds.ProtoReflect.Descriptor instead.
func (*JRJFunds) Descriptor() ([]byte, []int) {
	return file_jrj_proto_rawDescGZIP(), []int{5}
}

func (x *JRJFunds) GetCodes() []string {
	if x != nil {
		return x.Codes
	}
	return nil
}

// RequestJRJ - request jrj
type RequestJRJ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode     JRJMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.JRJMode" json:"mode,omitempty"`
	Fundcode string  `protobuf:"bytes,2,opt,name=fundcode,proto3" json:"fundcode,omitempty"`
	Year     string  `protobuf:"bytes,3,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *RequestJRJ) Reset() {
	*x = RequestJRJ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jrj_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestJRJ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestJRJ) ProtoMessage() {}

func (x *RequestJRJ) ProtoReflect() protoreflect.Message {
	mi := &file_jrj_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestJRJ.ProtoReflect.Descriptor instead.
func (*RequestJRJ) Descriptor() ([]byte, []int) {
	return file_jrj_proto_rawDescGZIP(), []int{6}
}

func (x *RequestJRJ) GetMode() JRJMode {
	if x != nil {
		return x.Mode
	}
	return JRJMode_JRJM_FUNDS
}

func (x *RequestJRJ) GetFundcode() string {
	if x != nil {
		return x.Fundcode
	}
	return ""
}

func (x *RequestJRJ) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

// ReplySteepAndCheap - reply steep&cheap
type ReplyJRJ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode JRJMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.JRJMode" json:"mode,omitempty"`
	// Types that are assignable to Reply:
	//	*ReplyJRJ_Funds
	//	*ReplyJRJ_Fund
	//	*ReplyJRJ_FundValue
	Reply isReplyJRJ_Reply `protobuf_oneof:"reply"`
}

func (x *ReplyJRJ) Reset() {
	*x = ReplyJRJ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jrj_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyJRJ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyJRJ) ProtoMessage() {}

func (x *ReplyJRJ) ProtoReflect() protoreflect.Message {
	mi := &file_jrj_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyJRJ.ProtoReflect.Descriptor instead.
func (*ReplyJRJ) Descriptor() ([]byte, []int) {
	return file_jrj_proto_rawDescGZIP(), []int{7}
}

func (x *ReplyJRJ) GetMode() JRJMode {
	if x != nil {
		return x.Mode
	}
	return JRJMode_JRJM_FUNDS
}

func (m *ReplyJRJ) GetReply() isReplyJRJ_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (x *ReplyJRJ) GetFunds() *JRJFunds {
	if x, ok := x.GetReply().(*ReplyJRJ_Funds); ok {
		return x.Funds
	}
	return nil
}

func (x *ReplyJRJ) GetFund() *JRJFund {
	if x, ok := x.GetReply().(*ReplyJRJ_Fund); ok {
		return x.Fund
	}
	return nil
}

func (x *ReplyJRJ) GetFundValue() *JRJFundValue {
	if x, ok := x.GetReply().(*ReplyJRJ_FundValue); ok {
		return x.FundValue
	}
	return nil
}

type isReplyJRJ_Reply interface {
	isReplyJRJ_Reply()
}

type ReplyJRJ_Funds struct {
	Funds *JRJFunds `protobuf:"bytes,100,opt,name=funds,proto3,oneof"`
}

type ReplyJRJ_Fund struct {
	Fund *JRJFund `protobuf:"bytes,101,opt,name=fund,proto3,oneof"`
}

type ReplyJRJ_FundValue struct {
	FundValue *JRJFundValue `protobuf:"bytes,102,opt,name=fundValue,proto3,oneof"`
}

func (*ReplyJRJ_Funds) isReplyJRJ_Reply() {}

func (*ReplyJRJ_Fund) isReplyJRJ_Reply() {}

func (*ReplyJRJ_FundValue) isReplyJRJ_Reply() {}

var File_jrj_proto protoreflect.FileDescriptor

var file_jrj_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6a, 0x72, 0x6a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6a, 0x61, 0x72,
	0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x35,
	0x0a, 0x0b, 0x4a, 0x52, 0x4a, 0x46, 0x75, 0x6e, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xdc, 0x01, 0x0a, 0x0e, 0x4a, 0x52, 0x4a, 0x46, 0x75, 0x6e,
	0x64, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x59, 0x65, 0x61,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x59, 0x65,
	0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x73, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x0f, 0x4a, 0x52, 0x4a, 0x46, 0x75, 0x6e, 0x64, 0x44,
	0x61, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0xf6, 0x01, 0x0a, 0x0c, 0x4a, 0x52, 0x4a, 0x46, 0x75, 0x6e, 0x64, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x02, 0x42, 0x02, 0x18, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x22, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x02, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x06, 0x69, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x05, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x69, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x24, 0x0a, 0x0b, 0x69, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x05, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0b, 0x69, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3a,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x4a, 0x52, 0x4a, 0x46, 0x75, 0x6e, 0x64, 0x44, 0x61, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xf2, 0x01, 0x0a, 0x07, 0x4a,
	0x52, 0x4a, 0x46, 0x75, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4a, 0x52, 0x4a, 0x46, 0x75, 0x6e, 0x64, 0x53, 0x69, 0x7a, 0x65,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c,
	0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4a, 0x52, 0x4a, 0x46, 0x75, 0x6e, 0x64, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x08, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x22,
	0x20, 0x0a, 0x08, 0x4a, 0x52, 0x4a, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x22, 0x6c, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4a, 0x52, 0x4a, 0x12,
	0x2e, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x4a, 0x52, 0x4a, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x64, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x64, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22,
	0xeb, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4a, 0x52, 0x4a, 0x12, 0x2e, 0x0a, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6a, 0x61, 0x72,
	0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4a,
	0x52, 0x4a, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x05,
	0x66, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x61,
	0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x4a, 0x52, 0x4a, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x48, 0x00, 0x52, 0x05, 0x66, 0x75, 0x6e, 0x64,
	0x73, 0x12, 0x30, 0x0a, 0x04, 0x66, 0x75, 0x6e, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x4a, 0x52, 0x4a, 0x46, 0x75, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x04, 0x66,
	0x75, 0x6e, 0x64, 0x12, 0x3f, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4a, 0x52, 0x4a, 0x46, 0x75,
	0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x64, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x52, 0x0a,
	0x07, 0x4a, 0x52, 0x4a, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4a, 0x52, 0x4a, 0x4d,
	0x5f, 0x46, 0x55, 0x4e, 0x44, 0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4a, 0x52, 0x4a, 0x4d,
	0x5f, 0x46, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4a, 0x52, 0x4a, 0x4d, 0x5f,
	0x46, 0x55, 0x4e, 0x44, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x4a,
	0x52, 0x4a, 0x4d, 0x5f, 0x46, 0x55, 0x4e, 0x44, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10,
	0x03, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x7a, 0x68, 0x73, 0x30, 0x30, 0x37, 0x2f, 0x6a, 0x63, 0x63, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jrj_proto_rawDescOnce sync.Once
	file_jrj_proto_rawDescData = file_jrj_proto_rawDesc
)

func file_jrj_proto_rawDescGZIP() []byte {
	file_jrj_proto_rawDescOnce.Do(func() {
		file_jrj_proto_rawDescData = protoimpl.X.CompressGZIP(file_jrj_proto_rawDescData)
	})
	return file_jrj_proto_rawDescData
}

var file_jrj_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_jrj_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_jrj_proto_goTypes = []interface{}{
	(JRJMode)(0),            // 0: jarviscrawlercore.JRJMode
	(*JRJFundSize)(nil),     // 1: jarviscrawlercore.JRJFundSize
	(*JRJFundManager)(nil),  // 2: jarviscrawlercore.JRJFundManager
	(*JRJFundDayValue)(nil), // 3: jarviscrawlercore.JRJFundDayValue
	(*JRJFundValue)(nil),    // 4: jarviscrawlercore.JRJFundValue
	(*JRJFund)(nil),         // 5: jarviscrawlercore.JRJFund
	(*JRJFunds)(nil),        // 6: jarviscrawlercore.JRJFunds
	(*RequestJRJ)(nil),      // 7: jarviscrawlercore.RequestJRJ
	(*ReplyJRJ)(nil),        // 8: jarviscrawlercore.ReplyJRJ
}
var file_jrj_proto_depIdxs = []int32{
	3, // 0: jarviscrawlercore.JRJFundValue.values:type_name -> jarviscrawlercore.JRJFundDayValue
	1, // 1: jarviscrawlercore.JRJFund.size:type_name -> jarviscrawlercore.JRJFundSize
	2, // 2: jarviscrawlercore.JRJFund.managers:type_name -> jarviscrawlercore.JRJFundManager
	0, // 3: jarviscrawlercore.RequestJRJ.mode:type_name -> jarviscrawlercore.JRJMode
	0, // 4: jarviscrawlercore.ReplyJRJ.mode:type_name -> jarviscrawlercore.JRJMode
	6, // 5: jarviscrawlercore.ReplyJRJ.funds:type_name -> jarviscrawlercore.JRJFunds
	5, // 6: jarviscrawlercore.ReplyJRJ.fund:type_name -> jarviscrawlercore.JRJFund
	4, // 7: jarviscrawlercore.ReplyJRJ.fundValue:type_name -> jarviscrawlercore.JRJFundValue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_jrj_proto_init() }
func file_jrj_proto_init() {
	if File_jrj_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jrj_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JRJFundSize); i {
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
		file_jrj_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JRJFundManager); i {
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
		file_jrj_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JRJFundDayValue); i {
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
		file_jrj_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JRJFundValue); i {
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
		file_jrj_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JRJFund); i {
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
		file_jrj_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JRJFunds); i {
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
		file_jrj_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestJRJ); i {
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
		file_jrj_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyJRJ); i {
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
	file_jrj_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*ReplyJRJ_Funds)(nil),
		(*ReplyJRJ_Fund)(nil),
		(*ReplyJRJ_FundValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_jrj_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jrj_proto_goTypes,
		DependencyIndexes: file_jrj_proto_depIdxs,
		EnumInfos:         file_jrj_proto_enumTypes,
		MessageInfos:      file_jrj_proto_msgTypes,
	}.Build()
	File_jrj_proto = out.File
	file_jrj_proto_rawDesc = nil
	file_jrj_proto_goTypes = nil
	file_jrj_proto_depIdxs = nil
}
