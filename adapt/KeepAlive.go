//
package adapt

import (
	"time"

	"github.com/icehubin/futu-go/pb/keepalive"
	"google.golang.org/protobuf/proto"
)

type KeepAlive struct {
	request *keepalive.Request
	// response *keepalive.Response

	adaptBase
}

func CreateKeepAlive(dopts ...Option) AdaptInterface {
	adp := &KeepAlive{
		request: &keepalive.Request{
			C2S: &keepalive.C2S{
				Time: proto.Int64(time.Now().Unix()),
			},
		},
		// response: &keepalive.Response{},
	}
	adp.setProtoID(ProtoID_KeepAlive)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *KeepAlive) UnPackBody(body []byte) Response {
	rsp := &keepalive.Response{}
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

func (a *KeepAlive) GetC2S() interface{} {
	return a.request.C2S
}
func (a *KeepAlive) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
