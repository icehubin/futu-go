//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetbasicqot"
	"google.golang.org/protobuf/proto"
)

type QotGetBasicQot struct {
	request *qotgetbasicqot.Request

	adaptBase
}

func CreateQotGetBasicQot(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotGetBasicQot{
		request: &qotgetbasicqot.Request{
			C2S: &qotgetbasicqot.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetBasicQot)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//
func (a *QotGetBasicQot) SetC2SOption(protoKey string, val interface{}) {
	//Todo fix Options or remove
	//SecurityList []*qotcommon.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"` //股票
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("SecurityList"), strings.ToUpper("code_list"):
		if v, ok := val.([]string); ok {
			nv := StocksToSecurity(v)
			a.request.C2S.SecurityList = nv
		}
	}
}

//=== no need to modify
func (a *QotGetBasicQot) UnPackBody(body []byte) Response {
	rsp := &qotgetbasicqot.Response{}
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
func (a *QotGetBasicQot) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetBasicQot) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
