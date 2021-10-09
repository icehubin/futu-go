package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotmodifyusersecurity"
	"google.golang.org/protobuf/proto"
)

type QotModifyUserSecurity struct {
	request *qotmodifyusersecurity.Request

	adaptBase
}

func CreateQotModifyUserSecurity(dopts ...Option) AdaptInterface {
	adp := &QotModifyUserSecurity{
		request: &qotmodifyusersecurity.Request{
			C2S: &qotmodifyusersecurity.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_ModifyUserSecurity)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotModifyUserSecurity) SetC2SOption(protoKey string, val interface{}) {
	/*
		GroupName    *string               `protobuf:"bytes,1,req,name=groupName" json:"groupName,omitempty"`       //分组名,有同名的返回排序的首个
		Op           *int32                `protobuf:"varint,2,req,name=op" json:"op,omitempty"`                    //ModifyUserSecurityOp,操作类型
		SecurityList []*qotcommon.Security `protobuf:"bytes,3,rep,name=securityList" json:"securityList,omitempty"` //新增、删除或移出该分组下的股票
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("GroupName"), strings.ToUpper("Name"):
		if v, ok := val.(string); ok {
			a.request.C2S.GroupName = proto.String(v)
		}
	case strings.ToUpper("SecurityList"), strings.ToUpper("code_list"):
		if v, ok := val.([]string); ok {
			nv := StocksToSecurity(v)
			a.request.C2S.SecurityList = nv
		}
	case strings.ToUpper("Op"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Op = proto.Int32(v)
		} else if v, ok := val.(string); ok {
			switch strings.ToUpper(v) {
			case strings.ToUpper("add"):
				a.request.C2S.Op = proto.Int32(int32(qotmodifyusersecurity.ModifyUserSecurityOp_ModifyUserSecurityOp_Add))
			case strings.ToUpper("del"):
				a.request.C2S.Op = proto.Int32(int32(qotmodifyusersecurity.ModifyUserSecurityOp_ModifyUserSecurityOp_Del))
			case strings.ToUpper("moveout"), strings.ToUpper("remove"), strings.ToUpper("move"):
				a.request.C2S.Op = proto.Int32(int32(qotmodifyusersecurity.ModifyUserSecurityOp_ModifyUserSecurityOp_MoveOut))
			}
		}
	}
}

//=== no need to modify
func (a *QotModifyUserSecurity) UnPackBody(body []byte) Response {
	rsp := &qotmodifyusersecurity.Response{}
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
func (a *QotModifyUserSecurity) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotModifyUserSecurity) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
