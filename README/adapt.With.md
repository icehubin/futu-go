# 使用adapt.With输入参数

还是先看[clinet.Sync](client.Sync.md)中输入参数的例子，以此来说明clinet.With支持的几种参数输入方式

Go代码：
```
res := clt.Sync(adapt.ProtoID_Qot_GetKL,
		adapt.With("code", "SZ.300957"),
		adapt.With("ktype", "K_DAY"),
		adapt.With("", adapt.Message{
			"reqNum": proto.Int32(5),
		}),
	)
```

proto文件定义的C2S
```
message C2S
{
	required int32 rehabType = 1; //Qot_Common.RehabType,复权类型
	required int32 klType = 2; //Qot_Common.KLType,K线类型
	required Qot_Common.Security security = 3; //股票
	required int32 reqNum = 4; //请求K线根数
}
```

- ## 直接设置基本类型的参数
```
adapt.With("", adapt.Message{
	"reqNum": proto.Int32(5),
}
```
当adapt.With第一个参数传入空字符串""时，可以直接设置C2S的所有基本类型参数，包括intN/uintN/string/float64等
和基本类型的数组[]intN,[]uintN,[]string,[]float64 等。proto文件中定义的 **`repeated`**，就是数组类型
此时第二个参数入 **`adapt.Message`** 类型，他实际是一个 map[string]interface{}类型的map
**`adapt.Message`** 的key是proto中定义的key就是proto message中定义的字段key，注意这里严格区分大小写
**`adapt.Message`** 的值，需要用proto.{Type}(具体的值)来设置。上面例子中的proto定义的说int32，所以Go里用对应的 **`proto.Int32`**
**`adapt.Message`** 的类型如果不对，设置的值不会生效
此例中的 rehabType、klType也可以直接设置，如下代码：
```
adapt.With("", adapt.Message{
	"reqNum":      proto.Int32(5),
	"rehabType":   proto.Int32(1),
	"klType":      proto.Int32(1),
}
```
为什么没有直接这样做，往后面看

- ## 设置具体字段的值
如果第一个参数传入了字段值，则设置具体字段的值
```
adapt.With("code", "SZ.300957"),
adapt.With("ktype", "K_DAY"),
```
上例中的这两段代码就是设置具体字段的值
这里有疑问，code、ktype是什么？
第一个字段的值实际是proto message中定义的字段名，但是为了方便futu-go对这个值做了大小写适配，并且对某些字段进行了别名处理，看看源代码就明白了

```
...
case "SECURITY", "CODE":
		if v, ok := val.(string); ok {
			nv := Stock2Security(v)
			a.request.C2S.Security = nv
		}
	case "KTYPE", "KLTYPE", "K_TYPE":
		if v, ok := val.(int32); ok {
			a.request.C2S.KlType = proto.Int32(v)
		} else if v, ok := val.(string); ok {
		switch strings.ToUpper(v) {
			case "KLTYPE_1MIN", "KL_1MIN", "K_1MIN":
				a.request.C2S.KlType = proto.Int32(1)
			case "KLTYPE_DAY", "KL_DAY", "K_DAY":
				a.request.C2S.KlType = proto.Int32(2)
		...
```
所以，对于单独设置单个字段的值：

- 大小写不敏感
- 有一些别名，比如code是security的别名，ktype、k_type都是klType的别名
- 有一些类型转换，比如klType明明是int32类型的，但是可以使用字符串类型来便利设置
- 有一个特列是 **`Qot_Common.Security`** 这个类型因为用得太频繁，对他做了转换直接传入字符串类型即可，格式如：US.TSLA、HK.00700、SH.300957
- 另一个特列 **`repeated Qot_Common.Security`**，也经常使用，传入字符串数组即可 **`[]string{"US.TSLA","HK.00700","SH.300957"}`**
- 注意以上两个特列使用便捷方式只能单独设置

- ## 设置结构引用
以 **`Qot_GetOptionChain.proto`** 为例
```
message C2S
{
    ...
	optional DataFilter dataFilter = 7; //数据字段筛选
}
message DataFilter
{
...
	optional double deltaMin = 3; //希腊值 Delta过滤起点（精确到小数点后 3 位，超出部分会被舍弃）
	optional double deltaMax = 4; //希腊值 Delta过滤终点（精确到小数点后 3 位，超出部分会被舍弃）

	optional double gammaMin = 5; //希腊值 Gamma过滤起点（精确到小数点后 3 位，超出部分会被舍弃）
	optional double gammaMax = 6; //希腊值 Gamma过滤终点（精确到小数点后 3 位，超出部分会被舍弃）
...
}
```

此例中的 dataFilter 是另一个message，如何设置？
Go代码如下：
```
adapt.With("DataFilter", adapt.Message{
	"deltaMin": proto.Float64(0.1),
	"deltaMax": proto.Float64(0.9),
}),
```
- 前面说过，第一个参数是大小写不敏感的
- 第二个参数传入 **`adapt.Message`** 跟前面是一样的，对应的map列表，就去看 **`message DataFilter`** 定义对应的填进去就可以了
- 这里同样也要注意，数据类型需要匹配才能设置成功
- 至此，你应该已经知道怎么设置结构体参数了，用同样的方法也是可以设置Security的，但是没必要，market编号容易忘

- ## 设置重复的结构引用

```
message C2S
{
...
	repeated BaseFilter baseFilterList = 5; // 简单指标过滤器
	repeated AccumulateFilter accumulateFilterList = 6; // 累积指标过滤器
	repeated FinancialFilter financialFilterList = 7; // 财务指标过滤器
	repeated PatternFilter patternFilterList = 8; // 形态技术指标过滤器
	repeated CustomIndicatorFilter customIndicatorFilterList = 9; // 自定义技术指标过滤器
...	
}

message BaseFilter 
{ 
	required int32 fieldName = 1; // StockField 简单属性
	optional double filterMin = 2; // 区间下限（闭区间），不传代表下限为 -∞
	optional double filterMax = 3; // 区间上限（闭区间），不传代表上限为 +∞
	optional bool isNoFilter = 4; // 该字段是否不需要筛选，True：不筛选，False：筛选。不传默认不筛选
	optional int32 sortDir = 5; // SortDir 排序方向，默认不排序。
}
```
如何设置 baseFilterList？

- 这里跟设置结构引用没有什么差别，就是多次adapt.With即可

```
adapt.With("baseFilterList",adapt.Message{
    "fieldName" : proto.Int32(1),
    "filterMin" : proto.Float64(0.1),
    ...
}),
adapt.With("baseFilterList",adapt.Message{
    "fieldName" : proto.Int32(2),
    "filterMin" : proto.Float64(0.1),
    ...
}),
adapt.With("baseFilterList",adapt.Message{
    "fieldName" : proto.Int32(3),
    "filterMin" : proto.Float64(0.1),
    ...
}),
```
实际上，用这个方法，也可以设置 **`repeated Qot_Common.Security`** ，但是没有必要，高频参数已经做了便利化转换。

至此

所有的futu参数你都可以便利传入，而不用引入proto生成的那堆几十个package的go对象。

futuproto目录下的proto文件，当一个文档来看即可
