//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/trdcommon"

	"github.com/icehubin/futu-go/pb/trdgetfunds"
	"google.golang.org/protobuf/proto"
)

type TrdGetFunds struct {
	request *trdgetfunds.Request

	adaptBase
}

func CreateTrdGetFunds(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &TrdGetFunds{
		request: &trdgetfunds.Request{
			C2S: &trdgetfunds.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Trd_GetFunds)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *TrdGetFunds) SetC2SOption(protoKey string, val interface{}) {
	/*
		Header       *trdcommon.TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`              //交易公共参数头
		RefreshCache *bool                `protobuf:"varint,2,opt,name=refreshCache" json:"refreshCache,omitempty"` //立即刷新OpenD缓存的此数据，默认不填。true向服务器获取最新数据更新缓存并返回；flase或没填则返回OpenD缓存的数据，不会向服务器请求。
		//正常情况下，服务器有更新就会立即推送到OpenD，OpenD缓存着数据，API请求过来，返回同步的缓存数据，一般不需要指定刷新缓存，保证快速返回且减少对服务器的压力
		//如果遇到丢包等情况，可能出现缓存数据与服务器不一致，用户如果发现数据更新有异样，可指定刷新缓存，解决数据同步的问题。
		Currency *int32 `protobuf:"varint,3,opt,name=currency" json:"currency,omitempty"` //货币种类，参见Trd_Common.Currency。期货账户必填，其它账户忽略
	*/
	a.request.C2S.Reset()
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Header"), strings.ToUpper("Acc"):
		/*
			TrdEnv    *int32  `protobuf:"varint,1,req,name=trdEnv" json:"trdEnv,omitempty"`       //交易环境, 参见TrdEnv的枚举定义
			AccID     *uint64 `protobuf:"varint,2,req,name=accID" json:"accID,omitempty"`         //业务账号, 业务账号与交易环境、市场权限需要匹配，否则会返回错误
			TrdMarket *int32  `protobuf:"varint,3,req,name=trdMarket" json:"trdMarket,omitempty"` //交易市场, 参见TrdMarket的枚举定义
		*/
		if v, ok := val.(TrdHeader); ok {
			a.request.C2S.Header = &trdcommon.TrdHeader{
				TrdEnv:    proto.Int32(v.TrdEnv),
				AccID:     proto.Uint64(v.AccID),
				TrdMarket: proto.Int32(v.TrdMarket),
			}
		}
	case strings.ToUpper("Currency"), strings.ToUpper("Money"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Currency = proto.Int32(v)
		}
	case strings.ToUpper("RefreshCache"), strings.ToUpper("Refresh"):
		if v, ok := val.(bool); ok {
			a.request.C2S.RefreshCache = proto.Bool(v)
		}
	}
}

//=== no need to modify
func (a *TrdGetFunds) UnPackBody(body []byte) Response {
	rsp := &trdgetfunds.Response{}
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
func (a *TrdGetFunds) GetC2S() interface{} {
	return a.request.C2S
}
func (a *TrdGetFunds) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
