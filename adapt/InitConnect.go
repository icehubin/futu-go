//
package adapt

import (
	"github.com/icehubin/futu-go/pb/common"
	"github.com/icehubin/futu-go/pb/initconnect"
	"google.golang.org/protobuf/proto"
)

type InitConnect struct {
	request  *initconnect.Request
	response *initconnect.Response

	adaptBase

	res_file string
}

/*
* 创建InitConnect数据包对象
* PARAMS:
*	- dopts ...option :
*		调用with方法到可选参数列表
*		example With("ClientID", "powered by icehu")
* RWTURNS:
*	- *InitConnect:
*		返回一个数据包处理对象
 */
func CreateInitConnect(dopts ...Option) AdaptInterface {
	adp := &InitConnect{
		request: &initconnect.Request{
			C2S: &initconnect.C2S{
				ClientID:            proto.String(DEFULAT_CLIENT_ID),
				ClientVer:           proto.Int32(CLIENT_VERSION),
				RecvNotify:          proto.Bool(true),
				PacketEncAlgo:       proto.Int32(int32(common.PacketEncAlgo_PacketEncAlgo_None)),
				PushProtoFmt:        proto.Int32(0),
				ProgrammingLanguage: proto.String("GO"),
			},
		},
		response: &initconnect.Response{},
	}
	adp.setProtoID(ProtoID_InitConnect)
	for _, opt := range dopts {
		opt.setOption(adp)
	}
	return adp
}

func (a *InitConnect) SetC2SOption(protoKey string, val interface{}) {
	/*
		ClientVer  *int32  `protobuf:"varint,1,req,name=clientVer" json:"clientVer,omitempty"`   //客户端版本号，clientVer = "."以前的数 * 100 + "."以后的，举例：1.1版本的clientVer为1 * 100 + 1 = 101，2.21版本为2 * 100 + 21 = 221
		ClientID   *string `protobuf:"bytes,2,req,name=clientID" json:"clientID,omitempty"`      //客户端唯一标识，无生具体生成规则，客户端自己保证唯一性即可
		RecvNotify *bool   `protobuf:"varint,3,opt,name=recvNotify" json:"recvNotify,omitempty"` //此连接是否接收市场状态、交易需要重新解锁等等事件通知，true代表接收，FutuOpenD就会向此连接推送这些通知，反之false代表不接收不推送
		//如果通信要加密，首先得在FutuOpenD和客户端都配置RSA密钥，不配置始终不加密
		//如果配置了RSA密钥且指定的加密算法不为PacketEncAlgo_None则加密(即便这里不设置，配置了RSA密钥，也会采用默认加密方式)，默认采用FTAES_ECB算法
		PacketEncAlgo       *int32  `protobuf:"varint,4,opt,name=packetEncAlgo" json:"packetEncAlgo,omitempty"`            //指定包加密算法，参见Common.PacketEncAlgo的枚举定义
		PushProtoFmt        *int32  `protobuf:"varint,5,opt,name=pushProtoFmt" json:"pushProtoFmt,omitempty"`              //指定这条连接上的推送协议格式，若不指定则使用push_proto_type配置项
		ProgrammingLanguage *string `protobuf:"bytes,6,opt,name=programmingLanguage" json:"programmingLanguage,omitempty"` //接口编程语言，用于统计语言偏好

	*/
	switch protoKey {
	case "RecvNotify", "recvNotify":
		if v, ok := val.(bool); ok {
			a.SetRecvNotify(v)
		}
	case "PacketEncAlgo", "packetEncAlgo":
		if v, ok := val.(int32); ok {
			a.SetPacketEncAlgo(v)
		}
	case "PushProtoFmt", "pushProtoFmt":
		if v, ok := val.(int32); ok {
			a.SetPushProtoFmt(v)
		}
	case "ClientID", "clientID":
		if v, ok := val.(string); ok {
			a.request.C2S.ClientID = proto.String(v)
		}
	case "rsa_file":
		if v, ok := val.(string); ok {
			a.res_file = v
		}
	}
}

/*
*
 */
func (a *InitConnect) SetRecvNotify(recvNotify bool) {
	a.request.C2S.RecvNotify = proto.Bool(recvNotify)
}

func (a *InitConnect) SetPacketEncAlgo(packetEncAlgo int32) {
	a.request.C2S.PacketEncAlgo = proto.Int32(packetEncAlgo)
}

func (a *InitConnect) SetPushProtoFmt(pushProtoFmt int32) {
	a.request.C2S.PushProtoFmt = proto.Int32(pushProtoFmt)
}

//=== no need to modify
func (a *InitConnect) UnPackBody(body []byte) Response {
	rsp := &initconnect.Response{}
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
func (a *InitConnect) GetC2S() interface{} {
	return a.request.C2S
}
func (a *InitConnect) PackBody() ([]byte, bool) {
	body_pack, ok := proto.Marshal(a.request)
	if ok != nil {
		return nil, false
	}
	return body_pack, true
}
