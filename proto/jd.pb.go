// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jd.proto

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

// JDMode - jd mode
type JDMode int32

const (
	// JDM_PRODUCT - product
	JDMode_JDM_PRODUCT JDMode = 0
	// JDM_ACTIVE - active
	JDMode_JDM_ACTIVE JDMode = 1
	// JDM_ACTIVEPAGE - active page
	JDMode_JDM_ACTIVEPAGE JDMode = 2
)

var JDMode_name = map[int32]string{
	0: "JDM_PRODUCT",
	1: "JDM_ACTIVE",
	2: "JDM_ACTIVEPAGE",
}
var JDMode_value = map[string]int32{
	"JDM_PRODUCT":    0,
	"JDM_ACTIVE":     1,
	"JDM_ACTIVEPAGE": 2,
}

func (x JDMode) String() string {
	return proto.EnumName(JDMode_name, int32(x))
}
func (JDMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_jd_0eac6b16f93b35be, []int{0}
}

// JDPingou - jd pingou
type JDPingou struct {
	PreOrders            int32    `protobuf:"varint,1,opt,name=preOrders,proto3" json:"preOrders,omitempty"`
	StrLastTime          string   `protobuf:"bytes,2,opt,name=strLastTime,proto3" json:"strLastTime,omitempty"`
	ScheduledPrice       float32  `protobuf:"fixed32,3,opt,name=scheduledPrice,proto3" json:"scheduledPrice,omitempty"`
	Price                float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JDPingou) Reset()         { *m = JDPingou{} }
func (m *JDPingou) String() string { return proto.CompactTextString(m) }
func (*JDPingou) ProtoMessage()    {}
func (*JDPingou) Descriptor() ([]byte, []int) {
	return fileDescriptor_jd_0eac6b16f93b35be, []int{0}
}
func (m *JDPingou) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JDPingou.Unmarshal(m, b)
}
func (m *JDPingou) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JDPingou.Marshal(b, m, deterministic)
}
func (dst *JDPingou) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JDPingou.Merge(dst, src)
}
func (m *JDPingou) XXX_Size() int {
	return xxx_messageInfo_JDPingou.Size(m)
}
func (m *JDPingou) XXX_DiscardUnknown() {
	xxx_messageInfo_JDPingou.DiscardUnknown(m)
}

var xxx_messageInfo_JDPingou proto.InternalMessageInfo

func (m *JDPingou) GetPreOrders() int32 {
	if m != nil {
		return m.PreOrders
	}
	return 0
}

func (m *JDPingou) GetStrLastTime() string {
	if m != nil {
		return m.StrLastTime
	}
	return ""
}

func (m *JDPingou) GetScheduledPrice() float32 {
	if m != nil {
		return m.ScheduledPrice
	}
	return 0
}

func (m *JDPingou) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

// JDShangou - jd shangou
type JDShangou struct {
	OldPrice             float32  `protobuf:"fixed32,1,opt,name=oldPrice,proto3" json:"oldPrice,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	StrLastTime          string   `protobuf:"bytes,3,opt,name=strLastTime,proto3" json:"strLastTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JDShangou) Reset()         { *m = JDShangou{} }
func (m *JDShangou) String() string { return proto.CompactTextString(m) }
func (*JDShangou) ProtoMessage()    {}
func (*JDShangou) Descriptor() ([]byte, []int) {
	return fileDescriptor_jd_0eac6b16f93b35be, []int{1}
}
func (m *JDShangou) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JDShangou.Unmarshal(m, b)
}
func (m *JDShangou) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JDShangou.Marshal(b, m, deterministic)
}
func (dst *JDShangou) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JDShangou.Merge(dst, src)
}
func (m *JDShangou) XXX_Size() int {
	return xxx_messageInfo_JDShangou.Size(m)
}
func (m *JDShangou) XXX_DiscardUnknown() {
	xxx_messageInfo_JDShangou.DiscardUnknown(m)
}

var xxx_messageInfo_JDShangou proto.InternalMessageInfo

func (m *JDShangou) GetOldPrice() float32 {
	if m != nil {
		return m.OldPrice
	}
	return 0
}

func (m *JDShangou) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *JDShangou) GetStrLastTime() string {
	if m != nil {
		return m.StrLastTime
	}
	return ""
}

