// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: Qot_GetCodeChange.proto

package qotgetcodechange

import (
	_ "github.com/icehubin/futu-go/pb/common"
	qotcommon "github.com/icehubin/futu-go/pb/qotcommon"
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

type CodeChangeType int32

const (
	CodeChangeType_CodeChangeType_Unkown     CodeChangeType = 0 //未知
	CodeChangeType_CodeChangeType_GemToMain  CodeChangeType = 1 //创业板转主板
	CodeChangeType_CodeChangeType_Unpaid     CodeChangeType = 2 //买卖未缴款供股权
	CodeChangeType_CodeChangeType_ChangeLot  CodeChangeType = 3 //更改买卖单位
	CodeChangeType_CodeChangeType_Split      CodeChangeType = 4 //拆股
	CodeChangeType_CodeChangeType_Joint      CodeChangeType = 5 //合股
	CodeChangeType_CodeChangeType_JointSplit CodeChangeType = 6 //股份先并后拆
	CodeChangeType_CodeChangeType_SplitJoint CodeChangeType = 7 //股份先拆后并
	CodeChangeType_CodeChangeType_Other      CodeChangeType = 8 //其他
)

// Enum value maps for CodeChangeType.
var (
	CodeChangeType_name = map[int32]string{
		0: "CodeChangeType_Unkown",
		1: "CodeChangeType_GemToMain",
		2: "CodeChangeType_Unpaid",
		3: "CodeChangeType_ChangeLot",
		4: "CodeChangeType_Split",
		5: "CodeChangeType_Joint",
		6: "CodeChangeType_JointSplit",
		7: "CodeChangeType_SplitJoint",
		8: "CodeChangeType_Other",
	}
	CodeChangeType_value = map[string]int32{
		"CodeChangeType_Unkown":     0,
		"CodeChangeType_GemToMain":  1,
		"CodeChangeType_Unpaid":     2,
		"CodeChangeType_ChangeLot":  3,
		"CodeChangeType_Split":      4,
		"CodeChangeType_Joint":      5,
		"CodeChangeType_JointSplit": 6,
		"CodeChangeType_SplitJoint": 7,
		"CodeChangeType_Other":      8,
	}
)

func (x CodeChangeType) Enum() *CodeChangeType {
	p := new(CodeChangeType)
	*p = x
	return p
}

func (x CodeChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodeChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_Qot_GetCodeChange_proto_enumTypes[0].Descriptor()
}

func (CodeChangeType) Type() protoreflect.EnumType {
	return &file_Qot_GetCodeChange_proto_enumTypes[0]
}

func (x CodeChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CodeChangeType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CodeChangeType(num)
	return nil
}

// Deprecated: Use CodeChangeType.Descriptor instead.
func (CodeChangeType) EnumDescriptor() ([]byte, []int) {
	return file_Qot_GetCodeChange_proto_rawDescGZIP(), []int{0}
}

type TimeFilterType int32

const (
	TimeFilterType_TimeFilterType_Unknow    TimeFilterType = 0
	TimeFilterType_TimeFilterType_Public    TimeFilterType = 1 //根据公布时间过滤
	TimeFilterType_TimeFilterType_Effective TimeFilterType = 2 //根据生效时间过滤
	TimeFilterType_TimeFilterType_End       TimeFilterType = 3 //根据结束时间过滤
)

// Enum value maps for TimeFilterType.
var (
	TimeFilterType_name = map[int32]string{
		0: "TimeFilterType_Unknow",
		1: "TimeFilterType_Public",
		2: "TimeFilterType_Effective",
		3: "TimeFilterType_End",
	}
	TimeFilterType_value = map[string]int32{
		"TimeFilterType_Unknow":    0,
		"TimeFilterType_Public":    1,
		"TimeFilterType_Effective": 2,
		"TimeFilterType_End":       3,
	}
)

func (x TimeFilterType) Enum() *TimeFilterType {
	p := new(TimeFilterType)
	*p = x
	return p
}

func (x TimeFilterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeFilterType) Descriptor() protoreflect.EnumDescriptor {
	return file_Qot_GetCodeChange_proto_enumTypes[1].Descriptor()
}

func (TimeFilterType) Type() protoreflect.EnumType {
	return &file_Qot_GetCodeChange_proto_enumTypes[1]
}

func (x TimeFilterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TimeFilterType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TimeFilterType(num)
	return nil
}

// Deprecated: Use TimeFilterType.Descriptor instead.
func (TimeFilterType) EnumDescriptor() ([]byte, []int) {
	return file_Qot_GetCodeChange_proto_rawDescGZIP(), []int{1}
}

type CodeChangeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type               *int32              `protobuf:"varint,1,req,name=type" json:"type,omitempty"`                              //CodeChangeType,代码变化或者新增临时代码的事件类型
	Security           *qotcommon.Security `protobuf:"bytes,2,req,name=security" json:"security,omitempty"`                       //主代码，在创业板转主板中表示主板
	RelatedSecurity    *qotcommon.Security `protobuf:"bytes,3,req,name=relatedSecurity" json:"relatedSecurity,omitempty"`         //关联代码，在创业板转主板中表示创业板，在剩余事件中表示临时代码
	PublicTime         *string             `protobuf:"bytes,4,opt,name=publicTime" json:"publicTime,omitempty"`                   //公布时间
	PublicTimestamp    *float64            `protobuf:"fixed64,5,opt,name=publicTimestamp" json:"publicTimestamp,omitempty"`       //公布时间戳
	EffectiveTime      *string             `protobuf:"bytes,6,opt,name=effectiveTime" json:"effectiveTime,omitempty"`             //生效时间
	EffectiveTimestamp *float64            `protobuf:"fixed64,7,opt,name=effectiveTimestamp" json:"effectiveTimestamp,omitempty"` //生效时间戳
	EndTime            *string             `protobuf:"bytes,8,opt,name=endTime" json:"endTime,omitempty"`                         //结束时间，在创业板转主板事件不存在该字段，在剩余事件表示临时代码交易结束时间
	EndTimestamp       *float64            `protobuf:"fixed64,9,opt,name=endTimestamp" json:"endTimestamp,omitempty"`             //结束时间戳，在创业板转主板事件不存在该字段，在剩余事件表示临时代码交易结束时间
}

