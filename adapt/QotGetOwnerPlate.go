package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetownerplate"
	"google.golang.org/protobuf/proto"
)

type QotGetOwnerPlate struct {
	request *qotgetownerplate.Request

	adaptBase
}

func CreateQotGetOwnerPlate(dopts ...Option) AdaptInterface {
	adp := &QotGetOwnerPlate{
		request: &qotgetownerplate.Request{
			C2S: &qotgetownerplate.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetOwnerPlate)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetOwnerPlate) SetC2SOption(protoKey string, val interface{}) {
	/*
		SecurityList []*qotcommon.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"` //股票
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("SecurityList"), strings.ToUpper("code_list"):
		if v, ok := val.([]string); ok {
			nv := StocksToSecurity(v)
			a.request.C2S.SecurityList = nv
		}
	}
}

//=== no need to modify
func (a *QotGetOwnerPlate) UnPackBody(body []byte) Response {
	rsp := &qotgetownerplate.Response{}
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
func (a *QotGetOwnerPlate) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetOwnerPlate) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
