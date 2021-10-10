# client.ResPack
ResPack定义了回包结构体，一个包包括Header和Response

```
type ResPack struct {
	Header   *adapt.Header
	Response *adapt.Response
}
```

client.ResPak只在异步模式的回调中用到
client.Worker的各种handler接受的参数