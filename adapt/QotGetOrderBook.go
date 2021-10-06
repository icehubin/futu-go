//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetorderbook"
	"google.golang.org/protobuf/proto"
)

type QotGetOrderBook struct {
	request *qotgetorderbook.Request

	adaptBase
}

func CreateQotGetOrderBook(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &QotGetOrderBook{
		request: &qotgetorderbook.Request{
			C2S: &qotgetorderbook.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetOrderBook)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetOrderBook) SetC2SOption(protoKey string, val interface{}) {
	//Todo fix Options or remove
	/*
		Security *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"` //股票
		Num      *int32              `protobuf:"varint,2,req,name=num" json:"num,omitempty"`          //请求的摆盘个数
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Security"), strings.ToUpper("code"):
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Security = nv
		}
	case strings.ToUpper("num"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Num = proto.Int32(v)
		}
	}
}

//=== no need to modify
func (a *QotGetOrderBook) UnPackBody(body []byte) Response {
	rsp := &qotgetorderbook.Response{}
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
func (a *QotGetOrderBook) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetOrderBook) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
