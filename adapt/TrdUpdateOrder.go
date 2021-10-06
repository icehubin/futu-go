//
package adapt

import (
	"github.com/icehubin/futu-go/pb/trdupdateorder"
	"google.golang.org/protobuf/proto"
)

type TrdUpdateOrder struct {
	adaptBase
}

func CreateTrdUpdateOrder(dopts ...Option) AdaptInterface {
	adp := &TrdUpdateOrder{}
	adp.setProtoID(ProtoID_Trd_UpdateOrder)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *TrdUpdateOrder) UnPackBody(body []byte) Response {
	rsp := &trdupdateorder.Response{}
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
