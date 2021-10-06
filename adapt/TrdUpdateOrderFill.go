//
package adapt

import (
	"github.com/icehubin/futu-go/pb/trdupdateorderfill"
	"google.golang.org/protobuf/proto"
)

type TrdUpdateOrderFill struct {
	adaptBase
}

func CreateTrdUpdateOrderFill(dopts ...Option) AdaptInterface {
	adp := &TrdUpdateOrderFill{}
	adp.setProtoID(ProtoID_Trd_UpdateOrderFill)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *TrdUpdateOrderFill) UnPackBody(body []byte) Response {
	rsp := &trdupdateorderfill.Response{}
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
