# client.New/NewEncrypt/Create()
以上三个方法都是创建一个clent.Client对象

> 需要注意的是，OpenD是否有开启API协议加密，在配置了rsa_private_key启用了API协议加密后需要使用NewEncrypt来创建Client

```
clt, err := client.New("127.0.0.1:11111")
if err != nil {
	panic("Client 创建失败")
}
```
使用以上代码即可创建一个client.Client对象

```
clt, err := client.EncryptNew("127.0.0.1:11111","{path_to_rsa}")
if err != nil {
	panic("Client 创建失败")
}
```
开启了API加密的情况下使用这样的代码创建

以上两种方式默认都会开启系统通知，可以使用Create方法关闭

```
rsa_file := "path_to_rsa" //空字符串不加密
clt, err := client.Create("127.0.0.1:11111",rsa_file,false)
if err != nil {
	panic("Client 创建失败")
}
```
以上方法关闭系统通知

为什么需要关闭系统通知？
>在采用client.Sync来同步调用并等待结果返回的方式来使用的话，如果开启了系统通知，不能保证读到的返回结果就是自己需要的
因此通常在异步回调使用的情况下开启通知，同步使用的情况下关闭