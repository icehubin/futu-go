package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotgetholdingchangelist"
	"google.golang.org/protobuf/proto"
)

type QotGetHoldingChangeList struct {
	request *qotgetholdingchangelist.Request

	adaptBase
}

func CreateQotGetHoldingChangeList(dopts ...Option) AdaptInterface {
	adp := &QotGetHoldingChangeList{
		request: &qotgetholdingchangelist.Request{
			C2S: &qotgetholdingchangelist.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Qot_GetHoldingChangeList)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *QotGetHoldingChangeList) SetC2SOption(protoKey string, val interface{}) {
	/*
		Security       *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`              //股票
		HolderCategory *int32              `protobuf:"varint,2,req,name=holderCategory" json:"holderCategory,omitempty"` //持有者类别（1机构、2基金、3高管）
		//以下是发布时间筛选，不传返回所有数据，传了返回发布时间属于开始时间到结束时间段内的数据
		BeginTime *string `protobuf:"bytes,3,opt,name=beginTime" json:"beginTime,omitempty"` //开始时间，严格按YYYY-MM-DD HH:MM:SS或YYYY-MM-DD HH:MM:SS.MS格式传
		EndTime   *string `protobuf:"bytes,4,opt,name=endTime" json:"endTime,omitempty"`     //结束时间，严格按YYYY-MM-DD HH:MM:SS或YYYY-MM-DD HH:MM:SS.MS格式传
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("Security"), strings.ToUpper("code"):
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Security = nv
		}
	case strings.ToUpper("HolderCategory"), strings.ToUpper("Category"):
		if v, ok := val.(int32); ok {
			a.request.C2S.HolderCategory = proto.Int32(v)
		}
	case strings.ToUpper("BeginTime"), strings.ToUpper("Begin"):
		if v, ok := val.(string); ok {
			a.request.C2S.BeginTime = proto.String(v)
		}
	case strings.ToUpper("EndTime"), strings.ToUpper("End"):
		if v, ok := val.(string); ok {
			a.request.C2S.EndTime = proto.String(v)
		}
	}
}

//=== no need to modify
func (a *QotGetHoldingChangeList) UnPackBody(body []byte) Response {
	rsp := &qotgetholdingchangelist.Response{}
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
func (a *QotGetHoldingChangeList) GetC2S() interface{} {
	return a.request.C2S
}
func (a *QotGetHoldingChangeList) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
