package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotstockfilter"
	"google.golang.org/protobuf/proto"
)

type QotStockFilter struct {
	request *qotstockfilter.Request

	adaptBase
}

func CreateQotStockFilter(dopts ...Option) AdaptInterface {
	adp := &QotStockFilter{
		request: &qotstockfilter.Request{
			C2S: &qotstockfilter.C2S{
				Begin:  proto.Int32(0),  //从0开始
				Num:    proto.Int32(20), //默认20条
				Market: proto.Int32(21), //沪股市场
			},
		},
	}
	adp.setProtoID(ProtoID_Qot_StockFilter)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotStockFilter) SetC2SOption(protoKey string, val interface{}) {
	/*
		Begin  *int32 `protobuf:"varint,1,req,name=begin" json:"begin,omitempty"`   // 数据起始点
		Num    *int32 `protobuf:"varint,2,req,name=num" json:"num,omitempty"`       // 请求数据个数，最大200
		Market *int32 `protobuf:"varint,3,req,name=market" json:"market,omitempty"` // Qot_Common::QotMarket股票市场，支持沪股和深股，且沪股和深股不做区分都代表A股市场。
		// 以下为筛选条件，可选字段，不填表示不过滤
		Plate                     *qotcommon.Security      `protobuf:"bytes,4,opt,name=plate" json:"plate,omitempty"`                                         // 板块
		BaseFilterList            []*BaseFilter            `protobuf:"bytes,5,rep,name=baseFilterList" json:"baseFilterList,omitempty"`                       // 简单指标过滤器
		AccumulateFilterList      []*AccumulateFilter      `protobuf:"bytes,6,rep,name=accumulateFilterList" json:"accumulateFilterList,omitempty"`           // 累积指标过滤器
		FinancialFilterList       []*FinancialFilter       `protobuf:"bytes,7,rep,name=financialFilterList" json:"financialFilterList,omitempty"`             // 财务指标过滤器
		PatternFilterList         []*PatternFilter         `protobuf:"bytes,8,rep,name=patternFilterList" json:"patternFilterList,omitempty"`                 // 形态技术指标过滤器
		CustomIndicatorFilterList []*CustomIndicatorFilter `protobuf:"bytes,9,rep,name=customIndicatorFilterList" json:"customIndicatorFilterList,omitempty"` // 自定义技术指标过滤器
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Begin"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Begin = proto.Int32(v)
		}
	case strings.ToUpper("Num"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Num = proto.Int32(v)
		}
	case strings.ToUpper("Market"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Market = proto.Int32(v)
		}
	case strings.ToUpper("Plate"), strings.ToUpper("code"):
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Plate = nv
		}
	case strings.ToUpper("BaseFilterList"):
		if v, ok := val.([]*qotstockfilter.BaseFilter); ok {
			a.request.C2S.BaseFilterList = v
		}
		//todo
	case strings.ToUpper("AccumulateFilterList"):
		if v, ok := val.([]*qotstockfilter.AccumulateFilter); ok {
			a.request.C2S.AccumulateFilterList = v
		}
		//todo
	case strings.ToUpper("FinancialFilterList"):
		if v, ok := val.([]*qotstockfilter.FinancialFilter); ok {
			a.request.C2S.FinancialFilterList = v
		}
		//todo
	case strings.ToUpper("PatternFilterList"):
		if v, ok := val.([]*qotstockfilter.PatternFilter); ok {
			a.request.C2S.PatternFilterList = v
		}
		//todo
	case strings.ToUpper("CustomIndicatorFilterList"):
		if v, ok := val.([]*qotstockfilter.CustomIndicatorFilter); ok {
			a.request.C2S.CustomIndicatorFilterList = v
		}
		//todo
	}
}

//=== no need to modify
func (a *QotStockFilter) UnPackBody(body []byte) Response {
	rsp := &qotstockfilter.Response{}
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
func (a *QotStockFilter) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotStockFilter) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
