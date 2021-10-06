//
package adapt

import (
	"github.com/icehubin/futu-go/pb/getglobalstate"
	"google.golang.org/protobuf/proto"
)

type GetGlobalState struct {
	request *getglobalstate.Request

	adaptBase
}

func CreateGetGlobalState(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &GetGlobalState{
		request: &getglobalstate.Request{
			C2S: &getglobalstate.C2S{
				UserID: proto.Uint64(0),
			},
		},
	}
	adp.setProtoID(ProtoID_GetGlobalState)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *GetGlobalState) UnPackBody(body []byte) Response {
	rsp := &getglobalstate.Response{}
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
func (a *GetGlobalState) GetC2S() interface{} {
	return a.request.C2S
}
func (a *GetGlobalState) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
