// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jrj.proto

package jarviscrawlercore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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

var JRJMode_name = map[int32]string{
	0: "JRJM_FUNDS",
	1: "JRJM_FUND",
	2: "JRJM_FUNDVALUE",
	3: "JRJM_FUNDMANAGER",
}
var JRJMode_value = map[string]int32{
	"JRJM_FUNDS":       0,
	"JRJM_FUND":        1,
	"JRJM_FUNDVALUE":   2,
	"JRJM_FUNDMANAGER": 3,
}

func (x JRJMode) String() string {
	return proto.EnumName(JRJMode_name, int32(x))
}
func (JRJMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_jrj_5707464fb197bfbd, []int{0}
}

// JRJFundSize - jrj fund size
type JRJFundSize struct {
	Size                 float32  `protobuf:"fixed32,1,opt,name=size,proto3" json:"size,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JRJFundSize) Reset()         { *m = JRJFundSize{} }
func (m *JRJFundSize) String() string { return proto.CompactTextString(m) }
func (*JRJFundSize) ProtoMessage()    {}
func (*JRJFundSize) Descriptor() ([]byte, []int) {
	return fileDescriptor_jrj_5707464fb197bfbd, []int{0}
}
func (m *JRJFundSize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JRJFundSize.Unmarshal(m, b)
}
func (m *JRJFundSize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JRJFundSize.Marshal(b, m, deterministic)
}
func (dst *JRJFundSize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JRJFundSize.Merge(dst, src)
}
func (m *JRJFundSize) XXX_Size() int {
	return xxx_messageInfo_JRJFundSize.Size(m)
}
func (m *JRJFundSize) XXX_DiscardUnknown() {
	xxx_messageInfo_JRJFundSize.DiscardUnknown(m)
}

var xxx_messageInfo_JRJFundSize proto.InternalMessageInfo

func (m *JRJFundSize) GetSize() float32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *JRJFundSize) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

// JRJFundManager - jrj fund manager
type JRJFundManager struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StartTime            int64    `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              int64    `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	BirthYear            int64    `protobuf:"varint,4,opt,name=birthYear,proto3" json:"birthYear,omitempty"`
	Sex                  bool     `protobuf:"varint,5,opt,name=sex,proto3" json:"sex,omitempty"`
	Education            string   `protobuf:"bytes,6,opt,name=education,proto3" json:"education,omitempty"`
	Country              string   `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty"`
	Resume               string   `protobuf:"bytes,8,opt,name=resume,proto3" json:"resume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JRJFundManager) Reset()         { *m = JRJFundManager{} }
func (m *JRJFundManager) String() string { return proto.CompactTextString(m) }
func (*JRJFundManager) ProtoMessage()    {}
func (*JRJFundManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_jrj_5707464fb197bfbd, []int{1}
}
func (m *JRJFundManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JRJFundManager.Unmarshal(m, b)
}
func (m *JRJFundManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JRJFundManager.Marshal(b, m, deterministic)
}
func (dst *JRJFundManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JRJFundManager.Merge(dst, src)
}
func (m *JRJFundManager) XXX_Size() int {
	return xxx_messageInfo_JRJFundManager.Size(m)
}
func (m *JRJFundManager) XXX_DiscardUnknown() {
	xxx_messageInfo_JRJFundManager.DiscardUnknown(m)
}

var xxx_messageInfo_JRJFundManager proto.InternalMessageInfo

func (m *JRJFundManager) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JRJFundManager) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *JRJFundManager) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *JRJFundManager) GetBirthYear() int64 {
	if m != nil {
		return m.BirthYear
	}
	return 0
}

func (m *JRJFundManager) GetSex() bool {
	if m != nil {
		return m.Sex
	}
	return false
}

func (m *JRJFundManager) GetEducation() string {
	if m != nil {
		return m.Education
	}
	return ""
}

func (m *JRJFundManager) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *JRJFundManager) GetResume() string {
	if m != nil {
		return m.Resume
	}
	return ""
}

type JRJFundDayValue struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	TotalValue           int32    `protobuf:"varint,3,opt,name=totalValue,proto3" json:"totalValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JRJFundDayValue) Reset()         { *m = JRJFundDayValue{} }
func (m *JRJFundDayValue) String() string { return proto.CompactTextString(m) }
func (*JRJFundDayValue) ProtoMessage()    {}
func (*JRJFundDayValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_jrj_5707464fb197bfbd, []int{2}
}
func (m *JRJFundDayValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JRJFundDayValue.Unmarshal(m, b)
}
func (m *JRJFundDayValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JRJFundDayValue.Marshal(b, m, deterministic)
}
func (dst *JRJFundDayValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JRJFundDayValue.Merge(dst, src)
}
func (m *JRJFundDayValue) XXX_Size() int {
	return xxx_messageInfo_JRJFundDayValue.Size(m)
}
func (m *JRJFundDayValue) XXX_DiscardUnknown() {
	xxx_messageInfo_JRJFundDayValue.DiscardUnknown(m)
}

var xxx_messageInfo_JRJFundDayValue proto.InternalMessageInfo

func (m *JRJFundDayValue) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *JRJFundDayValue) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *JRJFundDayValue) GetTotalValue() int32 {
	if m != nil {
		return m.TotalValue
	}
	return 0
}

