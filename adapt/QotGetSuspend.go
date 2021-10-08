package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetsuspend"
	"google.golang.org/protobuf/proto"
)

type QotGetSuspend struct {
	request *qotgetsuspend.Request

	adaptBase
}

func CreateQotGetSuspend(dopts ...Option) AdaptInterface {
	adp := &QotGetSuspend{
		request: &qotgetsuspend.Request{
			C2S: &qotgetsuspend.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetSuspend)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetSuspend) SetC2SOption(protoKey string, val interface{}) {
	/*
		SecurityList []*qotcommon.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"` //股票
		BeginTime    *string               `protobuf:"bytes,2,req,name=beginTime" json:"beginTime,omitempty"`       //开始时间字符串
		EndTime      *string               `protobuf:"bytes,3,req,name=endTime" json:"endTime,omitempty"`           //结束时间字符串
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("SecurityList"), strings.ToUpper("code_list"):
		if v, ok := val.([]string); ok {
			nv := StocksToSecurity(v)
			a.request.C2S.SecurityList = nv
		}
	case strings.ToUpper("BeginTime"), strings.ToUpper("Begin"):
		if v, ok := val.(string); ok {
			a.request.C2S.BeginTime = proto.String(v)
		}
	case strings.ToUpper("EndTime"), strings.ToUpper("End"):
		if v, ok := val.(string); ok {
			a.request.C2S.EndTime = proto.String(v)
		}
	}
}

//=== no need to modify
func (a *QotGetSuspend) UnPackBody(body []byte) Response {
	rsp := &qotgetsuspend.Response{}
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
func (a *QotGetSuspend) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetSuspend) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
