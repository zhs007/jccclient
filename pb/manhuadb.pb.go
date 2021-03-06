// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: manhuadb.proto

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

// ManhuaDBMode - manhuadb mode
type ManhuaDBMode int32

const (
	// MHDB_MANHUA - manhua
	ManhuaDBMode_MHDB_MANHUA ManhuaDBMode = 0
	// MHDB_BOOK - book
	ManhuaDBMode_MHDB_BOOK ManhuaDBMode = 1
	// MHDB_AUTHOR - author
	ManhuaDBMode_MHDB_AUTHOR ManhuaDBMode = 2
)

// Enum value maps for ManhuaDBMode.
var (
	ManhuaDBMode_name = map[int32]string{
		0: "MHDB_MANHUA",
		1: "MHDB_BOOK",
		2: "MHDB_AUTHOR",
	}
	ManhuaDBMode_value = map[string]int32{
		"MHDB_MANHUA": 0,
		"MHDB_BOOK":   1,
		"MHDB_AUTHOR": 2,
	}
)

func (x ManhuaDBMode) Enum() *ManhuaDBMode {
	p := new(ManhuaDBMode)
	*p = x
	return p
}

func (x ManhuaDBMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ManhuaDBMode) Descriptor() protoreflect.EnumDescriptor {
	return file_manhuadb_proto_enumTypes[0].Descriptor()
}

func (ManhuaDBMode) Type() protoreflect.EnumType {
	return &file_manhuadb_proto_enumTypes[0]
}

func (x ManhuaDBMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ManhuaDBMode.Descriptor instead.
func (ManhuaDBMode) EnumDescriptor() ([]byte, []int) {
	return file_manhuadb_proto_rawDescGZIP(), []int{0}
}

// ManhuaDBBook - manhuadb book
type ManhuaDBBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	RootType int32  `protobuf:"varint,4,opt,name=rootType,proto3" json:"rootType,omitempty"`
	RootName string `protobuf:"bytes,5,opt,name=rootName,proto3" json:"rootName,omitempty"`
}

func (x *ManhuaDBBook) Reset() {
	*x = ManhuaDBBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manhuadb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManhuaDBBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManhuaDBBook) ProtoMessage() {}

func (x *ManhuaDBBook) ProtoReflect() protoreflect.Message {
	mi := &file_manhuadb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManhuaDBBook.ProtoReflect.Descriptor instead.
func (*ManhuaDBBook) Descriptor() ([]byte, []int) {
	return file_manhuadb_proto_rawDescGZIP(), []int{0}
}

func (x *ManhuaDBBook) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ManhuaDBBook) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ManhuaDBBook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ManhuaDBBook) GetRootType() int32 {
	if x != nil {
		return x.RootType
	}
	return 0
}

func (x *ManhuaDBBook) GetRootName() string {
	if x != nil {
		return x.RootName
	}
	return ""
}

// ManhuaDBManhua - manhuadb manhua
type ManhuaDBManhua struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comicid string          `protobuf:"bytes,1,opt,name=comicid,proto3" json:"comicid,omitempty"`
	Name    string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Books   []*ManhuaDBBook `protobuf:"bytes,3,rep,name=books,proto3" json:"books,omitempty"`
	Authors []string        `protobuf:"bytes,4,rep,name=authors,proto3" json:"authors,omitempty"`
	Url     string          `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Cover   string          `protobuf:"bytes,6,opt,name=cover,proto3" json:"cover,omitempty"`
}

func (x *ManhuaDBManhua) Reset() {
	*x = ManhuaDBManhua{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manhuadb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManhuaDBManhua) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManhuaDBManhua) ProtoMessage() {}

func (x *ManhuaDBManhua) ProtoReflect() protoreflect.Message {
	mi := &file_manhuadb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManhuaDBManhua.ProtoReflect.Descriptor instead.
func (*ManhuaDBManhua) Descriptor() ([]byte, []int) {
	return file_manhuadb_proto_rawDescGZIP(), []int{1}
}

func (x *ManhuaDBManhua) GetComicid() string {
	if x != nil {
		return x.Comicid
	}
	return ""
}

func (x *ManhuaDBManhua) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ManhuaDBManhua) GetBooks() []*ManhuaDBBook {
	if x != nil {
		return x.Books
	}
	return nil
}

func (x *ManhuaDBManhua) GetAuthors() []string {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *ManhuaDBManhua) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ManhuaDBManhua) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

// ManhuaDBAuthor - manhuadb author
type ManhuaDBAuthor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OtherName []string          `protobuf:"bytes,2,rep,name=otherName,proto3" json:"otherName,omitempty"`
	Detail    string            `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	Manhua    []*ManhuaDBManhua `protobuf:"bytes,4,rep,name=manhua,proto3" json:"manhua,omitempty"`
	Authorid  string            `protobuf:"bytes,5,opt,name=authorid,proto3" json:"authorid,omitempty"`
}