// JRJFundValue - jrj fund value
type JRJFundValue struct {
	Code                 string             `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Value                []float32          `protobuf:"fixed32,2,rep,packed,name=value,proto3" json:"value,omitempty"`            // Deprecated: Do not use.
	TotalValue           []float32          `protobuf:"fixed32,3,rep,packed,name=totalValue,proto3" json:"totalValue,omitempty"`  // Deprecated: Do not use.
	IValue               []int32            `protobuf:"varint,4,rep,packed,name=iValue,proto3" json:"iValue,omitempty"`           // Deprecated: Do not use.
	ITotalValue          []int32            `protobuf:"varint,5,rep,packed,name=iTotalValue,proto3" json:"iTotalValue,omitempty"` // Deprecated: Do not use.
	Date                 []string           `protobuf:"bytes,6,rep,name=date,proto3" json:"date,omitempty"`                       // Deprecated: Do not use.
	Values               []*JRJFundDayValue `protobuf:"bytes,7,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *JRJFundValue) Reset()         { *m = JRJFundValue{} }
func (m *JRJFundValue) String() string { return proto.CompactTextString(m) }
func (*JRJFundValue) ProtoMessage()    {}
func (*JRJFundValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_jrj_5707464fb197bfbd, []int{3}
}
func (m *JRJFundValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JRJFundValue.Unmarshal(m, b)
}
func (m *JRJFundValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JRJFundValue.Marshal(b, m, deterministic)
}
func (dst *JRJFundValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JRJFundValue.Merge(dst, src)
}
func (m *JRJFundValue) XXX_Size() int {
	return xxx_messageInfo_JRJFundValue.Size(m)
}
func (m *JRJFundValue) XXX_DiscardUnknown() {
	xxx_messageInfo_JRJFundValue.DiscardUnknown(m)
}

var xxx_messageInfo_JRJFundValue proto.InternalMessageInfo

func (m *JRJFundValue) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

// Deprecated: Do not use.
func (m *JRJFundValue) GetValue() []float32 {
	if m != nil {
		return m.Value
	}
	return nil
}

// Deprecated: Do not use.
func (m *JRJFundValue) GetTotalValue() []float32 {
	if m != nil {
		return m.TotalValue
	}
	return nil
}

// Deprecated: Do not use.
func (m *JRJFundValue) GetIValue() []int32 {
	if m != nil {
		return m.IValue
	}
	return nil
}

// Deprecated: Do not use.
func (m *JRJFundValue) GetITotalValue() []int32 {
	if m != nil {
		return m.ITotalValue
	}
	return nil
}

// Deprecated: Do not use.
func (m *JRJFundValue) GetDate() []string {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *JRJFundValue) GetValues() []*JRJFundDayValue {
	if m != nil {
		return m.Values
	}
	return nil
}

// JRJFund - jrj fund
type JRJFund struct {
	Code                 string            `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tags                 []string          `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	CreateTime           int64             `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Size                 []*JRJFundSize    `protobuf:"bytes,5,rep,name=size,proto3" json:"size,omitempty"`
	Company              string            `protobuf:"bytes,6,opt,name=company,proto3" json:"company,omitempty"`
	Managers             []*JRJFundManager `protobuf:"bytes,7,rep,name=managers,proto3" json:"managers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JRJFund) Reset()         { *m = JRJFund{} }
func (m *JRJFund) String() string { return proto.CompactTextString(m) }
func (*JRJFund) ProtoMessage()    {}
func (*JRJFund) Descriptor() ([]byte, []int) {
	return fileDescriptor_jrj_5707464fb197bfbd, []int{4}
}
func (m *JRJFund) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JRJFund.Unmarshal(m, b)
}
func (m *JRJFund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JRJFund.Marshal(b, m, deterministic)
}
func (dst *JRJFund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JRJFund.Merge(dst, src)
}
func (m *JRJFund) XXX_Size() int {
	return xxx_messageInfo_JRJFund.Size(m)
}
func (m *JRJFund) XXX_DiscardUnknown() {
	xxx_messageInfo_JRJFund.DiscardUnknown(m)
}

var xxx_messageInfo_JRJFund proto.InternalMessageInfo

func (m *JRJFund) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *JRJFund) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JRJFund) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *JRJFund) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *JRJFund) GetSize() []*JRJFundSize {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *JRJFund) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *JRJFund) GetManagers() []*JRJFundManager {
	if m != nil {
		return m.Managers
	}
	return nil
}

