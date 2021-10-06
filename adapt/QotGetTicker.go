//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetticker"
	"google.golang.org/protobuf/proto"
)

type QotGetTicker struct {
	request *qotgetticker.Request

	adaptBase
}

func CreateQotGetTicker(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotGetTicker{
		request: &qotgetticker.Request{
			C2S: &qotgetticker.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetTicker)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetTicker) SetC2SOption(protoKey string, val interface{}) {
	//Todo fix Options or remove
	/*
		Security  *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`    //股票
		MaxRetNum *int32              `protobuf:"varint,2,req,name=maxRetNum" json:"maxRetNum,omitempty"` //最多返回的逐笔个数,实际返回数量不一定会返回这么多,最多返回1000个

	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Security"), strings.ToUpper("code"):
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Security = nv
		}
	case strings.ToUpper("MaxRetNum"), strings.ToUpper("maxNum"), strings.ToUpper("num"):
		if v, ok := val.(int32); ok {
			a.request.C2S.MaxRetNum = proto.Int32(v)
		}
	}
}

//=== no need to modify
func (a *QotGetTicker) UnPackBody(body []byte) Response {
	rsp := &qotgetticker.Response{}
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
func (a *QotGetTicker) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetTicker) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
