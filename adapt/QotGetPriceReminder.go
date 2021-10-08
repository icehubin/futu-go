package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetpricereminder"
	"google.golang.org/protobuf/proto"
)

type QotGetPriceReminder struct {
	request *qotgetpricereminder.Request

	adaptBase
}

func CreateQotGetPriceReminder(dopts ...Option) AdaptInterface {
	adp := &QotGetPriceReminder{
		request: &qotgetpricereminder.Request{
			C2S: &qotgetpricereminder.C2S{
				Market: proto.Int32(1), //默认香港
			},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetPriceReminder)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetPriceReminder) SetC2SOption(protoKey string, val interface{}) {
	/*
		Security *qotcommon.Security `protobuf:"bytes,1,opt,name=security" json:"security,omitempty"` // 查询股票下的到价提醒项，security和market二选一，都存在的情况下security优先。
		Market   *int32              `protobuf:"varint,2,opt,name=market" json:"market,omitempty"`    //Qot_Common::QotMarket 市场，查询市场下的到价提醒项，不区分沪深
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Security"), strings.ToUpper("code"):
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Security = nv
		}
	case strings.ToUpper("Market"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Market = proto.Int32(v)
		}
	}
}

//=== no need to modify
func (a *QotGetPriceReminder) UnPackBody(body []byte) Response {
	rsp := &qotgetpricereminder.Response{}
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
func (a *QotGetPriceReminder) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetPriceReminder) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