func (x *CodeChangeInfo) Reset() {
	*x = CodeChangeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetCodeChange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeChangeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeChangeInfo) ProtoMessage() {}

func (x *CodeChangeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetCodeChange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeChangeInfo.ProtoReflect.Descriptor instead.
func (*CodeChangeInfo) Descriptor() ([]byte, []int) {
	return file_Qot_GetCodeChange_proto_rawDescGZIP(), []int{0}
}

func (x *CodeChangeInfo) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *CodeChangeInfo) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *CodeChangeInfo) GetRelatedSecurity() *qotcommon.Security {
	if x != nil {
		return x.RelatedSecurity
	}
	return nil
}

func (x *CodeChangeInfo) GetPublicTime() string {
	if x != nil && x.PublicTime != nil {
		return *x.PublicTime
	}
	return ""
}

func (x *CodeChangeInfo) GetPublicTimestamp() float64 {
	if x != nil && x.PublicTimestamp != nil {
		return *x.PublicTimestamp
	}
	return 0
}

func (x *CodeChangeInfo) GetEffectiveTime() string {
	if x != nil && x.EffectiveTime != nil {
		return *x.EffectiveTime
	}
	return ""
}

func (x *CodeChangeInfo) GetEffectiveTimestamp() float64 {
	if x != nil && x.EffectiveTimestamp != nil {
		return *x.EffectiveTimestamp
	}
	return 0
}

func (x *CodeChangeInfo) GetEndTime() string {
	if x != nil && x.EndTime != nil {
		return *x.EndTime
	}
	return ""
}

func (x *CodeChangeInfo) GetEndTimestamp() float64 {
	if x != nil && x.EndTimestamp != nil {
		return *x.EndTimestamp
	}
	return 0
}

type TimeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      *int32  `protobuf:"varint,1,req,name=type" json:"type,omitempty"`          //TimeFilterType, 过滤类型
	BeginTime *string `protobuf:"bytes,2,opt,name=beginTime" json:"beginTime,omitempty"` //开始时间点
	EndTime   *string `protobuf:"bytes,3,opt,name=endTime" json:"endTime,omitempty"`     //结束时间点
}

