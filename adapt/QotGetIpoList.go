package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetipolist"
	"google.golang.org/protobuf/proto"
)

type QotGetIpoList struct {
	request *qotgetipolist.Request

	adaptBase
}

func CreateQotGetIpoList(dopts ...Option) AdaptInterface {
	adp := &QotGetIpoList{
		request: &qotgetipolist.Request{
			C2S: &qotgetipolist.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetIpoList)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetIpoList) SetC2SOption(protoKey string, val interface{}) {
	/*
		Market *int32 `protobuf:"varint,1,req,name=market" json:"market,omitempty"` // Qot_Common::QotMarket股票市场，支持沪股和深股，且沪股和深股不做区分都代表A股市场。
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("Market"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Market = proto.Int32(v)
		}
		//todo
	}
}

//=== no need to modify
func (a *QotGetIpoList) UnPackBody(body []byte) Response {
	rsp := &qotgetipolist.Response{}
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
func (a *QotGetIpoList) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetIpoList) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
