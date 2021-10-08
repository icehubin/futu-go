package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetsecuritysnapshot"
	"google.golang.org/protobuf/proto"
)

type QotGetSecuritySnapshot struct {
	request *qotgetsecuritysnapshot.Request

	adaptBase
}

func CreateQotGetSecuritySnapshot(dopts ...Option) AdaptInterface {
	adp := &QotGetSecuritySnapshot{
		request: &qotgetsecuritysnapshot.Request{
			C2S: &qotgetsecuritysnapshot.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetSecuritySnapshot)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetSecuritySnapshot) SetC2SOption(protoKey string, val interface{}) {
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
func (a *QotGetSecuritySnapshot) UnPackBody(body []byte) Response {
	rsp := &qotgetsecuritysnapshot.Response{}
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
func (a *QotGetSecuritySnapshot) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetSecuritySnapshot) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
