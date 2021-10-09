package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/trdcommon"
	"github.com/icehubin/futu-go/pb/trdgetmarginratio"
	"google.golang.org/protobuf/proto"
)

type TrdGetMarginRatio struct {
	request *trdgetmarginratio.Request

	adaptBase
}

func CreateTrdGetMarginRatio(dopts ...Option) AdaptInterface {
	adp := &TrdGetMarginRatio{
		request: &trdgetmarginratio.Request{
			C2S: &trdgetmarginratio.C2S{
				Header: &trdcommon.TrdHeader{},
			},
		},
	}
	adp.setProtoID(ProtoID_Trd_GetMarginRatio)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *TrdGetMarginRatio) SetC2SOption(protoKey string, val interface{}) {
	/*
		Header       *trdcommon.TrdHeader  `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`             //交易公共参数头
		SecurityList []*qotcommon.Security `protobuf:"bytes,2,rep,name=securityList" json:"securityList,omitempty"` //股票
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("Header"), strings.ToUpper("Acc"):
		/*
			TrdEnv    *int32  `protobuf:"varint,1,req,name=trdEnv" json:"trdEnv,omitempty"`       //交易环境, 参见TrdEnv的枚举定义
			AccID     *uint64 `protobuf:"varint,2,req,name=accID" json:"accID,omitempty"`         //业务账号, 业务账号与交易环境、市场权限需要匹配，否则会返回错误
			TrdMarket *int32  `protobuf:"varint,3,req,name=trdMarket" json:"trdMarket,omitempty"` //交易市场, 参见TrdMarket的枚举定义
		*/
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S.Header, v)
		}
	case strings.ToUpper("SecurityList"), strings.ToUpper("code_list"):
		if v, ok := val.([]string); ok {
			nv := StocksToSecurity(v)
			a.request.C2S.SecurityList = nv
		}
	}
}

//=== no need to modify
func (a *TrdGetMarginRatio) UnPackBody(body []byte) Response {
	rsp := &trdgetmarginratio.Response{}
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
func (a *TrdGetMarginRatio) GetC2S() interface{} {
	return a.request.C2S
}
func (a *TrdGetMarginRatio) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
