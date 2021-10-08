package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/verification"
	"google.golang.org/protobuf/proto"
)

type Verification struct {
	request *verification.Request

	adaptBase
}

func CreateVerification(dopts ...Option) AdaptInterface {
	adp := &Verification{
		request: &verification.Request{
			C2S: &verification.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Verification)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *Verification) SetC2SOption(protoKey string, val interface{}) {
	/*
		Type *int32  `protobuf:"varint,1,req,name=type" json:"type,omitempty"` //验证码类型, VerificationType
		Op   *int32  `protobuf:"varint,2,req,name=op" json:"op,omitempty"`     //操作, VerificationOp
		Code *string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`  //验证码，请求验证码时忽略该字段，输入时必填
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Type"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Type = proto.Int32(v)
		} else if v, ok := val.(string); ok {
			switch strings.ToUpper(v) {
			case strings.ToUpper("Phone"):
				a.request.C2S.Type = proto.Int32(int32(verification.VerificationType_VerificationType_Phone))
			case strings.ToUpper("Picture"), strings.ToUpper("Pic"):
				a.request.C2S.Type = proto.Int32(int32(verification.VerificationType_VerificationType_Picture))
			}
		}
	case strings.ToUpper("Op"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Op = proto.Int32(v)
		} else if v, ok := val.(string); ok {
			switch strings.ToUpper(v) {
			case strings.ToUpper("Request"), strings.ToUpper("req"):
				a.request.C2S.Op = proto.Int32(int32(verification.VerificationOp_VerificationOp_Request))
			case strings.ToUpper("Input"):
				a.request.C2S.Op = proto.Int32(int32(verification.VerificationOp_VerificationOp_InputAndLogin))
			}
		}
	case strings.ToUpper("Code"):
		if v, ok := val.(string); ok {
			a.request.C2S.Code = proto.String(v)
		}
	}
}

//=== no need to modify
func (a *Verification) UnPackBody(body []byte) Response {
	rsp := &verification.Response{}
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
func (a *Verification) GetC2S() interface{} {
	return a.request.C2S
}
func (a *Verification) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
