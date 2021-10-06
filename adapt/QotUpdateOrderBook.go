//
package adapt

import (
	"github.com/icehubin/futu-go/pb/qotupdateorderbook"
	"google.golang.org/protobuf/proto"
)

type QotUpdateOrderBook struct {
	adaptBase
}

func CreateQotUpdateOrderBook(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotUpdateOrderBook{}
	adp.setProtoID(ProtoID_Qot_UpdateOrderBook)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *QotUpdateOrderBook) UnPackBody(body []byte) Response {
	rsp := &qotupdateorderbook.Response{}
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
