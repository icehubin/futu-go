//
package adapt

import (
	"github.com/icehubin/futu-go/pb/qotupdatebroker"
	"google.golang.org/protobuf/proto"
)

type QotUpdateBroker struct {
	adaptBase
}

func CreateQotUpdateBroker(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotUpdateBroker{}
	adp.setProtoID(ProtoID_Qot_UpdateBroker)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *QotUpdateBroker) UnPackBody(body []byte) Response {
	rsp := &qotupdatebroker.Response{}
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
