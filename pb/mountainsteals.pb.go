// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: mountainsteals.proto

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

// MountainStealsMode - mountain steals mode
type MountainStealsMode int32

const (
	// MSM_SALE - sale
	MountainStealsMode_MSM_SALE MountainStealsMode = 0
	// MSM_PRODUCT - product
	MountainStealsMode_MSM_PRODUCT MountainStealsMode = 1
)

// Enum value maps for MountainStealsMode.
var (
	MountainStealsMode_name = map[int32]string{
		0: "MSM_SALE",
		1: "MSM_PRODUCT",
	}
	MountainStealsMode_value = map[string]int32{
		"MSM_SALE":    0,
		"MSM_PRODUCT": 1,
	}
)

func (x MountainStealsMode) Enum() *MountainStealsMode {
	p := new(MountainStealsMode)
	*p = x
	return p
}

func (x MountainStealsMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MountainStealsMode) Descriptor() protoreflect.EnumDescriptor {
	return file_mountainsteals_proto_enumTypes[0].Descriptor()
}

func (MountainStealsMode) Type() protoreflect.EnumType {
	return &file_mountainsteals_proto_enumTypes[0]
}

func (x MountainStealsMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MountainStealsMode.Descriptor instead.
func (MountainStealsMode) EnumDescriptor() ([]byte, []int) {
	return file_mountainsteals_proto_rawDescGZIP(), []int{0}
}

// MountainStealsSale - mountain steals sale
type MountainStealsSale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products        []string `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Saleurl         string   `protobuf:"bytes,2,opt,name=saleurl,proto3" json:"saleurl,omitempty"`
	Code            string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	LastUpdatedTime int64    `protobuf:"varint,4,opt,name=lastUpdatedTime,proto3" json:"lastUpdatedTime,omitempty"`
	CreateTime      int64    `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
}

func (x *MountainStealsSale) Reset() {
	*x = MountainStealsSale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mountainsteals_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountainStealsSale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountainStealsSale) ProtoMessage() {}

func (x *MountainStealsSale) ProtoReflect() protoreflect.Message {
	mi := &file_mountainsteals_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountainStealsSale.ProtoReflect.Descriptor instead.
func (*MountainStealsSale) Descriptor() ([]byte, []int) {
	return file_mountainsteals_proto_rawDescGZIP(), []int{0}
}

func (x *MountainStealsSale) GetProducts() []string {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *MountainStealsSale) GetSaleurl() string {
	if x != nil {
		return x.Saleurl
	}
	return ""
}

func (x *MountainStealsSale) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MountainStealsSale) GetLastUpdatedTime() int64 {
	if x != nil {
		return x.LastUpdatedTime
	}
	return 0
}

func (x *MountainStealsSale) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

// MountainStealsColorSize - mountain steals color size
type MountainStealsColorSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color string  `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	Size  string  `protobuf:"bytes,2,opt,name=size,proto3" json:"size,omitempty"`
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Img   string  `protobuf:"bytes,4,opt,name=img,proto3" json:"img,omitempty"`
	Sku   string  `protobuf:"bytes,5,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *MountainStealsColorSize) Reset() {
	*x = MountainStealsColorSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mountainsteals_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountainStealsColorSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountainStealsColorSize) ProtoMessage() {}

func (x *MountainStealsColorSize) ProtoReflect() protoreflect.Message {
	mi := &file_mountainsteals_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountainStealsColorSize.ProtoReflect.Descriptor instead.
func (*MountainStealsColorSize) Descriptor() ([]byte, []int) {
	return file_mountainsteals_proto_rawDescGZIP(), []int{1}
}

func (x *MountainStealsColorSize) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *MountainStealsColorSize) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *MountainStealsColorSize) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *MountainStealsColorSize) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *MountainStealsColorSize) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

// MountainStealsHistory - MountainSteals history
type MountainStealsHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateTime int64                      `protobuf:"varint,1,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Price      float32                    `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	CurPrice   float32                    `protobuf:"fixed32,3,opt,name=curPrice,proto3" json:"curPrice,omitempty"`
	StockLevel int32                      `protobuf:"varint,4,opt,name=stockLevel,proto3" json:"stockLevel,omitempty"`
	SaleCode   string                     `protobuf:"bytes,5,opt,name=saleCode,proto3" json:"saleCode,omitempty"`
	Offers     []*MountainStealsColorSize `protobuf:"bytes,6,rep,name=offers,proto3" json:"offers,omitempty"`
}

func (x *MountainStealsHistory) Reset() {
	*x = MountainStealsHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mountainsteals_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountainStealsHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountainStealsHistory) ProtoMessage() {}

