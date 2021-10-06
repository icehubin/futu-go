//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetrt"
	"google.golang.org/protobuf/proto"
)

type QotGetRT struct {
	request *qotgetrt.Request

	adaptBase
}

func CreateQotGetRT(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotGetRT{
		request: &qotgetrt.Request{
			C2S: &qotgetrt.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetRT)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetRT) SetC2SOption(protoKey string, val interface{}) {
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
func (a *QotGetRT) UnPackBody(body []byte) Response {
	rsp := &qotgetrt.Response{}
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
func (a *QotGetRT) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetRT) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
