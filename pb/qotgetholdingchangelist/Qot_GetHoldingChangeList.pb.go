// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: Qot_GetHoldingChangeList.proto

package qotgetholdingchangelist

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

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Security       *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`              //股票
	HolderCategory *int32              `protobuf:"varint,2,req,name=holderCategory" json:"holderCategory,omitempty"` //持有者类别（1机构、2基金、3高管）
	//以下是发布时间筛选，不传返回所有数据，传了返回发布时间属于开始时间到结束时间段内的数据
	BeginTime *string `protobuf:"bytes,3,opt,name=beginTime" json:"beginTime,omitempty"` //开始时间，严格按YYYY-MM-DD HH:MM:SS或YYYY-MM-DD HH:MM:SS.MS格式传
	EndTime   *string `protobuf:"bytes,4,opt,name=endTime" json:"endTime,omitempty"`     //结束时间，严格按YYYY-MM-DD HH:MM:SS或YYYY-MM-DD HH:MM:SS.MS格式传
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetHoldingChangeList_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHoldingChangeList_proto_msgTypes[0]
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
	return file_Qot_GetHoldingChangeList_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *C2S) GetHolderCategory() int32 {
	if x != nil && x.HolderCategory != nil {
		return *x.HolderCategory
	}
	return 0
}

func (x *C2S) GetBeginTime() string {
	if x != nil && x.BeginTime != nil {
		return *x.BeginTime
	}
	return ""
}

func (x *C2S) GetEndTime() string {
	if x != nil && x.EndTime != nil {
		return *x.EndTime
	}
	return ""
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Security          *qotcommon.Security             `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`                   //股票
	HoldingChangeList []*qotcommon.ShareHoldingChange `protobuf:"bytes,2,rep,name=holdingChangeList" json:"holdingChangeList,omitempty"` //对应类别的持股变化列表（最多返回前100大股东的变化）
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetHoldingChangeList_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHoldingChangeList_proto_msgTypes[1]
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
	return file_Qot_GetHoldingChangeList_proto_rawDescGZIP(), []int{1}
}

func (x *S2C) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *S2C) GetHoldingChangeList() []*qotcommon.ShareHoldingChange {
	if x != nil {
		return x.HoldingChangeList
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
		mi := &file_Qot_GetHoldingChangeList_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHoldingChangeList_proto_msgTypes[2]
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
	return file_Qot_GetHoldingChangeList_proto_rawDescGZIP(), []int{2}
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
		mi := &file_Qot_GetHoldingChangeList_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHoldingChangeList_proto_msgTypes[3]
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
	return file_Qot_GetHoldingChangeList_proto_rawDescGZIP(), []int{3}
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

var File_Qot_GetHoldingChangeList_proto protoreflect.FileDescriptor

var file_Qot_GetHoldingChangeList_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x03, 0x43,
	0x32, 0x53, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0e, 0x68, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x03, 0x53, 0x32, 0x43, 0x12, 0x30, 0x0a, 0x08,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x4c,
	0x0a, 0x11, 0x68, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x51, 0x6f, 0x74, 0x5f,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x48, 0x6f, 0x6c, 0x64,
	0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x11, 0x68, 0x6f, 0x6c, 0x64, 0x69,
	0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x03, 0x63, 0x32, 0x73, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x48, 0x6f,
	0x6c, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e,
	0x43, 0x32, 0x53, 0x52, 0x03, 0x63, 0x32, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x48, 0x6f,
	0x6c, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e,
	0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x42, 0x4d, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x63, 0x65, 0x68,
	0x75, 0x62, 0x69, 0x6e, 0x2f, 0x66, 0x75, 0x74, 0x75, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x71, 0x6f, 0x74, 0x67, 0x65, 0x74, 0x68, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74,
}

var (
	file_Qot_GetHoldingChangeList_proto_rawDescOnce sync.Once
	file_Qot_GetHoldingChangeList_proto_rawDescData = file_Qot_GetHoldingChangeList_proto_rawDesc
)

func file_Qot_GetHoldingChangeList_proto_rawDescGZIP() []byte {
	file_Qot_GetHoldingChangeList_proto_rawDescOnce.Do(func() {
		file_Qot_GetHoldingChangeList_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_GetHoldingChangeList_proto_rawDescData)
	})
	return file_Qot_GetHoldingChangeList_proto_rawDescData
}

var file_Qot_GetHoldingChangeList_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Qot_GetHoldingChangeList_proto_goTypes = []interface{}{
	(*C2S)(nil),                          // 0: Qot_GetHoldingChangeList.C2S
	(*S2C)(nil),                          // 1: Qot_GetHoldingChangeList.S2C
	(*Request)(nil),                      // 2: Qot_GetHoldingChangeList.Request
	(*Response)(nil),                     // 3: Qot_GetHoldingChangeList.Response
	(*qotcommon.Security)(nil),           // 4: Qot_Common.Security
	(*qotcommon.ShareHoldingChange)(nil), // 5: Qot_Common.ShareHoldingChange
}
var file_Qot_GetHoldingChangeList_proto_depIdxs = []int32{
	4, // 0: Qot_GetHoldingChangeList.C2S.security:type_name -> Qot_Common.Security
	4, // 1: Qot_GetHoldingChangeList.S2C.security:type_name -> Qot_Common.Security
	5, // 2: Qot_GetHoldingChangeList.S2C.holdingChangeList:type_name -> Qot_Common.ShareHoldingChange
	0, // 3: Qot_GetHoldingChangeList.Request.c2s:type_name -> Qot_GetHoldingChangeList.C2S
	1, // 4: Qot_GetHoldingChangeList.Response.s2c:type_name -> Qot_GetHoldingChangeList.S2C
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_Qot_GetHoldingChangeList_proto_init() }
func file_Qot_GetHoldingChangeList_proto_init() {
	if File_Qot_GetHoldingChangeList_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_GetHoldingChangeList_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Qot_GetHoldingChangeList_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_Qot_GetHoldingChangeList_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_Qot_GetHoldingChangeList_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_Qot_GetHoldingChangeList_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_GetHoldingChangeList_proto_goTypes,
		DependencyIndexes: file_Qot_GetHoldingChangeList_proto_depIdxs,
		MessageInfos:      file_Qot_GetHoldingChangeList_proto_msgTypes,
	}.Build()
	File_Qot_GetHoldingChangeList_proto = out.File
	file_Qot_GetHoldingChangeList_proto_rawDesc = nil
	file_Qot_GetHoldingChangeList_proto_goTypes = nil
	file_Qot_GetHoldingChangeList_proto_depIdxs = nil
}