func (x *TimeFilter) Reset() {
	*x = TimeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetCodeChange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeFilter) ProtoMessage() {}

func (x *TimeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetCodeChange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeFilter.ProtoReflect.Descriptor instead.
func (*TimeFilter) Descriptor() ([]byte, []int) {
	return file_Qot_GetCodeChange_proto_rawDescGZIP(), []int{1}
}

func (x *TimeFilter) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *TimeFilter) GetBeginTime() string {
	if x != nil && x.BeginTime != nil {
		return *x.BeginTime
	}
	return ""
}

func (x *TimeFilter) GetEndTime() string {
	if x != nil && x.EndTime != nil {
		return *x.EndTime
	}
	return ""
}

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaceHolder    *int32                `protobuf:"varint,1,opt,name=placeHolder" json:"placeHolder,omitempty"`      //占位
	SecurityList   []*qotcommon.Security `protobuf:"bytes,2,rep,name=securityList" json:"securityList,omitempty"`     //根据股票筛选
	TimeFilterList []*TimeFilter         `protobuf:"bytes,3,rep,name=timeFilterList" json:"timeFilterList,omitempty"` //根据时间筛选
	TypeList       []int32               `protobuf:"varint,4,rep,name=typeList" json:"typeList,omitempty"`            //CodeChangeType，根据类型筛选
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetCodeChange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetCodeChange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_Qot_GetCodeChange_proto_rawDescGZIP(), []int{2}
}

func (x *C2S) GetPlaceHolder() int32 {
	if x != nil && x.PlaceHolder != nil {
		return *x.PlaceHolder
	}
	return 0
}

func (x *C2S) GetSecurityList() []*qotcommon.Security {
	if x != nil {
		return x.SecurityList
	}
	return nil
}

func (x *C2S) GetTimeFilterList() []*TimeFilter {
	if x != nil {
		return x.TimeFilterList
	}
	return nil
}

func (x *C2S) GetTypeList() []int32 {
	if x != nil {
		return x.TypeList
	}
	return nil
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodeChangeList []*CodeChangeInfo `protobuf:"bytes,1,rep,name=codeChangeList" json:"codeChangeList,omitempty"` //股票代码更换信息，目前仅有港股数据
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetCodeChange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetCodeChange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_Qot_GetCodeChange_proto_rawDescGZIP(), []int{3}
}

