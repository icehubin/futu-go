//
package adapt

import (
	"github.com/icehubin/futu-go/pb/qotupdatepricereminder"
	"google.golang.org/protobuf/proto"
)

type QotUpdatePriceReminder struct {
	adaptBase
}

func CreateQotUpdatePriceReminder(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotUpdatePriceReminder{}
	adp.setProtoID(ProtoID_Qot_UpdatePriceReminder)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *QotUpdatePriceReminder) UnPackBody(body []byte) Response {
	rsp := &qotupdatepricereminder.Response{}
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
