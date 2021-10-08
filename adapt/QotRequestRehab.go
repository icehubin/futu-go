package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotrequestrehab"
	"google.golang.org/protobuf/proto"
)

type QotRequestRehab struct {
	request *qotrequestrehab.Request

	adaptBase
}

func CreateQotRequestRehab(dopts ...Option) AdaptInterface {
	adp := &QotRequestRehab{
		request: &qotrequestrehab.Request{
			C2S: &qotrequestrehab.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_RequestRehab)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotRequestRehab) SetC2SOption(protoKey string, val interface{}) {
	/*
		Security *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"` //股票
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Security"), strings.ToUpper("code"):
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Security = nv
		}
	}
}

//=== no need to modify
func (a *QotRequestRehab) UnPackBody(body []byte) Response {
	rsp := &qotrequestrehab.Response{}
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
func (a *QotRequestRehab) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotRequestRehab) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
