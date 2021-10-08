package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetusersecurity"
	"google.golang.org/protobuf/proto"
)

type QotGetUserSecurity struct {
	request *qotgetusersecurity.Request

	adaptBase
}

func CreateQotGetUserSecurity(dopts ...Option) AdaptInterface {
	adp := &QotGetUserSecurity{
		request: &qotgetusersecurity.Request{
			C2S: &qotgetusersecurity.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetUserSecurity)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetUserSecurity) SetC2SOption(protoKey string, val interface{}) {
	/*
		GroupName *string `protobuf:"bytes,1,req,name=groupName" json:"groupName,omitempty"` //分组名,有同名的返回排序首个
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("GroupName"), strings.ToUpper("Name"):
		if v, ok := val.(string); ok {
			a.request.C2S.GroupName = proto.String(v)
		}
	}
}

//=== no need to modify
func (a *QotGetUserSecurity) UnPackBody(body []byte) Response {
	rsp := &qotgetusersecurity.Response{}
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
func (a *QotGetUserSecurity) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetUserSecurity) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
