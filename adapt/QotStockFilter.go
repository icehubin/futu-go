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
				Begin:                     proto.Int32(0),  //从0开始
				Num:                       proto.Int32(20), //默认20条
				Market:                    proto.Int32(21), //沪股市场
				BaseFilterList:            make([]*qotstockfilter.BaseFilter, 0),
				AccumulateFilterList:      make([]*qotstockfilter.AccumulateFilter, 0),
				FinancialFilterList:       make([]*qotstockfilter.FinancialFilter, 0),
				PatternFilterList:         make([]*qotstockfilter.PatternFilter, 0),
				CustomIndicatorFilterList: make([]*qotstockfilter.CustomIndicatorFilter, 0),
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
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
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
	case strings.ToUpper("BaseFilterList"), strings.ToUpper("BaseFilter"):
		/*
			FieldName  *int32   `protobuf:"varint,1,req,name=fieldName" json:"fieldName,omitempty"`   // StockField 简单属性
			FilterMin  *float64 `protobuf:"fixed64,2,opt,name=filterMin" json:"filterMin,omitempty"`  // 区间下限（闭区间），不传代表下限为 -∞
			FilterMax  *float64 `protobuf:"fixed64,3,opt,name=filterMax" json:"filterMax,omitempty"`  // 区间上限（闭区间），不传代表上限为 +∞
			IsNoFilter *bool    `protobuf:"varint,4,opt,name=isNoFilter" json:"isNoFilter,omitempty"` // 该字段是否不需要筛选，True：不筛选，False：筛选。不传默认不筛选
			SortDir    *int32   `protobuf:"varint,5,opt,name=sortDir" json:"sortDir,omitempty"`       // SortDir 排序方向，默认不排序。
		*/
		if v, ok := val.(Message); ok {
			protoAppend(&a.request.C2S.BaseFilterList, v)
		}
	case strings.ToUpper("AccumulateFilterList"), strings.ToUpper("AccumulateFilter"):
		/*
			FieldName  *int32   `protobuf:"varint,1,req,name=fieldName" json:"fieldName,omitempty"`   // AccumulateField 累积属性
			FilterMin  *float64 `protobuf:"fixed64,2,opt,name=filterMin" json:"filterMin,omitempty"`  // 区间下限（闭区间），不传代表下限为 -∞
			FilterMax  *float64 `protobuf:"fixed64,3,opt,name=filterMax" json:"filterMax,omitempty"`  // 区间上限（闭区间），不传代表上限为 +∞
			IsNoFilter *bool    `protobuf:"varint,4,opt,name=isNoFilter" json:"isNoFilter,omitempty"` // 该字段是否不需要筛选，True：不筛选，False：筛选。不传默认不筛选
			SortDir    *int32   `protobuf:"varint,5,opt,name=sortDir" json:"sortDir,omitempty"`       // SortDir 排序方向，默认不排序。
			Days       *int32   `protobuf:"varint,6,req,name=days" json:"days,omitempty"`             // 近几日，累积时间
		*/
		if v, ok := val.(Message); ok {
			protoAppend(&a.request.C2S.AccumulateFilterList, v)
		}
	case strings.ToUpper("FinancialFilterList"), strings.ToUpper("FinancialFilter"):
		/*
			FieldName  *int32   `protobuf:"varint,1,req,name=fieldName" json:"fieldName,omitempty"`   // FinancialField 财务属性
			FilterMin  *float64 `protobuf:"fixed64,2,opt,name=filterMin" json:"filterMin,omitempty"`  // 区间下限（闭区间），不传代表下限为 -∞
			FilterMax  *float64 `protobuf:"fixed64,3,opt,name=filterMax" json:"filterMax,omitempty"`  // 区间上限（闭区间），不传代表上限为 +∞
			IsNoFilter *bool    `protobuf:"varint,4,opt,name=isNoFilter" json:"isNoFilter,omitempty"` // 该字段是否不需要筛选，True：不筛选，False：筛选。不传默认不筛选
			SortDir    *int32   `protobuf:"varint,5,opt,name=sortDir" json:"sortDir,omitempty"`       // SortDir 排序方向，默认不排序。
			Quarter    *int32   `protobuf:"varint,6,req,name=quarter" json:"quarter,omitempty"`       // FinancialQuarter 财报累积时间
		*/
		if v, ok := val.(Message); ok {
			protoAppend(&a.request.C2S.FinancialFilterList, v)
		}
	case strings.ToUpper("PatternFilterList"), strings.ToUpper("PatternFilter"):
		/*
			FieldName  *int32 `protobuf:"varint,1,req,name=fieldName" json:"fieldName,omitempty"`   // PatternField 形态技术指标属性
			KlType     *int32 `protobuf:"varint,2,req,name=klType" json:"klType,omitempty"`         // Qot_Common.KLType，K线类型，仅支持K_60M，K_DAY，K_WEEK，K_MON 四种时间周期
			IsNoFilter *bool  `protobuf:"varint,3,opt,name=isNoFilter" json:"isNoFilter,omitempty"` // 该字段是否不需要筛选，True代表不筛选，False代表筛选。不传默认为不筛选
		*/
		if v, ok := val.(Message); ok {
			protoAppend(&a.request.C2S.PatternFilterList, v)
		}
	case strings.ToUpper("CustomIndicatorFilterList"), strings.ToUpper("CustomIndicatorFilter"):
		/*
			FirstFieldName   *int32   `protobuf:"varint,1,req,name=firstFieldName" json:"firstFieldName,omitempty"`     // CustomIndicatorField 自定义技术指标属性
			SecondFieldName  *int32   `protobuf:"varint,2,req,name=secondFieldName" json:"secondFieldName,omitempty"`   // CustomIndicatorField 自定义技术指标属性
			RelativePosition *int32   `protobuf:"varint,3,req,name=relativePosition" json:"relativePosition,omitempty"` // RelativePosition 相对位置,主要用于MA，EMA，RSI指标做比较
			FieldValue       *float64 `protobuf:"fixed64,4,opt,name=fieldValue" json:"fieldValue,omitempty"`            // 自定义数值，用于与RSI进行比较
			KlType           *int32   `protobuf:"varint,5,req,name=klType" json:"klType,omitempty"`                     // Qot_Common.KLType，K线类型，仅支持K_60M，K_DAY，K_WEEK，K_MON 四种时间周期
			IsNoFilter       *bool    `protobuf:"varint,6,opt,name=isNoFilter" json:"isNoFilter,omitempty"`             // 该字段是否不需要筛选，True代表不筛选，False代表筛选。不传默认为不筛选
		*/
		if v, ok := val.(Message); ok {
			protoAppend(&a.request.C2S.CustomIndicatorFilterList, v)
		}
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