func (x *MountainStealsHistory) ProtoReflect() protoreflect.Message {
	mi := &file_mountainsteals_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountainStealsHistory.ProtoReflect.Descriptor instead.
func (*MountainStealsHistory) Descriptor() ([]byte, []int) {
	return file_mountainsteals_proto_rawDescGZIP(), []int{2}
}

func (x *MountainStealsHistory) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *MountainStealsHistory) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *MountainStealsHistory) GetCurPrice() float32 {
	if x != nil {
		return x.CurPrice
	}
	return 0
}

func (x *MountainStealsHistory) GetStockLevel() int32 {
	if x != nil {
		return x.StockLevel
	}
	return 0
}

func (x *MountainStealsHistory) GetSaleCode() string {
	if x != nil {
		return x.SaleCode
	}
	return ""
}

func (x *MountainStealsHistory) GetOffers() []*MountainStealsColorSize {
	if x != nil {
		return x.Offers
	}
	return nil
}

// MountainStealsProduct - mountain steals product
type MountainStealsProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code            string                     `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name            string                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category        []string                   `protobuf:"bytes,3,rep,name=category,proto3" json:"category,omitempty"`
	Brand           string                     `protobuf:"bytes,4,opt,name=brand,proto3" json:"brand,omitempty"`
	Rating          float32                    `protobuf:"fixed32,5,opt,name=rating,proto3" json:"rating,omitempty"`
	ColorSize       []*MountainStealsColorSize `protobuf:"bytes,6,rep,name=colorSize,proto3" json:"colorSize,omitempty"`
	Images          []string                   `protobuf:"bytes,7,rep,name=images,proto3" json:"images,omitempty"`
	Price           float32                    `protobuf:"fixed32,8,opt,name=price,proto3" json:"price,omitempty"`
	MinPrice        float32                    `protobuf:"fixed32,9,opt,name=minPrice,proto3" json:"minPrice,omitempty"`
	MaxPrice        float32                    `protobuf:"fixed32,10,opt,name=maxPrice,proto3" json:"maxPrice,omitempty"`
	Details         string                     `protobuf:"bytes,11,opt,name=details,proto3" json:"details,omitempty"`
	Spec            string                     `protobuf:"bytes,12,opt,name=spec,proto3" json:"spec,omitempty"`
	RatingCount     []int32                    `protobuf:"varint,13,rep,packed,name=ratingCount,proto3" json:"ratingCount,omitempty"`
	MapRading       map[string]float32         `protobuf:"bytes,14,rep,name=mapRading,proto3" json:"mapRading,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	SizeGiud        string                     `protobuf:"bytes,15,opt,name=sizeGiud,proto3" json:"sizeGiud,omitempty"`
	LastUpdatedTime int64                      `protobuf:"varint,16,opt,name=lastUpdatedTime,proto3" json:"lastUpdatedTime,omitempty"`
	LstHistory      []*MountainStealsHistory   `protobuf:"bytes,17,rep,name=lstHistory,proto3" json:"lstHistory,omitempty"`
}

func (x *MountainStealsProduct) Reset() {
	*x = MountainStealsProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mountainsteals_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountainStealsProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountainStealsProduct) ProtoMessage() {}

func (x *MountainStealsProduct) ProtoReflect() protoreflect.Message {
	mi := &file_mountainsteals_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountainStealsProduct.ProtoReflect.Descriptor instead.
func (*MountainStealsProduct) Descriptor() ([]byte, []int) {
	return file_mountainsteals_proto_rawDescGZIP(), []int{3}
}

func (x *MountainStealsProduct) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MountainStealsProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MountainStealsProduct) GetCategory() []string {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *MountainStealsProduct) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *MountainStealsProduct) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *MountainStealsProduct) GetColorSize() []*MountainStealsColorSize {
	if x != nil {
		return x.ColorSize
	}
	return nil
}

func (x *MountainStealsProduct) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *MountainStealsProduct) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *MountainStealsProduct) GetMinPrice() float32 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *MountainStealsProduct) GetMaxPrice() float32 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *MountainStealsProduct) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *MountainStealsProduct) GetSpec() string {
	if x != nil {
		return x.Spec
	}
	return ""
}

func (x *MountainStealsProduct) GetRatingCount() []int32 {
	if x != nil {
		return x.RatingCount
	}
	return nil
}

func (x *MountainStealsProduct) GetMapRading() map[string]float32 {
	if x != nil {
		return x.MapRading
	}
	return nil
}