func (x *ManhuaDBAuthor) Reset() {
	*x = ManhuaDBAuthor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manhuadb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManhuaDBAuthor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManhuaDBAuthor) ProtoMessage() {}

func (x *ManhuaDBAuthor) ProtoReflect() protoreflect.Message {
	mi := &file_manhuadb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManhuaDBAuthor.ProtoReflect.Descriptor instead.
func (*ManhuaDBAuthor) Descriptor() ([]byte, []int) {
	return file_manhuadb_proto_rawDescGZIP(), []int{2}
}

func (x *ManhuaDBAuthor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ManhuaDBAuthor) GetOtherName() []string {
	if x != nil {
		return x.OtherName
	}
	return nil
}

func (x *ManhuaDBAuthor) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *ManhuaDBAuthor) GetManhua() []*ManhuaDBManhua {
	if x != nil {
		return x.Manhua
	}
	return nil
}

func (x *ManhuaDBAuthor) GetAuthorid() string {
	if x != nil {
		return x.Authorid
	}
	return ""
}

// ManhuaDBPage - manhuadb page
type ManhuaDBPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url       string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	PageIndex int32  `protobuf:"varint,2,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ManhuaDBPage) Reset() {
	*x = ManhuaDBPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manhuadb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManhuaDBPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManhuaDBPage) ProtoMessage() {}

func (x *ManhuaDBPage) ProtoReflect() protoreflect.Message {
	mi := &file_manhuadb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManhuaDBPage.ProtoReflect.Descriptor instead.
func (*ManhuaDBPage) Descriptor() ([]byte, []int) {
	return file_manhuadb_proto_rawDescGZIP(), []int{3}
}

func (x *ManhuaDBPage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ManhuaDBPage) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *ManhuaDBPage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// ManhuaDBBookInfo - manhuadb book
type ManhuaDBBookInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comicid  string          `protobuf:"bytes,1,opt,name=comicid,proto3" json:"comicid,omitempty"`
	Bookid   string          `protobuf:"bytes,2,opt,name=bookid,proto3" json:"bookid,omitempty"`
	Pages    []*ManhuaDBPage `protobuf:"bytes,3,rep,name=pages,proto3" json:"pages,omitempty"`
	PageNums int32           `protobuf:"varint,4,opt,name=pageNums,proto3" json:"pageNums,omitempty"`
}

func (x *ManhuaDBBookInfo) Reset() {
	*x = ManhuaDBBookInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manhuadb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManhuaDBBookInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManhuaDBBookInfo) ProtoMessage() {}

func (x *ManhuaDBBookInfo) ProtoReflect() protoreflect.Message {
	mi := &file_manhuadb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManhuaDBBookInfo.ProtoReflect.Descriptor instead.
func (*ManhuaDBBookInfo) Descriptor() ([]byte, []int) {
	return file_manhuadb_proto_rawDescGZIP(), []int{4}
}

func (x *ManhuaDBBookInfo) GetComicid() string {
	if x != nil {
		return x.Comicid
	}
	return ""
}

func (x *ManhuaDBBookInfo) GetBookid() string {
	if x != nil {
		return x.Bookid
	}
	return ""
}

func (x *ManhuaDBBookInfo) GetPages() []*ManhuaDBPage {
	if x != nil {
		return x.Pages
	}
	return nil
}

func (x *ManhuaDBBookInfo) GetPageNums() int32 {
	if x != nil {
		return x.PageNums
	}
	return 0
}

// RequestManhuaDB - request manhuadb
type RequestManhuaDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode      ManhuaDBMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.ManhuaDBMode" json:"mode,omitempty"`
	Comicid   string       `protobuf:"bytes,2,opt,name=comicid,proto3" json:"comicid,omitempty"`
	Bookid    string       `protobuf:"bytes,3,opt,name=bookid,proto3" json:"bookid,omitempty"`
	PageIndex int32        `protobuf:"varint,4,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	Authorid  string       `protobuf:"bytes,5,opt,name=authorid,proto3" json:"authorid,omitempty"`
}

func (x *RequestManhuaDB) Reset() {
	*x = RequestManhuaDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manhuadb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestManhuaDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestManhuaDB) ProtoMessage() {}

func (x *RequestManhuaDB) ProtoReflect() protoreflect.Message {
	mi := &file_manhuadb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestManhuaDB.ProtoReflect.Descriptor instead.
func (*RequestManhuaDB) Descriptor() ([]byte, []int) {
	return file_manhuadb_proto_rawDescGZIP(), []int{5}
}

func (x *RequestManhuaDB) GetMode() ManhuaDBMode {
	if x != nil {
		return x.Mode
	}
	return ManhuaDBMode_MHDB_MANHUA
}

func (x *RequestManhuaDB) GetComicid() string {
	if x != nil {
		return x.Comicid
	}
	return ""
}

func (x *RequestManhuaDB) GetBookid() string {
	if x != nil {
		return x.Bookid
	}
	return ""
}

func (x *RequestManhuaDB) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *RequestManhuaDB) GetAuthorid() string {
	if x != nil {
		return x.Authorid
	}
	return ""
}

// ReplyManhuaDB - reply manhuadb
type ReplyManhuaDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode ManhuaDBMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.ManhuaDBMode" json:"mode,omitempty"`
	// Types that are assignable to Reply:
	//	*ReplyManhuaDB_Manhua
	//	*ReplyManhuaDB_Book
	//	*ReplyManhuaDB_Author
	Reply isReplyManhuaDB_Reply `protobuf_oneof:"reply"`
}

