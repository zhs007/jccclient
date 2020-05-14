// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: geoip.proto

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

// RequestGeoIP - request geoip
type RequestGeoIP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ip - ip address
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// platform - it's like ipvoid, default is ipvoid
	Platform string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
}

func (x *RequestGeoIP) Reset() {
	*x = RequestGeoIP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geoip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGeoIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGeoIP) ProtoMessage() {}

func (x *RequestGeoIP) ProtoReflect() protoreflect.Message {
	mi := &file_geoip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGeoIP.ProtoReflect.Descriptor instead.
func (*RequestGeoIP) Descriptor() ([]byte, []int) {
	return file_geoip_proto_rawDescGZIP(), []int{0}
}

func (x *RequestGeoIP) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *RequestGeoIP) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

// ReplyGeoIP - reply geoip
type ReplyGeoIP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude     float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude    float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Organization string  `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty"`
	Asn          string  `protobuf:"bytes,4,opt,name=asn,proto3" json:"asn,omitempty"`
	Continent    string  `protobuf:"bytes,5,opt,name=continent,proto3" json:"continent,omitempty"`
	Country      string  `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	Region       string  `protobuf:"bytes,7,opt,name=region,proto3" json:"region,omitempty"`
	City         string  `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty"`
	Hostname     string  `protobuf:"bytes,9,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *ReplyGeoIP) Reset() {
	*x = ReplyGeoIP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geoip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyGeoIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyGeoIP) ProtoMessage() {}

func (x *ReplyGeoIP) ProtoReflect() protoreflect.Message {
	mi := &file_geoip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyGeoIP.ProtoReflect.Descriptor instead.
func (*ReplyGeoIP) Descriptor() ([]byte, []int) {
	return file_geoip_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyGeoIP) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *ReplyGeoIP) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *ReplyGeoIP) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *ReplyGeoIP) GetAsn() string {
	if x != nil {
		return x.Asn
	}
	return ""
}

func (x *ReplyGeoIP) GetContinent() string {
	if x != nil {
		return x.Continent
	}
	return ""
}

func (x *ReplyGeoIP) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *ReplyGeoIP) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ReplyGeoIP) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ReplyGeoIP) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

var File_geoip_proto protoreflect.FileDescriptor

var file_geoip_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6a,
	0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65,
	0x22, 0x3a, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6f, 0x49, 0x50,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0xfc, 0x01, 0x0a,
	0x0a, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x47, 0x65, 0x6f, 0x49, 0x50, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x73, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x73, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x73, 0x30, 0x30, 0x37,
	0x2f, 0x6a, 0x63, 0x63, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_geoip_proto_rawDescOnce sync.Once
	file_geoip_proto_rawDescData = file_geoip_proto_rawDesc
)

func file_geoip_proto_rawDescGZIP() []byte {
	file_geoip_proto_rawDescOnce.Do(func() {
		file_geoip_proto_rawDescData = protoimpl.X.CompressGZIP(file_geoip_proto_rawDescData)
	})
	return file_geoip_proto_rawDescData
}

var file_geoip_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_geoip_proto_goTypes = []interface{}{
	(*RequestGeoIP)(nil), // 0: jarviscrawlercore.RequestGeoIP
	(*ReplyGeoIP)(nil),   // 1: jarviscrawlercore.ReplyGeoIP
}
var file_geoip_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_geoip_proto_init() }
func file_geoip_proto_init() {
	if File_geoip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_geoip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGeoIP); i {
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
		file_geoip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyGeoIP); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_geoip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geoip_proto_goTypes,
		DependencyIndexes: file_geoip_proto_depIdxs,
		MessageInfos:      file_geoip_proto_msgTypes,
	}.Build()
	File_geoip_proto = out.File
	file_geoip_proto_rawDesc = nil
	file_geoip_proto_goTypes = nil
	file_geoip_proto_depIdxs = nil
}