func (x *MountainStealsProduct) GetSizeGiud() string {
	if x != nil {
		return x.SizeGiud
	}
	return ""
}

func (x *MountainStealsProduct) GetLastUpdatedTime() int64 {
	if x != nil {
		return x.LastUpdatedTime
	}
	return 0
}

func (x *MountainStealsProduct) GetLstHistory() []*MountainStealsHistory {
	if x != nil {
		return x.LstHistory
	}
	return nil
}

// RequestMountainSteals - request mountain steals
type RequestMountainSteals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode MountainStealsMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.MountainStealsMode" json:"mode,omitempty"`
	Url  string             `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *RequestMountainSteals) Reset() {
	*x = RequestMountainSteals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mountainsteals_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMountainSteals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMountainSteals) ProtoMessage() {}

func (x *RequestMountainSteals) ProtoReflect() protoreflect.Message {
	mi := &file_mountainsteals_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMountainSteals.ProtoReflect.Descriptor instead.
func (*RequestMountainSteals) Descriptor() ([]byte, []int) {
	return file_mountainsteals_proto_rawDescGZIP(), []int{4}
}

func (x *RequestMountainSteals) GetMode() MountainStealsMode {
	if x != nil {
		return x.Mode
	}
	return MountainStealsMode_MSM_SALE
}

func (x *RequestMountainSteals) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// ReplyMountainSteals - reply mountain steals
type ReplyMountainSteals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode MountainStealsMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.MountainStealsMode" json:"mode,omitempty"`
	// Types that are assignable to Reply:
	//	*ReplyMountainSteals_Sale
	//	*ReplyMountainSteals_Product
	Reply isReplyMountainSteals_Reply `protobuf_oneof:"reply"`
}

func (x *ReplyMountainSteals) Reset() {
	*x = ReplyMountainSteals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mountainsteals_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyMountainSteals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMountainSteals) ProtoMessage() {}

func (x *ReplyMountainSteals) ProtoReflect() protoreflect.Message {
	mi := &file_mountainsteals_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMountainSteals.ProtoReflect.Descriptor instead.
func (*ReplyMountainSteals) Descriptor() ([]byte, []int) {
	return file_mountainsteals_proto_rawDescGZIP(), []int{5}
}

func (x *ReplyMountainSteals) GetMode() MountainStealsMode {
	if x != nil {
		return x.Mode
	}
	return MountainStealsMode_MSM_SALE
}

func (m *ReplyMountainSteals) GetReply() isReplyMountainSteals_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (x *ReplyMountainSteals) GetSale() *MountainStealsSale {
	if x, ok := x.GetReply().(*ReplyMountainSteals_Sale); ok {
		return x.Sale
	}
	return nil
}

func (x *ReplyMountainSteals) GetProduct() *MountainStealsProduct {
	if x, ok := x.GetReply().(*ReplyMountainSteals_Product); ok {
		return x.Product
	}
	return nil
}

type isReplyMountainSteals_Reply interface {
	isReplyMountainSteals_Reply()
}

type ReplyMountainSteals_Sale struct {
	Sale *MountainStealsSale `protobuf:"bytes,100,opt,name=sale,proto3,oneof"`
}

type ReplyMountainSteals_Product struct {
	Product *MountainStealsProduct `protobuf:"bytes,101,opt,name=product,proto3,oneof"`
}

func (*ReplyMountainSteals_Sale) isReplyMountainSteals_Reply() {}

func (*ReplyMountainSteals_Product) isReplyMountainSteals_Reply() {}

var File_mountainsteals_proto protoreflect.FileDescriptor

var file_mountainsteals_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x65, 0x61, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72,
	0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x12, 0x4d, 0x6f,
	0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x73, 0x53, 0x61, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x61, 0x6c, 0x65, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x61, 0x6c, 0x65, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x7d, 0x0a, 0x17, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x53, 0x74, 0x65, 0x61, 0x6c, 0x73, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x6b, 0x75, 0x22, 0xe9, 0x01, 0x0a, 0x15, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x53, 0x74, 0x65, 0x61, 0x6c, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x75, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x61, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x61, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6a, 0x61,
	0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x73, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x22,
	0xae, 0x05, 0x0a, 0x15, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x61,
	0x6c, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x48, 0x0a, 0x09, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x61, 0x6c,
	0x73, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x55, 0x0a, 0x09, 0x6d,
	0x61, 0x70, 0x52, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37,
	0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x61, 0x6c,
	0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6d, 0x61, 0x70, 0x52, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x7a, 0x65, 0x47, 0x69, 0x75, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69, 0x7a, 0x65, 0x47, 0x69, 0x75, 0x64, 0x12, 0x28,
	0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x6c, 0x73, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6a,
	0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x73, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x6c, 0x73, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x1a, 0x3c, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x52, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x64, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x73, 0x12, 0x39, 0x0a, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73,
	0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x6f, 0x75, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xdc, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x73, 0x12, 0x39,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6a,
	0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x73, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x73, 0x61, 0x6c,
	0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73,
	0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x6f, 0x75, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x73, 0x53, 0x61, 0x6c, 0x65, 0x48, 0x00,
	0x52, 0x04, 0x73, 0x61, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73,
	0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x6f, 0x75, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x07, 0x0a, 0x05,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x33, 0x0a, 0x12, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4d,
	0x53, 0x4d, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x53, 0x4d,
	0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x10, 0x01, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x73, 0x30, 0x30, 0x37, 0x2f,
	0x6a, 0x63, 0x63, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mountainsteals_proto_rawDescOnce sync.Once
	file_mountainsteals_proto_rawDescData = file_mountainsteals_proto_rawDesc
)