func (x *ReplyManhuaDB) Reset() {
	*x = ReplyManhuaDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manhuadb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyManhuaDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyManhuaDB) ProtoMessage() {}

func (x *ReplyManhuaDB) ProtoReflect() protoreflect.Message {
	mi := &file_manhuadb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyManhuaDB.ProtoReflect.Descriptor instead.
func (*ReplyManhuaDB) Descriptor() ([]byte, []int) {
	return file_manhuadb_proto_rawDescGZIP(), []int{6}
}

func (x *ReplyManhuaDB) GetMode() ManhuaDBMode {
	if x != nil {
		return x.Mode
	}
	return ManhuaDBMode_MHDB_MANHUA
}

func (m *ReplyManhuaDB) GetReply() isReplyManhuaDB_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (x *ReplyManhuaDB) GetManhua() *ManhuaDBManhua {
	if x, ok := x.GetReply().(*ReplyManhuaDB_Manhua); ok {
		return x.Manhua
	}
	return nil
}

func (x *ReplyManhuaDB) GetBook() *ManhuaDBBookInfo {
	if x, ok := x.GetReply().(*ReplyManhuaDB_Book); ok {
		return x.Book
	}
	return nil
}

func (x *ReplyManhuaDB) GetAuthor() *ManhuaDBAuthor {
	if x, ok := x.GetReply().(*ReplyManhuaDB_Author); ok {
		return x.Author
	}
	return nil
}

type isReplyManhuaDB_Reply interface {
	isReplyManhuaDB_Reply()
}

type ReplyManhuaDB_Manhua struct {
	Manhua *ManhuaDBManhua `protobuf:"bytes,100,opt,name=manhua,proto3,oneof"`
}

type ReplyManhuaDB_Book struct {
	Book *ManhuaDBBookInfo `protobuf:"bytes,101,opt,name=book,proto3,oneof"`
}

type ReplyManhuaDB_Author struct {
	Author *ManhuaDBAuthor `protobuf:"bytes,102,opt,name=author,proto3,oneof"`
}

func (*ReplyManhuaDB_Manhua) isReplyManhuaDB_Reply() {}

func (*ReplyManhuaDB_Book) isReplyManhuaDB_Reply() {}

func (*ReplyManhuaDB_Author) isReplyManhuaDB_Reply() {}

var File_manhuadb_proto protoreflect.FileDescriptor

var file_manhuadb_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63,
	0x6f, 0x72, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x44, 0x42,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x6f, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x6f, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x6e,
	0x68, 0x75, 0x61, 0x44, 0x42, 0x4d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x69, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x69, 0x63, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69,
	0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61, 0x6e,
	0x68, 0x75, 0x61, 0x44, 0x42, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x22, 0xb1, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x44, 0x42, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x39, 0x0a, 0x06, 0x6d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x44, 0x42, 0x4d, 0x61, 0x6e, 0x68,
	0x75, 0x61, 0x52, 0x06, 0x6d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x64, 0x22, 0x52, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x68, 0x75, 0x61,
	0x44, 0x42, 0x50, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x97, 0x01, 0x0a, 0x10, 0x4d,
	0x61, 0x6e, 0x68, 0x75, 0x61, 0x44, 0x42, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x64, 0x12, 0x35, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x44, 0x42, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x44, 0x42, 0x12, 0x33, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61, 0x6e, 0x68, 0x75,
	0x61, 0x44, 0x42, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x69, 0x63, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x64, 0x22, 0x82, 0x02, 0x0a, 0x0d, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x4d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x44, 0x42, 0x12, 0x33, 0x0a, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6a, 0x61, 0x72, 0x76,
	0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61,
	0x6e, 0x68, 0x75, 0x61, 0x44, 0x42, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x12, 0x3b, 0x0a, 0x06, 0x6d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x44, 0x42, 0x4d, 0x61, 0x6e,
	0x68, 0x75, 0x61, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x12, 0x39, 0x0a,
	0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6a, 0x61,
	0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x4d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x44, 0x42, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x3b, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69,
	0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61, 0x6e,
	0x68, 0x75, 0x61, 0x44, 0x42, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x3f,
	0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x68, 0x75, 0x61, 0x44, 0x42, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x4d, 0x48, 0x44, 0x42, 0x5f, 0x4d, 0x41, 0x4e, 0x48, 0x55, 0x41, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x4d, 0x48, 0x44, 0x42, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x4d, 0x48, 0x44, 0x42, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x10, 0x02, 0x42,
	0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68,
	0x73, 0x30, 0x30, 0x37, 0x2f, 0x6a, 0x63, 0x63, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manhuadb_proto_rawDescOnce sync.Once
	file_manhuadb_proto_rawDescData = file_manhuadb_proto_rawDesc
)

