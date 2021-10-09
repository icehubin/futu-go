package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/testcmd"
	"google.golang.org/protobuf/proto"
)

type TestCmd struct {
	request *testcmd.Request

	adaptBase
}

func CreateTestCmd(dopts ...Option) AdaptInterface {
	adp := &TestCmd{
		request: &testcmd.Request{
			C2S: &testcmd.C2S{},
		},
	}
	adp.setProtoID(ProtoID_TestCmd)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *TestCmd) SetC2SOption(protoKey string, val interface{}) {
	/*
		Cmd    *string `protobuf:"bytes,1,req,name=cmd" json:"cmd,omitempty"`
		Params *string `protobuf:"bytes,2,opt,name=params" json:"params,omitempty"`
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("Cmd"):
		if v, ok := val.(string); ok {
			a.request.C2S.Cmd = proto.String(v)
		}
	case strings.ToUpper("Params"), strings.ToUpper("Param"):
		if v, ok := val.(string); ok {
			a.request.C2S.Params = proto.String(v)
		}
	}
}

//=== no need to modify
func (a *TestCmd) UnPackBody(body []byte) Response {
	rsp := &testcmd.Response{}
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
func (a *TestCmd) GetC2S() interface{} {
	return a.request.C2S
}
func (a *TestCmd) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
