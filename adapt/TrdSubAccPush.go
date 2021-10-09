//
package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/trdsubaccpush"
	"google.golang.org/protobuf/proto"
)

type TrdSubAccPush struct {
	request *trdsubaccpush.Request

	adaptBase
}

func CreateTrdSubAccPush(dopts ...Option) AdaptInterface {
	//Todo fix request format
	adp := &TrdSubAccPush{
		request: &trdsubaccpush.Request{
			C2S: &trdsubaccpush.C2S{},
		},
	}
	adp.setProtoID(ProtoID_Trd_SubAccPush)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *TrdSubAccPush) SetC2SOption(protoKey string, val interface{}) {
	/*
		AccIDList []uint64 `protobuf:"varint,1,rep,name=accIDList" json:"accIDList,omitempty"` //要接收推送数据的业务账号列表，全量非增量，即使用者请每次传需要接收推送数据的所有业务账号
	*/
	switch strings.ToUpper(protoKey) {
	case "":
		//尝试直接设置所有普调变量
		if v, ok := val.(Message); ok {
			protoFill(a.request.C2S, v)
		}
	case strings.ToUpper("AccIDList"), strings.ToUpper("AccIDs"):
		if v, ok := val.([]uint64); ok {
			a.request.C2S.AccIDList = v
		}
	}
}

//=== no need to modify
func (a *TrdSubAccPush) UnPackBody(body []byte) Response {
	rsp := &trdsubaccpush.Response{}
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
func (a *TrdSubAccPush) GetC2S() interface{} {
	return a.request.C2S
}
func (a *TrdSubAccPush) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
