//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotcommon"
	"github.com/icehubin/futu-go/pb/qotregqotpush"
	"google.golang.org/protobuf/proto"
)

type QotRegQotPush struct {
	request *qotregqotpush.Request

	adaptBase
}

func CreateQotRegQotPush(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotRegQotPush{
		request: &qotregqotpush.Request{
			C2S: &qotregqotpush.C2S{
				IsRegOrUnReg: proto.Bool(true),
			},
		},
	}
	adp.setProtoID(ProtoID_Qot_RegQotPush)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotRegQotPush) SetC2SOption(protoKey string, val interface{}) {
	/*
		SecurityList  []*qotcommon.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"`    //股票
		SubTypeList   []int32               `protobuf:"varint,2,rep,name=subTypeList" json:"subTypeList,omitempty"`     //Qot_Common.SubType,要注册到该连接的订阅类型
		RehabTypeList []int32               `protobuf:"varint,3,rep,name=rehabTypeList" json:"rehabTypeList,omitempty"` //Qot_Common.RehabType,复权类型,注册K线类型才生效,其他订阅类型忽略该参数,注册K线时该参数不指定默认前复权
		IsRegOrUnReg  *bool                 `protobuf:"varint,4,req,name=isRegOrUnReg" json:"isRegOrUnReg,omitempty"`   //注册或取消
		IsFirstPush   *bool                 `protobuf:"varint,5,opt,name=isFirstPush" json:"isFirstPush,omitempty"`     //注册后如果本地已有数据是否首推一次已存在数据,该参数不指定则默认true
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("SecurityList"):
		if v, ok := val.([]string); ok {
			nv := StocksToSecurity(v)
			a.request.C2S.SecurityList = nv
		}
	case strings.ToUpper("IsRegOrUnReg"):
		if v, ok := val.(bool); ok {
			a.request.C2S.IsRegOrUnReg = proto.Bool(v)
		}
	case strings.ToUpper("IsFirstPush"):
		if v, ok := val.(bool); ok {
			a.request.C2S.IsFirstPush = proto.Bool(v)
		}
	case strings.ToUpper("RehabTypeList"):
		if v, ok := val.([]int32); ok {
			a.request.C2S.RehabTypeList = v
		}
		//字符串类型，不常用，算了
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
	}
}

//=== no need to modify
func (a *QotRegQotPush) UnPackBody(body []byte) Response {
	rsp := &qotregqotpush.Response{}
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
func (a *QotRegQotPush) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotRegQotPush) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