// JDPromotional - jd promotional
type JDPromotional struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JDPromotional) Reset()         { *m = JDPromotional{} }
func (m *JDPromotional) String() string { return proto.CompactTextString(m) }
func (*JDPromotional) ProtoMessage()    {}
func (*JDPromotional) Descriptor() ([]byte, []int) {
	return fileDescriptor_jd_0eac6b16f93b35be, []int{2}
}
func (m *JDPromotional) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JDPromotional.Unmarshal(m, b)
}
func (m *JDPromotional) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JDPromotional.Marshal(b, m, deterministic)
}
func (dst *JDPromotional) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JDPromotional.Merge(dst, src)
}
func (m *JDPromotional) XXX_Size() int {
	return xxx_messageInfo_JDPromotional.Size(m)
}
func (m *JDPromotional) XXX_DiscardUnknown() {
	xxx_messageInfo_JDPromotional.DiscardUnknown(m)
}

var xxx_messageInfo_JDPromotional proto.InternalMessageInfo

func (m *JDPromotional) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *JDPromotional) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

// JDNormalPrice - jd normal price
type JDNormalPrice struct {
	OldPrice             float32          `protobuf:"fixed32,1,opt,name=oldPrice,proto3" json:"oldPrice,omitempty"`
	Price                float32          `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Promotionals         []*JDPromotional `protobuf:"bytes,3,rep,name=promotionals,proto3" json:"promotionals,omitempty"`
	Coupons              []string         `protobuf:"bytes,4,rep,name=coupons,proto3" json:"coupons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *JDNormalPrice) Reset()         { *m = JDNormalPrice{} }
func (m *JDNormalPrice) String() string { return proto.CompactTextString(m) }
func (*JDNormalPrice) ProtoMessage()    {}
func (*JDNormalPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_jd_0eac6b16f93b35be, []int{3}
}
func (m *JDNormalPrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JDNormalPrice.Unmarshal(m, b)
}
func (m *JDNormalPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JDNormalPrice.Marshal(b, m, deterministic)
}
func (dst *JDNormalPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JDNormalPrice.Merge(dst, src)
}
func (m *JDNormalPrice) XXX_Size() int {
	return xxx_messageInfo_JDNormalPrice.Size(m)
}
func (m *JDNormalPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_JDNormalPrice.DiscardUnknown(m)
}

var xxx_messageInfo_JDNormalPrice proto.InternalMessageInfo

func (m *JDNormalPrice) GetOldPrice() float32 {
	if m != nil {
		return m.OldPrice
	}
	return 0
}

func (m *JDNormalPrice) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *JDNormalPrice) GetPromotionals() []*JDPromotional {
	if m != nil {
		return m.Promotionals
	}
	return nil
}

func (m *JDNormalPrice) GetCoupons() []string {
	if m != nil {
		return m.Coupons
	}
	return nil
}

// JDSKUInfo - jd sku info
type JDSKUInfo struct {
	SkuID                string   `protobuf:"bytes,1,opt,name=skuID,proto3" json:"skuID,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Color                string   `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Series               string   `protobuf:"bytes,4,opt,name=series,proto3" json:"series,omitempty"`
	Variety              string   `protobuf:"bytes,5,opt,name=variety,proto3" json:"variety,omitempty"`
	Size                 string   `protobuf:"bytes,6,opt,name=size,proto3" json:"size,omitempty"`
	Model                string   `protobuf:"bytes,7,opt,name=model,proto3" json:"model,omitempty"`
	Purchase             string   `protobuf:"bytes,8,opt,name=purchase,proto3" json:"purchase,omitempty"`
	Disabled             bool     `protobuf:"varint,9,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Selected             bool     `protobuf:"varint,10,opt,name=selected,proto3" json:"selected,omitempty"`
	Category             string   `protobuf:"bytes,11,opt,name=category,proto3" json:"category,omitempty"`
	ProductType          string   `protobuf:"bytes,12,opt,name=productType,proto3" json:"productType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JDSKUInfo) Reset()         { *m = JDSKUInfo{} }
func (m *JDSKUInfo) String() string { return proto.CompactTextString(m) }
func (*JDSKUInfo) ProtoMessage()    {}
func (*JDSKUInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_jd_0eac6b16f93b35be, []int{4}
}
func (m *JDSKUInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JDSKUInfo.Unmarshal(m, b)
}
func (m *JDSKUInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JDSKUInfo.Marshal(b, m, deterministic)
}
func (dst *JDSKUInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JDSKUInfo.Merge(dst, src)
}
func (m *JDSKUInfo) XXX_Size() int {
	return xxx_messageInfo_JDSKUInfo.Size(m)
}
func (m *JDSKUInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_JDSKUInfo.DiscardUnknown(m)
}

var xxx_messageInfo_JDSKUInfo proto.InternalMessageInfo

func (m *JDSKUInfo) GetSkuID() string {
	if m != nil {
		return m.SkuID
	}
	return ""
}

func (m *JDSKUInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *JDSKUInfo) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *JDSKUInfo) GetSeries() string {
	if m != nil {
		return m.Series
	}
	return ""
}

