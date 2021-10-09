package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetstaticinfo"
	"google.golang.org/protobuf/proto"
)

type QotGetStaticInfo struct {
	request *qotgetstaticinfo.Request

	adaptBase
}

func CreateQotGetStaticInfo(dopts ...Option) AdaptInterface {
	adp := &QotGetStaticInfo{
		request: &qotgetstaticinfo.Request{
			C2S: &qotgetstaticinfo.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetStaticInfo)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetStaticInfo) SetC2SOption(protoKey string, val interface{}) {
	/*
		// 注：当 market 和 code_list 同时存在时，会忽略 market，仅对 code_list 进行查询。
		Market       *int32                `protobuf:"varint,1,opt,name=market" json:"market,omitempty"`            //Qot_Common.QotMarket,股票市场
		SecType      *int32                `protobuf:"varint,2,opt,name=secType" json:"secType,omitempty"`          //Qot_Common.SecurityType,股票类型
		SecurityList []*qotcommon.Security `protobuf:"bytes,3,rep,name=securityList" json:"securityList,omitempty"` //股票，若该字段存在，忽略其他字段，只返回该字段股票的静态信息
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
	case strings.ToUpper("Market"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Market = proto.Int32(v)
		}
	case strings.ToUpper("SecType"):
		if v, ok := val.(int32); ok {
			a.request.C2S.SecType = proto.Int32(v)
		}
	}
}

//=== no need to modify
func (a *QotGetStaticInfo) UnPackBody(body []byte) Response {
	rsp := &qotgetstaticinfo.Response{}
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
func (a *QotGetStaticInfo) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetStaticInfo) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
