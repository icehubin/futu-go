//
package adapt

import (
	"github.com/icehubin/futu-go/pb/trdgetacclist"
	"google.golang.org/protobuf/proto"
)

type TrdGetAccList struct {
	request *trdgetacclist.Request

	adaptBase
}

func CreateTrdGetAccList(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &TrdGetAccList{
		request: &trdgetacclist.Request{
			C2S: &trdgetacclist.C2S{
				UserID: proto.Uint64(0),
			},
		},
	}
	adp.setProtoID(ProtoID_Trd_GetAccList)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

//=== no need to modify
func (a *TrdGetAccList) UnPackBody(body []byte) Response {
	rsp := &trdgetacclist.Response{}
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
func (a *TrdGetAccList) GetC2S() interface{} {
	return a.request.C2S
}
func (a *TrdGetAccList) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