func file_mountainsteals_proto_rawDescGZIP() []byte {
	file_mountainsteals_proto_rawDescOnce.Do(func() {
		file_mountainsteals_proto_rawDescData = protoimpl.X.CompressGZIP(file_mountainsteals_proto_rawDescData)
	})
	return file_mountainsteals_proto_rawDescData
}

var file_mountainsteals_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mountainsteals_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mountainsteals_proto_goTypes = []interface{}{
	(MountainStealsMode)(0),         // 0: jarviscrawlercore.MountainStealsMode
	(*MountainStealsSale)(nil),      // 1: jarviscrawlercore.MountainStealsSale
	(*MountainStealsColorSize)(nil), // 2: jarviscrawlercore.MountainStealsColorSize
	(*MountainStealsHistory)(nil),   // 3: jarviscrawlercore.MountainStealsHistory
	(*MountainStealsProduct)(nil),   // 4: jarviscrawlercore.MountainStealsProduct
	(*RequestMountainSteals)(nil),   // 5: jarviscrawlercore.RequestMountainSteals
	(*ReplyMountainSteals)(nil),     // 6: jarviscrawlercore.ReplyMountainSteals
	nil,                             // 7: jarviscrawlercore.MountainStealsProduct.MapRadingEntry
}
var file_mountainsteals_proto_depIdxs = []int32{
	2, // 0: jarviscrawlercore.MountainStealsHistory.offers:type_name -> jarviscrawlercore.MountainStealsColorSize
	2, // 1: jarviscrawlercore.MountainStealsProduct.colorSize:type_name -> jarviscrawlercore.MountainStealsColorSize
	7, // 2: jarviscrawlercore.MountainStealsProduct.mapRading:type_name -> jarviscrawlercore.MountainStealsProduct.MapRadingEntry
	3, // 3: jarviscrawlercore.MountainStealsProduct.lstHistory:type_name -> jarviscrawlercore.MountainStealsHistory
	0, // 4: jarviscrawlercore.RequestMountainSteals.mode:type_name -> jarviscrawlercore.MountainStealsMode
	0, // 5: jarviscrawlercore.ReplyMountainSteals.mode:type_name -> jarviscrawlercore.MountainStealsMode
	1, // 6: jarviscrawlercore.ReplyMountainSteals.sale:type_name -> jarviscrawlercore.MountainStealsSale
	4, // 7: jarviscrawlercore.ReplyMountainSteals.product:type_name -> jarviscrawlercore.MountainStealsProduct
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_mountainsteals_proto_init() }
func file_mountainsteals_proto_init() {
	if File_mountainsteals_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mountainsteals_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountainStealsSale); i {
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
		file_mountainsteals_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountainStealsColorSize); i {
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
		file_mountainsteals_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountainStealsHistory); i {
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
		file_mountainsteals_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountainStealsProduct); i {
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
		file_mountainsteals_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMountainSteals); i {
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
		file_mountainsteals_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyMountainSteals); i {
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
	file_mountainsteals_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ReplyMountainSteals_Sale)(nil),
		(*ReplyMountainSteals_Product)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mountainsteals_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mountainsteals_proto_goTypes,
		DependencyIndexes: file_mountainsteals_proto_depIdxs,
		EnumInfos:         file_mountainsteals_proto_enumTypes,
		MessageInfos:      file_mountainsteals_proto_msgTypes,
	}.Build()
	File_mountainsteals_proto = out.File
	file_mountainsteals_proto_rawDesc = nil
	file_mountainsteals_proto_goTypes = nil
	file_mountainsteals_proto_depIdxs = nil
}
