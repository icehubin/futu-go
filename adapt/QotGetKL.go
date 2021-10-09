//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetkl"
	"google.golang.org/protobuf/proto"
)

type QotGetKL struct {
	request *qotgetkl.Request

	adaptBase
}

func CreateQotGetKL(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotGetKL{
		request: &qotgetkl.Request{
			C2S: &qotgetkl.C2S{
				RehabType: proto.Int32(1),
			},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetKL)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetKL) SetC2SOption(protoKey string, val interface{}) {
	//Todo fix Options or remove
	//notify or push proto remove this method
	/*
		RehabType *int32              `protobuf:"varint,1,req,name=rehabType" json:"rehabType,omitempty"` //Qot_Common.RehabType,复权类型
		KlType    *int32              `protobuf:"varint,2,req,name=klType" json:"klType,omitempty"`       //Qot_Common.KLType,K线类型
		Security  *qotcommon.Security `protobuf:"bytes,3,req,name=security" json:"security,omitempty"`    //股票
		ReqNum    *int32              `protobuf:"varint,4,req,name=reqNum" json:"reqNum,omitempty"`       //请求K线根数
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case "SECURITY", "CODE":
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Security = nv
		}
	case "REQNUM", "NUM":
		if v, ok := val.(int32); ok {
			a.request.C2S.ReqNum = proto.Int32(v)
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
	case "REHABTYPE":
		if v, ok := val.(int32); ok {
			a.request.C2S.RehabType = proto.Int32(v)
		}
	}
}

//=== no need to modify
func (a *QotGetKL) UnPackBody(body []byte) Response {
	rsp := &qotgetkl.Response{}
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
func (a *QotGetKL) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetKL) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
