package main

import (
	"time"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func main() {

	worker := client.NewWorker()

	worker.PrepareClient(func() *client.Client {
		//创建连接
		//
		clt, err := client.New("127.0.0.1:11111")
		if err != nil {
			panic("Client 创建失败")
		}

		//订阅数据
		clt.Async(adapt.ProtoID_Qot_Sub,
			adapt.With("code_list", []string{"SH.600519", "SZ.300957"}),
			adapt.With("subtype_list", []string{"QUOTE", "TICKER"}),
			adapt.With("IsFirstPush", true),
			adapt.With("push", true),
		)
		time.Sleep(time.Microsecond * 500)
		clt.Async(adapt.ProtoID_Qot_Sub,
			adapt.With("code_list", []string{"SH.600519"}),
			adapt.With("subtype_list", []string{"QUOTE", "ORDERBOOK"}),
			adapt.With("IsFirstPush", true),
			adapt.With("push", true),
		)
		time.Sleep(time.Microsecond * 500)
		clt.Async(adapt.ProtoID_Qot_Sub,
			adapt.With("code_list", []string{"HK.00700"}),
			adapt.With("subtype_list", []string{"QUOTE", "BROKER"}),
			adapt.With("IsFirstPush", true),
			adapt.With("push", true),
		)
		time.Sleep(time.Microsecond * 500)

		return clt
	})

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

	//设置回调方法
	worker.SetQuoteNotifyHandle(QuoteNotifyHand)
	worker.SetSysNotifyHandle(SysNotifyHand)
	worker.SetTrdNotifyHandle(TrdNotifyHand)
	worker.SetDefaultHandle(DefaultHand)

	//开始主循环
	worker.Work()

}
