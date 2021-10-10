package adapt

import (
	"google.golang.org/protobuf/proto"
)

type Response struct {
	RetType int32  //返回结果，参见adapt.RetType_XXX的枚举定义
	RetMsg  string //返回结果描述
	ErrCode int32  //错误码，客户端一般通过retType和retMsg来判断结果和详情，errCode只做日志记录，仅在个别协议失败时对账用
	S2C     proto.Message
	Data    map[string]interface{}
}

func DefaultErr() Response {
	return Response{
		RetType: RetType_Unknown,
		RetMsg:  "unKnow",
		ErrCode: 110,
	}
}

func PackErr() Response {
	return Response{
		RetType: RetType_Invalid,
		RetMsg:  "packErr",
		ErrCode: 110,
	}
}
