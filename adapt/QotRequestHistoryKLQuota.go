package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotrequesthistoryklquota"
	"google.golang.org/protobuf/proto"
)

type QotRequestHistoryKLQuota struct {
	request *qotrequesthistoryklquota.Request

	adaptBase
}

func CreateQotRequestHistoryKLQuota(dopts ...Option) AdaptInterface {
	adp := &QotRequestHistoryKLQuota{
		request: &qotrequesthistoryklquota.Request{
			C2S: &qotrequesthistoryklquota.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_RequestHistoryKLQuota)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotRequestHistoryKLQuota) SetC2SOption(protoKey string, val interface{}) {
	/*
		BGetDetail *bool `protobuf:"varint,2,opt,name=bGetDetail" json:"bGetDetail,omitempty"` //是否返回详细拉取过的历史纪录
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("BGetDetail"), strings.ToUpper("Detail"):
		if v, ok := val.(bool); ok {
			a.request.C2S.BGetDetail = proto.Bool(v)
		}
	}
}

//=== no need to modify
func (a *QotRequestHistoryKLQuota) UnPackBody(body []byte) Response {
	rsp := &qotrequesthistoryklquota.Response{}
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
func (a *QotRequestHistoryKLQuota) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotRequestHistoryKLQuota) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
