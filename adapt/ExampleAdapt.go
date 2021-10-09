package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/keepalive"
	"google.golang.org/protobuf/proto"
)

type ExampleAdapt struct {
	request *keepalive.Request

	adaptBase
}

func CreateExampleAdapt(dopts ...Option) AdaptInterface {
	adp := &ExampleAdapt{
		request: &keepalive.Request{
			C2S: &keepalive.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Example_Adapt)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *ExampleAdapt) SetC2SOption(protoKey string, val interface{}) {
	//Todo fix Options or remove
	//notify or push proto remove this method
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("Security"), strings.ToUpper("code"):
		// if v, ok := val.(string); ok {
		// 	nv := Stock2Security(v)
		// 	a.request.C2S.Security = nv
		// }
	}
}

//=== no need to modify
func (a *ExampleAdapt) UnPackBody(body []byte) Response {
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
func (a *ExampleAdapt) GetC2S() interface{} {
	return a.request.C2S
}
func (a *ExampleAdapt) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
