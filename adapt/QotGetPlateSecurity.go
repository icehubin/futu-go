package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetplatesecurity"
	"google.golang.org/protobuf/proto"
)

type QotGetPlateSecurity struct {
	request *qotgetplatesecurity.Request

	adaptBase
}

func CreateQotGetPlateSecurity(dopts ...Option) AdaptInterface {
	adp := &QotGetPlateSecurity{
		request: &qotgetplatesecurity.Request{
			C2S: &qotgetplatesecurity.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetPlateSecurity)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetPlateSecurity) SetC2SOption(protoKey string, val interface{}) {
	/*
		Plate     *qotcommon.Security `protobuf:"bytes,1,req,name=plate" json:"plate,omitempty"`          //板块
		SortField *int32              `protobuf:"varint,2,opt,name=sortField" json:"sortField,omitempty"` //Qot_Common.SortField,根据哪个字段排序,不填默认Code排序
		Ascend    *bool               `protobuf:"varint,3,opt,name=ascend" json:"ascend,omitempty"`       //升序ture, 降序false, 不填默认升序
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Plate"), strings.ToUpper("code"):
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Plate = nv
		}
	case strings.ToUpper("SortField"), strings.ToUpper("Sort"):
		if v, ok := val.(int32); ok {
			a.request.C2S.SortField = proto.Int32(v)
		}
	case strings.ToUpper("Ascend"), strings.ToUpper("Asc"):
		if v, ok := val.(bool); ok {
			a.request.C2S.Ascend = proto.Bool(v)
		}
	}
}

//=== no need to modify
func (a *QotGetPlateSecurity) UnPackBody(body []byte) Response {
	rsp := &qotgetplatesecurity.Response{}
	err := proto.Unmarshal(body, rsp)
	if err != nil {
		return PackErr()
	}
	return Response{
		RetType: rsp.GetRetType(),
		RetMsg:  rsp.GetRetMsg(),
		ErrCode: rsp.GetErrCode(),
		S2C:     rsp.GetS2C(),
	}
}
func (a *QotGetPlateSecurity) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetPlateSecurity) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
