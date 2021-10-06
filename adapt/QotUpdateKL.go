//
package adapt

import (
	"github.com/icehubin/futu-go/pb/qotupdatekl"
	"google.golang.org/protobuf/proto"
)

type QotUpdateKL struct {
	adaptBase
}

func CreateQotUpdateKL(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotUpdateKL{}
	adp.setProtoID(ProtoID_Qot_UpdateKL)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *QotUpdateKL) UnPackBody(body []byte) Response {
	rsp := &qotupdatekl.Response{}
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