func (m *JDSKUInfo) GetVariety() string {
	if m != nil {
		return m.Variety
	}
	return ""
}

func (m *JDSKUInfo) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *JDSKUInfo) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *JDSKUInfo) GetPurchase() string {
	if m != nil {
		return m.Purchase
	}
	return ""
}

func (m *JDSKUInfo) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *JDSKUInfo) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

func (m *JDSKUInfo) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *JDSKUInfo) GetProductType() string {
	if m != nil {
		return m.ProductType
	}
	return ""
}

// JDCommentType - jd comments type
type JDCommentsType struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Nums                 int32    `protobuf:"varint,2,opt,name=nums,proto3" json:"nums,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JDCommentsType) Reset()         { *m = JDCommentsType{} }
func (m *JDCommentsType) String() string { return proto.CompactTextString(m) }
func (*JDCommentsType) ProtoMessage()    {}
func (*JDCommentsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_jd_0eac6b16f93b35be, []int{5}
}
func (m *JDCommentsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JDCommentsType.Unmarshal(m, b)
}
func (m *JDCommentsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JDCommentsType.Marshal(b, m, deterministic)
}
func (dst *JDCommentsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JDCommentsType.Merge(dst, src)
}
func (m *JDCommentsType) XXX_Size() int {
	return xxx_messageInfo_JDCommentsType.Size(m)
}
func (m *JDCommentsType) XXX_DiscardUnknown() {
	xxx_messageInfo_JDCommentsType.DiscardUnknown(m)
}

var xxx_messageInfo_JDCommentsType proto.InternalMessageInfo

func (m *JDCommentsType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *JDCommentsType) GetNums() int32 {
	if m != nil {
		return m.Nums
	}
	return 0
}

// JDCommentsInfo - jd comments infomation
type JDCommentsInfo struct {
	Percent              float32           `protobuf:"fixed32,1,opt,name=percent,proto3" json:"percent,omitempty"`
	Tags                 []string          `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	Lst                  []*JDCommentsType `protobuf:"bytes,3,rep,name=lst,proto3" json:"lst,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JDCommentsInfo) Reset()         { *m = JDCommentsInfo{} }
func (m *JDCommentsInfo) String() string { return proto.CompactTextString(m) }
func (*JDCommentsInfo) ProtoMessage()    {}
func (*JDCommentsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_jd_0eac6b16f93b35be, []int{6}
}
func (m *JDCommentsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JDCommentsInfo.Unmarshal(m, b)
}
func (m *JDCommentsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JDCommentsInfo.Marshal(b, m, deterministic)
}
func (dst *JDCommentsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JDCommentsInfo.Merge(dst, src)
}
func (m *JDCommentsInfo) XXX_Size() int {
	return xxx_messageInfo_JDCommentsInfo.Size(m)
}
func (m *JDCommentsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_JDCommentsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_JDCommentsInfo proto.InternalMessageInfo

func (m *JDCommentsInfo) GetPercent() float32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *JDCommentsInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *JDCommentsInfo) GetLst() []*JDCommentsType {
	if m != nil {
		return m.Lst
	}
	return nil
}

// JDActive - jd active
type JDActive struct {
	UrlActive            []string `protobuf:"bytes,1,rep,name=urlActive,proto3" json:"urlActive,omitempty"`
	UrlProduct           []string `protobuf:"bytes,2,rep,name=urlProduct,proto3" json:"urlProduct,omitempty"`
	LastUpdatedTime      int64    `protobuf:"varint,3,opt,name=lastUpdatedTime,proto3" json:"lastUpdatedTime,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JDActive) Reset()         { *m = JDActive{} }
func (m *JDActive) String() string { return proto.CompactTextString(m) }
func (*JDActive) ProtoMessage()    {}
func (*JDActive) Descriptor() ([]byte, []int) {
	return fileDescriptor_jd_0eac6b16f93b35be, []int{7}
}
func (m *JDActive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JDActive.Unmarshal(m, b)
}
func (m *JDActive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JDActive.Marshal(b, m, deterministic)
}
func (dst *JDActive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JDActive.Merge(dst, src)
}
func (m *JDActive) XXX_Size() int {
	return xxx_messageInfo_JDActive.Size(m)
}
func (m *JDActive) XXX_DiscardUnknown() {
	xxx_messageInfo_JDActive.DiscardUnknown(m)
}

var xxx_messageInfo_JDActive proto.InternalMessageInfo

