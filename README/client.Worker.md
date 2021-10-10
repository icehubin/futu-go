# client.Worker是什么？

> client.Worker是一个超级超级轻量的调度框架，配合client.Client可以非常方便的实现你的策略逻辑

# client.Worker为什么要这样设计？

> client.Worker这样设计是为了让你可以任意细粒度的拆分不同的数据策略到不同的goroutine里执行不同的任务，只要你愿意。

> 这需要配合QotSub协议，细粒度的订阅行情数据

> 这样做的好处是可以让你避免混合数据订阅的相互堵塞

> 更多的好处可以自己发掘

1. client.NewWorker()

创建client.Worker对象

2. Worker.PrepareClient()

PrepareClient接受一个 **`func () *client.Client`** 类型的参数，可以传入一个返回值是 *client.Client类型的匿名函数

这个方法需要为Worker准备一个连接好的Client

3. Worker.SetDefaultHandle/SetQuoteNotifyHandle/SetSysNotifyHandle/SetTrdNotifyHandle()

设置回调处理方法

这些回调方法接受的参数类型一样，都是 **`func(*client.ResPack)`**，这些方法里处理OpenD推送的消息，接受到的消息已经被处理为*ResPack类型，只用处理这个数据即可

SetQuoteNotifyHandle，设置行情推送回调

SetSysNotifyHandle，系统通知回调

SetTrdNotifyHandle，交易通知回调

SetDefaultHandle，以上三个除外的默认回调

具体示例可以看 examples/single 和 examples/multi

4. Worker.Work()

Worker开始干活，Work()方法会阻塞程序不退出，并且会监控回调goroutine的运行状态，goroutine挂掉会重新拉起，直接启动的话不用再执行for循环。具体可以看 **`examples/single`**

如果想用不同的goroutine独立运行不同的回调，分离回到数据处理，使用 **`go Worker.Work()`** 启动一个goroutine来运行，这样做有很多好处，比如可以避免行情回调太多阻塞了交易回调
这种方式运行，需要在程序最后加入for循环让主程序不退出。具体可以看 **`examples/multi`**