func file_manhuadb_proto_rawDescGZIP() []byte {
	file_manhuadb_proto_rawDescOnce.Do(func() {
		file_manhuadb_proto_rawDescData = protoimpl.X.CompressGZIP(file_manhuadb_proto_rawDescData)
	})
	return file_manhuadb_proto_rawDescData
}

var file_manhuadb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_manhuadb_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_manhuadb_proto_goTypes = []interface{}{
	(ManhuaDBMode)(0),        // 0: jarviscrawlercore.ManhuaDBMode
	(*ManhuaDBBook)(nil),     // 1: jarviscrawlercore.ManhuaDBBook
	(*ManhuaDBManhua)(nil),   // 2: jarviscrawlercore.ManhuaDBManhua
	(*ManhuaDBAuthor)(nil),   // 3: jarviscrawlercore.ManhuaDBAuthor
	(*ManhuaDBPage)(nil),     // 4: jarviscrawlercore.ManhuaDBPage
	(*ManhuaDBBookInfo)(nil), // 5: jarviscrawlercore.ManhuaDBBookInfo
	(*RequestManhuaDB)(nil),  // 6: jarviscrawlercore.RequestManhuaDB
	(*ReplyManhuaDB)(nil),    // 7: jarviscrawlercore.ReplyManhuaDB
}
var file_manhuadb_proto_depIdxs = []int32{
	1, // 0: jarviscrawlercore.ManhuaDBManhua.books:type_name -> jarviscrawlercore.ManhuaDBBook
	2, // 1: jarviscrawlercore.ManhuaDBAuthor.manhua:type_name -> jarviscrawlercore.ManhuaDBManhua
	4, // 2: jarviscrawlercore.ManhuaDBBookInfo.pages:type_name -> jarviscrawlercore.ManhuaDBPage
	0, // 3: jarviscrawlercore.RequestManhuaDB.mode:type_name -> jarviscrawlercore.ManhuaDBMode
	0, // 4: jarviscrawlercore.ReplyManhuaDB.mode:type_name -> jarviscrawlercore.ManhuaDBMode
	2, // 5: jarviscrawlercore.ReplyManhuaDB.manhua:type_name -> jarviscrawlercore.ManhuaDBManhua
	5, // 6: jarviscrawlercore.ReplyManhuaDB.book:type_name -> jarviscrawlercore.ManhuaDBBookInfo
	3, // 7: jarviscrawlercore.ReplyManhuaDB.author:type_name -> jarviscrawlercore.ManhuaDBAuthor
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_manhuadb_proto_init() }
func file_manhuadb_proto_init() {
	if File_manhuadb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manhuadb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManhuaDBBook); i {
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
		file_manhuadb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManhuaDBManhua); i {
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
		file_manhuadb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManhuaDBAuthor); i {
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
		file_manhuadb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManhuaDBPage); i {
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
		file_manhuadb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManhuaDBBookInfo); i {
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
		file_manhuadb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestManhuaDB); i {
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
		file_manhuadb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyManhuaDB); i {
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
	file_manhuadb_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*ReplyManhuaDB_Manhua)(nil),
		(*ReplyManhuaDB_Book)(nil),
		(*ReplyManhuaDB_Author)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_manhuadb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_manhuadb_proto_goTypes,
		DependencyIndexes: file_manhuadb_proto_depIdxs,
		EnumInfos:         file_manhuadb_proto_enumTypes,
		MessageInfos:      file_manhuadb_proto_msgTypes,
	}.Build()
	File_manhuadb_proto = out.File
	file_manhuadb_proto_rawDesc = nil
	file_manhuadb_proto_goTypes = nil
	file_manhuadb_proto_depIdxs = nil
}
