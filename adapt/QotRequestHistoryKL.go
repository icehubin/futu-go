//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotrequesthistorykl"
	"google.golang.org/protobuf/proto"
)

type QotRequestHistoryKL struct {
	request *qotrequesthistorykl.Request

	adaptBase
}

func CreateQotRequestHistoryKL(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotRequestHistoryKL{
		request: &qotrequesthistorykl.Request{
			C2S: &qotrequesthistorykl.C2S{
				RehabType: proto.Int32(1), //默认前复权
			},
		},
	}
	adp.setProtoID(ProtoID_Qot_RequestHistoryKL)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotRequestHistoryKL) SetC2SOption(protoKey string, val interface{}) {
	/*
		RehabType        *int32              `protobuf:"varint,1,req,name=rehabType" json:"rehabType,omitempty"`               //Qot_Common.RehabType,复权类型
		KlType           *int32              `protobuf:"varint,2,req,name=klType" json:"klType,omitempty"`                     //Qot_Common.KLType,K线类型
		Security         *qotcommon.Security `protobuf:"bytes,3,req,name=security" json:"security,omitempty"`                  //股票市场以及股票代码
		BeginTime        *string             `protobuf:"bytes,4,req,name=beginTime" json:"beginTime,omitempty"`                //开始时间字符串
		EndTime          *string             `protobuf:"bytes,5,req,name=endTime" json:"endTime,omitempty"`                    //结束时间字符串
		MaxAckKLNum      *int32              `protobuf:"varint,6,opt,name=maxAckKLNum" json:"maxAckKLNum,omitempty"`           //最多返回多少根K线，如果未指定表示不限制
		NeedKLFieldsFlag *int64              `protobuf:"varint,7,opt,name=needKLFieldsFlag" json:"needKLFieldsFlag,omitempty"` //指定返回K线结构体特定某几项数据，KLFields枚举值或组合，如果未指定返回全部字段
		NextReqKey       []byte              `protobuf:"bytes,8,opt,name=nextReqKey" json:"nextReqKey,omitempty"`              //分页请求key
		ExtendedTime     *bool               `protobuf:"varint,9,opt,name=extendedTime" json:"extendedTime,omitempty"`         //是否获取美股盘前盘后数据，当前仅支持1分k。
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
	case "KTYPE", "KLTYPE", "K_TYPE":
		if v, ok := val.(int32); ok {
			a.request.C2S.KlType = proto.Int32(v)
		} else if v, ok := val.(string); ok {
			/*
				KLType_KLType_Unknown KLType = 0  //未知
				KLType_KLType_1Min    KLType = 1  //1分K
				KLType_KLType_Day     KLType = 2  //日K
				KLType_KLType_Week    KLType = 3  //周K
				KLType_KLType_Month   KLType = 4  //月K
				KLType_KLType_Year    KLType = 5  //年K
				KLType_KLType_5Min    KLType = 6  //5分K
				KLType_KLType_15Min   KLType = 7  //15分K
				KLType_KLType_30Min   KLType = 8  //30分K
				KLType_KLType_60Min   KLType = 9  //60分K
				KLType_KLType_3Min    KLType = 10 //3分K
				KLType_KLType_Quarter KLType = 11 //季K
			*/
			switch strings.ToUpper(v) {
			case "KLTYPE_1MIN", "KL_1MIN", "K_1MIN":
				a.request.C2S.KlType = proto.Int32(1)
			case "KLTYPE_DAY", "KL_DAY", "K_DAY":
				a.request.C2S.KlType = proto.Int32(2)
			case "KLTYPE_WEEK", "KL_WEEK", "K_WEEK":
				a.request.C2S.KlType = proto.Int32(3)
			case "KLTYPE_MONTH", "KL_MONTH", "K_MONTH":
				a.request.C2S.KlType = proto.Int32(4)
			case "KLTYPE_YEAR", "KL_YEAR", "K_YEAR":
				a.request.C2S.KlType = proto.Int32(5)
			case "KLTYPE_5MIN", "KL_5MIN", "K_5MIN":
				a.request.C2S.KlType = proto.Int32(6)
			case "KLTYPE_15MIN", "KL_15MIN", "K_15MIN":
				a.request.C2S.KlType = proto.Int32(7)
			case "KLTYPE_30MIN", "KL_30MIN", "K_30MIN":
				a.request.C2S.KlType = proto.Int32(8)
			case "KLTYPE_60MIN", "KL_60MIN", "K_60MIN":
				a.request.C2S.KlType = proto.Int32(9)
			case "KLTYPE_3MIN", "KL_3MIN", "K_3MIN":
				a.request.C2S.KlType = proto.Int32(10)
			case "KLTYPE_QUARTER", "KL_QUARTER", "K_QUARTER":
				a.request.C2S.KlType = proto.Int32(11)
			}
		}
	case strings.ToUpper("RehabType"):
		if v, ok := val.(int32); ok {
			a.request.C2S.RehabType = proto.Int32(v)
		}
	case strings.ToUpper("BeginTime"), strings.ToUpper("begin"):
		if v, ok := val.(string); ok {
			a.request.C2S.BeginTime = proto.String(v)
		}
	case strings.ToUpper("EndTime"), strings.ToUpper("end"):
		if v, ok := val.(string); ok {
			a.request.C2S.EndTime = proto.String(v)
		}
	case strings.ToUpper("MaxAckKLNum"), strings.ToUpper("Max"):
		if v, ok := val.(int32); ok {
			a.request.C2S.MaxAckKLNum = proto.Int32(v)
		}
	case strings.ToUpper("NextReqKey"), strings.ToUpper("next"):
		if v, ok := val.([]byte); ok {
			a.request.C2S.NextReqKey = v
		}
	}
}

//=== no need to modify
func (a *QotRequestHistoryKL) UnPackBody(body []byte) Response {
	rsp := &qotrequesthistorykl.Response{}
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
func (a *QotRequestHistoryKL) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotRequestHistoryKL) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
