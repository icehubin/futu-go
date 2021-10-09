package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetplateset"
	"google.golang.org/protobuf/proto"
)

type QotGetPlateSet struct {
	request *qotgetplateset.Request

	adaptBase
}

func CreateQotGetPlateSet(dopts ...Option) AdaptInterface {
	adp := &QotGetPlateSet{
		request: &qotgetplateset.Request{
			C2S: &qotgetplateset.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetPlateSet)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetPlateSet) SetC2SOption(protoKey string, val interface{}) {
	/*
		Market       *int32 `protobuf:"varint,1,req,name=market" json:"market,omitempty"`             //Qot_Common.QotMarket,股票市场
		PlateSetType *int32 `protobuf:"varint,2,req,name=plateSetType" json:"plateSetType,omitempty"` //Qot_Common.PlateSetType,板块集合的类型
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("Market"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Market = proto.Int32(v)
		}
	case strings.ToUpper("PlateSetType"), strings.ToUpper("Type"):
		if v, ok := val.(int32); ok {
			a.request.C2S.PlateSetType = proto.Int32(v)
		}
	}
}

//=== no need to modify
func (a *QotGetPlateSet) UnPackBody(body []byte) Response {
	rsp := &qotgetplateset.Response{}
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
func (a *QotGetPlateSet) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetPlateSet) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
