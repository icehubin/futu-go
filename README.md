# futu-go简介

futu-go是富途牛牛OpenApi的Go语言SDK

>[富途牛牛OpenApi官方地址
](https://www.futunn.com/OpenAPI) https://www.futunn.com/OpenAPI

>[富途牛牛OpenApi官方文档地址](https://openapi.futunn.com/futu-api-doc/)https://openapi.futunn.com/futu-api-doc/

futu-go被极简设计，方便开发者对照现有官方文档就能方便的使用，以下是简单介绍

# futu-go的目录和包结构
| 目录名  | 包结构  | 备注  |
|:----------|:----------|:----------|
| adapt    | package adapt    | 富途协议适配代码目录，适配富途proto协议，文件毕竟多但需要关系的并不多，下面会具体介绍adapt的用途|	
|client|package client|futu-go的client代码，提供了便利的调用方法，下面会介绍几种典型的调用方式|
|logger|package logger|简单封装了一个日志模块，引用了github.com/sirupsen/logrus|
|futuproto|-|富途的protobuf定义文件，从富途Python SDK代码库中copy过来，不要修改他，目录内的build.sh用通过protoc生成GO语言的pb类定义文件|
|pb|子目录下分布定义|使用build.sh生产的GO语言的pb定义|
|examples/xxx|main|SDK调用的示例|

# logger介绍

logger是futu-go封装的日志打印接口，采用logrus格式

为了开发调试方便，默认的日志级别是DebugLevel级别，这样可以打印出所有交付日志方便开发过程测试，这样futu-go的单测程序都没有做判定，而是通过观察日志来验证结果。

正式运行时，可以通过 `logger.SetLevel(logger.WarnLevel)` 来改变日志级别，减少交互日志打印，但实际没有必要，不建议这样做。

# adapt介绍
adapt主要是协议适配，主要用到的有

1. adapt.ProtoID_{xxxx} 常量，调用proto协议编号
2. adapt.With() 方法，调用proto协议时设置参数

通常是配合client包来使用，比如以下代码：

首先创建一个client
	
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		panic("Client 创建失败")
	}


然后调用Sync方法调用一个adapt.ProtoID_GetGlobalState，这个调用没有参数

	//获取全局状态
	res := clt.Sync(adapt.ProtoID_GetGlobalState)

	fmt.Println(res)

然后调用Sync方法订阅数据，这个调用有参数列表

	//订阅数据
	res := clt.Sync(adapt.ProtoID_Qot_Sub,
		adapt.With("code_list", []string{"SH.600519", "SZ.300957"}),
		adapt.With("subtype_list", []string{"QUOTE", "K_5Min"}),
		adapt.With("IsFirstPush", true),
		adapt.With("push", true),
	)
	fmt.Println(res)

Sync这个方法接受的第一个参数是调用的协议编号，后面是可变参数adapt.Option，他是adapt.With方法的返回值。可以传入0-n个adapt.Option参数

3. adapt.Header

富途交互协议头的封装，主要需要关心的是：

Header.GetProtoID() 获取协议编号，返回的包里可以知道是哪个协议的返回
Header.GetSeriaNo() 获取包序号，通常在Async方式发送命令时候有用

4. adapt.Response

Response是响应结构体，定义了返回结构

	type Response struct {
		RetType int32  //返回结果，参见Common.RetType的枚举定义
		RetMsg  string //返回结果描述
		ErrCode int32  //错误码，客户端一般通过retType和retMsg来判断结果和详情，errCode只做日志记录，仅在个别协议失败时对账用
		S2C     proto.Message
	}
注意：S2C就是富途proto定义中的S2C，有些接口没有返回S2C，这个字段的值有可能是nil。这个字段的类型是proto.Message类型，使用时候需要注意，先判定不是nil后，在使用GetXxx()方法获取对应的值。或者可以使用adapt.PbParser()返回的PBMessageParser的Map方法，将其转成map[string]interface{}，这个类型可以直接被gota载入，类似Python的dataframe，做K线分析非常方便。

	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	//需要先订阅数据
	fmt.Println(clt.Sync(adapt.ProtoID_Qot_Sub,
		adapt.With("code_list", []string{"SH.600519", "SZ.300957"}),
		adapt.With("subtype_list", []string{"QUOTE", "TICKER", "K_DAY"}),
		adapt.With("IsFirstPush", true),
	))
	time.Sleep(time.Microsecond * 500)
	//获取K线列表
	res := clt.Sync(adapt.ProtoID_Qot_GetKL,
		adapt.With("code", "SZ.300957"),
		adapt.With("ktype", "K_DAY"),
		adapt.With("reqNum", int32(10)),
	)
	//拿到KlList节点数据
	Klist := res.S2C.(*qotgetkl.S2C).GetKlList()
	mp := make([]map[string]interface{}, 0)
	//循环转换每个节点
	for _, v := range Klist {
		mp = append(mp, adapt.PbParser().Map(v))
	}
	//需要 import github.com/go-gota/gota/dataframe
	df := dataframe.LoadMaps(mp)
	fmt.Println(df)


# client介绍

1. client.New/NewEncrypt/Create()

New创建一个没有加密的Client对象

NewEncrypt创建一个有加密的Client对象

Create按照自己的参数创建一个Client对象，New/NewEncrypt都是调用的他，可以设置Notify参数

2. client.ResPack

ResPack定义了回包结构体，一个包包括Header和Response

	type ResPack struct {
		Header   *adapt.Header
		Response *adapt.Response
	}



## client.Client
1. Client.KeepAlive()

定时发送KeepAlive包，避免被OpenD关闭链接

2. Sync()/Async()

同步调用/异步调用

3. Read()

读一个包，返回ResPack，异步场景使用，使用Worker框架的情况下自己代码基本用不到

## client.Worker

**client.Worker是什么？**

> client.Worker是一个超级超级轻量的调度框架，配合client.Client可以非常方便的实现你的策略逻辑

**client.Worker为什么要这样设计？**

> client.Worker这样设计是为了让你可以任意细粒度的拆分不同的数据策略到不同的goroutine里执行不同的任务，只要你愿意。

> 这需要配合QotSub协议，细粒度的订阅行情数据

> 这样做的好处是可以让你避免混合数据订阅的相互堵塞

> 更多的好处可以自己发掘

1. client.NewWorker()

创建client.Worker对象

2. Worker.PrepareClient()

PrepareClient接受一个 `func () \*client.Client` 类型的参数，可以传入一个返回值是 *client.Client类型的匿名函数

为Worker准备一个连接好的Client

3. Worker.SetDefaultHandle/SetQuoteNotifyHandle/SetSysNotifyHandle/SetTrdNotifyHandle()

设置回调处理方法

这些回调方法接受的参数类型一样，都是`func(*ResPack)`，这些方法里处理OpenD推送的消息，接受到的消息已经被处理为*ResPack类型，只用处理这个数据即可

SetQuoteNotifyHandle，设置行情推送回调

SetSysNotifyHandle，系统通知回调

SetTrdNotifyHandle，交易通知回调

SetDefaultHandle，以上三个除外的默认回调

具体示例可以看 examples/single 和 examples/multi

4. Worker.Work()

Worker开始干活，Work()方法会阻塞程序不退出，并且会监控回调goroutine的运行状态，goroutine挂掉会重新拉起，直接启动的话不用再执行for循环。具体可以看`examples/single`

如果想用不同的goroutine独立运行不同的回调分离处理，使用`go Worker.Work()`启动一个goroutine来运行，这样做有很多好处，比如可以避免行情回调太多阻塞了交易回调
这种方式运行，需要在程序最后加入for循环让主程序不退出。具体可以看`examples/multi`
