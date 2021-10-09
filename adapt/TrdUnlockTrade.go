//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/trdunlocktrade"
	"google.golang.org/protobuf/proto"
)

type TrdUnlockTrade struct {
	request *trdunlocktrade.Request

	adaptBase
}

func CreateTrdUnlockTrade(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &TrdUnlockTrade{
		request: &trdunlocktrade.Request{
			C2S: &trdunlocktrade.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Trd_UnlockTrade)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *TrdUnlockTrade) SetC2SOption(protoKey string, val interface{}) {
	/*
		Unlock       *bool   `protobuf:"varint,1,req,name=unlock" json:"unlock,omitempty"`             //true解锁交易，false锁定交易
		PwdMD5       *string `protobuf:"bytes,2,opt,name=pwdMD5" json:"pwdMD5,omitempty"`              //交易密码的MD5转16进制(全小写)，解锁交易必须要填密码，锁定交易不需要验证密码，可不填
		SecurityFirm *int32  `protobuf:"varint,3,opt,name=securityFirm" json:"securityFirm,omitempty"` //券商标识，取值见Trd_Common.SecurityFirm
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("lock"):
		if v, ok := val.(bool); ok {
			a.request.C2S.Unlock = proto.Bool(!v)
		}
	case strings.ToUpper("Unlock"):
		if v, ok := val.(bool); ok {
			a.request.C2S.Unlock = proto.Bool(v)
		}
	case strings.ToUpper("PwdMD5"), strings.ToUpper("pwd"):
		if v, ok := val.(string); ok {
			a.request.C2S.PwdMD5 = proto.String(v)
		}
	case strings.ToUpper("SecurityFirm"):
		if v, ok := val.(int32); ok {
			a.request.C2S.SecurityFirm = proto.Int32(v)
		}
	}
}

//=== no need to modify
func (a *TrdUnlockTrade) UnPackBody(body []byte) Response {
	rsp := &trdunlocktrade.Response{}
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
func (a *TrdUnlockTrade) GetC2S() interface{} {
	return a.request.C2S
}
func (a *TrdUnlockTrade) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
