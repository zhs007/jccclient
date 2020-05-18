// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: db.proto

package dbpb

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

// HostInfo - host info
type HostInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostName       string `protobuf:"bytes,1,opt,name=hostName,proto3" json:"hostName,omitempty"`
	TaskNums       int32  `protobuf:"varint,2,opt,name=taskNums,proto3" json:"taskNums,omitempty"`
	FailNums       int32  `protobuf:"varint,3,opt,name=failNums,proto3" json:"failNums,omitempty"`
	LastTime       int64  `protobuf:"varint,4,opt,name=lastTime,proto3" json:"lastTime,omitempty"`
	LastFailTime   int64  `protobuf:"varint,5,opt,name=lastFailTime,proto3" json:"lastFailTime,omitempty"`
	StartSleepTime int64  `protobuf:"varint,6,opt,name=startSleepTime,proto3" json:"startSleepTime,omitempty"`
	SleepTime      int64  `protobuf:"varint,7,opt,name=sleepTime,proto3" json:"sleepTime,omitempty"`
	MultiNums      int32  `protobuf:"varint,8,opt,name=multiNums,proto3" json:"multiNums,omitempty"`
	// startTime - 这个用来统计这次错误之前，请求次数用的
	StartTime int64 `protobuf:"varint,9,opt,name=startTime,proto3" json:"startTime,omitempty"`
	// curNums - 同上
	CurNums int64 `protobuf:"varint,10,opt,name=curNums,proto3" json:"curNums,omitempty"`
	// sleepTimeAtStart - sleep time at start task
	//
	// Deprecated: Do not use.
	SleepTimeAtStart int32 `protobuf:"varint,11,opt,name=sleepTimeAtStart,proto3" json:"sleepTimeAtStart,omitempty"`
	// sleepTimeMsAtStart - sleep time (ms) at start task
	SleepTimeMsAtStart int32 `protobuf:"varint,12,opt,name=sleepTimeMsAtStart,proto3" json:"sleepTimeMsAtStart,omitempty"`
}

func (x *HostInfo) Reset() {
	*x = HostInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostInfo) ProtoMessage() {}

func (x *HostInfo) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostInfo.ProtoReflect.Descriptor instead.
func (*HostInfo) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{0}
}

func (x *HostInfo) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *HostInfo) GetTaskNums() int32 {
	if x != nil {
		return x.TaskNums
	}
	return 0
}

func (x *HostInfo) GetFailNums() int32 {
	if x != nil {
		return x.FailNums
	}
	return 0
}

func (x *HostInfo) GetLastTime() int64 {
	if x != nil {
		return x.LastTime
	}
	return 0
}

func (x *HostInfo) GetLastFailTime() int64 {
	if x != nil {
		return x.LastFailTime
	}
	return 0
}

func (x *HostInfo) GetStartSleepTime() int64 {
	if x != nil {
		return x.StartSleepTime
	}
	return 0
}

func (x *HostInfo) GetSleepTime() int64 {
	if x != nil {
		return x.SleepTime
	}
	return 0
}

func (x *HostInfo) GetMultiNums() int32 {
	if x != nil {
		return x.MultiNums
	}
	return 0
}

func (x *HostInfo) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *HostInfo) GetCurNums() int64 {
	if x != nil {
		return x.CurNums
	}
	return 0
}

// Deprecated: Do not use.
func (x *HostInfo) GetSleepTimeAtStart() int32 {
	if x != nil {
		return x.SleepTimeAtStart
	}
	return 0
}

func (x *HostInfo) GetSleepTimeMsAtStart() int32 {
	if x != nil {
		return x.SleepTimeMsAtStart
	}
	return 0
}

var File_db_proto protoreflect.FileDescriptor

var file_db_proto_rawDesc = []byte{
	0x0a, 0x08, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6a, 0x63, 0x63, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x64, 0x62, 0x70, 0x62, 0x22, 0x9a, 0x03, 0x0a, 0x08, 0x48, 0x6f,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x75, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x4e, 0x75, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x46, 0x61,
	0x69, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x4e, 0x75, 0x6d, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x75, 0x72, 0x4e, 0x75, 0x6d, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63,
	0x75, 0x72, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x2e, 0x0a, 0x10, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x54,
	0x69, 0x6d, 0x65, 0x41, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x10, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x41,
	0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x54,
	0x69, 0x6d, 0x65, 0x4d, 0x73, 0x41, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x12, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x41,
	0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x73, 0x30, 0x30, 0x37, 0x2f, 0x6a, 0x63, 0x63, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x62, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_db_proto_rawDescOnce sync.Once
	file_db_proto_rawDescData = file_db_proto_rawDesc
)

func file_db_proto_rawDescGZIP() []byte {
	file_db_proto_rawDescOnce.Do(func() {
		file_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_proto_rawDescData)
	})
	return file_db_proto_rawDescData
}

var file_db_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_db_proto_goTypes = []interface{}{
	(*HostInfo)(nil), // 0: jccclientdbpb.HostInfo
}
var file_db_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_db_proto_init() }
func file_db_proto_init() {
	if File_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostInfo); i {
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
			RawDescriptor: file_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_proto_goTypes,
		DependencyIndexes: file_db_proto_depIdxs,
		MessageInfos:      file_db_proto_msgTypes,
	}.Build()
	File_db_proto = out.File
	file_db_proto_rawDesc = nil
	file_db_proto_goTypes = nil
	file_db_proto_depIdxs = nil
}
