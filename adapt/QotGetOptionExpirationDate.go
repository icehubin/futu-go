package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetoptionexpirationdate"
	"google.golang.org/protobuf/proto"
)

type QotGetOptionExpirationDate struct {
	request *qotgetoptionexpirationdate.Request

	adaptBase
}

func CreateQotGetOptionExpirationDate(dopts ...Option) AdaptInterface {
	adp := &QotGetOptionExpirationDate{
		request: &qotgetoptionexpirationdate.Request{
			C2S: &qotgetoptionexpirationdate.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetOptionExpirationDate)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetOptionExpirationDate) SetC2SOption(protoKey string, val interface{}) {
	/*
		Owner           *qotcommon.Security `protobuf:"bytes,1,req,name=owner" json:"owner,omitempty"`                      //期权标的股，目前仅支持传入港美正股以及恒指国指
		IndexOptionType *int32              `protobuf:"varint,2,opt,name=indexOptionType" json:"indexOptionType,omitempty"` //Qot_Common.IndexOptionType，指数期权的类型，仅用于恒指国指
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Owner"), strings.ToUpper("code"):
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Owner = nv
		}
	case strings.ToUpper("IndexOptionType"), strings.ToUpper("type"):
		/*
			IndexOptionType_IndexOptionType_Unknown IndexOptionType = 0 //未知
			IndexOptionType_IndexOptionType_Normal  IndexOptionType = 1 //正常普通的指数期权
			IndexOptionType_IndexOptionType_Small   IndexOptionType = 2 //小型指数期权
		*/
		if v, ok := val.(int32); ok {
			a.request.C2S.IndexOptionType = proto.Int32(v)
		}
	}
}

//=== no need to modify
func (a *QotGetOptionExpirationDate) UnPackBody(body []byte) Response {
	rsp := &qotgetoptionexpirationdate.Response{}
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
func (a *QotGetOptionExpirationDate) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetOptionExpirationDate) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
