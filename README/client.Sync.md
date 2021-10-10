# client.Sync/Async
client.Sync 同步调用，有返回值
client.Async 异步调用，无返回值
两个方法调用参数完全一样

一个典型的调用方式如下：
```
res := clt.Sync(adapt.ProtoID_Qot_GetKL,
		adapt.With("code", "SZ.300957"),
		adapt.With("ktype", "K_DAY"),
		adapt.With("", adapt.Message{
			"reqNum": proto.Int32(5),
		}),
	)
```

- 第一个参数 **`adapt.ProtoID_Qot_GetKL`** 是具体的命令号，具体看[支持的命令列表](../adapt/consts_ProtoID.go)
- **`adapt.With`** 指定可变参数列表，不同的命令接受不同参数，具体查看[如何输入参数](adapt.With.md)
- 具体接受参数在proto文件中定义由 message C2S
- 以上方法返回一个[adapt.Response](adapt.Response.md)类型

Qot_GetKL.proto是如下定义的C2S
```
message C2S
{
	required int32 rehabType = 1; //Qot_Common.RehabType,复权类型
	required int32 klType = 2; //Qot_Common.KLType,K线类型
	required Qot_Common.Security security = 3; //股票
	required int32 reqNum = 4; //请求K线根数
}
```
