package main

import (
	"time"

	"github.com/icehubin/futu-go/pb/trdgetacclist"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func main() {

	var QuoteQuote = func() *client.Client {
		//创建连接
		//
		clt, err := client.New("127.0.0.1:11111")
		if err != nil {
			panic("Client 创建失败")
		}

		//订阅数据
		clt.Async(adapt.ProtoID_Qot_Sub,
			adapt.With("code_list", []string{"SH.600519", "SZ.300957"}),
			adapt.With("subtype_list", []string{"QUOTE", "K_5Min"}),
			adapt.With("IsFirstPush", true),
			adapt.With("push", true),
		)

		return clt
	}

	var QuoteOrderbook = func() *client.Client {
		//创建连接
		//
		clt, err := client.New("127.0.0.1:11111")
		if err != nil {
			panic("Client 创建失败")
		}

		//订阅数据
		clt.Async(adapt.ProtoID_Qot_Sub,
			adapt.With("code_list", []string{"SH.600519"}),
			adapt.With("subtype_list", []string{"QUOTE", "ORDERBOOK"}),
			adapt.With("IsFirstPush", true),
			adapt.With("push", true),
		)

		return clt
	}
	var QuoteBroker = func() *client.Client {
		//创建连接
		//
		clt, err := client.New("127.0.0.1:11111")
		if err != nil {
			panic("Client 创建失败")
		}

		//订阅数据
		clt.Async(adapt.ProtoID_Qot_Sub,
			adapt.With("code_list", []string{"HK.00700"}),
			adapt.With("subtype_list", []string{"QUOTE", "BROKER"}),
			adapt.With("IsFirstPush", true),
			adapt.With("push", true),
		)

		return clt
	}
	var TradeBroker = func() *client.Client {
		//创建连接
		//
		clt, err := client.New("127.0.0.1:11111")
		if err != nil {
			panic("Client 创建失败")
		}

		//获取账户列表
		res := clt.Sync(adapt.ProtoID_Trd_GetAccList)
		if res.RetType == adapt.RetType_Succeed {
			//注册交易推送
			accs := make([]uint64, 0)

			if _, ok := res.S2C.(*trdgetacclist.S2C); ok {
				for _, acc := range res.S2C.(*trdgetacclist.S2C).GetAccList() {
					accs = append(accs, acc.GetAccID())
				}
			}

			clt.Sync(adapt.ProtoID_Trd_SubAccPush,
				adapt.With("accids", accs),
			)
		}

		return clt
	}
	//行情通知回调方法
	var QuoteNotifyHand = func(res *client.ResPack) {
		//do sth.
		//your code
	}
	//交易回调方法
	var TrdNotifyHand = func(res *client.ResPack) {
		//do sth.
		//your code
	}
	//系统通知回调
	var SysNotifyHand = func(res *client.ResPack) {
		//do sth.
		//your code
	}
	//默认回调
	var DefaultHand = func(res *client.ResPack) {
		//do sth.
		//your code
	}

	//起一个单独线程监控行情
	worker := client.NewWorker()
	worker.PrepareClient(QuoteQuote)
	worker.SetQuoteNotifyHandle(QuoteNotifyHand)
	worker.OpenSysHand = false    //fale，关掉系统回调
	worker.OpenTradeHand = false  //fale，关掉系统回调
	worker.OpenDefaultHand = true //默认就是true，可以不写
	worker.OpenQuoteHand = true   //默认就是true，可以不写
	go worker.Work()

	//起一个单独线程监控挂单摆盘
	worker = client.NewWorker()
	worker.PrepareClient(QuoteOrderbook)
	worker.SetQuoteNotifyHandle(QuoteNotifyHand)
	worker.OpenSysHand = false    //fale，关掉系统回调
	worker.OpenTradeHand = false  //fale，关掉系统回调
	worker.OpenDefaultHand = true //默认就是true，可以不写
	worker.OpenQuoteHand = true   //默认就是true，可以不写
	go worker.Work()

	//起一个单独线程监控券商经纪
	worker = client.NewWorker()
	worker.PrepareClient(QuoteBroker)
	worker.SetQuoteNotifyHandle(QuoteNotifyHand)
	worker.OpenSysHand = false    //fale，关掉系统回调
	worker.OpenTradeHand = false  //fale，关掉系统回调
	worker.OpenDefaultHand = true //默认就是true，可以不写
	worker.OpenQuoteHand = true   //默认就是true，可以不写
	go worker.Work()

	//起一个单独线程监控交易信息
	worker = client.NewWorker()
	worker.PrepareClient(TradeBroker)
	worker.SetSysNotifyHandle(SysNotifyHand)
	worker.SetTrdNotifyHandle(TrdNotifyHand)
	worker.SetDefaultHandle(DefaultHand)
	worker.OpenSysHand = true     //默认就是true，可以不写
	worker.OpenTradeHand = true   //默认就是true，可以不写
	worker.OpenDefaultHand = true //默认就是true，可以不写
	worker.OpenQuoteHand = false  //fale，关掉行情回调
	go worker.Work()
	//设置回调方法

	for {
		//do noth. sleep
		time.Sleep(time.Second * 10)
	}

}
