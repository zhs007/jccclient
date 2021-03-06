// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: techinasia.proto

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

type TechInAsiaMode int32

const (
	// TIAM_COMPANY - company
	TechInAsiaMode_TIAM_COMPANY TechInAsiaMode = 0
	// TIAM_JOB - job
	TechInAsiaMode_TIAM_JOB TechInAsiaMode = 1
	// TIAM_JOBLIST - job list
	TechInAsiaMode_TIAM_JOBLIST TechInAsiaMode = 2
	// TIAM_JOBTAG - job tag
	TechInAsiaMode_TIAM_JOBTAG TechInAsiaMode = 3
)

// Enum value maps for TechInAsiaMode.
var (
	TechInAsiaMode_name = map[int32]string{
		0: "TIAM_COMPANY",
		1: "TIAM_JOB",
		2: "TIAM_JOBLIST",
		3: "TIAM_JOBTAG",
	}
	TechInAsiaMode_value = map[string]int32{
		"TIAM_COMPANY": 0,
		"TIAM_JOB":     1,
		"TIAM_JOBLIST": 2,
		"TIAM_JOBTAG":  3,
	}
)

func (x TechInAsiaMode) Enum() *TechInAsiaMode {
	p := new(TechInAsiaMode)
	*p = x
	return p
}

func (x TechInAsiaMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TechInAsiaMode) Descriptor() protoreflect.EnumDescriptor {
	return file_techinasia_proto_enumTypes[0].Descriptor()
}

func (TechInAsiaMode) Type() protoreflect.EnumType {
	return &file_techinasia_proto_enumTypes[0]
}

func (x TechInAsiaMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TechInAsiaMode.Descriptor instead.
func (TechInAsiaMode) EnumDescriptor() ([]byte, []int) {
	return file_techinasia_proto_rawDescGZIP(), []int{0}
}

// TechInAsiaCompany - TechInAsia Company
type TechInAsiaCompany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Avatar          string   `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Location        []string `protobuf:"bytes,3,rep,name=location,proto3" json:"location,omitempty"`
	Categories      []string `protobuf:"bytes,4,rep,name=categories,proto3" json:"categories,omitempty"`
	Employees       int32    `protobuf:"varint,5,opt,name=employees,proto3" json:"employees,omitempty"`
	Introduction    string   `protobuf:"bytes,6,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Links           []string `protobuf:"bytes,7,rep,name=links,proto3" json:"links,omitempty"`
	CompanyCode     string   `protobuf:"bytes,8,opt,name=companyCode,proto3" json:"companyCode,omitempty"`
	LastUpdatedTime int64    `protobuf:"varint,9,opt,name=lastUpdatedTime,proto3" json:"lastUpdatedTime,omitempty"`
}

func (x *TechInAsiaCompany) Reset() {
	*x = TechInAsiaCompany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techinasia_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TechInAsiaCompany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TechInAsiaCompany) ProtoMessage() {}

func (x *TechInAsiaCompany) ProtoReflect() protoreflect.Message {
	mi := &file_techinasia_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TechInAsiaCompany.ProtoReflect.Descriptor instead.
func (*TechInAsiaCompany) Descriptor() ([]byte, []int) {
	return file_techinasia_proto_rawDescGZIP(), []int{0}
}

func (x *TechInAsiaCompany) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TechInAsiaCompany) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *TechInAsiaCompany) GetLocation() []string {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *TechInAsiaCompany) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *TechInAsiaCompany) GetEmployees() int32 {
	if x != nil {
		return x.Employees
	}
	return 0
}

func (x *TechInAsiaCompany) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *TechInAsiaCompany) GetLinks() []string {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *TechInAsiaCompany) GetCompanyCode() string {
	if x != nil {
		return x.CompanyCode
	}
	return ""
}

func (x *TechInAsiaCompany) GetLastUpdatedTime() int64 {
	if x != nil {
		return x.LastUpdatedTime
	}
	return 0
}

