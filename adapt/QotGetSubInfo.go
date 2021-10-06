//
package adapt

import (
	"github.com/icehubin/futu-go/pb/qotgetsubinfo"
	"google.golang.org/protobuf/proto"
)

type QotGetSubInfo struct {
	request *qotgetsubinfo.Request

	adaptBase
}

func CreateQotGetSubInfo(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotGetSubInfo{
		request: &qotgetsubinfo.Request{
			C2S: &qotgetsubinfo.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetSubInfo)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *QotGetSubInfo) UnPackBody(body []byte) Response {
	rsp := &qotgetsubinfo.Response{}
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
func (a *QotGetSubInfo) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetSubInfo) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
