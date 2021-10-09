package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotsetpricereminder"
	"google.golang.org/protobuf/proto"
)

type QotSetPriceReminder struct {
	request *qotsetpricereminder.Request

	adaptBase
}

func CreateQotSetPriceReminder(dopts ...Option) AdaptInterface {
	adp := &QotSetPriceReminder{
		request: &qotsetpricereminder.Request{
			C2S: &qotsetpricereminder.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_SetPriceReminder)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotSetPriceReminder) SetC2SOption(protoKey string, val interface{}) {
	/*
		Security *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"` // 股票
		Op       *int32              `protobuf:"varint,2,req,name=op" json:"op,omitempty"`            // ，操作类型
		Key      *int64              `protobuf:"varint,3,opt,name=key" json:"key,omitempty"`          // 到价提醒的标识，GetPriceReminder协议可获得，用于指定要操作的到价提醒项，对于新增的情况不需要填
		Type     *int32              `protobuf:"varint,4,opt,name=type" json:"type,omitempty"`        // Qot_Common::PriceReminderType，提醒类型，删除、启用、禁用的情况下会忽略该字段
		Freq     *int32              `protobuf:"varint,7,opt,name=freq" json:"freq,omitempty"`        // Qot_Common::PriceReminderFreq，提醒频率类型，删除、启用、禁用的情况下会忽略该字段
		Value    *float64            `protobuf:"fixed64,5,opt,name=value" json:"value,omitempty"`     // 提醒值，删除、启用、禁用的情况下会忽略该字段（精确到小数点后 3 位，超出部分会被舍弃）
		Note     *string             `protobuf:"bytes,6,opt,name=note" json:"note,omitempty"`         // 用户设置到价提醒时的标注，仅支持 20 个以内的中文字符，删除、启用、禁用的情况下会忽略该字段
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
	case strings.ToUpper("Op"):
		/*
			SetPriceReminderOp_SetPriceReminderOp_Add     SetPriceReminderOp = 1 //新增
			SetPriceReminderOp_SetPriceReminderOp_Del     SetPriceReminderOp = 2 //删除
			SetPriceReminderOp_SetPriceReminderOp_Enable  SetPriceReminderOp = 3 //启用
			SetPriceReminderOp_SetPriceReminderOp_Disable SetPriceReminderOp = 4 //禁用
			SetPriceReminderOp_SetPriceReminderOp_Modify  SetPriceReminderOp = 5 //修改
			SetPriceReminderOp_SetPriceReminderOp_DelAll  SetPriceReminderOp = 6 //删除该支股票下所有到价提醒
		*/
		if v, ok := val.(int32); ok {
			a.request.C2S.Op = proto.Int32(v)
		} //todo
	case strings.ToUpper("Key"):
		if v, ok := val.(int64); ok {
			a.request.C2S.Key = proto.Int64(v)
		}
	case strings.ToUpper("Type"):
		/*
			PriceReminderType_PriceReminderType_Unknown            PriceReminderType = 0  // 未知
			PriceReminderType_PriceReminderType_PriceUp            PriceReminderType = 1  // 价格涨到
			PriceReminderType_PriceReminderType_PriceDown          PriceReminderType = 2  // 价格跌到
			PriceReminderType_PriceReminderType_ChangeRateUp       PriceReminderType = 3  // 日涨幅超（该字段为百分比字段，设置时填 20 表示 20%）
			PriceReminderType_PriceReminderType_ChangeRateDown     PriceReminderType = 4  // 日跌幅超（该字段为百分比字段，设置时填 20 表示 20%）
			PriceReminderType_PriceReminderType_5MinChangeRateUp   PriceReminderType = 5  // 5 分钟涨幅超（该字段为百分比字段，设置时填 20 表示 20%）
			PriceReminderType_PriceReminderType_5MinChangeRateDown PriceReminderType = 6  // 5 分钟跌幅超（该字段为百分比字段，设置时填 20 表示 20%）
			PriceReminderType_PriceReminderType_VolumeUp           PriceReminderType = 7  // 成交量超过
			PriceReminderType_PriceReminderType_TurnoverUp         PriceReminderType = 8  // 成交额超过
			PriceReminderType_PriceReminderType_TurnoverRateUp     PriceReminderType = 9  // 换手率超过（该字段为百分比字段，设置时填 20 表示 20%）
			PriceReminderType_PriceReminderType_BidPriceUp         PriceReminderType = 10 // 买一价高于
			PriceReminderType_PriceReminderType_AskPriceDown       PriceReminderType = 11 // 卖一价低于
			PriceReminderType_PriceReminderType_BidVolUp           PriceReminderType = 12 // 买一量高于
			PriceReminderType_PriceReminderType_AskVolUp           PriceReminderType = 13 // 卖一量高于
			PriceReminderType_PriceReminderType_3MinChangeRateUp   PriceReminderType = 14 // 3 分钟涨幅超（该字段为百分比字段，设置时填 20 表示 20%）
			PriceReminderType_PriceReminderType_3MinChangeRateDown PriceReminderType = 15 // 3 分钟跌幅超（该字段为百分比字段，设置时填 20 表示 20%）
		*/
		if v, ok := val.(int32); ok {
			a.request.C2S.Type = proto.Int32(v)
		}
	case strings.ToUpper("Freq"):
		/*
			PriceReminderFreq_PriceReminderFreq_Unknown  PriceReminderFreq = 0 // 未知
			PriceReminderFreq_PriceReminderFreq_Always   PriceReminderFreq = 1 // 持续提醒
			PriceReminderFreq_PriceReminderFreq_OnceADay PriceReminderFreq = 2 // 每日一次
			PriceReminderFreq_PriceReminderFreq_OnlyOnce PriceReminderFreq = 3 // 仅提醒一次
		*/
		if v, ok := val.(int32); ok {
			a.request.C2S.Freq = proto.Int32(v)
		}
	case strings.ToUpper("Value"):
		if v, ok := val.(float64); ok {
			a.request.C2S.Value = proto.Float64(v)
		}
	case strings.ToUpper("Note"):
		if v, ok := val.(string); ok {
			a.request.C2S.Note = proto.String(v)
		}
	}
}

//=== no need to modify
func (a *QotSetPriceReminder) UnPackBody(body []byte) Response {
	rsp := &qotsetpricereminder.Response{}
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
func (a *QotSetPriceReminder) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotSetPriceReminder) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