func (m *JDActive) GetUrlActive() []string {
	if m != nil {
		return m.UrlActive
	}
	return nil
}

func (m *JDActive) GetUrlProduct() []string {
	if m != nil {
		return m.UrlProduct
	}
	return nil
}

func (m *JDActive) GetLastUpdatedTime() int64 {
	if m != nil {
		return m.LastUpdatedTime
	}
	return 0
}

func (m *JDActive) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *JDActive) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

// JDProduct - jd product
type JDProduct struct {
	Url                  string          `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BreadCrumbs          []string        `protobuf:"bytes,3,rep,name=breadCrumbs,proto3" json:"breadCrumbs,omitempty"`
	Info                 string          `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	NameTag              string          `protobuf:"bytes,5,opt,name=nameTag,proto3" json:"nameTag,omitempty"`
	Pingou               *JDPingou       `protobuf:"bytes,6,opt,name=pingou,proto3" json:"pingou,omitempty"`
	SummaryService       string          `protobuf:"bytes,7,opt,name=summaryService,proto3" json:"summaryService,omitempty"`
	StrShipTime          string          `protobuf:"bytes,8,opt,name=strShipTime,proto3" json:"strShipTime,omitempty"`
	StrWeight            string          `protobuf:"bytes,9,opt,name=strWeight,proto3" json:"strWeight,omitempty"`
	BrandChs             string          `protobuf:"bytes,10,opt,name=brandChs,proto3" json:"brandChs,omitempty"`
	BrandEng             string          `protobuf:"bytes,11,opt,name=brandEng,proto3" json:"brandEng,omitempty"`
	SKUs                 []*JDSKUInfo    `protobuf:"bytes,12,rep,name=SKUs,proto3" json:"SKUs,omitempty"`
	Comment              *JDCommentsInfo `protobuf:"bytes,13,opt,name=comment,proto3" json:"comment,omitempty"`
	LastUpdatedTime      int64           `protobuf:"varint,14,opt,name=lastUpdatedTime,proto3" json:"lastUpdatedTime,omitempty"`
	Price                *JDNormalPrice  `protobuf:"bytes,15,opt,name=price,proto3" json:"price,omitempty"`
	Shangou              *JDShangou      `protobuf:"bytes,16,opt,name=shangou,proto3" json:"shangou,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *JDProduct) Reset()         { *m = JDProduct{} }
func (m *JDProduct) String() string { return proto.CompactTextString(m) }
func (*JDProduct) ProtoMessage()    {}
func (*JDProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_jd_0eac6b16f93b35be, []int{8}
}
func (m *JDProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JDProduct.Unmarshal(m, b)
}
func (m *JDProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JDProduct.Marshal(b, m, deterministic)
}
func (dst *JDProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JDProduct.Merge(dst, src)
}
func (m *JDProduct) XXX_Size() int {
	return xxx_messageInfo_JDProduct.Size(m)
}
func (m *JDProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_JDProduct.DiscardUnknown(m)
}

var xxx_messageInfo_JDProduct proto.InternalMessageInfo

func (m *JDProduct) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *JDProduct) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JDProduct) GetBreadCrumbs() []string {
	if m != nil {
		return m.BreadCrumbs
	}
	return nil
}

func (m *JDProduct) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *JDProduct) GetNameTag() string {
	if m != nil {
		return m.NameTag
	}
	return ""
}

func (m *JDProduct) GetPingou() *JDPingou {
	if m != nil {
		return m.Pingou
	}
	return nil
}

func (m *JDProduct) GetSummaryService() string {
	if m != nil {
		return m.SummaryService
	}
	return ""
}

func (m *JDProduct) GetStrShipTime() string {
	if m != nil {
		return m.StrShipTime
	}
	return ""
}

func (m *JDProduct) GetStrWeight() string {
	if m != nil {
		return m.StrWeight
	}
	return ""
}

func (m *JDProduct) GetBrandChs() string {
	if m != nil {
		return m.BrandChs
	}
	return ""
}

func (m *JDProduct) GetBrandEng() string {
	if m != nil {
		return m.BrandEng
	}
	return ""
}

func (m *JDProduct) GetSKUs() []*JDSKUInfo {
	if m != nil {
		return m.SKUs
	}
	return nil
}

func (m *JDProduct) GetComment() *JDCommentsInfo {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *JDProduct) GetLastUpdatedTime() int64 {
	if m != nil {
		return m.LastUpdatedTime
	}
	return 0
}

func (m *JDProduct) GetPrice() *JDNormalPrice {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *JDProduct) GetShangou() *JDShangou {
	if m != nil {
		return m.Shangou
	}
	return nil
}

// RequestJD - request jd
type RequestJD struct {
	Mode                 JDMode   `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.JDMode" json:"mode,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestJD) Reset()         { *m = RequestJD{} }
