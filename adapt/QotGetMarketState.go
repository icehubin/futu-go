package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetmarketstate"
	"google.golang.org/protobuf/proto"
)

type QotGetMarketState struct {
	request *qotgetmarketstate.Request

	adaptBase
}

func CreateQotGetMarketState(dopts ...Option) AdaptInterface {
	adp := &QotGetMarketState{
		request: &qotgetmarketstate.Request{
			C2S: &qotgetmarketstate.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetMarketState)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetMarketState) SetC2SOption(protoKey string, val interface{}) {
	/*
		SecurityList []*qotcommon.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"` //股票列表
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
func (a *QotGetMarketState) UnPackBody(body []byte) Response {
	rsp := &qotgetmarketstate.Response{}
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
func (a *QotGetMarketState) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetMarketState) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
