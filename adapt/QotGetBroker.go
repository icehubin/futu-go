//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetbroker"
	"google.golang.org/protobuf/proto"
)

type QotGetBroker struct {
	request *qotgetbroker.Request

	adaptBase
}

func CreateQotGetBroker(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotGetBroker{
		request: &qotgetbroker.Request{
			C2S: &qotgetbroker.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetBroker)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetBroker) SetC2SOption(protoKey string, val interface{}) {
	/*
		Security *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"` //股票
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Security"), strings.ToUpper("code"):
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Security = nv
		}
	}
}

//=== no need to modify
func (a *QotGetBroker) UnPackBody(body []byte) Response {
	rsp := &qotgetbroker.Response{}
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
func (a *QotGetBroker) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetBroker) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