func (m *RequestJD) String() string { return proto.CompactTextString(m) }
func (*RequestJD) ProtoMessage()    {}
func (*RequestJD) Descriptor() ([]byte, []int) {
	return fileDescriptor_jd_0eac6b16f93b35be, []int{9}
}
func (m *RequestJD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestJD.Unmarshal(m, b)
}
func (m *RequestJD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestJD.Marshal(b, m, deterministic)
}
func (dst *RequestJD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestJD.Merge(dst, src)
}
func (m *RequestJD) XXX_Size() int {
	return xxx_messageInfo_RequestJD.Size(m)
}
func (m *RequestJD) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestJD.DiscardUnknown(m)
}

var xxx_messageInfo_RequestJD proto.InternalMessageInfo

func (m *RequestJD) GetMode() JDMode {
	if m != nil {
		return m.Mode
	}
	return JDMode_JDM_PRODUCT
}

func (m *RequestJD) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// ReplyJD - reply jd
type ReplyJD struct {
	Mode JDMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.JDMode" json:"mode,omitempty"`
	// Types that are valid to be assigned to Reply:
	//	*ReplyJD_Product
	//	*ReplyJD_Active
	Reply                isReplyJD_Reply `protobuf_oneof:"reply"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReplyJD) Reset()         { *m = ReplyJD{} }
func (m *ReplyJD) String() string { return proto.CompactTextString(m) }
func (*ReplyJD) ProtoMessage()    {}
func (*ReplyJD) Descriptor() ([]byte, []int) {
	return fileDescriptor_jd_0eac6b16f93b35be, []int{10}
}
func (m *ReplyJD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyJD.Unmarshal(m, b)
}
func (m *ReplyJD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyJD.Marshal(b, m, deterministic)
}
func (dst *ReplyJD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyJD.Merge(dst, src)
}
func (m *ReplyJD) XXX_Size() int {
	return xxx_messageInfo_ReplyJD.Size(m)
}
func (m *ReplyJD) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyJD.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyJD proto.InternalMessageInfo

type isReplyJD_Reply interface {
	isReplyJD_Reply()
}

type ReplyJD_Product struct {
	Product *JDProduct `protobuf:"bytes,100,opt,name=product,proto3,oneof"`
}
type ReplyJD_Active struct {
	Active *JDActive `protobuf:"bytes,101,opt,name=active,proto3,oneof"`
}

func (*ReplyJD_Product) isReplyJD_Reply() {}
func (*ReplyJD_Active) isReplyJD_Reply()  {}

func (m *ReplyJD) GetReply() isReplyJD_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *ReplyJD) GetMode() JDMode {
	if m != nil {
		return m.Mode
	}
	return JDMode_JDM_PRODUCT
}

func (m *ReplyJD) GetProduct() *JDProduct {
	if x, ok := m.GetReply().(*ReplyJD_Product); ok {
		return x.Product
	}
	return nil
}

func (m *ReplyJD) GetActive() *JDActive {
	if x, ok := m.GetReply().(*ReplyJD_Active); ok {
		return x.Active
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReplyJD) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReplyJD_OneofMarshaler, _ReplyJD_OneofUnmarshaler, _ReplyJD_OneofSizer, []interface{}{
		(*ReplyJD_Product)(nil),
		(*ReplyJD_Active)(nil),
	}
}

func _ReplyJD_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReplyJD)
	// reply
	switch x := m.Reply.(type) {
	case *ReplyJD_Product:
		b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Product); err != nil {
			return err
		}
	case *ReplyJD_Active:
		b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Active); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ReplyJD.Reply has unexpected type %T", x)
	}
	return nil
}

func _ReplyJD_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReplyJD)
	switch tag {
	case 100: // reply.product
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JDProduct)
		err := b.DecodeMessage(msg)
		m.Reply = &ReplyJD_Product{msg}
		return true, err
	case 101: // reply.active
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JDActive)
		err := b.DecodeMessage(msg)
		m.Reply = &ReplyJD_Active{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ReplyJD_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReplyJD)
	// reply
	switch x := m.Reply.(type) {
	case *ReplyJD_Product:
		s := proto.Size(x.Product)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ReplyJD_Active:
		s := proto.Size(x.Active)
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
	proto.RegisterType((*JDPingou)(nil), "jarviscrawlercore.JDPingou")
	proto.RegisterType((*JDShangou)(nil), "jarviscrawlercore.JDShangou")
	proto.RegisterType((*JDPromotional)(nil), "jarviscrawlercore.JDPromotional")
	proto.RegisterType((*JDNormalPrice)(nil), "jarviscrawlercore.JDNormalPrice")
	proto.RegisterType((*JDSKUInfo)(nil), "jarviscrawlercore.JDSKUInfo")
	proto.RegisterType((*JDCommentsType)(nil), "jarviscrawlercore.JDCommentsType")
	proto.RegisterType((*JDCommentsInfo)(nil), "jarviscrawlercore.JDCommentsInfo")
	proto.RegisterType((*JDActive)(nil), "jarviscrawlercore.JDActive")
	proto.RegisterType((*JDProduct)(nil), "jarviscrawlercore.JDProduct")
	proto.RegisterType((*RequestJD)(nil), "jarviscrawlercore.RequestJD")
	proto.RegisterType((*ReplyJD)(nil), "jarviscrawlercore.ReplyJD")
	proto.RegisterEnum("jarviscrawlercore.JDMode", JDMode_name, JDMode_value)
}

func init() { proto.RegisterFile("jd.proto", fileDescriptor_jd_0eac6b16f93b35be) }

var fileDescriptor_jd_0eac6b16f93b35be = []byte{
	// 905 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x5e, 0xc7, 0xf9, 0xf3, 0x64, 0x9b, 0x0d, 0x23, 0x84, 0x06, 0xa8, 0x50, 0xf0, 0x05, 0x8a,
	0x90, 0x58, 0xa1, 0x5d, 0x51, 0x15, 0x21, 0x2e, 0x96, 0x78, 0x45, 0x37, 0x6d, 0x69, 0x34, 0x9b,
	0x85, 0xcb, 0x6a, 0x62, 0x4f, 0x13, 0x17, 0xdb, 0x63, 0x66, 0xc6, 0x8b, 0xc2, 0x13, 0xf0, 0x0c,
	0xdc, 0xf1, 0x0a, 0xdc, 0xf1, 0x34, 0xbc, 0x0a, 0x9a, 0x33, 0x1e, 0xc7, 0xdd, 0xa6, 0x2b, 0xf5,
	0x6e, 0xbe, 0xf3, 0xe7, 0x33, 0x67, 0xbe, 0xf3, 0xc9, 0x68, 0xf8, 0x3a, 0x39, 0x2d, 0xa5, 0xd0,
	0x02, 0x7f, 0xf0, 0x9a, 0xc9, 0xdb, 0x54, 0xc5, 0x92, 0xfd, 0x9e, 0x71, 0x19, 0x0b, 0xc9, 0xc3,
	0x3f, 0x3d, 0x34, 0x5c, 0x44, 0xcb, 0xb4, 0xd8, 0x88, 0x0a, 0x3f, 0x44, 0x41, 0x29, 0xf9, 0x0b,
	0x99, 0x70, 0xa9, 0x88, 0x37, 0xf5, 0x66, 0x3d, 0xba, 0x37, 0xe0, 0x29, 0x1a, 0x29, 0x2d, 0x9f,
	0x31, 0xa5, 0x57, 0x69, 0xce, 0x49, 0x67, 0xea, 0xcd, 0x02, 0xda, 0x36, 0xe1, 0x2f, 0xd0, 0x58,
	0xc5, 0x5b, 0x9e, 0x54, 0x19, 0x4f, 0x96, 0x32, 0x8d, 0x39, 0xf1, 0xa7, 0xde, 0xac, 0x43, 0xef,
	0x58, 0xf1, 0x87, 0xa8, 0x57, 0x82, 0xbb, 0x0b, 0x6e, 0x0b, 0xc2, 0x97, 0x28, 0x58, 0x44, 0xd7,
	0x5b, 0x06, 0xad, 0x7c, 0x82, 0x86, 0x22, 0xab, 0x8b, 0x78, 0x10, 0xd5, 0xe0, 0x7d, 0x7a, 0xa7,
	0x95, 0x7e, 0xb7, 0x3d, 0xff, 0xad, 0xf6, 0xc2, 0x6f, 0xd1, 0x83, 0x45, 0xb4, 0x94, 0x22, 0x17,
	0x3a, 0x15, 0x05, 0xcb, 0x4c, 0x21, 0x9d, 0xea, 0xcc, 0x7e, 0x21, 0xa0, 0x16, 0x60, 0x8c, 0xba,
	0x69, 0xf1, 0x4a, 0xd4, 0x17, 0x84, 0x73, 0xf8, 0xb7, 0x67, 0x72, 0x7f, 0x12, 0x32, 0x67, 0x99,
	0x6d, 0xe2, 0xfd, 0x1b, 0x8c, 0xd0, 0x71, 0xb9, 0xff, 0xb8, 0x22, 0xfe, 0xd4, 0x9f, 0x8d, 0xce,
	0xa6, 0xa7, 0x6f, 0x3d, 0xca, 0xe9, 0x1b, 0x5d, 0xd2, 0x37, 0xb2, 0x30, 0x41, 0x83, 0x58, 0x54,
	0xa5, 0x28, 0x14, 0xe9, 0x4e, 0xfd, 0x59, 0x40, 0x1d, 0x0c, 0xff, 0xed, 0xc0, 0x00, 0x9f, 0xde,
	0x5c, 0x15, 0xaf, 0x84, 0xe9, 0x41, 0xfd, 0x5a, 0x5d, 0x45, 0xee, 0x6e, 0x00, 0xcc, 0xdd, 0xf4,
	0xae, 0x74, 0x8f, 0x07, 0x67, 0x13, 0x19, 0x8b, 0x4c, 0xc8, 0x7a, 0x64, 0x16, 0xe0, 0x8f, 0x50,
	0x5f, 0x71, 0x99, 0x72, 0x05, 0x8f, 0x14, 0xd0, 0x1a, 0x99, 0xef, 0xdf, 0x32, 0x99, 0x72, 0xbd,
	0x23, 0x3d, 0x70, 0x38, 0x68, 0x6a, 0xab, 0xf4, 0x0f, 0x4e, 0xfa, 0xb6, 0xb6, 0x39, 0x9b, 0xda,
	0xb9, 0x48, 0x78, 0x46, 0x06, 0xb6, 0x36, 0x00, 0x33, 0xbb, 0xb2, 0x92, 0xf1, 0x96, 0x29, 0x4e,
	0x86, 0xe0, 0x68, 0xb0, 0xf1, 0x25, 0xa9, 0x62, 0xeb, 0x8c, 0x27, 0x24, 0x98, 0x7a, 0xb3, 0x21,
	0x6d, 0xb0, 0xf1, 0x29, 0x9e, 0xf1, 0x58, 0xf3, 0x84, 0x20, 0xeb, 0x73, 0xd8, 0xf8, 0x62, 0xa6,
	0xf9, 0x46, 0xc8, 0x1d, 0x19, 0xd9, 0x9a, 0x0e, 0x1b, 0x6a, 0x94, 0x52, 0x24, 0x55, 0xac, 0x57,
	0xe6, 0xf2, 0xc7, 0x96, 0x1a, 0x2d, 0x53, 0xf8, 0x18, 0x8d, 0x17, 0xd1, 0x5c, 0xe4, 0x39, 0x2f,
	0xb4, 0x32, 0x96, 0x66, 0x52, 0x5e, 0x6b, 0x52, 0x18, 0x75, 0x8b, 0x2a, 0x57, 0x30, 0xbd, 0x1e,
	0x85, 0x73, 0xa8, 0xda, 0x99, 0x30, 0x79, 0x82, 0x06, 0x25, 0x97, 0x31, 0x2f, 0x74, 0x4d, 0x0c,
	0x07, 0xa1, 0x26, 0xdb, 0x98, 0x7c, 0x1f, 0x6a, 0xb2, 0x8d, 0xc2, 0xe7, 0xc8, 0xcf, 0x94, 0xae,
	0xc9, 0xf0, 0xf9, 0x41, 0x32, 0xb4, 0xfb, 0xa2, 0x26, 0x3a, 0xfc, 0x0b, 0xb6, 0xf6, 0x22, 0xd6,
	0xe9, 0x2d, 0x37, 0x5b, 0x5b, 0xc9, 0xcc, 0x02, 0xe2, 0x41, 0xe9, 0xbd, 0x01, 0x7f, 0x86, 0x50,
	0x25, 0xb3, 0xa5, 0xbd, 0x6b, 0xfd, 0xe5, 0x96, 0x05, 0xcf, 0xd0, 0x49, 0xc6, 0x94, 0xbe, 0x29,
	0x13, 0xa6, 0x79, 0xd2, 0xac, 0x8e, 0x4f, 0xef, 0x9a, 0xf1, 0x04, 0xf9, 0x95, 0xcc, 0x6a, 0x3a,
	0x98, 0xe3, 0x7e, 0x7f, 0x7a, 0xad, 0xfd, 0x09, 0xff, 0xeb, 0x1a, 0x1e, 0xba, 0xfa, 0x75, 0x96,
	0xb7, 0xcf, 0x32, 0x53, 0x64, 0x8d, 0x80, 0xc0, 0xd9, 0xbc, 0xd0, 0x5a, 0x72, 0x96, 0xcc, 0x65,
	0x95, 0xaf, 0xed, 0x6a, 0x04, 0xb4, 0x6d, 0x6a, 0xb6, 0xb2, 0xbb, 0xdf, 0x4a, 0x33, 0x69, 0x93,
	0xbd, 0x62, 0x1b, 0xc7, 0xc5, 0x1a, 0xe2, 0x73, 0xd4, 0x2f, 0x41, 0xd3, 0x80, 0x8d, 0xa3, 0xb3,
	0x4f, 0x0f, 0x6f, 0x19, 0x84, 0xd0, 0x3a, 0x14, 0xe4, 0xab, 0xca, 0x73, 0x26, 0x77, 0xd7, 0x5c,
	0xde, 0x9a, 0xfd, 0xb5, 0xac, 0xbd, 0x63, 0xad, 0x95, 0xe6, 0x7a, 0x9b, 0x96, 0x30, 0xae, 0x61,
	0xa3, 0x34, 0xce, 0x64, 0x9e, 0x44, 0x69, 0xf9, 0x0b, 0x4f, 0x37, 0x5b, 0x0d, 0x2c, 0x0e, 0xe8,
	0xde, 0x60, 0xa8, 0xba, 0x96, 0xac, 0x48, 0xe6, 0x5b, 0x05, 0x34, 0x0e, 0x68, 0x83, 0x1b, 0xdf,
	0x65, 0xb1, 0x71, 0x34, 0x76, 0x18, 0x7f, 0x8d, 0xba, 0xd7, 0x4f, 0x6f, 0x14, 0x39, 0x06, 0xae,
	0x3c, 0x3c, 0x78, 0xa5, 0x7a, 0xfd, 0x29, 0x44, 0xe2, 0xef, 0x8c, 0x58, 0x00, 0x79, 0xc8, 0x03,
	0x98, 0xc3, 0xfd, 0x04, 0x83, 0x4c, 0x97, 0x71, 0x88, 0x19, 0xe3, 0xc3, 0xcc, 0x78, 0xe4, 0xf4,
	0xee, 0x04, 0x3e, 0x72, 0x58, 0xd2, 0x5a, 0xe2, 0xe9, 0x14, 0xf1, 0x11, 0x1a, 0x28, 0xab, 0xf7,
	0x64, 0x02, 0x99, 0xef, 0xb8, 0x93, 0x8d, 0xa1, 0x2e, 0x38, 0x7c, 0x86, 0x02, 0xca, 0x7f, 0xab,
	0xb8, 0xd2, 0x8b, 0x08, 0x7f, 0x85, 0xba, 0x46, 0x55, 0x80, 0x61, 0xe3, 0xb3, 0x8f, 0x0f, 0x56,
	0x78, 0x2e, 0x12, 0x4e, 0x21, 0xcc, 0xf1, 0xb1, 0xd3, 0xf0, 0x31, 0xfc, 0xc7, 0x43, 0x03, 0xca,
	0xcb, 0x6c, 0xf7, 0xfe, 0xc5, 0x1e, 0xa3, 0x41, 0xad, 0x22, 0x24, 0xb9, 0xe7, 0x02, 0xf5, 0x2e,
	0x3c, 0x39, 0xa2, 0x2e, 0x1c, 0x7f, 0x83, 0xfa, 0xcc, 0x6e, 0x2c, 0xbf, 0x87, 0xa0, 0x76, 0x87,
	0x9f, 0x1c, 0xd1, 0x3a, 0xf8, 0x87, 0x01, 0xea, 0x49, 0xd3, 0xea, 0x97, 0xdf, 0xa3, 0xbe, 0xed,
	0x04, 0x9f, 0xa0, 0xd1, 0x22, 0x7a, 0xfe, 0x72, 0x49, 0x5f, 0x44, 0x37, 0xf3, 0xd5, 0xe4, 0x08,
	0x8f, 0x11, 0x32, 0x86, 0x8b, 0xf9, 0xea, 0xea, 0xe7, 0xcb, 0x89, 0x87, 0xb1, 0x51, 0x28, 0x87,
	0x97, 0x17, 0x3f, 0x5e, 0x4e, 0x3a, 0xeb, 0x3e, 0xfc, 0x10, 0x9c, 0xff, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0x4d, 0x3d, 0x73, 0x75, 0x1c, 0x08, 0x00, 0x00,
}