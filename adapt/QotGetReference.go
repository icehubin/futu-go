package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetreference"
	"google.golang.org/protobuf/proto"
)

type QotGetReference struct {
	request *qotgetreference.Request

	adaptBase
}

func CreateQotGetReference(dopts ...Option) AdaptInterface {
	adp := &QotGetReference{
		request: &qotgetreference.Request{
			C2S: &qotgetreference.C2S{
				ReferenceType: proto.Int32(1),
			},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetReference)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetReference) SetC2SOption(protoKey string, val interface{}) {
	/*
		Security      *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`            //股票
		ReferenceType *int32              `protobuf:"varint,2,req,name=referenceType" json:"referenceType,omitempty"` // ReferenceType, 相关类型
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Security"), strings.ToUpper("code"):
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Security = nv
		}
	case strings.ToUpper("ReferenceType"), strings.ToUpper("Type"):
		if v, ok := val.(int32); ok {
			a.request.C2S.ReferenceType = proto.Int32(v)
		}
	}
}

//=== no need to modify
func (a *QotGetReference) UnPackBody(body []byte) Response {
	rsp := &qotgetreference.Response{}
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
func (a *QotGetReference) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetReference) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
