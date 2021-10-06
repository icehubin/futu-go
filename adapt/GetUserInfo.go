//
package adapt

import (
	"github.com/icehubin/futu-go/pb/getuserinfo"
	"google.golang.org/protobuf/proto"
)

type GetUserInfo struct {
	request *getuserinfo.Request

	adaptBase
}

func CreateGetUserInfo(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &GetUserInfo{
		request: &getuserinfo.Request{
			C2S: &getuserinfo.C2S{},
		},
	}
	adp.setProtoID(ProtoID_GetUserInfo)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *GetUserInfo) UnPackBody(body []byte) Response {
	rsp := &getuserinfo.Response{}
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
func (a *GetUserInfo) GetC2S() interface{} {
	return a.request.C2S
}
func (a *GetUserInfo) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
