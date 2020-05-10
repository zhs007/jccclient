// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p6vdy.proto

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

// P6vdyMode - 6vdy mode
type P6VdyMode int32

const (
	// P6VDY_MOVIES - 视频列表
	P6VdyMode_P6VDY_MOVIES P6VdyMode = 0
	// P6VDY_MOVIE - 视频页面
	P6VdyMode_P6VDY_MOVIE P6VdyMode = 1
)

var P6VdyMode_name = map[int32]string{
	0: "P6VDY_MOVIES",
	1: "P6VDY_MOVIE",
}
var P6VdyMode_value = map[string]int32{
	"P6VDY_MOVIES": 0,
	"P6VDY_MOVIE":  1,
}

func (x P6VdyMode) String() string {
	return proto.EnumName(P6VdyMode_name, int32(x))
}
func (P6VdyMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_p6vdy_518afa5c6727bdbb, []int{0}
}

// 6vdy Resource infomation
type P6VdyResInfo struct {
	// name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// url
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *P6VdyResInfo) Reset()         { *m = P6VdyResInfo{} }
func (m *P6VdyResInfo) String() string { return proto.CompactTextString(m) }
func (*P6VdyResInfo) ProtoMessage()    {}
func (*P6VdyResInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_p6vdy_518afa5c6727bdbb, []int{0}
}
func (m *P6VdyResInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_P6VdyResInfo.Unmarshal(m, b)
}
func (m *P6VdyResInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_P6VdyResInfo.Marshal(b, m, deterministic)
}
func (dst *P6VdyResInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_P6VdyResInfo.Merge(dst, src)
}
func (m *P6VdyResInfo) XXX_Size() int {
	return xxx_messageInfo_P6VdyResInfo.Size(m)
}
func (m *P6VdyResInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_P6VdyResInfo.DiscardUnknown(m)
}

var xxx_messageInfo_P6VdyResInfo proto.InternalMessageInfo

func (m *P6VdyResInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *P6VdyResInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// 6vdy Movie
type P6VdyMovie struct {
	// fullname
	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	// resource id
	Resid string `protobuf:"bytes,2,opt,name=resid,proto3" json:"resid,omitempty"`
	// 片名
	Title []string `protobuf:"bytes,3,rep,name=title,proto3" json:"title,omitempty"`
	// 导演
	Director []string `protobuf:"bytes,4,rep,name=director,proto3" json:"director,omitempty"`
	// url
	Url string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	// cover
	Cover string `protobuf:"bytes,6,opt,name=cover,proto3" json:"cover,omitempty"`
	// 导演
	Fulldirector string `protobuf:"bytes,7,opt,name=fulldirector,proto3" json:"fulldirector,omitempty"`
	// 类型
	Category string `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
	// 季
	Season int32 `protobuf:"varint,9,opt,name=season,proto3" json:"season,omitempty"`
	// 集
	Episode int32 `protobuf:"varint,10,opt,name=episode,proto3" json:"episode,omitempty"`
	// 视频列表
	Lst                  []*P6VdyResInfo `protobuf:"bytes,11,rep,name=lst,proto3" json:"lst,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *P6VdyMovie) Reset()         { *m = P6VdyMovie{} }
func (m *P6VdyMovie) String() string { return proto.CompactTextString(m) }
func (*P6VdyMovie) ProtoMessage()    {}
func (*P6VdyMovie) Descriptor() ([]byte, []int) {
	return fileDescriptor_p6vdy_518afa5c6727bdbb, []int{1}
}
func (m *P6VdyMovie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_P6VdyMovie.Unmarshal(m, b)
}
func (m *P6VdyMovie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_P6VdyMovie.Marshal(b, m, deterministic)
}
func (dst *P6VdyMovie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_P6VdyMovie.Merge(dst, src)
}
func (m *P6VdyMovie) XXX_Size() int {
	return xxx_messageInfo_P6VdyMovie.Size(m)
}
func (m *P6VdyMovie) XXX_DiscardUnknown() {
	xxx_messageInfo_P6VdyMovie.DiscardUnknown(m)
}

var xxx_messageInfo_P6VdyMovie proto.InternalMessageInfo

func (m *P6VdyMovie) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *P6VdyMovie) GetResid() string {
	if m != nil {
		return m.Resid
	}
	return ""
}

func (m *P6VdyMovie) GetTitle() []string {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *P6VdyMovie) GetDirector() []string {
	if m != nil {
		return m.Director
	}
	return nil
}

func (m *P6VdyMovie) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *P6VdyMovie) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *P6VdyMovie) GetFulldirector() string {
	if m != nil {
		return m.Fulldirector
	}
	return ""
}

