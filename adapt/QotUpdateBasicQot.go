//
package adapt

import (
	"github.com/icehubin/futu-go/pb/qotupdatebasicqot"
	"google.golang.org/protobuf/proto"
)

type QotUpdateBasicQot struct {
	adaptBase
}

func CreateQotUpdateBasicQot(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotUpdateBasicQot{}
	adp.setProtoID(ProtoID_Qot_UpdateBasicQot)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *QotUpdateBasicQot) UnPackBody(body []byte) Response {
	rsp := &qotupdatebasicqot.Response{}
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
