package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetcapitalflow"
	"google.golang.org/protobuf/proto"
)

type QotGetCapitalFlow struct {
	request *qotgetcapitalflow.Request

	adaptBase
}

func CreateQotGetCapitalFlow(dopts ...Option) AdaptInterface {
	adp := &QotGetCapitalFlow{
		request: &qotgetcapitalflow.Request{
			C2S: &qotgetcapitalflow.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetCapitalFlow)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetCapitalFlow) SetC2SOption(protoKey string, val interface{}) {
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
func (a *QotGetCapitalFlow) UnPackBody(body []byte) Response {
	rsp := &qotgetcapitalflow.Response{}
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
func (a *QotGetCapitalFlow) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetCapitalFlow) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
