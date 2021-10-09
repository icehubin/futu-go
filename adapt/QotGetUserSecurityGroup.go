package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetusersecuritygroup"
	"google.golang.org/protobuf/proto"
)

type QotGetUserSecurityGroup struct {
	request *qotgetusersecuritygroup.Request

	adaptBase
}

func CreateQotGetUserSecurityGroup(dopts ...Option) AdaptInterface {
	adp := &QotGetUserSecurityGroup{
		request: &qotgetusersecuritygroup.Request{
			C2S: &qotgetusersecuritygroup.C2S{
				GroupType: proto.Int32(3), //默认全部
			},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetUserSecurityGroup)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetUserSecurityGroup) SetC2SOption(protoKey string, val interface{}) {
	/*
		GroupType *int32 `protobuf:"varint,1,req,name=groupType" json:"groupType,omitempty"` // GroupType,自选股分组类型。
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("GroupType"), strings.ToUpper("Group"):
		if v, ok := val.(int32); ok {
			a.request.C2S.GroupType = proto.Int32(v)
		}
	}
}

//=== no need to modify
func (a *QotGetUserSecurityGroup) UnPackBody(body []byte) Response {
	rsp := &qotgetusersecuritygroup.Response{}
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
func (a *QotGetUserSecurityGroup) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetUserSecurityGroup) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
