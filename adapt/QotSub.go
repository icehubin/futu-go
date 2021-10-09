//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotcommon"

	"github.com/icehubin/futu-go/pb/qotsub"
	"google.golang.org/protobuf/proto"
)

type QotSub struct {
	request  *qotsub.Request
	response *qotsub.Response

	adaptBase
}

func CreateQotSub(dopts ...Option) AdaptInterface {
	adp := &QotSub{
		request: &qotsub.Request{
			C2S: &qotsub.C2S{
				IsSubOrUnSub: proto.Bool(true),
			},
		},
		response: &qotsub.Response{},
	}
	adp.setProtoID(ProtoID_Qot_Sub)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotSub) SetC2SOption(protoKey string, val interface{}) {
	/*
		SecurityList         []*qotcommon.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"`                  //股票
		SubTypeList          []int32               `protobuf:"varint,2,rep,name=subTypeList" json:"subTypeList,omitempty"`                   //Qot_Common.SubType,订阅数据类型
		IsSubOrUnSub         *bool                 `protobuf:"varint,3,req,name=isSubOrUnSub" json:"isSubOrUnSub,omitempty"`                 //ture表示订阅,false表示反订阅
		IsRegOrUnRegPush     *bool                 `protobuf:"varint,4,opt,name=isRegOrUnRegPush" json:"isRegOrUnRegPush,omitempty"`         //是否注册或反注册该连接上面行情的推送,该参数不指定不做注册反注册操作
		RegPushRehabTypeList []int32               `protobuf:"varint,5,rep,name=regPushRehabTypeList" json:"regPushRehabTypeList,omitempty"` //Qot_Common.RehabType,复权类型,注册推送并且是K线类型才生效,其他订阅类型忽略该参数,注册K线推送时该参数不指定默认前复权
		IsFirstPush          *bool                 `protobuf:"varint,6,opt,name=isFirstPush" json:"isFirstPush,omitempty"`                   //注册后如果本地已有数据是否首推一次已存在数据,该参数不指定则默认true
		IsUnsubAll           *bool                 `protobuf:"varint,7,opt,name=isUnsubAll" json:"isUnsubAll,omitempty"`                     //当被设置为True时忽略其他参数，取消当前连接的所有订阅，并且反注册推送。
		IsSubOrderBookDetail *bool                 `protobuf:"varint,8,opt,name=isSubOrderBookDetail" json:"isSubOrderBookDetail,omitempty"` //订阅摆盘可用,是否订阅摆盘明细,仅支持SF行情,该参数不指定则默认false
		ExtendedTime         *bool                 `protobuf:"varint,9,opt,name=extendedTime" json:"extendedTime,omitempty"`                 // 是否允许美股盘前盘后数据（仅用于订阅美股的实时K线、实时分时、实时逐笔）
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("SecurityList"), strings.ToUpper("code_list"):
		if v, ok := val.([]string); ok {
			nv := StocksToSecurity(v)
			a.request.C2S.SecurityList = nv
		}
	case strings.ToUpper("IsRegOrUnRegPush"), strings.ToUpper("Push"):
		if v, ok := val.(bool); ok {
			a.request.C2S.IsRegOrUnRegPush = proto.Bool(v)
		}
	case strings.ToUpper("IsSubOrUnSub"), strings.ToUpper("sub"):
		if v, ok := val.(bool); ok {
			a.request.C2S.IsSubOrUnSub = proto.Bool(v)
		}
	case strings.ToUpper("SubTypeList"), strings.ToUpper("subtype_list"):
		/*
			None       SubType = 0
			SubType_SubType_Basic      SubType = 1  //基础报价
			SubType_SubType_OrderBook  SubType = 2  //摆盘
			SubType_SubType_Ticker     SubType = 4  //逐笔
			SubType_SubType_RT         SubType = 5  //分时
			SubType_SubType_KL_Day     SubType = 6  //日K
			SubType_SubType_KL_5Min    SubType = 7  //5分K
			SubType_SubType_KL_15Min   SubType = 8  //15分K
			SubType_SubType_KL_30Min   SubType = 9  //30分K
			SubType_SubType_KL_60Min   SubType = 10 //60分K
			SubType_SubType_KL_1Min    SubType = 11 //1分K
			SubType_SubType_KL_Week    SubType = 12 //周K
			SubType_SubType_KL_Month   SubType = 13 //月K
			SubType_SubType_Broker     SubType = 14 //经纪队列
			SubType_SubType_KL_Qurater SubType = 15 //季K
			SubType_SubType_KL_Year    SubType = 16 //年K
			SubType_SubType_KL_3Min    SubType = 17 //3分K
		*/
		if v, ok := val.([]int32); ok {
			a.request.C2S.SubTypeList = v
		} else if v, ok := val.([]string); ok {
			//字符串类型
			nv := make([]int32, 0)
			for _, v_ := range v {
				switch strings.ToUpper(v_) {
				case "QUOTE", "BASIC":
					nv = append(nv, int32(qotcommon.SubType_SubType_Basic))
				case "ORDER_BOOK", "ORDERBOOK":
					nv = append(nv, int32(qotcommon.SubType_SubType_OrderBook))
				case "TICKER":
					nv = append(nv, int32(qotcommon.SubType_SubType_Ticker))
				case "RT":
					nv = append(nv, int32(qotcommon.SubType_SubType_RT))
				case "BROKER":
					nv = append(nv, int32(qotcommon.SubType_SubType_Broker))
				case "K_DAY", "KL_DAY":
					nv = append(nv, int32(qotcommon.SubType_SubType_KL_Day))
				case "K_5M", "KL_5M":
					nv = append(nv, int32(qotcommon.SubType_SubType_KL_5Min))
				case "K_15M", "KL_15M":
					nv = append(nv, int32(qotcommon.SubType_SubType_KL_15Min))
				case "K_30M", "KL_30M":
					nv = append(nv, int32(qotcommon.SubType_SubType_KL_30Min))
				case "K_60M", "KL_60M":
					nv = append(nv, int32(qotcommon.SubType_SubType_KL_60Min))
				case "K_1M", "KL_1M":
					nv = append(nv, int32(qotcommon.SubType_SubType_KL_1Min))
				case "K_3M", "KL_3M":
					nv = append(nv, int32(qotcommon.SubType_SubType_KL_3Min))
				case "K_WEEK", "KL_WEEK":
					nv = append(nv, int32(qotcommon.SubType_SubType_KL_Week))
				case "K_MON", "KL_MON":
					nv = append(nv, int32(qotcommon.SubType_SubType_KL_Month))
				case "K_QURATER", "KL_QURATER":
					nv = append(nv, int32(qotcommon.SubType_SubType_KL_Qurater))
				case "K_YEAR", "KL_YEAR":
					nv = append(nv, int32(qotcommon.SubType_SubType_KL_Year))
				}
			}
			a.request.C2S.SubTypeList = nv
		}
	case strings.ToUpper("ExtendedTime"):
		if v, ok := val.(bool); ok {
			a.request.C2S.ExtendedTime = proto.Bool(v)
		}
	case strings.ToUpper("IsSubOrderBookDetail"):
		if v, ok := val.(bool); ok {
			a.request.C2S.IsSubOrderBookDetail = proto.Bool(v)
		}
	case strings.ToUpper("IsUnsubAll"):
		if v, ok := val.(bool); ok {
			a.request.C2S.IsUnsubAll = proto.Bool(v)
		}
	case strings.ToUpper("IsFirstPush"):
		if v, ok := val.(bool); ok {
			a.request.C2S.IsFirstPush = proto.Bool(v)
		}
	case strings.ToUpper("RegPushRehabTypeList"):
		//字符串类型，不常用，算了
		if v, ok := val.([]int32); ok {
			a.request.C2S.RegPushRehabTypeList = v
		}

	}
}

//=== no need to modify
func (a *QotSub) UnPackBody(body []byte) Response {
	rsp := &qotsub.Response{}
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
func (a *QotSub) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotSub) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
