# futu-go简介

futu-go是富途牛牛OpenApi的Go语言SDK

>[富途牛牛OpenApi官方地址
](https://www.futunn.com/OpenAPI) https://www.futunn.com/OpenAPI

>[富途牛牛OpenApi官方文档地址](https://openapi.futunn.com/futu-api-doc/)https://openapi.futunn.com/futu-api-doc/

futu-go被极简设计，一切为了方便开发者使用，方便开发者理解，最大限度减少开发者的依赖，把复杂的细节都cover在SDK内部。
目标是开发在花10-20分钟阅读完本文档，然后对照官方的.proto文件定义，再结合其他一些信息即可快速上手。
以下是介绍

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

1. [adapt.ProtoID_{xxxx}](adapt.consts_ProtoID.go) 常量，调用proto协议编号
2. [adapt.With()](README/adapt.With.md) 方法，调用proto协议时设置参数

3. [adapt.Header](README/adapt.Header.md) 富途协议头封装

4. [adapt.Response](README/adapt.Response.md) 调用client.Sync的返回值

# client介绍

1. [client.New/NewEncrypt/Create()](README/client.New.md)
    创建Client对象
2. [client.ResPack](README/client.ResPack.md)
    异步回调返回的包体
3. [client.Worker](README/client.Worker.md)
    一个极简多goroutine调度框架
    使用client.Worker可以轻松实现复杂策略逻辑
4. Client.KeepAlive()
    富途OpenD需要周期性发送KeepAlive()请求，否则会被关闭链接，使用client.Worker的调度已经自动调用不用关系。