// JRJFunds - jrj funds
type JRJFunds struct {
	Codes                []string `protobuf:"bytes,1,rep,name=codes,proto3" json:"codes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JRJFunds) Reset()         { *m = JRJFunds{} }
func (m *JRJFunds) String() string { return proto.CompactTextString(m) }
func (*JRJFunds) ProtoMessage()    {}
func (*JRJFunds) Descriptor() ([]byte, []int) {
	return fileDescriptor_jrj_5707464fb197bfbd, []int{5}
}
func (m *JRJFunds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JRJFunds.Unmarshal(m, b)
}
func (m *JRJFunds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JRJFunds.Marshal(b, m, deterministic)
}
func (dst *JRJFunds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JRJFunds.Merge(dst, src)
}
func (m *JRJFunds) XXX_Size() int {
	return xxx_messageInfo_JRJFunds.Size(m)
}
func (m *JRJFunds) XXX_DiscardUnknown() {
	xxx_messageInfo_JRJFunds.DiscardUnknown(m)
}

var xxx_messageInfo_JRJFunds proto.InternalMessageInfo

func (m *JRJFunds) GetCodes() []string {
	if m != nil {
		return m.Codes
	}
	return nil
}

// RequestJRJ - request jrj
type RequestJRJ struct {
	Mode                 JRJMode  `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.JRJMode" json:"mode,omitempty"`
	Fundcode             string   `protobuf:"bytes,2,opt,name=fundcode,proto3" json:"fundcode,omitempty"`
	Year                 string   `protobuf:"bytes,3,opt,name=year,proto3" json:"year,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestJRJ) Reset()         { *m = RequestJRJ{} }
func (m *RequestJRJ) String() string { return proto.CompactTextString(m) }
func (*RequestJRJ) ProtoMessage()    {}
func (*RequestJRJ) Descriptor() ([]byte, []int) {
	return fileDescriptor_jrj_5707464fb197bfbd, []int{6}
}
func (m *RequestJRJ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestJRJ.Unmarshal(m, b)
}
func (m *RequestJRJ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestJRJ.Marshal(b, m, deterministic)
}
func (dst *RequestJRJ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestJRJ.Merge(dst, src)
}
func (m *RequestJRJ) XXX_Size() int {
	return xxx_messageInfo_RequestJRJ.Size(m)
}
func (m *RequestJRJ) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestJRJ.DiscardUnknown(m)
}

var xxx_messageInfo_RequestJRJ proto.InternalMessageInfo

func (m *RequestJRJ) GetMode() JRJMode {
	if m != nil {
		return m.Mode
	}
	return JRJMode_JRJM_FUNDS
}

func (m *RequestJRJ) GetFundcode() string {
	if m != nil {
		return m.Fundcode
	}
	return ""
}

func (m *RequestJRJ) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

// ReplySteepAndCheap - reply steep&cheap
type ReplyJRJ struct {
	Mode JRJMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.JRJMode" json:"mode,omitempty"`
	// Types that are valid to be assigned to Reply:
	//	*ReplyJRJ_Funds
	//	*ReplyJRJ_Fund
	//	*ReplyJRJ_FundValue
	Reply                isReplyJRJ_Reply `protobuf_oneof:"reply"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReplyJRJ) Reset()         { *m = ReplyJRJ{} }
func (m *ReplyJRJ) String() string { return proto.CompactTextString(m) }
func (*ReplyJRJ) ProtoMessage()    {}
func (*ReplyJRJ) Descriptor() ([]byte, []int) {
	return fileDescriptor_jrj_5707464fb197bfbd, []int{7}
}
func (m *ReplyJRJ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyJRJ.Unmarshal(m, b)
}
func (m *ReplyJRJ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyJRJ.Marshal(b, m, deterministic)
}
func (dst *ReplyJRJ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyJRJ.Merge(dst, src)
}
func (m *ReplyJRJ) XXX_Size() int {
	return xxx_messageInfo_ReplyJRJ.Size(m)
}
func (m *ReplyJRJ) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyJRJ.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyJRJ proto.InternalMessageInfo

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

func (*ReplyJRJ_Funds) isReplyJRJ_Reply()     {}
func (*ReplyJRJ_Fund) isReplyJRJ_Reply()      {}
func (*ReplyJRJ_FundValue) isReplyJRJ_Reply() {}

func (m *ReplyJRJ) GetReply() isReplyJRJ_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *ReplyJRJ) GetMode() JRJMode {
	if m != nil {
		return m.Mode
	}
	return JRJMode_JRJM_FUNDS
}

func (m *ReplyJRJ) GetFunds() *JRJFunds {
	if x, ok := m.GetReply().(*ReplyJRJ_Funds); ok {
		return x.Funds
	}
	return nil
}

func (m *ReplyJRJ) GetFund() *JRJFund {
	if x, ok := m.GetReply().(*ReplyJRJ_Fund); ok {
		return x.Fund
	}
	return nil
}

func (m *ReplyJRJ) GetFundValue() *JRJFundValue {
	if x, ok := m.GetReply().(*ReplyJRJ_FundValue); ok {
		return x.FundValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReplyJRJ) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReplyJRJ_OneofMarshaler, _ReplyJRJ_OneofUnmarshaler, _ReplyJRJ_OneofSizer, []interface{}{
		(*ReplyJRJ_Funds)(nil),
		(*ReplyJRJ_Fund)(nil),
		(*ReplyJRJ_FundValue)(nil),
	}
}

func _ReplyJRJ_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReplyJRJ)
	// reply
	switch x := m.Reply.(type) {
	case *ReplyJRJ_Funds:
		b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Funds); err != nil {
			return err
		}
	case *ReplyJRJ_Fund:
		b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Fund); err != nil {
			return err
		}
	case *ReplyJRJ_FundValue:
		b.EncodeVarint(102<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FundValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ReplyJRJ.Reply has unexpected type %T", x)
	}
	return nil
}

func _ReplyJRJ_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReplyJRJ)
	switch tag {
	case 100: // reply.funds
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JRJFunds)
		err := b.DecodeMessage(msg)
		m.Reply = &ReplyJRJ_Funds{msg}
		return true, err
	case 101: // reply.fund
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JRJFund)
		err := b.DecodeMessage(msg)
		m.Reply = &ReplyJRJ_Fund{msg}
		return true, err
	case 102: // reply.fundValue
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JRJFundValue)
		err := b.DecodeMessage(msg)
		m.Reply = &ReplyJRJ_FundValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ReplyJRJ_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReplyJRJ)
	// reply
	switch x := m.Reply.(type) {
	case *ReplyJRJ_Funds:
		s := proto.Size(x.Funds)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ReplyJRJ_Fund:
		s := proto.Size(x.Fund)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ReplyJRJ_FundValue:
		s := proto.Size(x.FundValue)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*JRJFundSize)(nil), "jarviscrawlercore.JRJFundSize")
	proto.RegisterType((*JRJFundManager)(nil), "jarviscrawlercore.JRJFundManager")
	proto.RegisterType((*JRJFundDayValue)(nil), "jarviscrawlercore.JRJFundDayValue")
	proto.RegisterType((*JRJFundValue)(nil), "jarviscrawlercore.JRJFundValue")
	proto.RegisterType((*JRJFund)(nil), "jarviscrawlercore.JRJFund")
	proto.RegisterType((*JRJFunds)(nil), "jarviscrawlercore.JRJFunds")
	proto.RegisterType((*RequestJRJ)(nil), "jarviscrawlercore.RequestJRJ")
	proto.RegisterType((*ReplyJRJ)(nil), "jarviscrawlercore.ReplyJRJ")
	proto.RegisterEnum("jarviscrawlercore.JRJMode", JRJMode_name, JRJMode_value)
}

func init() { proto.RegisterFile("jrj.proto", fileDescriptor_jrj_5707464fb197bfbd) }

var fileDescriptor_jrj_5707464fb197bfbd = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x8d, 0xe3, 0x38, 0x89, 0x27, 0xbf, 0x5f, 0x09, 0xa3, 0xaa, 0x5a, 0x05, 0x54, 0x8c, 0xc5,
	0xc1, 0xe2, 0x10, 0xa1, 0x54, 0x5c, 0x90, 0x10, 0x6a, 0xd5, 0x96, 0x2a, 0xa2, 0x3d, 0x6c, 0xff,
	0x48, 0x88, 0x03, 0xda, 0xc6, 0xdb, 0xe2, 0x2a, 0xb1, 0xc3, 0xda, 0x2e, 0xa4, 0x9f, 0x95, 0x6f,
	0xc0, 0x8d, 0x03, 0x67, 0xb4, 0xb3, 0x6b, 0x27, 0x2a, 0xcd, 0x85, 0xdb, 0xcc, 0xbc, 0x79, 0xb3,
	0xfb, 0x46, 0x6f, 0x17, 0xfc, 0x1b, 0x75, 0x33, 0x9c, 0xab, 0xac, 0xc8, 0xf0, 0xf1, 0x8d, 0x50,
	0xb7, 0x49, 0x3e, 0x51, 0xe2, 0xdb, 0x54, 0xaa, 0x49, 0xa6, 0x64, 0xf8, 0x1a, 0x7a, 0x63, 0x3e,
	0x3e, 0x2c, 0xd3, 0xf8, 0x34, 0xb9, 0x93, 0x88, 0xd0, 0xca, 0x93, 0x3b, 0xc9, 0x9c, 0xc0, 0x89,
	0x9a, 0x9c, 0x62, 0x5d, 0x2b, 0x92, 0x99, 0x64, 0xcd, 0xc0, 0x89, 0x5c, 0x4e, 0x71, 0xf8, 0xc3,
	0x81, 0x0d, 0xcb, 0x3b, 0x16, 0xa9, 0xb8, 0x96, 0x4a, 0xb7, 0xa5, 0x62, 0x66, 0xa8, 0x3e, 0xa7,
	0x18, 0x9f, 0x82, 0x9f, 0x17, 0x42, 0x15, 0x67, 0x4b, 0xfe, 0xb2, 0x80, 0x0c, 0x3a, 0x32, 0x8d,
	0x09, 0x73, 0x09, 0xab, 0x52, 0xcd, 0xbb, 0x4c, 0x54, 0xf1, 0xe5, 0xa3, 0x14, 0x8a, 0xb5, 0x0c,
	0xaf, 0x2e, 0x60, 0x1f, 0xdc, 0x5c, 0x7e, 0x67, 0x5e, 0xe0, 0x44, 0x5d, 0xae, 0x43, 0xdd, 0x2f,
	0xe3, 0x72, 0x22, 0x8a, 0x24, 0x4b, 0x59, 0x9b, 0x2e, 0xb0, 0x2c, 0xe8, 0x73, 0x26, 0x59, 0x99,
	0x16, 0x6a, 0xc1, 0x3a, 0x84, 0x55, 0x29, 0x6e, 0x41, 0x5b, 0xc9, 0xbc, 0x9c, 0x49, 0xd6, 0x25,
	0xc0, 0x66, 0xe1, 0x27, 0x78, 0x64, 0xd5, 0xed, 0x8b, 0xc5, 0x85, 0x98, 0x96, 0xb4, 0x85, 0x58,
	0x14, 0xb5, 0x3c, 0x1d, 0xe3, 0x26, 0x78, 0xb7, 0x1a, 0x24, 0x69, 0x1e, 0x37, 0x09, 0x6e, 0x03,
	0x14, 0x59, 0x21, 0xa6, 0xc4, 0x23, 0x65, 0x1e, 0x5f, 0xa9, 0x84, 0xbf, 0x1d, 0xf8, 0xcf, 0x4e,
	0xaf, 0x47, 0x4f, 0xb2, 0xb8, 0x1e, 0xad, 0x63, 0x64, 0xcb, 0xd1, 0x6e, 0xd4, 0xdc, 0x6b, 0x32,
	0xa7, 0x1a, 0x1f, 0xde, 0x1b, 0x5f, 0xc1, 0x2b, 0x55, 0x1c, 0x40, 0x3b, 0x31, 0x78, 0x2b, 0x70,
	0x23, 0x8f, 0x70, 0x5b, 0xc1, 0x17, 0xd0, 0x4b, 0xce, 0x96, 0x03, 0xbc, 0xba, 0x61, 0xb5, 0x8c,
	0x5b, 0x56, 0x6e, 0x3b, 0x70, 0x23, 0x9f, 0x60, 0x23, 0xf9, 0x0d, 0xb4, 0xe9, 0x1a, 0x39, 0xeb,
	0x04, 0x6e, 0xd4, 0x1b, 0x85, 0xc3, 0xbf, 0x3c, 0x35, 0xbc, 0xb7, 0x3a, 0x6e, 0x19, 0xe1, 0x2f,
	0x07, 0x3a, 0x16, 0x7b, 0x50, 0x73, 0xe5, 0xa0, 0xe6, 0x8a, 0x83, 0xb4, 0xf9, 0xc4, 0x75, 0x4e,
	0x3a, 0x7d, 0x4e, 0xb1, 0x5e, 0xf0, 0x44, 0x49, 0x51, 0x48, 0xb2, 0x8e, 0xb1, 0xc7, 0x4a, 0x05,
	0x47, 0xd6, 0xc4, 0x1e, 0xdd, 0x70, 0x7b, 0xfd, 0x0d, 0xb5, 0xe5, 0xad, 0xc9, 0xc9, 0x23, 0xb3,
	0xb9, 0x48, 0x17, 0xd6, 0x3f, 0x55, 0x8a, 0x6f, 0xa1, 0x3b, 0x33, 0x16, 0xaf, 0x34, 0x3f, 0x5f,
	0x3f, 0xd1, 0x3e, 0x06, 0x5e, 0x53, 0xc2, 0x00, 0xba, 0x16, 0xcb, 0xb5, 0x5f, 0xb4, 0xd0, 0x9c,
	0x39, 0xa4, 0xc6, 0x24, 0xe1, 0x14, 0x80, 0xcb, 0xaf, 0xa5, 0xcc, 0x8b, 0x31, 0x1f, 0xe3, 0x10,
	0x5a, 0xb3, 0x6a, 0x31, 0x1b, 0xa3, 0xc1, 0xc3, 0x47, 0x1d, 0x67, 0xb1, 0xe4, 0xd4, 0x87, 0x03,
	0xe8, 0x5e, 0x95, 0x69, 0x4c, 0xcb, 0x34, 0x8b, 0xab, 0x73, 0xbd, 0xbc, 0x85, 0x7e, 0x41, 0xae,
	0x59, 0xa8, 0x8e, 0xc3, 0x9f, 0x0e, 0x74, 0xb9, 0x9c, 0x4f, 0x17, 0xff, 0x72, 0xd8, 0x0e, 0x78,
	0x7a, 0x78, 0xce, 0xe2, 0xc0, 0x89, 0x7a, 0xa3, 0x27, 0xeb, 0x17, 0x91, 0x1f, 0x35, 0xb8, 0xe9,
	0xc5, 0x57, 0xd0, 0xd2, 0x01, 0x93, 0xc4, 0x19, 0xac, 0xe7, 0x1c, 0x35, 0x38, 0x75, 0xe2, 0x3b,
	0xf0, 0xaf, 0xaa, 0xd7, 0xc1, 0xae, 0x88, 0xf6, 0x6c, 0x3d, 0x8d, 0xda, 0x8e, 0x1a, 0x7c, 0xc9,
	0xd9, 0xeb, 0x80, 0xa7, 0xb4, 0xc6, 0x97, 0x9c, 0x1c, 0xa7, 0x15, 0xe0, 0x06, 0x80, 0x0e, 0x3f,
	0x1f, 0x9e, 0x9f, 0xec, 0x9f, 0xf6, 0x1b, 0xf8, 0x3f, 0xf8, 0x75, 0xde, 0x77, 0x10, 0xe9, 0x43,
	0x33, 0xe9, 0xc5, 0xee, 0x87, 0xf3, 0x83, 0x7e, 0x13, 0x37, 0xa1, 0x5f, 0xd7, 0x8e, 0x77, 0x4f,
	0x76, 0xdf, 0x1f, 0xf0, 0xbe, 0x7b, 0xd9, 0xa6, 0xcf, 0x74, 0xe7, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8a, 0xa2, 0x7c, 0xad, 0x59, 0x05, 0x00, 0x00,
}