func (m *P6VdyMovie) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *P6VdyMovie) GetSeason() int32 {
	if m != nil {
		return m.Season
	}
	return 0
}

func (m *P6VdyMovie) GetEpisode() int32 {
	if m != nil {
		return m.Episode
	}
	return 0
}

func (m *P6VdyMovie) GetLst() []*P6VdyResInfo {
	if m != nil {
		return m.Lst
	}
	return nil
}

// P6vdyMovies - movies
type P6VdyMovies struct {
	Lst                  []*P6VdyMovie `protobuf:"bytes,1,rep,name=lst,proto3" json:"lst,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *P6VdyMovies) Reset()         { *m = P6VdyMovies{} }
func (m *P6VdyMovies) String() string { return proto.CompactTextString(m) }
func (*P6VdyMovies) ProtoMessage()    {}
func (*P6VdyMovies) Descriptor() ([]byte, []int) {
	return fileDescriptor_p6vdy_518afa5c6727bdbb, []int{2}
}
func (m *P6VdyMovies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_P6VdyMovies.Unmarshal(m, b)
}
func (m *P6VdyMovies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_P6VdyMovies.Marshal(b, m, deterministic)
}
func (dst *P6VdyMovies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_P6VdyMovies.Merge(dst, src)
}
func (m *P6VdyMovies) XXX_Size() int {
	return xxx_messageInfo_P6VdyMovies.Size(m)
}
func (m *P6VdyMovies) XXX_DiscardUnknown() {
	xxx_messageInfo_P6VdyMovies.DiscardUnknown(m)
}

var xxx_messageInfo_P6VdyMovies proto.InternalMessageInfo

func (m *P6VdyMovies) GetLst() []*P6VdyMovie {
	if m != nil {
		return m.Lst
	}
	return nil
}

// RequestP6vdy - request 6vdy
type RequestP6Vdy struct {
	Mode                 P6VdyMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.P6VdyMode" json:"mode,omitempty"`
	Url                  string    `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RequestP6Vdy) Reset()         { *m = RequestP6Vdy{} }
func (m *RequestP6Vdy) String() string { return proto.CompactTextString(m) }
func (*RequestP6Vdy) ProtoMessage()    {}
func (*RequestP6Vdy) Descriptor() ([]byte, []int) {
	return fileDescriptor_p6vdy_518afa5c6727bdbb, []int{3}
}
func (m *RequestP6Vdy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestP6Vdy.Unmarshal(m, b)
}
func (m *RequestP6Vdy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestP6Vdy.Marshal(b, m, deterministic)
}
func (dst *RequestP6Vdy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestP6Vdy.Merge(dst, src)
}
func (m *RequestP6Vdy) XXX_Size() int {
	return xxx_messageInfo_RequestP6Vdy.Size(m)
}
func (m *RequestP6Vdy) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestP6Vdy.DiscardUnknown(m)
}

var xxx_messageInfo_RequestP6Vdy proto.InternalMessageInfo

func (m *RequestP6Vdy) GetMode() P6VdyMode {
	if m != nil {
		return m.Mode
	}
	return P6VdyMode_P6VDY_MOVIES
}

func (m *RequestP6Vdy) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// ReplyP6vdy - reply 6vdy
type ReplyP6Vdy struct {
	Mode P6VdyMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.P6VdyMode" json:"mode,omitempty"`
	// Types that are valid to be assigned to Reply:
	//	*ReplyP6Vdy_Movies
	//	*ReplyP6Vdy_Movie
	Reply                isReplyP6Vdy_Reply `protobuf_oneof:"reply"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ReplyP6Vdy) Reset()         { *m = ReplyP6Vdy{} }
func (m *ReplyP6Vdy) String() string { return proto.CompactTextString(m) }
func (*ReplyP6Vdy) ProtoMessage()    {}
func (*ReplyP6Vdy) Descriptor() ([]byte, []int) {
	return fileDescriptor_p6vdy_518afa5c6727bdbb, []int{4}
}
func (m *ReplyP6Vdy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyP6Vdy.Unmarshal(m, b)
}
func (m *ReplyP6Vdy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyP6Vdy.Marshal(b, m, deterministic)
}
func (dst *ReplyP6Vdy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyP6Vdy.Merge(dst, src)
}
func (m *ReplyP6Vdy) XXX_Size() int {
	return xxx_messageInfo_ReplyP6Vdy.Size(m)
}
func (m *ReplyP6Vdy) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyP6Vdy.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyP6Vdy proto.InternalMessageInfo

type isReplyP6Vdy_Reply interface {
	isReplyP6Vdy_Reply()
}

type ReplyP6Vdy_Movies struct {
	Movies *P6VdyMovies `protobuf:"bytes,100,opt,name=movies,proto3,oneof"`
}
type ReplyP6Vdy_Movie struct {
	Movie *P6VdyMovie `protobuf:"bytes,101,opt,name=movie,proto3,oneof"`
}

func (*ReplyP6Vdy_Movies) isReplyP6Vdy_Reply() {}
func (*ReplyP6Vdy_Movie) isReplyP6Vdy_Reply()  {}

func (m *ReplyP6Vdy) GetReply() isReplyP6Vdy_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *ReplyP6Vdy) GetMode() P6VdyMode {
	if m != nil {
		return m.Mode
	}
	return P6VdyMode_P6VDY_MOVIES
}

func (m *ReplyP6Vdy) GetMovies() *P6VdyMovies {
	if x, ok := m.GetReply().(*ReplyP6Vdy_Movies); ok {
		return x.Movies
	}
	return nil
}

func (m *ReplyP6Vdy) GetMovie() *P6VdyMovie {
	if x, ok := m.GetReply().(*ReplyP6Vdy_Movie); ok {
		return x.Movie
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReplyP6Vdy) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReplyP6Vdy_OneofMarshaler, _ReplyP6Vdy_OneofUnmarshaler, _ReplyP6Vdy_OneofSizer, []interface{}{
		(*ReplyP6Vdy_Movies)(nil),
		(*ReplyP6Vdy_Movie)(nil),
	}
}

func _ReplyP6Vdy_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReplyP6Vdy)
	// reply
	switch x := m.Reply.(type) {
	case *ReplyP6Vdy_Movies:
		b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Movies); err != nil {
			return err
		}
	case *ReplyP6Vdy_Movie:
		b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Movie); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ReplyP6Vdy.Reply has unexpected type %T", x)
	}
	return nil
}

func _ReplyP6Vdy_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReplyP6Vdy)
	switch tag {
	case 100: // reply.movies
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(P6VdyMovies)
		err := b.DecodeMessage(msg)
		m.Reply = &ReplyP6Vdy_Movies{msg}
		return true, err
	case 101: // reply.movie
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(P6VdyMovie)
		err := b.DecodeMessage(msg)
		m.Reply = &ReplyP6Vdy_Movie{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ReplyP6Vdy_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReplyP6Vdy)
	// reply
	switch x := m.Reply.(type) {
	case *ReplyP6Vdy_Movies:
		s := proto.Size(x.Movies)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ReplyP6Vdy_Movie:
		s := proto.Size(x.Movie)
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
	proto.RegisterType((*P6VdyResInfo)(nil), "jarviscrawlercore.P6vdyResInfo")
	proto.RegisterType((*P6VdyMovie)(nil), "jarviscrawlercore.P6vdyMovie")
	proto.RegisterType((*P6VdyMovies)(nil), "jarviscrawlercore.P6vdyMovies")
	proto.RegisterType((*RequestP6Vdy)(nil), "jarviscrawlercore.RequestP6vdy")
	proto.RegisterType((*ReplyP6Vdy)(nil), "jarviscrawlercore.ReplyP6vdy")
	proto.RegisterEnum("jarviscrawlercore.P6VdyMode", P6VdyMode_name, P6VdyMode_value)
}

func init() { proto.RegisterFile("p6vdy.proto", fileDescriptor_p6vdy_518afa5c6727bdbb) }

var fileDescriptor_p6vdy_518afa5c6727bdbb = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcf, 0xcb, 0xd3, 0x40,
	0x10, 0x6d, 0x9a, 0x26, 0x6d, 0x27, 0x45, 0xe3, 0x22, 0xb2, 0x88, 0x3f, 0x42, 0x4e, 0xc1, 0x43,
	0xd4, 0xfa, 0x03, 0x4f, 0x1e, 0x44, 0xa1, 0x3d, 0x14, 0x65, 0x85, 0x82, 0x27, 0x89, 0xc9, 0x54,
	0x22, 0xdb, 0x6e, 0xdc, 0x4d, 0x23, 0xf9, 0xcf, 0xbc, 0xf8, 0xbf, 0xc9, 0x4e, 0xda, 0x58, 0xf9,
	0xbe, 0xf6, 0xf0, 0xdd, 0xe6, 0xcd, 0xbc, 0x79, 0xfb, 0xf6, 0x65, 0x03, 0x41, 0xf5, 0xba, 0x29,
	0xda, 0xb4, 0xd2, 0xaa, 0x56, 0xec, 0xce, 0x8f, 0x4c, 0x37, 0xa5, 0xc9, 0x75, 0xf6, 0x4b, 0xa2,
	0xce, 0x95, 0xc6, 0xf8, 0x25, 0xcc, 0x3e, 0x59, 0x86, 0x40, 0xb3, 0xdc, 0x6d, 0x14, 0x63, 0x30,
	0xda, 0x65, 0x5b, 0xe4, 0x4e, 0xe4, 0x24, 0x53, 0x41, 0x35, 0x0b, 0xc1, 0xdd, 0x6b, 0xc9, 0x87,
	0xd4, 0xb2, 0x65, 0xfc, 0x67, 0x08, 0x40, 0x6b, 0x2b, 0xd5, 0x94, 0xc8, 0xee, 0xc3, 0x64, 0xb3,
	0x97, 0xf2, 0x64, 0xb1, 0xc7, 0xec, 0x2e, 0x78, 0x1a, 0x4d, 0x59, 0x1c, 0xd6, 0x3b, 0x60, 0xbb,
	0x75, 0x59, 0x4b, 0xe4, 0x6e, 0xe4, 0xda, 0x2e, 0x01, 0xab, 0x53, 0x94, 0x1a, 0xf3, 0x5a, 0x69,
	0x3e, 0xa2, 0x41, 0x8f, 0x8f, 0x26, 0xbc, 0xde, 0x84, 0xd5, 0xc8, 0x55, 0x83, 0x9a, 0xfb, 0x9d,
	0x32, 0x01, 0x16, 0xc3, 0xcc, 0x9e, 0xdd, 0xeb, 0x8c, 0x69, 0xf8, 0x5f, 0xcf, 0x9e, 0x93, 0x67,
	0x35, 0x7e, 0x57, 0xba, 0xe5, 0x93, 0xce, 0xef, 0x11, 0xb3, 0x7b, 0xe0, 0x1b, 0xcc, 0x8c, 0xda,
	0xf1, 0x69, 0xe4, 0x24, 0x9e, 0x38, 0x20, 0xc6, 0x61, 0x8c, 0x55, 0x69, 0x54, 0x81, 0x1c, 0x68,
	0x70, 0x84, 0xec, 0x39, 0xb8, 0xd2, 0xd4, 0x3c, 0x88, 0xdc, 0x24, 0x98, 0x3f, 0x4e, 0xaf, 0x64,
	0x9c, 0x9e, 0x06, 0x2c, 0x2c, 0x37, 0x7e, 0x0b, 0xc1, 0xbf, 0xf8, 0x0c, 0x7b, 0xda, 0x29, 0x38,
	0xa4, 0xf0, 0xf0, 0x9c, 0x02, 0x91, 0xbb, 0x7d, 0x01, 0x33, 0x81, 0x3f, 0xf7, 0x68, 0x6a, 0x9a,
	0xb0, 0x67, 0x30, 0xda, 0x5a, 0x67, 0x36, 0xfc, 0x5b, 0xf3, 0x07, 0xe7, 0x15, 0x0a, 0x14, 0xc4,
	0xbc, 0xe6, 0x9b, 0xfe, 0x76, 0x00, 0x04, 0x56, 0xb2, 0xbd, 0xa9, 0xe4, 0x1b, 0xf0, 0xb7, 0x74,
	0x1f, 0x5e, 0x44, 0x4e, 0x12, 0xcc, 0x1f, 0x5d, 0xbc, 0x88, 0x59, 0x0c, 0xc4, 0x81, 0xcf, 0x5e,
	0x81, 0x47, 0x15, 0x47, 0x5a, 0xbc, 0x9c, 0xc0, 0x62, 0x20, 0x3a, 0xf6, 0xbb, 0xb1, 0x7d, 0x5a,
	0x95, 0x6c, 0x9f, 0xa4, 0x30, 0xed, 0xcd, 0xb0, 0xd0, 0xbe, 0xe8, 0xf5, 0xfb, 0x2f, 0x5f, 0x57,
	0x1f, 0xd7, 0xcb, 0x0f, 0x9f, 0xc3, 0x01, 0xbb, 0x6d, 0xd3, 0xee, 0x3b, 0xa1, 0xf3, 0xcd, 0xa7,
	0xdf, 0xe1, 0xc5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x0c, 0x43, 0x1d, 0x1d, 0x03, 0x00,
	0x00,
}
