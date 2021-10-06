package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/trdcommon"
	"github.com/icehubin/futu-go/pb/trdgetmaxtrdqtys"
	"google.golang.org/protobuf/proto"
)

type TrdGetMaxTrdQtys struct {
	request *trdgetmaxtrdqtys.Request

	adaptBase
}

func CreateTrdGetMaxTrdQtys(dopts ...Option) AdaptInterface {
	adp := &TrdGetMaxTrdQtys{
		request: &trdgetmaxtrdqtys.Request{
			C2S: &trdgetmaxtrdqtys.C2S{
				OrderType: proto.Int32(1), //默认限价单
			},
		},
	}
	adp.setProtoID(ProtoID_Trd_GetMaxTrdQtys)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *TrdGetMaxTrdQtys) SetC2SOption(protoKey string, val interface{}) {
	/*
		Header    *trdcommon.TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`        //交易公共参数头
		OrderType *int32               `protobuf:"varint,2,req,name=orderType" json:"orderType,omitempty"` //订单类型, 参见Trd_Common.OrderType的枚举定义
		Code      *string              `protobuf:"bytes,3,req,name=code" json:"code,omitempty"`            //代码，港股必须是5位数字，A股必须是6位数字，美股没限制
		Price     *float64             `protobuf:"fixed64,4,req,name=price" json:"price,omitempty"`        //价格，（证券账户精确到小数点后 3 位，期货账户精确到小数点后 9 位，超出部分会被舍弃）。如果是竞价、市价单，请也填入一个当前价格，服务器才好计算
		OrderID   *uint64              `protobuf:"varint,5,opt,name=orderID" json:"orderID,omitempty"`     //订单号，新下订单不需要，如果是修改订单就需要把原订单号带上才行，因为改单的最大买卖数量会包含原订单数量。
		//为保证与下单的价格同步，也提供调整价格选项，以下2个为调整价格使用，对港、A股有意义，因为港股有价位，A股2位精度，美股可不传
		AdjustPrice        *bool    `protobuf:"varint,6,opt,name=adjustPrice" json:"adjustPrice,omitempty"`                //是否调整价格，如果价格不合法，是否调整到合法价位，true调整，false不调整
		AdjustSideAndLimit *float64 `protobuf:"fixed64,7,opt,name=adjustSideAndLimit" json:"adjustSideAndLimit,omitempty"` //调整方向和调整幅度百分比限制，正数代表向上调整，负数代表向下调整，具体值代表调整幅度限制，如：0.015代表向上调整且幅度不超过1.5%；-0.01代表向下调整且幅度不超过1%
		SecMarket          *int32   `protobuf:"varint,8,opt,name=secMarket" json:"secMarket,omitempty"`                    //证券所属市场，参见TrdSecMarket的枚举定义
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Header"), strings.ToUpper("Acc"):
		/*
			TrdEnv    *int32  `protobuf:"varint,1,req,name=trdEnv" json:"trdEnv,omitempty"`       //交易环境, 参见TrdEnv的枚举定义
			AccID     *uint64 `protobuf:"varint,2,req,name=accID" json:"accID,omitempty"`         //业务账号, 业务账号与交易环境、市场权限需要匹配，否则会返回错误
			TrdMarket *int32  `protobuf:"varint,3,req,name=trdMarket" json:"trdMarket,omitempty"` //交易市场, 参见TrdMarket的枚举定义
		*/
		if v, ok := val.(TrdHeader); ok {
			a.request.C2S.Header = &trdcommon.TrdHeader{
				TrdEnv:    proto.Int32(v.TrdEnv),
				AccID:     proto.Uint64(v.AccID),
				TrdMarket: proto.Int32(v.TrdMarket),
			}
		}
	case strings.ToUpper("OrderType"), strings.ToUpper("Type"):
		if v, ok := val.(int32); ok {
			a.request.C2S.OrderType = proto.Int32(v)
		}
	case strings.ToUpper("Code"):
		if v, ok := val.(string); ok {
			a.request.C2S.Code = proto.String(v)
		}
	case strings.ToUpper("Price"):
		if v, ok := val.(float64); ok {
			a.request.C2S.Price = proto.Float64(v)
		}
	case strings.ToUpper("OrderID"), strings.ToUpper("Order"):
		if v, ok := val.(uint64); ok {
			a.request.C2S.OrderID = proto.Uint64(v)
		}
	}
}

//=== no need to modify
func (a *TrdGetMaxTrdQtys) UnPackBody(body []byte) Response {
	rsp := &trdgetmaxtrdqtys.Response{}
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
func (a *TrdGetMaxTrdQtys) GetC2S() interface{} {
	return a.request.C2S
}
func (a *TrdGetMaxTrdQtys) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
