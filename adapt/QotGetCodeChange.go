package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetcodechange"
	"google.golang.org/protobuf/proto"
)

type QotGetCodeChange struct {
	request *qotgetcodechange.Request

	adaptBase
}

func CreateQotGetCodeChange(dopts ...Option) AdaptInterface {
	adp := &QotGetCodeChange{
		request: &qotgetcodechange.Request{
			C2S: &qotgetcodechange.C2S{
				TimeFilterList: make([]*qotgetcodechange.TimeFilter, 0),
			},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetCodeChange)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetCodeChange) SetC2SOption(protoKey string, val interface{}) {
	/*
		PlaceHolder    *int32                `protobuf:"varint,1,opt,name=placeHolder" json:"placeHolder,omitempty"`      //占位
		SecurityList   []*qotcommon.Security `protobuf:"bytes,2,rep,name=securityList" json:"securityList,omitempty"`     //根据股票筛选
		TimeFilterList []*TimeFilter         `protobuf:"bytes,3,rep,name=timeFilterList" json:"timeFilterList,omitempty"` //根据时间筛选
		TypeList       []int32               `protobuf:"varint,4,rep,name=typeList" json:"typeList,omitempty"`            //CodeChangeType，根据类型筛选
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("SecurityList"), strings.ToUpper("code_list"):
		if v, ok := val.([]string); ok {
			nv := StocksToSecurity(v)
			a.request.C2S.SecurityList = nv
		}
	case strings.ToUpper("PlaceHolder"):
		if v, ok := val.(int32); ok {
			a.request.C2S.PlaceHolder = proto.Int32(v)
		}
	case strings.ToUpper("TypeList"):
		if v, ok := val.([]int32); ok {
			a.request.C2S.TypeList = v
		}
	case strings.ToUpper("TimeFilterList"), strings.ToUpper("TimeFilter"), strings.ToUpper("Filter"):
		/*
			Type      *int32  `protobuf:"varint,1,req,name=type" json:"type,omitempty"`          //TimeFilterType, 过滤类型
			BeginTime *string `protobuf:"bytes,2,opt,name=beginTime" json:"beginTime,omitempty"` //开始时间点
			EndTime   *string `protobuf:"bytes,3,opt,name=endTime" json:"endTime,omitempty"`     //结束时间点
		*/
		if v, ok := val.(Message); ok {
			protoAppend(&a.request.C2S.TimeFilterList, v)
		}
	}
}

//=== no need to modify
func (a *QotGetCodeChange) UnPackBody(body []byte) Response {
	rsp := &qotgetcodechange.Response{}
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
func (a *QotGetCodeChange) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetCodeChange) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