// TechInAsiaJob - TechInAsia Job
type TechInAsiaJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyName     string   `protobuf:"bytes,1,opt,name=companyName,proto3" json:"companyName,omitempty"`
	Title           string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Location        []string `protobuf:"bytes,3,rep,name=location,proto3" json:"location,omitempty"`
	MinSalary       int32    `protobuf:"varint,4,opt,name=minSalary,proto3" json:"minSalary,omitempty"`
	MaxSalary       int32    `protobuf:"varint,5,opt,name=maxSalary,proto3" json:"maxSalary,omitempty"`
	Currency        string   `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	CreateTime      int64    `protobuf:"varint,7,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime      int64    `protobuf:"varint,8,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	JobFunction     string   `protobuf:"bytes,9,opt,name=jobFunction,proto3" json:"jobFunction,omitempty"`
	JobType         string   `protobuf:"bytes,10,opt,name=jobType,proto3" json:"jobType,omitempty"`
	Experience      string   `protobuf:"bytes,11,opt,name=experience,proto3" json:"experience,omitempty"`
	Vacancies       int32    `protobuf:"varint,12,opt,name=vacancies,proto3" json:"vacancies,omitempty"`
	Description     string   `protobuf:"bytes,13,opt,name=description,proto3" json:"description,omitempty"`
	RequiredSkills  []string `protobuf:"bytes,14,rep,name=requiredSkills,proto3" json:"requiredSkills,omitempty"`
	Culture         string   `protobuf:"bytes,15,opt,name=culture,proto3" json:"culture,omitempty"`
	CompanyCode     string   `protobuf:"bytes,16,opt,name=companyCode,proto3" json:"companyCode,omitempty"`
	JobCode         string   `protobuf:"bytes,17,opt,name=jobCode,proto3" json:"jobCode,omitempty"`
	SubType         []string `protobuf:"bytes,18,rep,name=subType,proto3" json:"subType,omitempty"`
	LastUpdatedTime int64    `protobuf:"varint,19,opt,name=lastUpdatedTime,proto3" json:"lastUpdatedTime,omitempty"`
	Tags            []string `protobuf:"bytes,20,rep,name=tags,proto3" json:"tags,omitempty"`
	Err             string   `protobuf:"bytes,21,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *TechInAsiaJob) Reset() {
	*x = TechInAsiaJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techinasia_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TechInAsiaJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TechInAsiaJob) ProtoMessage() {}

func (x *TechInAsiaJob) ProtoReflect() protoreflect.Message {
	mi := &file_techinasia_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TechInAsiaJob.ProtoReflect.Descriptor instead.
func (*TechInAsiaJob) Descriptor() ([]byte, []int) {
	return file_techinasia_proto_rawDescGZIP(), []int{1}
}

func (x *TechInAsiaJob) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *TechInAsiaJob) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TechInAsiaJob) GetLocation() []string {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *TechInAsiaJob) GetMinSalary() int32 {
	if x != nil {
		return x.MinSalary
	}
	return 0
}

func (x *TechInAsiaJob) GetMaxSalary() int32 {
	if x != nil {
		return x.MaxSalary
	}
	return 0
}

func (x *TechInAsiaJob) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TechInAsiaJob) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *TechInAsiaJob) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *TechInAsiaJob) GetJobFunction() string {
	if x != nil {
		return x.JobFunction
	}
	return ""
}

func (x *TechInAsiaJob) GetJobType() string {
	if x != nil {
		return x.JobType
	}
	return ""
}

func (x *TechInAsiaJob) GetExperience() string {
	if x != nil {
		return x.Experience
	}
	return ""
}

func (x *TechInAsiaJob) GetVacancies() int32 {
	if x != nil {
		return x.Vacancies
	}
	return 0
}

func (x *TechInAsiaJob) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TechInAsiaJob) GetRequiredSkills() []string {
	if x != nil {
		return x.RequiredSkills
	}
	return nil
}

func (x *TechInAsiaJob) GetCulture() string {
	if x != nil {
		return x.Culture
	}
	return ""
}

func (x *TechInAsiaJob) GetCompanyCode() string {
	if x != nil {
		return x.CompanyCode
	}
	return ""
}

func (x *TechInAsiaJob) GetJobCode() string {
	if x != nil {
		return x.JobCode
	}
	return ""
}

func (x *TechInAsiaJob) GetSubType() []string {
	if x != nil {
		return x.SubType
	}
	return nil
}

func (x *TechInAsiaJob) GetLastUpdatedTime() int64 {
	if x != nil {
		return x.LastUpdatedTime
	}
	return 0
}

func (x *TechInAsiaJob) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *TechInAsiaJob) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

// TechInAsiaJobTag - TechInAsia Job tag
type TechInAsiaJobTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag     string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	SubTags []string `protobuf:"bytes,2,rep,name=subTags,proto3" json:"subTags,omitempty"`
}

func (x *TechInAsiaJobTag) Reset() {
	*x = TechInAsiaJobTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techinasia_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TechInAsiaJobTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TechInAsiaJobTag) ProtoMessage() {}

func (x *TechInAsiaJobTag) ProtoReflect() protoreflect.Message {
	mi := &file_techinasia_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TechInAsiaJobTag.ProtoReflect.Descriptor instead.
func (*TechInAsiaJobTag) Descriptor() ([]byte, []int) {
	return file_techinasia_proto_rawDescGZIP(), []int{2}
}

func (x *TechInAsiaJobTag) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *TechInAsiaJobTag) GetSubTags() []string {
	if x != nil {
		return x.SubTags
	}
	return nil
}

// TechInAsiaJobTagList - TechInAsia Job tag list
type TechInAsiaJobTagList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []*TechInAsiaJobTag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *TechInAsiaJobTagList) Reset() {
	*x = TechInAsiaJobTagList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techinasia_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TechInAsiaJobTagList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TechInAsiaJobTagList) ProtoMessage() {}

func (x *TechInAsiaJobTagList) ProtoReflect() protoreflect.Message {
	mi := &file_techinasia_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TechInAsiaJobTagList.ProtoReflect.Descriptor instead.
func (*TechInAsiaJobTagList) Descriptor() ([]byte, []int) {
	return file_techinasia_proto_rawDescGZIP(), []int{3}
}

func (x *TechInAsiaJobTagList) GetTags() []*TechInAsiaJobTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

// TechInAsiaJobList - TechInAsia JobList
type TechInAsiaJobList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*TechInAsiaJob `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	Tags []string         `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *TechInAsiaJobList) Reset() {
	*x = TechInAsiaJobList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techinasia_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TechInAsiaJobList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TechInAsiaJobList) ProtoMessage() {}

func (x *TechInAsiaJobList) ProtoReflect() protoreflect.Message {
	mi := &file_techinasia_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TechInAsiaJobList.ProtoReflect.Descriptor instead.
func (*TechInAsiaJobList) Descriptor() ([]byte, []int) {
	return file_techinasia_proto_rawDescGZIP(), []int{4}
}

func (x *TechInAsiaJobList) GetJobs() []*TechInAsiaJob {
	if x != nil {
		return x.Jobs
	}
	return nil
}

func (x *TechInAsiaJobList) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// RequestTechInAsia - request techinasia
type RequestTechInAsia struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode        TechInAsiaMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.TechInAsiaMode" json:"mode,omitempty"`
	CompanyCode string         `protobuf:"bytes,2,opt,name=companyCode,proto3" json:"companyCode,omitempty"`
	JobCode     string         `protobuf:"bytes,3,opt,name=jobCode,proto3" json:"jobCode,omitempty"`
	JobNums     int32          `protobuf:"varint,4,opt,name=jobNums,proto3" json:"jobNums,omitempty"`
	JobTag      string         `protobuf:"bytes,5,opt,name=jobTag,proto3" json:"jobTag,omitempty"`
	JobSubTag   string         `protobuf:"bytes,6,opt,name=jobSubTag,proto3" json:"jobSubTag,omitempty"`
}

