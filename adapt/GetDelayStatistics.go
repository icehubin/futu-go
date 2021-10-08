package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/getdelaystatistics"
	"google.golang.org/protobuf/proto"
)

type GetDelayStatistics struct {
	request *getdelaystatistics.Request

	adaptBase
}

func CreateGetDelayStatistics(dopts ...Option) AdaptInterface {
	adp := &GetDelayStatistics{
		request: &getdelaystatistics.Request{
			C2S: &getdelaystatistics.C2S{
				TypeList: []int32{1, 2, 3},
			},
		},
	}
	adp.setProtoID(ProtoID_GetDelayStatistics)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *GetDelayStatistics) SetC2SOption(protoKey string, val interface{}) {
	/*
		TypeList     []int32 `protobuf:"varint,1,rep,name=typeList" json:"typeList,omitempty"`         //统计数据类型，DelayStatisticsType
		QotPushStage *int32  `protobuf:"varint,2,opt,name=qotPushStage" json:"qotPushStage,omitempty"` //行情推送统计的区间，行情推送统计时有效，QotPushStage
		SegmentList  []int32 `protobuf:"varint,3,rep,name=segmentList" json:"segmentList,omitempty"`   //统计分段，默认100ms以下以2ms分段，100ms以上以500，1000，2000，-1分段，-1表示无穷大。
	*/
	switch strings.ToUpper(protoKey) {
	case strings.ToUpper("TypeList"), strings.ToUpper("Types"):
		/*
			DelayStatisticsType_DelayStatisticsType_Unkonw     DelayStatisticsType = 0 //未知类型
			DelayStatisticsType_DelayStatisticsType_QotPush    DelayStatisticsType = 1 //行情推送统计
			DelayStatisticsType_DelayStatisticsType_ReqReply   DelayStatisticsType = 2 //请求回应统计
			DelayStatisticsType_DelayStatisticsType_PlaceOrder DelayStatisticsType = 3 //下单统计
		*/
		if v, ok := val.([]int32); ok {
			a.request.C2S.TypeList = v
		}
	case strings.ToUpper("SegmentList"), strings.ToUpper("Segments"):
		if v, ok := val.([]int32); ok {
			a.request.C2S.SegmentList = v
		}
	case strings.ToUpper("QotPushStage"):
		/*
			QotPushStage_QotPushStage_Unkonw QotPushStage = 0 // 未知
			QotPushStage_QotPushStage_SR2SS  QotPushStage = 1 //统计服务端处理耗时
			QotPushStage_QotPushStage_SS2CR  QotPushStage = 2 //统计网络耗时
			QotPushStage_QotPushStage_CR2CS  QotPushStage = 3 //统计OpenD处理耗时
			QotPushStage_QotPushStage_SS2CS  QotPushStage = 4 //统计服务器发出到OpenD发出的处理耗时
			QotPushStage_QotPushStage_SR2CS  QotPushStage = 5 //统计服务器收到数据到OpenD发出的处理耗时
		*/
		if v, ok := val.(int32); ok {
			a.request.C2S.QotPushStage = proto.Int32(v)
		}
	}
}

//=== no need to modify
func (a *GetDelayStatistics) UnPackBody(body []byte) Response {
	rsp := &getdelaystatistics.Response{}
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
func (a *GetDelayStatistics) GetC2S() interface{} {
	return a.request.C2S
}
func (a *GetDelayStatistics) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