func (x *S2C) GetCodeChangeList() []*CodeChangeInfo {
	if x != nil {
		return x.CodeChangeList
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetCodeChange_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetCodeChange_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_Qot_GetCodeChange_proto_rawDescGZIP(), []int{4}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //RetType,返回结果
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetCodeChange_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetCodeChange_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_Qot_GetCodeChange_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_Qot_GetCodeChange_proto protoreflect.FileDescriptor

var file_Qot_GetCodeChange_proto_rawDesc = []byte{
	0x0a, 0x17, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x51, 0x6f, 0x74, 0x5f, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x0c, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x51, 0x6f, 0x74, 0x5f,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x02, 0x0a,
	0x0e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x3e, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x24, 0x0a, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x12, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x58, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xc4, 0x01,
	0x0a, 0x03, 0x43, 0x32, 0x53, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x48, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x45, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x51, 0x6f, 0x74, 0x5f,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x50, 0x0a, 0x03, 0x53, 0x32, 0x43, 0x12, 0x49, 0x0a, 0x0e, 0x63,
	0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x03, 0x63, 0x32, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63, 0x32, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x03, 0x73, 0x32,
	0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x53, 0x32, 0x43, 0x52,
	0x03, 0x73, 0x32, 0x63, 0x2a, 0x8e, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x6f, 0x64, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x5f, 0x47, 0x65, 0x6d, 0x54, 0x6f, 0x4d, 0x61, 0x69, 0x6e, 0x10, 0x01,
	0x12, 0x19, 0x0a, 0x15, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x5f, 0x55, 0x6e, 0x70, 0x61, 0x69, 0x64, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x43,
	0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x74, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x6f, 0x64,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x53, 0x70, 0x6c, 0x69,
	0x74, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x10, 0x05, 0x12, 0x1d, 0x0a,
	0x19, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f,
	0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x10, 0x06, 0x12, 0x1d, 0x0a, 0x19,
	0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x53,
	0x70, 0x6c, 0x69, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x43,
	0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4f, 0x74,
	0x68, 0x65, 0x72, 0x10, 0x08, 0x2a, 0x7c, 0x0a, 0x0e, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x69, 0x6d, 0x65, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x10, 0x01, 0x12, 0x1c, 0x0a,
	0x18, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x5f,
	0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x54,
	0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x45, 0x6e,
	0x64, 0x10, 0x03, 0x42, 0x46, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x63, 0x65, 0x68, 0x75, 0x62, 0x69, 0x6e, 0x2f,
	0x66, 0x75, 0x74, 0x75, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x71, 0x6f, 0x74, 0x67, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
}

var (
	file_Qot_GetCodeChange_proto_rawDescOnce sync.Once
	file_Qot_GetCodeChange_proto_rawDescData = file_Qot_GetCodeChange_proto_rawDesc
)

func file_Qot_GetCodeChange_proto_rawDescGZIP() []byte {
	file_Qot_GetCodeChange_proto_rawDescOnce.Do(func() {
		file_Qot_GetCodeChange_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_GetCodeChange_proto_rawDescData)
	})
	return file_Qot_GetCodeChange_proto_rawDescData
}

var file_Qot_GetCodeChange_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_Qot_GetCodeChange_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Qot_GetCodeChange_proto_goTypes = []interface{}{
	(CodeChangeType)(0),        // 0: Qot_GetCodeChange.CodeChangeType
	(TimeFilterType)(0),        // 1: Qot_GetCodeChange.TimeFilterType
	(*CodeChangeInfo)(nil),     // 2: Qot_GetCodeChange.CodeChangeInfo
	(*TimeFilter)(nil),         // 3: Qot_GetCodeChange.TimeFilter
	(*C2S)(nil),                // 4: Qot_GetCodeChange.C2S
	(*S2C)(nil),                // 5: Qot_GetCodeChange.S2C
	(*Request)(nil),            // 6: Qot_GetCodeChange.Request
	(*Response)(nil),           // 7: Qot_GetCodeChange.Response
	(*qotcommon.Security)(nil), // 8: Qot_Common.Security
}
var file_Qot_GetCodeChange_proto_depIdxs = []int32{
	8, // 0: Qot_GetCodeChange.CodeChangeInfo.security:type_name -> Qot_Common.Security
	8, // 1: Qot_GetCodeChange.CodeChangeInfo.relatedSecurity:type_name -> Qot_Common.Security
	8, // 2: Qot_GetCodeChange.C2S.securityList:type_name -> Qot_Common.Security
	3, // 3: Qot_GetCodeChange.C2S.timeFilterList:type_name -> Qot_GetCodeChange.TimeFilter
	2, // 4: Qot_GetCodeChange.S2C.codeChangeList:type_name -> Qot_GetCodeChange.CodeChangeInfo
	4, // 5: Qot_GetCodeChange.Request.c2s:type_name -> Qot_GetCodeChange.C2S
	5, // 6: Qot_GetCodeChange.Response.s2c:type_name -> Qot_GetCodeChange.S2C
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_Qot_GetCodeChange_proto_init() }
func file_Qot_GetCodeChange_proto_init() {
	if File_Qot_GetCodeChange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_GetCodeChange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeChangeInfo); i {
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
		file_Qot_GetCodeChange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeFilter); i {
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
		file_Qot_GetCodeChange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S); i {
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
		file_Qot_GetCodeChange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C); i {
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
		file_Qot_GetCodeChange_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_Qot_GetCodeChange_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_Qot_GetCodeChange_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_GetCodeChange_proto_goTypes,
		DependencyIndexes: file_Qot_GetCodeChange_proto_depIdxs,
		EnumInfos:         file_Qot_GetCodeChange_proto_enumTypes,
		MessageInfos:      file_Qot_GetCodeChange_proto_msgTypes,
	}.Build()
	File_Qot_GetCodeChange_proto = out.File
	file_Qot_GetCodeChange_proto_rawDesc = nil
	file_Qot_GetCodeChange_proto_goTypes = nil
	file_Qot_GetCodeChange_proto_depIdxs = nil
}
