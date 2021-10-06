//
package adapt

import (
	"github.com/icehubin/futu-go/pb/qotupdatert"
	"google.golang.org/protobuf/proto"
)

type QotUpdateRT struct {
	adaptBase
}

func CreateQotUpdateRT(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotUpdateRT{}
	adp.setProtoID(ProtoID_Qot_UpdateRT)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *QotUpdateRT) UnPackBody(body []byte) Response {
	rsp := &qotupdatert.Response{}
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
