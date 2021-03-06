//
package adapt

type AdaptInterface interface {
	PackBody() ([]byte, bool)                      //打包body
	UnPackBody(body []byte) Response               //解包 body 到response
	GetHeader() *Header                            //获取header
	GetProtoID() uint32                            //获取protoID
	GetC2S() interface{}                           //获取请求变量
	SetC2SOption(protoKey string, val interface{}) //设置C2S参数
	SetPacketID(packetID PacketID)                 //设置PacketID
}

type OptionInterface interface {
	setOption(ddp AdaptInterface, protoKey string, val int32)
}

type Option struct {
	protoKey string
	val      interface{}
}

func (o *Option) setOption(ddp AdaptInterface) {
	ddp.SetC2SOption(o.protoKey, o.val)
}

/*
* 创建一个数据包对象时传入的参数
* PARAMS:
* 	- protoKey: proto 文件中定义的 C2S的key
*	- val:	初始化值，注意：值的类型需要匹配，如果类型不匹配会被丢弃设置不成功
* RETURNS:
*	- Option 参数设置struct，通过他来实现到可变参数
 */
func With(protoKey string, val interface{}) Option {
	do := Option{
		protoKey: protoKey,
		val:      val,
	}
	return do
}

type adaptBase struct {
	header   *Header
	protoID  uint32
	packetID PacketID
	// server  *Server
}

func (a *adaptBase) GetHeader() *Header {
	return a.header
}

func (a *adaptBase) GetProtoID() uint32 {
	return a.protoID
}

func (a *adaptBase) GetC2S() interface{} {
	return make(map[string]interface{})
}

func (a *adaptBase) SetC2SOption(protoKey string, val interface{}) {
}

func (a *adaptBase) PackBody() ([]byte, bool) {
	return nil, false
}

func (a *adaptBase) SetPacketID(packetID PacketID) {
	a.packetID = packetID
}

func (a *adaptBase) setProtoID(protoID uint32) {
	a.header = &Header{}
	a.protoID = protoID
	a.header.nProtoID = protoID
}

type Server struct {
	ConnID            uint64
	KeepAliveInterval int32
	ConnAESKey        string
	AesCBCiv          string
	LoginUserID       uint64
	ServerVer         int32
	UserAttribution   int32

	Encrypt bool
}

type PacketID struct {
	ConnID   uint64 //
	SerialNo uint32
}

