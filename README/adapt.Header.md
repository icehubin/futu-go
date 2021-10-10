# adapt.Header
富途交互协议头的封装，主要需要关心的是：

**`Header.GetProtoID() `** 获取协议编号，返回的包里可以知道是哪个协议的返回
**`Header.GetSeriaNo() `** 获取包序号，通常在Async方式发送命令时候有用

对于Header通常通常只在异步方式下的Handler回调中用来区分是什么通知