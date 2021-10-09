//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/common"
	"github.com/icehubin/futu-go/pb/trdcommon"

	"github.com/icehubin/futu-go/pb/trdplaceorder"
	"google.golang.org/protobuf/proto"
)

type TrdPlaceOrder struct {
	request *trdplaceorder.Request

	adaptBase
}

func CreateTrdPlaceOrder(dopts ...Option) AdaptInterface {
	adp := &TrdPlaceOrder{
		request: &trdplaceorder.Request{
			C2S: &trdplaceorder.C2S{
				OrderType: proto.Int32(1), //普通订单
				Header:    &trdcommon.TrdHeader{},
			},
		},
	}
	adp.setProtoID(ProtoID_Trd_PlaceOrder)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *TrdPlaceOrder) SetC2SOption(protoKey string, val interface{}) {
	/*
		//PacketID  *common.PacketID     `protobuf:"bytes,1,req,name=packetID" json:"packetID,omitempty"`    //交易写操作防重放攻击
		Header    *trdcommon.TrdHeader `protobuf:"bytes,2,req,name=header" json:"header,omitempty"`        //交易公共参数头
		TrdSide   *int32               `protobuf:"varint,3,req,name=trdSide" json:"trdSide,omitempty"`     //交易方向, 参见Trd_Common.TrdSide的枚举定义
		OrderType *int32               `protobuf:"varint,4,req,name=orderType" json:"orderType,omitempty"` //订单类型, 参见Trd_Common.OrderType的枚举定义
		Code      *string              `protobuf:"bytes,5,req,name=code" json:"code,omitempty"`            //代码，港股必须是5位数字，A股必须是6位数字，美股没限制
		Qty       *float64             `protobuf:"fixed64,6,req,name=qty" json:"qty,omitempty"`            //数量，期权单位是"张"（精确到小数点后 0 位，超出部分会被舍弃。期权期货单位是"张"）
		Price     *float64             `protobuf:"fixed64,7,opt,name=price" json:"price,omitempty"`        //价格，（证券账户精确到小数点后 3 位，期货账户精确到小数点后 9 位，超出部分会被舍弃）
		//以下2个为调整价格使用，都传才有效，对港、A股有意义，因为港股有价位，A股2位精度，美股可不传
		AdjustPrice        *bool    `protobuf:"varint,8,opt,name=adjustPrice" json:"adjustPrice,omitempty"`                //是否调整价格，如果价格不合法，是否调整到合法价位，true调整，false不调整
		AdjustSideAndLimit *float64 `protobuf:"fixed64,9,opt,name=adjustSideAndLimit" json:"adjustSideAndLimit,omitempty"` //调整方向和调整幅度百分比限制，正数代表向上调整，负数代表向下调整，具体值代表调整幅度限制，如：0.015代表向上调整且幅度不超过1.5%；-0.01代表向下调整且幅度不超过1%
		SecMarket          *int32   `protobuf:"varint,10,opt,name=secMarket" json:"secMarket,omitempty"`                   //证券所属市场，参见TrdSecMarket的枚举定义
		Remark             *string  `protobuf:"bytes,11,opt,name=remark" json:"remark,omitempty"`                          //用户备注字符串，最多只能传64字节。可用于标识订单唯一信息等，下单填上，订单结构就会带上。
		TimeInForce        *int32   `protobuf:"varint,12,opt,name=timeInForce" json:"timeInForce,omitempty"`               //订单有效期限，参见TrdCommon_TimeInForce的枚举定义
		FillOutsideRTH     *bool    `protobuf:"varint,13,opt,name=fillOutsideRTH" json:"fillOutsideRTH,omitempty"`         //是否允许盘前盘后成交。仅适用于美股限价单。默认false
		AuxPrice           *float64 `protobuf:"fixed64,14,opt,name=auxPrice" json:"auxPrice,omitempty"`                    //触发价格
		TrailType          *int32   `protobuf:"varint,15,opt,name=trailType" json:"trailType,omitempty"`                   //跟踪类型, 参见Trd_Common.TrailType的枚举定义
		TrailValue         *float64 `protobuf:"fixed64,16,opt,name=trailValue" json:"trailValue,omitempty"`                //跟踪金额/百分比
		TrailSpread        *float64 `protobuf:"fixed64,17,opt,name=trailSpread" json:"trailSpread,omitempty"`              //指定价差
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("Header"), strings.ToUpper("Acc"):
		/*
			TrdEnv    *int32  `protobuf:"varint,1,req,name=trdEnv" json:"trdEnv,omitempty"`       //交易环境, 参见TrdEnv的枚举定义
			AccID     *uint64 `protobuf:"varint,2,req,name=accID" json:"accID,omitempty"`         //业务账号, 业务账号与交易环境、市场权限需要匹配，否则会返回错误
			TrdMarket *int32  `protobuf:"varint,3,req,name=trdMarket" json:"trdMarket,omitempty"` //交易市场, 参见TrdMarket的枚举定义
		*/
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S.Header, v)
		}
	case strings.ToUpper("TrdSide"), strings.ToUpper("Side"):
		if v, ok := val.(int32); ok {
			a.request.C2S.TrdSide = proto.Int32(v)
		} else if v, ok := val.(string); ok {
			switch strings.ToUpper(v) {
			case "BUY": //买入
				a.request.C2S.TrdSide = proto.Int32(int32(trdcommon.TrdSide_TrdSide_Buy))
			case "SELL": //卖出
				a.request.C2S.TrdSide = proto.Int32(int32(trdcommon.TrdSide_TrdSide_Sell))
			case "SELLSHORT": //卖空
				a.request.C2S.TrdSide = proto.Int32(int32(trdcommon.TrdSide_TrdSide_SellShort))
			case "BUYBACK": //买回
				a.request.C2S.TrdSide = proto.Int32(int32(trdcommon.TrdSide_TrdSide_BuyBack))
			}
		}
	case strings.ToUpper("OrderType"), strings.ToUpper("Type"):
		if v, ok := val.(int32); ok {
			a.request.C2S.OrderType = proto.Int32(v)
		}
	case strings.ToUpper("Code"):
		if v, ok := val.(string); ok {
			trdStock := StockToTrd(v)
			if nil != trdStock {
				a.request.C2S.Code = proto.String(trdStock.Code)
				a.request.C2S.SecMarket = proto.Int32(trdStock.SecMarket)
			}
		}
	case strings.ToUpper("Qty"):
		if v, ok := val.(float64); ok {
			a.request.C2S.Qty = proto.Float64(v)
		}
	case strings.ToUpper("Price"):
		if v, ok := val.(float64); ok {
			a.request.C2S.Price = proto.Float64(v)
		}
	}
}

func (a *TrdPlaceOrder) PackBody() ([]byte, bool) {
	a.request.C2S.PacketID = &common.PacketID{
		ConnID:   proto.Uint64(a.packetID.ConnID),
		SerialNo: proto.Uint32(a.packetID.SerialNo),
	}
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}

//=== no need to modify
func (a *TrdPlaceOrder) UnPackBody(body []byte) Response {
	rsp := &trdplaceorder.Response{}
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
func (a *TrdPlaceOrder) GetC2S() interface{} {
	return a.request.C2S
}
