package adapt

const (
	ProtoID_InitConnect        = 1001 // 初始化连接
	ProtoID_GetGlobalState     = 1002 // 获取全局状态
	ProtoID_Notify             = 1003 // 通知推送
	ProtoID_KeepAlive          = 1004 // 心跳保活
	ProtoID_GetUserInfo        = 1005 // 获取用户信息
	ProtoID_Verification       = 1006 // 请求或输入验证码
	ProtoID_GetDelayStatistics = 1007 // 获取延迟统计
	ProtoID_TestCmd            = 1008 // futu-go已支持，没有文档，不知道支持哪些
	ProtoID_InitQuantMode      = 1009 //不支持，不知道干嘛的，也没有看到proto文件

	ProtoID_Trd_GetAccList  = 2001 // 获取业务账户列表
	ProtoID_Trd_UnlockTrade = 2005 // 解锁或锁定交易
	ProtoID_Trd_SubAccPush  = 2008 // 订阅业务账户的交易推送数据

	ProtoID_Trd_GetFunds        = 2101 // 获取账户资金
	ProtoID_Trd_GetPositionList = 2102 // 获取账户持仓

	ProtoID_Trd_GetOrderList = 2201 // 获取订单列表
	ProtoID_Trd_PlaceOrder   = 2202 // 下单
	ProtoID_Trd_ModifyOrder  = 2205 // 修改订单
	ProtoID_Trd_UpdateOrder  = 2208 // 订单状态变动通知(推送)

	ProtoID_Trd_GetOrderFillList = 2211 // 获取成交列表
	ProtoID_Trd_UpdateOrderFill  = 2218 // 成交通知(推送)

	ProtoID_Trd_GetHistoryOrderList     = 2221 // 获取历史订单列表
	ProtoID_Trd_GetHistoryOrderFillList = 2222 // 获取历史成交列表
	ProtoID_Trd_GetMaxTrdQtys           = 2111 // 查询最大买卖数量
	ProtoID_Trd_GetMarginRatio          = 2223 // 获取融资融券数据

	// 订阅数据
	ProtoID_Qot_Sub                 = 3001 // 订阅或者反订阅
	ProtoID_Qot_RegQotPush          = 3002 // 注册推送
	ProtoID_Qot_GetSubInfo          = 3003 // 获取订阅信息
	ProtoID_Qot_GetBasicQot         = 3004 // 获取股票基本行情
	ProtoID_Qot_UpdateBasicQot      = 3005 // 推送股票基本行情
	ProtoID_Qot_GetKL               = 3006 // 获取K线
	ProtoID_Qot_UpdateKL            = 3007 // 推送K线
	ProtoID_Qot_GetRT               = 3008 // 获取分时
	ProtoID_Qot_UpdateRT            = 3009 // 推送分时
	ProtoID_Qot_GetTicker           = 3010 // 获取逐笔
	ProtoID_Qot_UpdateTicker        = 3011 // 推送逐笔
	ProtoID_Qot_GetOrderBook        = 3012 // 获取买卖盘
	ProtoID_Qot_UpdateOrderBook     = 3013 // 推送买卖盘
	ProtoID_Qot_GetBroker           = 3014 // 获取经纪队列
	ProtoID_Qot_UpdateBroker        = 3015 // 推送经纪队列
	ProtoID_Qot_UpdatePriceReminder = 3019 //到价提醒通知

	// 历史数据
	ProtoID_Qot_RequestHistoryKL      = 3103 // 拉取历史K线
	ProtoID_Qot_RequestHistoryKLQuota = 3104 // 拉取历史K线已经用掉的额度
	ProtoID_Qot_RequestRehab          = 3105 // 获取除权信息

	// 其他行情数据
	ProtoID_Qot_GetSuspend           = 3201 // 获取股票停牌信息
	ProtoID_Qot_GetStaticInfo        = 3202 // 获取股票列表
	ProtoID_Qot_GetSecuritySnapshot  = 3203 // 获取股票快照
	ProtoID_Qot_GetPlateSet          = 3204 // 获取板块集合下的板块
	ProtoID_Qot_GetPlateSecurity     = 3205 // 获取板块下的股票
	ProtoID_Qot_GetReference         = 3206 // 获取正股相关股票，暂时只有窝轮
	ProtoID_Qot_GetOwnerPlate        = 3207 // 获取股票所属板块
	ProtoID_Qot_GetHoldingChangeList = 3208 // 获取高管持股变动
	ProtoID_Qot_GetOptionChain       = 3209 // 获取期权链

	ProtoID_Qot_GetWarrant             = 3210 // 拉取窝轮信息
	ProtoID_Qot_GetCapitalFlow         = 3211 // 获取资金流向
	ProtoID_Qot_GetCapitalDistribution = 3212 // 获取资金分布

	ProtoID_Qot_GetUserSecurity    = 3213 // 获取自选股分组下的股票
	ProtoID_Qot_ModifyUserSecurity = 3214 // 修改自选股分组下的股票
	ProtoID_Qot_StockFilter        = 3215 // 条件选股
	ProtoID_Qot_GetCodeChange      = 3216 // 代码变换
	ProtoID_Qot_GetIpoList         = 3217 // 获取新股Ipo
	ProtoID_Qot_GetFutureInfo      = 3218 // 获取期货资料
	ProtoID_Qot_RequestTradeDate   = 3219 // 在线拉取交易日
	ProtoID_Qot_SetPriceReminder   = 3220 // 设置到价提醒
	ProtoID_Qot_GetPriceReminder   = 3221 // 获取到价提醒

	ProtoID_Qot_GetUserSecurityGroup    = 3222 // 获取自选股分组
	ProtoID_Qot_GetMarketState          = 3223 // 获取指定品种的市场状态
	ProtoID_Qot_GetOptionExpirationDate = 3224 // 获取期权到期日

	ProtoID_Example_Adapt = 111111111 //just a tag fix Example template
)
