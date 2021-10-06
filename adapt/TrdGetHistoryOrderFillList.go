package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/trdcommon"
	"github.com/icehubin/futu-go/pb/trdgethistoryorderfilllist"
	"google.golang.org/protobuf/proto"
)

type TrdGetHistoryOrderFillList struct {
	request *trdgethistoryorderfilllist.Request

	adaptBase
}

func CreateTrdGetHistoryOrderFillList(dopts ...Option) AdaptInterface {
	adp := &TrdGetHistoryOrderFillList{
		request: &trdgethistoryorderfilllist.Request{
			C2S: &trdgethistoryorderfilllist.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Trd_GetHistoryOrderFillList)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *TrdGetHistoryOrderFillList) SetC2SOption(protoKey string, val interface{}) {
	/*
		Header           *trdcommon.TrdHeader           `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`                     //交易公共参数头
		FilterConditions *trdcommon.TrdFilterConditions `protobuf:"bytes,2,req,name=filterConditions" json:"filterConditions,omitempty"` //过滤条件
	*/
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
	case strings.ToUpper("FilterConditions"), strings.ToUpper("Conditions"):
		/*
			CodeList  []string `protobuf:"bytes,1,rep,name=codeList" json:"codeList,omitempty"`   //代码过滤，只返回包含这些代码的数据，没传不过滤
			IdList    []uint64 `protobuf:"varint,2,rep,name=idList" json:"idList,omitempty"`      //ID主键过滤，只返回包含这些ID的数据，没传不过滤，订单是orderID、成交是fillID、持仓是positionID
			BeginTime *string  `protobuf:"bytes,3,opt,name=beginTime" json:"beginTime,omitempty"` //开始时间，严格按YYYY-MM-DD HH:MM:SS或YYYY-MM-DD HH:MM:SS.MS格式传，对持仓无效，拉历史数据必须填
			EndTime   *string  `protobuf:"bytes,4,opt,name=endTime" json:"endTime,omitempty"`     //结束时间，严格按YYYY-MM-DD HH:MM:SS或YYYY-MM-DD HH:MM:SS.MS格式传，对持仓无效，拉历史数据必须填
		*/
		if v, ok := val.(TrdFilterConditions); ok {
			trdCon := &trdcommon.TrdFilterConditions{
				CodeList:  v.CodeList,
				IdList:    v.IdList,
				BeginTime: proto.String(v.BeginTime),
				EndTime:   proto.String(v.EndTime),
			}
			if len(v.CodeList) > 0 {
				trdCon.CodeList = v.CodeList
			}
			if len(v.IdList) > 0 {
				trdCon.IdList = v.IdList
			}
			if len(v.BeginTime) > 0 {
				trdCon.BeginTime = proto.String(v.BeginTime)
			}
			if len(v.EndTime) > 0 {
				trdCon.EndTime = proto.String(v.EndTime)
			}

			a.request.C2S.FilterConditions = trdCon
		}
	}
}

//=== no need to modify
func (a *TrdGetHistoryOrderFillList) UnPackBody(body []byte) Response {
	rsp := &trdgethistoryorderfilllist.Response{}
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
func (a *TrdGetHistoryOrderFillList) GetC2S() interface{} {
	return a.request.C2S
}
func (a *TrdGetHistoryOrderFillList) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
