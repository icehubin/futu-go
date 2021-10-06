//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/common"
	"github.com/icehubin/futu-go/pb/trdcommon"
	"github.com/icehubin/futu-go/pb/trdmodifyorder"
	"google.golang.org/protobuf/proto"
)

type TrdModifyOrder struct {
	request *trdmodifyorder.Request

	adaptBase
}

func CreateTrdModifyOrder(dopts ...Option) AdaptInterface {
	adp := &TrdModifyOrder{
		request: &trdmodifyorder.Request{
			C2S: &trdmodifyorder.C2S{
				ForAll: proto.Bool(false),
			},
		},
	}
	adp.setProtoID(ProtoID_Trd_ModifyOrder)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *TrdModifyOrder) SetC2SOption(protoKey string, val interface{}) {
	/*
		//PacketID      *common.PacketID     `protobuf:"bytes,1,req,name=packetID" json:"packetID,omitempty"`            //交易写操作防重放攻击
		Header        *trdcommon.TrdHeader `protobuf:"bytes,2,req,name=header" json:"header,omitempty"`                //交易公共参数头
		OrderID       *uint64              `protobuf:"varint,3,req,name=orderID" json:"orderID,omitempty"`             //订单号，forAll为true时，传0
		ModifyOrderOp *int32               `protobuf:"varint,4,req,name=modifyOrderOp" json:"modifyOrderOp,omitempty"` //修改操作类型，参见Trd_Common.ModifyOrderOp的枚举定义
		ForAll        *bool                `protobuf:"varint,5,opt,name=forAll" json:"forAll,omitempty"`               //是否对此业务账户的全部订单操作，true是，false否(对单个订单)，无此字段代表false，仅对单个订单
		//下面的字段仅针对单个订单，且modifyOrderOp为ModifyOrderOp_Normal有效
		Qty   *float64 `protobuf:"fixed64,8,opt,name=qty" json:"qty,omitempty"`     //数量，期权单位是"张"（精确到小数点后 0 位，超出部分会被舍弃）
		Price *float64 `protobuf:"fixed64,9,opt,name=price" json:"price,omitempty"` //价格，（证券账户精确到小数点后 3 位，期货账户精确到小数点后 9 位，超出部分会被舍弃）
		//以下为调整价格使用，都传才有效，对港、A股有意义，因为港股有价位，A股2位精度，美股可不传
		AdjustPrice        *bool    `protobuf:"varint,10,opt,name=adjustPrice" json:"adjustPrice,omitempty"`                //是否调整价格，如果价格不合法，是否调整到合法价位，true调整，false不调整
		AdjustSideAndLimit *float64 `protobuf:"fixed64,11,opt,name=adjustSideAndLimit" json:"adjustSideAndLimit,omitempty"` //调整方向和调整幅度百分比限制，正数代表向上调整，负数代表向下调整，具体值代表调整幅度限制，如：0.015代表向上调整且幅度不超过1.5%；-0.01代表向下调整且幅度不超过1%
		AuxPrice           *float64 `protobuf:"fixed64,12,opt,name=auxPrice" json:"auxPrice,omitempty"`                     //触发价格
		TrailType          *int32   `protobuf:"varint,13,opt,name=trailType" json:"trailType,omitempty"`                    //跟踪类型, 参见Trd_Common.TrailType的枚举定义
		TrailValue         *float64 `protobuf:"fixed64,14,opt,name=trailValue" json:"trailValue,omitempty"`                 //跟踪金额/百分比
		TrailSpread        *float64 `protobuf:"fixed64,15,opt,name=trailSpread" json:"trailSpread,omitempty"`               //指定价差
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
	case strings.ToUpper("OrderID"), strings.ToUpper("Order"):
		if v, ok := val.(uint64); ok {
			a.request.C2S.OrderID = proto.Uint64(v)
		}
	case strings.ToUpper("ModifyOrderOp"), strings.ToUpper("op"):
		if v, ok := val.(int32); ok {
			a.request.C2S.ModifyOrderOp = proto.Int32(v)
		} else if v, ok := val.(string); ok {
			switch strings.ToUpper(v) {
			case "MODIFY": //改单
				a.request.C2S.ModifyOrderOp = proto.Int32(int32(trdcommon.ModifyOrderOp_ModifyOrderOp_Normal))
			case "CANCEL": //撤单
				a.request.C2S.ModifyOrderOp = proto.Int32(int32(trdcommon.ModifyOrderOp_ModifyOrderOp_Cancel))
			case "DISABLE": //失效
				a.request.C2S.ModifyOrderOp = proto.Int32(int32(trdcommon.ModifyOrderOp_ModifyOrderOp_Disable))
			case "ENABLE": //生效
				a.request.C2S.ModifyOrderOp = proto.Int32(int32(trdcommon.TrdSide_TrdSide_BuyBack))
			case "DELETE", "DEL": //删除
				a.request.C2S.ModifyOrderOp = proto.Int32(int32(trdcommon.ModifyOrderOp_ModifyOrderOp_Delete))
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

func (a *TrdModifyOrder) PackBody() ([]byte, bool) {
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
func (a *TrdModifyOrder) UnPackBody(body []byte) Response {
	rsp := &trdmodifyorder.Response{}
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
func (a *TrdModifyOrder) GetC2S() interface{} {
	return a.request.C2S
}
