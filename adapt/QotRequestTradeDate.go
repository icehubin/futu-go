package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotrequesttradedate"
	"google.golang.org/protobuf/proto"
)

type QotRequestTradeDate struct {
	request *qotrequesttradedate.Request

	adaptBase
}

func CreateQotRequestTradeDate(dopts ...Option) AdaptInterface {
	adp := &QotRequestTradeDate{
		request: &qotrequesttradedate.Request{
			C2S: &qotrequesttradedate.C2S{
				Market: proto.Int32(1),
			},
		},
	}
	adp.setProtoID(ProtoID_Qot_RequestTradeDate)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotRequestTradeDate) SetC2SOption(protoKey string, val interface{}) {
	/*
		//当 market 和 security 同时存在，会忽略 market，仅对 security 进行查询。
		Market    *int32              `protobuf:"varint,1,req,name=market" json:"market,omitempty"`      //Qot_Common.TradeDateMarket,要查询的市场
		BeginTime *string             `protobuf:"bytes,2,req,name=beginTime" json:"beginTime,omitempty"` //开始时间字符串
		EndTime   *string             `protobuf:"bytes,3,req,name=endTime" json:"endTime,omitempty"`     //结束时间字符串
		Security  *qotcommon.Security `protobuf:"bytes,4,opt,name=security" json:"security,omitempty"`   // 指定标的
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("Security"), strings.ToUpper("code"):
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Security = nv
		}
	case strings.ToUpper("Market"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Market = proto.Int32(v)
		} //todo
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
func (a *QotRequestTradeDate) UnPackBody(body []byte) Response {
	rsp := &qotrequesttradedate.Response{}
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
func (a *QotRequestTradeDate) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotRequestTradeDate) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