var DataAdaptMap = map[uint32]interface{}{

	ProtoID_InitConnect:        CreateInitConnect,        // 1001 // 初始化连接
	ProtoID_GetGlobalState:     CreateGetGlobalState,     // 1002 // 获取全局状态
	ProtoID_Notify:             CreateNotify,             // 1003 // 通知推送
	ProtoID_KeepAlive:          CreateKeepAlive,          // 1004 // 心跳保活
	ProtoID_GetUserInfo:        CreateGetUserInfo,        // 1005 // 获取用户信息
	ProtoID_Verification:       CreateVerification,       // 1006 // 请求或输入验证码
	ProtoID_GetDelayStatistics: CreateGetDelayStatistics, // 1007 // 获取延迟统计
	ProtoID_TestCmd:            CreateTestCmd,            // 1008
	/*
		Todo 没有proto文件
		ProtoID_InitQuantMode      = 1009
	*/

	// 交易相关
	ProtoID_Trd_GetAccList:              CreateTrdGetAccList,              //2001 // 获取业务账户列表
	ProtoID_Trd_UnlockTrade:             CreateTrdUnlockTrade,             //2005 // 解锁或锁定交易
	ProtoID_Trd_SubAccPush:              CreateTrdSubAccPush,              //2008 // 订阅业务账户的交易推送数据
	ProtoID_Trd_GetFunds:                CreateTrdGetFunds,                //2101 // 获取账户资金
	ProtoID_Trd_GetPositionList:         CreateTrdGetPositionList,         //2102 // 获取账户持仓
	ProtoID_Trd_GetOrderList:            CreateTrdGetOrderList,            //2201 // 获取订单列表
	ProtoID_Trd_PlaceOrder:              CreateTrdPlaceOrder,              //2202 // 下单
	ProtoID_Trd_ModifyOrder:             CreateTrdModifyOrder,             //2205 // 修改订单
	ProtoID_Trd_UpdateOrder:             CreateTrdUpdateOrder,             // 2208 // 订单状态变动通知(推送)
	ProtoID_Trd_GetOrderFillList:        CreateTrdGetOrderFillList,        // 2211 // 获取成交列表
	ProtoID_Trd_UpdateOrderFill:         CreateTrdUpdateOrderFill,         // 2218 // 成交通知(推送)
	ProtoID_Trd_GetHistoryOrderList:     CreateTrdGetHistoryOrderList,     // 2221 // 获取历史订单列表
	ProtoID_Trd_GetHistoryOrderFillList: CreateTrdGetHistoryOrderFillList, //2222 // 获取历史成交列表
	ProtoID_Trd_GetMaxTrdQtys:           CreateTrdGetMaxTrdQtys,           //2111 // 查询最大买卖数量
	ProtoID_Trd_GetMarginRatio:          CreateTrdGetMarginRatio,          //2223 // 获取融资融券数据

	// 行情相关
	ProtoID_Qot_Sub:                   CreateQotSub,                   // 3001 // 订阅或者反订阅
	ProtoID_Qot_RegQotPush:            CreateQotRegQotPush,            // 3002 // 注册推送
	ProtoID_Qot_GetSubInfo:            CreateQotGetSubInfo,            // 3003 // 获取订阅信息
	ProtoID_Qot_GetBasicQot:           CreateQotGetBasicQot,           // 3004 // 获取股票基本行情
	ProtoID_Qot_UpdateBasicQot:        CreateQotUpdateBasicQot,        // 3005 // 推送股票基本行情
	ProtoID_Qot_GetKL:                 CreateQotGetKL,                 // 3006 // 获取K线
	ProtoID_Qot_UpdateKL:              CreateQotUpdateKL,              // 3007 // 推送K线
	ProtoID_Qot_GetRT:                 CreateQotGetRT,                 // 3008 // 获取分时
	ProtoID_Qot_UpdateRT:              CreateQotUpdateRT,              // 3009 // 推送分时
	ProtoID_Qot_GetTicker:             CreateQotGetTicker,             // 3010 // 获取逐笔
	ProtoID_Qot_UpdateTicker:          CreateQotUpdateTicker,          // 3011 // 推送逐笔
	ProtoID_Qot_GetOrderBook:          CreateQotGetOrderBook,          // 3012 // 获取买卖盘
	ProtoID_Qot_UpdateOrderBook:       CreateQotUpdateOrderBook,       // 3013 // 推送买卖盘
	ProtoID_Qot_GetBroker:             CreateQotGetBroker,             // 3014 // 获取经纪队列
	ProtoID_Qot_UpdateBroker:          CreateQotUpdateBroker,          // 3015 // 推送经纪队列
	ProtoID_Qot_UpdatePriceReminder:   CreateQotUpdatePriceReminder,   // 3019 //到价提醒通知
	ProtoID_Qot_RequestHistoryKL:      CreateQotRequestHistoryKL,      // 3103 // 拉取历史K线
	ProtoID_Qot_RequestHistoryKLQuota: CreateQotRequestHistoryKLQuota, // 3104 // 拉取历史K线已经用掉的额度
	ProtoID_Qot_RequestRehab:          CreateQotRequestRehab,          // 3105 // 获取除权信息

	// 其他行情数据

	ProtoID_Qot_GetSuspend:              CreateQotGetSuspend,              // 3201 // 获取股票停牌信息
	ProtoID_Qot_GetStaticInfo:           CreateQotGetStaticInfo,           // 3202 // 获取股票列表
	ProtoID_Qot_GetSecuritySnapshot:     CreateQotGetSecuritySnapshot,     // 3203 // 获取股票快照
	ProtoID_Qot_GetPlateSet:             CreateQotGetPlateSet,             // 3204 // 获取板块集合下的板块
	ProtoID_Qot_GetPlateSecurity:        CreateQotGetPlateSecurity,        // 3205 // 获取板块下的股票
	ProtoID_Qot_GetReference:            CreateQotGetReference,            // 3206 // 获取正股相关股票，暂时只有窝轮
	ProtoID_Qot_GetOwnerPlate:           CreateQotGetOwnerPlate,           // 3207 // 获取股票所属板块
	ProtoID_Qot_GetHoldingChangeList:    CreateQotGetHoldingChangeList,    // 3208 // 获取高管持股变动
	ProtoID_Qot_GetOptionChain:          CreateQotGetOptionChain,          // 3209 // 获取期权链
	ProtoID_Qot_GetWarrant:              CreateQotGetWarrant,              // 3210 // 拉取窝轮信息
	ProtoID_Qot_GetCapitalFlow:          CreateQotGetCapitalFlow,          // 3211 // 获取资金流向
	ProtoID_Qot_GetCapitalDistribution:  CreateQotGetCapitalDistribution,  // 3212 // 获取资金分布
	ProtoID_Qot_GetUserSecurity:         CreateQotGetUserSecurity,         // 3213 // 获取自选股分组下的股票
	ProtoID_Qot_ModifyUserSecurity:      CreateQotModifyUserSecurity,      // 3214 // 修改自选股分组下的股票
	ProtoID_Qot_StockFilter:             CreateQotStockFilter,             // 3215 // 条件选股
	ProtoID_Qot_GetCodeChange:           CreateQotGetCodeChange,           // 3216 // 代码变换
	ProtoID_Qot_GetIpoList:              CreateQotGetIpoList,              // 3217 // 获取新股Ipo
	ProtoID_Qot_GetFutureInfo:           CreateQotGetFutureInfo,           // 3218 // 获取期货资料
	ProtoID_Qot_RequestTradeDate:        CreateQotRequestTradeDate,        // 3219 // 在线拉取交易日
	ProtoID_Qot_SetPriceReminder:        CreateQotSetPriceReminder,        // 3220 // 设置到价提醒
	ProtoID_Qot_GetPriceReminder:        CreateQotGetPriceReminder,        // 3221 // 获取到价提醒
	ProtoID_Qot_GetUserSecurityGroup:    CreateQotGetUserSecurityGroup,    // 3222 // 获取自选股分组
	ProtoID_Qot_GetMarketState:          CreateQotGetMarketState,          // 3223 // 获取指定品种的市场状态
	ProtoID_Qot_GetOptionExpirationDate: CreateQotGetOptionExpirationDate, // 3224 // 获取期权到期日
}
