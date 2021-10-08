package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetoptionchain"
	"google.golang.org/protobuf/proto"
)

type QotGetOptionChain struct {
	request *qotgetoptionchain.Request

	adaptBase
}

func CreateQotGetOptionChain(dopts ...Option) AdaptInterface {
	adp := &QotGetOptionChain{
		request: &qotgetoptionchain.Request{
			C2S: &qotgetoptionchain.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetOptionChain)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetOptionChain) SetC2SOption(protoKey string, val interface{}) {
	/*
		Owner           *qotcommon.Security `protobuf:"bytes,1,req,name=owner" json:"owner,omitempty"`                      //期权标的股，目前仅支持传入港美正股以及恒指国指
		IndexOptionType *int32              `protobuf:"varint,6,opt,name=indexOptionType" json:"indexOptionType,omitempty"` //Qot_Common.IndexOptionType，指数期权的类型，仅用于恒指国指
		Type            *int32              `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`                       //Qot_Common.OptionType，期权类型，可选字段，不指定则表示都返回
		Condition       *int32              `protobuf:"varint,3,opt,name=condition" json:"condition,omitempty"`             //OptionCondType，价内价外，可选字段，不指定则表示都返回
		BeginTime       *string             `protobuf:"bytes,4,req,name=beginTime" json:"beginTime,omitempty"`              //期权到期日开始时间
		EndTime         *string             `protobuf:"bytes,5,req,name=endTime" json:"endTime,omitempty"`                  //期权到期日结束时间，时间跨度最多一个月
		DataFilter      *DataFilter         `protobuf:"bytes,7,opt,name=dataFilter" json:"dataFilter,omitempty"`            //数据字段筛选
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Owner"), strings.ToUpper("Security"), strings.ToUpper("code"):
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Owner = nv
		}
	case strings.ToUpper("BeginTime"), strings.ToUpper("Begin"):
		if v, ok := val.(string); ok {
			a.request.C2S.BeginTime = proto.String(v)
		}
	case strings.ToUpper("EndTime"), strings.ToUpper("End"):
		if v, ok := val.(string); ok {
			a.request.C2S.EndTime = proto.String(v)
		}
	case strings.ToUpper("IndexOptionType"):
		if v, ok := val.(int32); ok {
			a.request.C2S.IndexOptionType = proto.Int32(v)
		}
	case strings.ToUpper("Type"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Type = proto.Int32(v)
		} //todo string
	case strings.ToUpper("Condition"):
		if v, ok := val.(int32); ok {
			a.request.C2S.Condition = proto.Int32(v)
		}
	case strings.ToUpper("DataFilter"):
		//todo
		if v, ok := val.(*qotgetoptionchain.DataFilter); ok {
			a.request.C2S.DataFilter = v
		}
	}
}

//=== no need to modify
func (a *QotGetOptionChain) UnPackBody(body []byte) Response {
	rsp := &qotgetoptionchain.Response{}
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
func (a *QotGetOptionChain) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetOptionChain) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