func (x *RequestTechInAsia) Reset() {
	*x = RequestTechInAsia{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techinasia_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestTechInAsia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestTechInAsia) ProtoMessage() {}

func (x *RequestTechInAsia) ProtoReflect() protoreflect.Message {
	mi := &file_techinasia_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestTechInAsia.ProtoReflect.Descriptor instead.
func (*RequestTechInAsia) Descriptor() ([]byte, []int) {
	return file_techinasia_proto_rawDescGZIP(), []int{5}
}

func (x *RequestTechInAsia) GetMode() TechInAsiaMode {
	if x != nil {
		return x.Mode
	}
	return TechInAsiaMode_TIAM_COMPANY
}

func (x *RequestTechInAsia) GetCompanyCode() string {
	if x != nil {
		return x.CompanyCode
	}
	return ""
}

func (x *RequestTechInAsia) GetJobCode() string {
	if x != nil {
		return x.JobCode
	}
	return ""
}

func (x *RequestTechInAsia) GetJobNums() int32 {
	if x != nil {
		return x.JobNums
	}
	return 0
}

func (x *RequestTechInAsia) GetJobTag() string {
	if x != nil {
		return x.JobTag
	}
	return ""
}

func (x *RequestTechInAsia) GetJobSubTag() string {
	if x != nil {
		return x.JobSubTag
	}
	return ""
}

// ReplyTechInAsia - reply techinasia
type ReplyTechInAsia struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode TechInAsiaMode `protobuf:"varint,1,opt,name=mode,proto3,enum=jarviscrawlercore.TechInAsiaMode" json:"mode,omitempty"`
	// Types that are assignable to Reply:
	//	*ReplyTechInAsia_Company
	//	*ReplyTechInAsia_Job
	//	*ReplyTechInAsia_Jobs
	//	*ReplyTechInAsia_Tags
	Reply isReplyTechInAsia_Reply `protobuf_oneof:"reply"`
}

func (x *ReplyTechInAsia) Reset() {
	*x = ReplyTechInAsia{}
	if protoimpl.UnsafeEnabled {
		mi := &file_techinasia_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyTechInAsia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyTechInAsia) ProtoMessage() {}

func (x *ReplyTechInAsia) ProtoReflect() protoreflect.Message {
	mi := &file_techinasia_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyTechInAsia.ProtoReflect.Descriptor instead.
func (*ReplyTechInAsia) Descriptor() ([]byte, []int) {
	return file_techinasia_proto_rawDescGZIP(), []int{6}
}

func (x *ReplyTechInAsia) GetMode() TechInAsiaMode {
	if x != nil {
		return x.Mode
	}
	return TechInAsiaMode_TIAM_COMPANY
}

func (m *ReplyTechInAsia) GetReply() isReplyTechInAsia_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (x *ReplyTechInAsia) GetCompany() *TechInAsiaCompany {
	if x, ok := x.GetReply().(*ReplyTechInAsia_Company); ok {
		return x.Company
	}
	return nil
}

func (x *ReplyTechInAsia) GetJob() *TechInAsiaJob {
	if x, ok := x.GetReply().(*ReplyTechInAsia_Job); ok {
		return x.Job
	}
	return nil
}

func (x *ReplyTechInAsia) GetJobs() *TechInAsiaJobList {
	if x, ok := x.GetReply().(*ReplyTechInAsia_Jobs); ok {
		return x.Jobs
	}
	return nil
}

func (x *ReplyTechInAsia) GetTags() *TechInAsiaJobTagList {
	if x, ok := x.GetReply().(*ReplyTechInAsia_Tags); ok {
		return x.Tags
	}
	return nil
}

type isReplyTechInAsia_Reply interface {
	isReplyTechInAsia_Reply()
}

type ReplyTechInAsia_Company struct {
	Company *TechInAsiaCompany `protobuf:"bytes,100,opt,name=company,proto3,oneof"`
}

type ReplyTechInAsia_Job struct {
	Job *TechInAsiaJob `protobuf:"bytes,101,opt,name=job,proto3,oneof"`
}

type ReplyTechInAsia_Jobs struct {
	Jobs *TechInAsiaJobList `protobuf:"bytes,102,opt,name=jobs,proto3,oneof"`
}

type ReplyTechInAsia_Tags struct {
	Tags *TechInAsiaJobTagList `protobuf:"bytes,103,opt,name=tags,proto3,oneof"`
}

func (*ReplyTechInAsia_Company) isReplyTechInAsia_Reply() {}

func (*ReplyTechInAsia_Job) isReplyTechInAsia_Reply() {}

func (*ReplyTechInAsia_Jobs) isReplyTechInAsia_Reply() {}

func (*ReplyTechInAsia_Tags) isReplyTechInAsia_Reply() {}

var File_techinasia_proto protoreflect.FileDescriptor

var file_techinasia_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x65, 0x63, 0x68, 0x69, 0x6e, 0x61, 0x73, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65,
	0x72, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x9f, 0x02, 0x0a, 0x11, 0x54, 0x65, 0x63, 0x68, 0x49, 0x6e,
	0x41, 0x73, 0x69, 0x61, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xff, 0x04, 0x0a, 0x0d, 0x54, 0x65, 0x63, 0x68,
	0x49, 0x6e, 0x41, 0x73, 0x69, 0x61, 0x4a, 0x6f, 0x62, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x6d, 0x69, 0x6e, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x6d, 0x69, 0x6e, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6d,
	0x61, 0x78, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6d, 0x61, 0x78, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x6f, 0x62, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x6c,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x75, 0x6c, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x3e, 0x0a, 0x10, 0x54, 0x65, 0x63,
	0x68, 0x49, 0x6e, 0x41, 0x73, 0x69, 0x61, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x54, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x54, 0x61, 0x67, 0x73, 0x22, 0x4f, 0x0a, 0x14, 0x54, 0x65, 0x63,
	0x68, 0x49, 0x6e, 0x41, 0x73, 0x69, 0x61, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x54, 0x65, 0x63, 0x68, 0x49, 0x6e, 0x41, 0x73, 0x69, 0x61, 0x4a, 0x6f,
	0x62, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x5d, 0x0a, 0x11, 0x54, 0x65,
	0x63, 0x68, 0x49, 0x6e, 0x41, 0x73, 0x69, 0x61, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x54, 0x65, 0x63, 0x68, 0x49, 0x6e, 0x41, 0x73, 0x69, 0x61, 0x4a, 0x6f, 0x62, 0x52,
	0x04, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x11, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x65, 0x63, 0x68, 0x49, 0x6e, 0x41, 0x73, 0x69, 0x61, 0x12,
	0x35, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e,
	0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x54, 0x65, 0x63, 0x68, 0x49, 0x6e, 0x41, 0x73, 0x69, 0x61, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x75, 0x6d, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x6a, 0x6f, 0x62, 0x54, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6a, 0x6f,
	0x62, 0x54, 0x61, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x53, 0x75, 0x62, 0x54, 0x61,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x53, 0x75, 0x62, 0x54,
	0x61, 0x67, 0x22, 0xc4, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x65, 0x63, 0x68,
	0x49, 0x6e, 0x41, 0x73, 0x69, 0x61, 0x12, 0x35, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61,
	0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x65, 0x63, 0x68, 0x49, 0x6e, 0x41,
	0x73, 0x69, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x54, 0x65, 0x63, 0x68, 0x49, 0x6e, 0x41, 0x73, 0x69, 0x61, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12,
	0x34, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6a,
	0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x54, 0x65, 0x63, 0x68, 0x49, 0x6e, 0x41, 0x73, 0x69, 0x61, 0x4a, 0x6f, 0x62, 0x48, 0x00,
	0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x3a, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x66, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77,
	0x6c, 0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x65, 0x63, 0x68, 0x49, 0x6e, 0x41, 0x73,
	0x69, 0x61, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x6a, 0x6f, 0x62,
	0x73, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x54, 0x65, 0x63, 0x68, 0x49, 0x6e, 0x41, 0x73, 0x69, 0x61, 0x4a, 0x6f,
	0x62, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x42, 0x07, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2a, 0x53, 0x0a, 0x0e, 0x54, 0x65, 0x63,
	0x68, 0x49, 0x6e, 0x41, 0x73, 0x69, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x54,
	0x49, 0x41, 0x4d, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x4e, 0x59, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x54, 0x49, 0x41, 0x4d, 0x5f, 0x4a, 0x4f, 0x42, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54,
	0x49, 0x41, 0x4d, 0x5f, 0x4a, 0x4f, 0x42, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x02, 0x12, 0x0f, 0x0a,
	0x0b, 0x54, 0x49, 0x41, 0x4d, 0x5f, 0x4a, 0x4f, 0x42, 0x54, 0x41, 0x47, 0x10, 0x03, 0x42, 0x20,
	0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x73,
	0x30, 0x30, 0x37, 0x2f, 0x6a, 0x63, 0x63, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_techinasia_proto_rawDescOnce sync.Once
	file_techinasia_proto_rawDescData = file_techinasia_proto_rawDesc
)

func file_techinasia_proto_rawDescGZIP() []byte {
	file_techinasia_proto_rawDescOnce.Do(func() {
		file_techinasia_proto_rawDescData = protoimpl.X.CompressGZIP(file_techinasia_proto_rawDescData)
	})
	return file_techinasia_proto_rawDescData
}

var file_techinasia_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_techinasia_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_techinasia_proto_goTypes = []interface{}{
	(TechInAsiaMode)(0),          // 0: jarviscrawlercore.TechInAsiaMode
	(*TechInAsiaCompany)(nil),    // 1: jarviscrawlercore.TechInAsiaCompany
	(*TechInAsiaJob)(nil),        // 2: jarviscrawlercore.TechInAsiaJob
	(*TechInAsiaJobTag)(nil),     // 3: jarviscrawlercore.TechInAsiaJobTag
	(*TechInAsiaJobTagList)(nil), // 4: jarviscrawlercore.TechInAsiaJobTagList
	(*TechInAsiaJobList)(nil),    // 5: jarviscrawlercore.TechInAsiaJobList
	(*RequestTechInAsia)(nil),    // 6: jarviscrawlercore.RequestTechInAsia
	(*ReplyTechInAsia)(nil),      // 7: jarviscrawlercore.ReplyTechInAsia
}
var file_techinasia_proto_depIdxs = []int32{
	3, // 0: jarviscrawlercore.TechInAsiaJobTagList.tags:type_name -> jarviscrawlercore.TechInAsiaJobTag
	2, // 1: jarviscrawlercore.TechInAsiaJobList.jobs:type_name -> jarviscrawlercore.TechInAsiaJob
	0, // 2: jarviscrawlercore.RequestTechInAsia.mode:type_name -> jarviscrawlercore.TechInAsiaMode
	0, // 3: jarviscrawlercore.ReplyTechInAsia.mode:type_name -> jarviscrawlercore.TechInAsiaMode
	1, // 4: jarviscrawlercore.ReplyTechInAsia.company:type_name -> jarviscrawlercore.TechInAsiaCompany
	2, // 5: jarviscrawlercore.ReplyTechInAsia.job:type_name -> jarviscrawlercore.TechInAsiaJob
	5, // 6: jarviscrawlercore.ReplyTechInAsia.jobs:type_name -> jarviscrawlercore.TechInAsiaJobList
	4, // 7: jarviscrawlercore.ReplyTechInAsia.tags:type_name -> jarviscrawlercore.TechInAsiaJobTagList
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_techinasia_proto_init() }
func file_techinasia_proto_init() {
	if File_techinasia_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_techinasia_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TechInAsiaCompany); i {
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
		file_techinasia_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TechInAsiaJob); i {
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
		file_techinasia_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TechInAsiaJobTag); i {
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
		file_techinasia_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TechInAsiaJobTagList); i {
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
		file_techinasia_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TechInAsiaJobList); i {
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
		file_techinasia_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestTechInAsia); i {
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
		file_techinasia_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyTechInAsia); i {
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
	file_techinasia_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*ReplyTechInAsia_Company)(nil),
		(*ReplyTechInAsia_Job)(nil),
		(*ReplyTechInAsia_Jobs)(nil),
		(*ReplyTechInAsia_Tags)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_techinasia_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_techinasia_proto_goTypes,
		DependencyIndexes: file_techinasia_proto_depIdxs,
		EnumInfos:         file_techinasia_proto_enumTypes,
		MessageInfos:      file_techinasia_proto_msgTypes,
	}.Build()
	File_techinasia_proto = out.File
	file_techinasia_proto_rawDesc = nil
	file_techinasia_proto_goTypes = nil
	file_techinasia_proto_depIdxs = nil
}
