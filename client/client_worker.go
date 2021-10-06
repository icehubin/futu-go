package client

import (
	"time"

	"github.com/icehubin/futu-go/logger"

	"github.com/icehubin/futu-go/adapt"
)

func in(e uint32, arr []uint32) bool {
	for _, v := range arr {
		if v == e {
			return true
		}
	}
	return false
}

func (w *Worker) Read() {
	pack := w.client.Read()
	header := pack.Header
	if nil == header {
		return
	}
	if header.GetProtoID() == adapt.ProtoID_Notify {
		if w.OpenSysHand {
			w.sChan <- pack
		}
		//未开启丢弃
	} else if in(header.GetProtoID(), []uint32{
		adapt.ProtoID_Qot_UpdateBasicQot,
		adapt.ProtoID_Qot_UpdateBroker,
		adapt.ProtoID_Qot_UpdateKL,
		adapt.ProtoID_Qot_UpdateOrderBook,
		adapt.ProtoID_Qot_UpdatePriceReminder,
		adapt.ProtoID_Qot_UpdateRT,
		adapt.ProtoID_Qot_UpdateTicker,
	}) {
		if w.OpenQuoteHand {
			w.qChan <- pack
		}
	} else if in(header.GetProtoID(), []uint32{
		adapt.ProtoID_Trd_UpdateOrder,
		adapt.ProtoID_Trd_UpdateOrderFill,
	}) {
		if w.OpenTradeHand {
			w.tChan <- pack
		}
	} else {
		if w.OpenDefaultHand {
			w.dChan <- pack
		}
	}
}

func (w *Worker) Work() {

	lc := make(chan interface{}, 1)

	var Reader = func() {
		defer func() {
			err := recover()
			w.client.Close()
			timeout := 3
			logger.Logger().Warn("Reader panic. waiting", timeout, "seconds for restart")
			time.Sleep(time.Second * time.Duration(timeout))
			lc <- err
		}()
		w.client = w.prepareClient()
		logger.Logger().Warn("Reader started")
		for {
			w.Read()
		}
	}

	sc, qc, tc, dc := make(chan int, 1), make(chan int, 1), make(chan int, 1), make(chan int, 1)
	var Handler = func(name string, listenC chan int, cb func(*ResPack), dataC chan *ResPack) {
		defer func() {
			_ = recover()
			timeout := 1
			logger.Logger().Warn("Handler", name, "panic. waiting ", timeout, "seconds for restart")
			time.Sleep(time.Second * time.Duration(timeout))
			listenC <- 1
		}()
		logger.Logger().Warn("Handler", name, "started")
		for {
			pack := <-dataC
			cb(pack)
		}
	}

	//发送启动消息
	lc <- "Please Reconnect" //启动读取线程
	if w.OpenSysHand {       //启动系统通知监控线程
		sc <- 1
	}
	if w.OpenQuoteHand { //启动行情推送监控线程
		qc <- 1
	}
	if w.OpenTradeHand { //启动交易通知监控线程
		tc <- 1
	}
	if w.OpenDefaultHand { //启动其他消息推送监控线程
		dc <- 1
	}
	t := time.NewTicker(time.Second * 5)
	for {
		select {
		case e := <-lc:
			msg, ok := e.(string)
			if ok && msg == "Please Reconnect" {
				go Reader()
			}
		case _ = <-sc:
			go Handler("sysNotifyHand", sc, w.sNotifyHand, w.sChan)
		case _ = <-qc:
			go Handler("quoteNotifyHand", qc, w.qNotifyHand, w.qChan)
		case _ = <-tc:
			go Handler("tradeNotifyHand", tc, w.tNotifyHand, w.tChan)
		case _ = <-dc:
			go Handler("defaultHandle", dc, w.defaultHandle, w.dChan)
		case <-t.C:
			//keepalive
			w.client.KeepAlive()
		}
	}
}

func (w *Worker) PrepareClient(cb func() *Client) {
	w.prepareClient = cb
}

func (w *Worker) SetQuoteNotifyHandle(cb func(*ResPack)) {
	w.qNotifyHand = cb
}

func (w *Worker) SetTrdNotifyHandle(cb func(*ResPack)) {
	w.tNotifyHand = cb
}

func (w *Worker) SetSysNotifyHandle(cb func(*ResPack)) {
	w.sNotifyHand = cb
}

func (w *Worker) SetDefaultHandle(cb func(*ResPack)) {
	w.defaultHandle = cb
}

func NewWorker() *Worker {
	worker := &Worker{
		qChan:           make(chan *ResPack, 1024),
		tChan:           make(chan *ResPack, 1024),
		sChan:           make(chan *ResPack, 1024),
		dChan:           make(chan *ResPack, 1024),
		OpenQuoteHand:   true, // 默认开启行情回调
		OpenSysHand:     true, // 默认开启系统回调
		OpenTradeHand:   true, // 默认开启交易回调
		OpenDefaultHand: true, // 默认开启其他回调
	}
	//empty handlers
	worker.qNotifyHand = func(*ResPack) {
		//do noth.
	}
	worker.tNotifyHand = func(*ResPack) {
		//do noth.
	}
	worker.sNotifyHand = func(*ResPack) {
		//do noth.
	}
	worker.defaultHandle = func(*ResPack) {
		//do noth.
	}
	return worker
}

type Worker struct {
	client          *Client
	prepareClient   func() *Client
	qNotifyHand     func(*ResPack)
	tNotifyHand     func(*ResPack)
	sNotifyHand     func(*ResPack)
	defaultHandle   func(*ResPack)
	qChan           chan *ResPack
	tChan           chan *ResPack
	sChan           chan *ResPack
	dChan           chan *ResPack
	OpenQuoteHand   bool
	OpenTradeHand   bool
	OpenSysHand     bool
	OpenDefaultHand bool
}
