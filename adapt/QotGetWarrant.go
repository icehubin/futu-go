package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetwarrant"
	"google.golang.org/protobuf/proto"
)

type QotGetWarrant struct {
	request *qotgetwarrant.Request

	adaptBase
}

func CreateQotGetWarrant(dopts ...Option) AdaptInterface {
	adp := &QotGetWarrant{
		request: &qotgetwarrant.Request{
			C2S: &qotgetwarrant.C2S{
				Ascend:    proto.Bool(true),
				Num:       proto.Int32(20),
				Begin:     proto.Int32(0),
				SortField: proto.Int32(1),
			},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetWarrant)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetWarrant) SetC2SOption(protoKey string, val interface{}) {
	/*
		Begin     *int32 `protobuf:"varint,1,req,name=begin" json:"begin,omitempty"`         //数据起始点
		Num       *int32 `protobuf:"varint,2,req,name=num" json:"num,omitempty"`             //请求数据个数，最大200
		SortField *int32 `protobuf:"varint,3,req,name=sortField" json:"sortField,omitempty"` //Qot_Common.SortField，根据哪个字段排序
		Ascend    *bool  `protobuf:"varint,4,req,name=ascend" json:"ascend,omitempty"`       //升序ture，降序false
		//以下为筛选条件，可选字段，不填表示不过滤
		Owner                 *qotcommon.Security `protobuf:"bytes,5,opt,name=owner" json:"owner,omitempty"`                                    //所属正股
		TypeList              []int32             `protobuf:"varint,6,rep,name=typeList" json:"typeList,omitempty"`                             //Qot_Common.WarrantType，窝轮类型过滤列表
		IssuerList            []int32             `protobuf:"varint,7,rep,name=issuerList" json:"issuerList,omitempty"`                         //Qot_Common.Issuer，发行人过滤列表
		MaturityTimeMin       *string             `protobuf:"bytes,8,opt,name=maturityTimeMin" json:"maturityTimeMin,omitempty"`                //到期日，到期日范围的开始时间戳
		MaturityTimeMax       *string             `protobuf:"bytes,9,opt,name=maturityTimeMax" json:"maturityTimeMax,omitempty"`                //到期日范围的结束时间戳
		IpoPeriod             *int32              `protobuf:"varint,10,opt,name=ipoPeriod" json:"ipoPeriod,omitempty"`                          //Qot_Common.IpoPeriod，上市日
		PriceType             *int32              `protobuf:"varint,11,opt,name=priceType" json:"priceType,omitempty"`                          //Qot_Common.PriceType，价内/价外（暂不支持界内证的界内外筛选）
		Status                *int32              `protobuf:"varint,12,opt,name=status" json:"status,omitempty"`                                //Qot_Common.WarrantStatus，窝轮状态
		CurPriceMin           *float64            `protobuf:"fixed64,13,opt,name=curPriceMin" json:"curPriceMin,omitempty"`                     //最新价的过滤下限（闭区间），不传代表下限为 -∞（精确到小数点后 3 位，超出部分会被舍弃）
		CurPriceMax           *float64            `protobuf:"fixed64,14,opt,name=curPriceMax" json:"curPriceMax,omitempty"`                     //最新价的过滤上限（闭区间），不传代表上限为 +∞（精确到小数点后 3 位，超出部分会被舍弃）
		StrikePriceMin        *float64            `protobuf:"fixed64,15,opt,name=strikePriceMin" json:"strikePriceMin,omitempty"`               //行使价的过滤下限（闭区间），不传代表下限为 -∞（精确到小数点后 3 位，超出部分会被舍弃）
		StrikePriceMax        *float64            `protobuf:"fixed64,16,opt,name=strikePriceMax" json:"strikePriceMax,omitempty"`               //行使价的过滤上限（闭区间），不传代表上限为 +∞（精确到小数点后 3 位，超出部分会被舍弃）
		StreetMin             *float64            `protobuf:"fixed64,17,opt,name=streetMin" json:"streetMin,omitempty"`                         //街货占比的过滤下限（闭区间），该字段为百分比字段，默认不展示 %，如 20 实际对应 20%。不传代表下限为 -∞（精确到小数点后 3 位，超出部分会被舍弃）
		StreetMax             *float64            `protobuf:"fixed64,18,opt,name=streetMax" json:"streetMax,omitempty"`                         //街货占比的过滤上限（闭区间），该字段为百分比字段，默认不展示 %，如 20 实际对应 20%。不传代表上限为 +∞（精确到小数点后 3 位，超出部分会被舍弃）
		ConversionMin         *float64            `protobuf:"fixed64,19,opt,name=conversionMin" json:"conversionMin,omitempty"`                 //换股比率的过滤下限（闭区间），不传代表下限为 -∞（精确到小数点后 3 位，超出部分会被舍弃）
		ConversionMax         *float64            `protobuf:"fixed64,20,opt,name=conversionMax" json:"conversionMax,omitempty"`                 //换股比率的过滤上限（闭区间），不传代表上限为 +∞（精确到小数点后 3 位，超出部分会被舍弃）
		VolMin                *uint64             `protobuf:"varint,21,opt,name=volMin" json:"volMin,omitempty"`                                //成交量的过滤下限（闭区间），不传代表下限为 -∞
		VolMax                *uint64             `protobuf:"varint,22,opt,name=volMax" json:"volMax,omitempty"`                                //成交量的过滤上限（闭区间），不传代表上限为 +∞
		PremiumMin            *float64            `protobuf:"fixed64,23,opt,name=premiumMin" json:"premiumMin,omitempty"`                       //溢价的过滤下限（闭区间），该字段为百分比字段，默认不展示 %，如 20 实际对应 20%。不传代表下限为 -∞（精确到小数点后 3 位，超出部分会被舍弃）
		PremiumMax            *float64            `protobuf:"fixed64,24,opt,name=premiumMax" json:"premiumMax,omitempty"`                       //溢价的过滤上限（闭区间），该字段为百分比字段，默认不展示 %，如 20 实际对应 20%。不传代表上限为 +∞（精确到小数点后 3 位，超出部分会被舍弃）
		LeverageRatioMin      *float64            `protobuf:"fixed64,25,opt,name=leverageRatioMin" json:"leverageRatioMin,omitempty"`           //杠杆比率的过滤下限（闭区间），不传代表下限为 -∞（精确到小数点后 3 位，超出部分会被舍弃）
		LeverageRatioMax      *float64            `protobuf:"fixed64,26,opt,name=leverageRatioMax" json:"leverageRatioMax,omitempty"`           //杠杆比率的过滤上限（闭区间），不传代表上限为 +∞（精确到小数点后 3 位，超出部分会被舍弃）
		DeltaMin              *float64            `protobuf:"fixed64,27,opt,name=deltaMin" json:"deltaMin,omitempty"`                           //对冲值的过滤下限（闭区间），仅认购认沽支持此字段过滤，不传代表下限为 -∞（精确到小数点后 3 位，超出部分会被舍弃）
		DeltaMax              *float64            `protobuf:"fixed64,28,opt,name=deltaMax" json:"deltaMax,omitempty"`                           //对冲值的过滤上限（闭区间），仅认购认沽支持此字段过滤，不传代表上限为 +∞（精确到小数点后 3 位，超出部分会被舍弃）
		ImpliedMin            *float64            `protobuf:"fixed64,29,opt,name=impliedMin" json:"impliedMin,omitempty"`                       //引伸波幅的过滤下限（闭区间），仅认购认沽支持此字段过滤，不传代表下限为 -∞（精确到小数点后 3 位，超出部分会被舍弃）
		ImpliedMax            *float64            `protobuf:"fixed64,30,opt,name=impliedMax" json:"impliedMax,omitempty"`                       //引伸波幅的过滤上限（闭区间），仅认购认沽支持此字段过滤，不传代表上限为 +∞（精确到小数点后 3 位，超出部分会被舍弃）
		RecoveryPriceMin      *float64            `protobuf:"fixed64,31,opt,name=recoveryPriceMin" json:"recoveryPriceMin,omitempty"`           //收回价的过滤下限（闭区间），仅牛熊证支持此字段过滤，不传代表下限为 -∞（精确到小数点后 3 位，超出部分会被舍弃）
		RecoveryPriceMax      *float64            `protobuf:"fixed64,32,opt,name=recoveryPriceMax" json:"recoveryPriceMax,omitempty"`           //收回价的过滤上限（闭区间），仅牛熊证支持此字段过滤，不传代表上限为 +∞（精确到小数点后 3 位，超出部分会被舍弃）
		PriceRecoveryRatioMin *float64            `protobuf:"fixed64,33,opt,name=priceRecoveryRatioMin" json:"priceRecoveryRatioMin,omitempty"` //正股距收回价，的过滤下限（闭区间），仅牛熊证支持此字段过滤。该字段为百分比字段，默认不展示 %，如 20 实际对应 20%。不传代表下限为 -∞（精确到小数点后 3 位，超出部分会被舍弃）
		PriceRecoveryRatioMax *float64            `protobuf:"fixed64,34,opt,name=priceRecoveryRatioMax" json:"priceRecoveryRatioMax,omitempty"` //正股距收回价，的过滤上限（闭区间），仅牛熊证支持此字段过滤。该字段为百分比字段，默认不展示 %，如 20 实际对应 20%。不传代表上限为 +∞（精确到小数点后 3 位，超出部分会被舍弃）
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("Owner"), strings.ToUpper("Security"), strings.ToUpper("code"):
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Owner = nv
		}
	case strings.ToUpper("Begin"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Begin = proto.Int32(v)
		}
	case strings.ToUpper("Num"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Num = proto.Int32(v)
		}
	case strings.ToUpper("SortField"), strings.ToUpper("Sort"):
		if v, ok := val.(int32); ok {
			a.request.C2S.SortField = proto.Int32(v)
		}
	case strings.ToUpper("Ascend"), strings.ToUpper("Asc"):
		if v, ok := val.(bool); ok {
			a.request.C2S.Ascend = proto.Bool(v)
		}
	}
}

//=== no need to modify
func (a *QotGetWarrant) UnPackBody(body []byte) Response {
	rsp := &qotgetwarrant.Response{}
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
func (a *QotGetWarrant) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetWarrant) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
