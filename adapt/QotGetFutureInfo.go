package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetfutureinfo"
	"google.golang.org/protobuf/proto"
)

type QotGetFutureInfo struct {
	request *qotgetfutureinfo.Request

	adaptBase
}

func CreateQotGetFutureInfo(dopts ...Option) AdaptInterface {
	adp := &QotGetFutureInfo{
		request: &qotgetfutureinfo.Request{
			C2S: &qotgetfutureinfo.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetFutureInfo)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetFutureInfo) SetC2SOption(protoKey string, val interface{}) {
	/*
		SecurityList []*qotcommon.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"` //股票列表
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("SecurityList"), strings.ToUpper("code_list"):
		if v, ok := val.([]string); ok {
			nv := StocksToSecurity(v)
			a.request.C2S.SecurityList = nv
		}
	}
}

//=== no need to modify
func (a *QotGetFutureInfo) UnPackBody(body []byte) Response {
	rsp := &qotgetfutureinfo.Response{}
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
func (a *QotGetFutureInfo) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetFutureInfo) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
