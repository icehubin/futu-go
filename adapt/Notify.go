//
package adapt

import (
	"github.com/icehubin/futu-go/pb/notify"
	"google.golang.org/protobuf/proto"
)

type Notify struct {
	// request  *notify.Request

	adaptBase
}

func CreateNotify(dopts ...Option) AdaptInterface {
	adp := &Notify{}
	adp.setProtoID(ProtoID_Notify)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *Notify) UnPackBody(body []byte) Response {
	rsp := &notify.Response{}
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
