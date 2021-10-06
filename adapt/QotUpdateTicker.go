//
package adapt

import (
	"github.com/icehubin/futu-go/pb/qotupdateticker"
	"google.golang.org/protobuf/proto"
)

type QotUpdateTicker struct {
	adaptBase
}

func CreateQotUpdateTicker(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotUpdateTicker{}
	adp.setProtoID(ProtoID_Qot_UpdateTicker)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *QotUpdateTicker) UnPackBody(body []byte) Response {
	rsp := &qotupdateticker.Response{}
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
