package adapt_test

import (
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func TestQotGetOptionChain(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	/*
		DataFilter:
			ImpliedVolatilityMin *float64 `protobuf:"fixed64,1,opt,name=impliedVolatilityMin" json:"impliedVolatilityMin,omitempty"` //隐含波动率过滤起点（精确到小数点后 0 位，超出部分会被舍弃）
			ImpliedVolatilityMax *float64 `protobuf:"fixed64,2,opt,name=impliedVolatilityMax" json:"impliedVolatilityMax,omitempty"` //隐含波动率过滤终点（精确到小数点后 0 位，超出部分会被舍弃）
			DeltaMin             *float64 `protobuf:"fixed64,3,opt,name=deltaMin" json:"deltaMin,omitempty"`                         //希腊值 Delta过滤起点（精确到小数点后 3 位，超出部分会被舍弃）
			DeltaMax             *float64 `protobuf:"fixed64,4,opt,name=deltaMax" json:"deltaMax,omitempty"`                         //希腊值 Delta过滤终点（精确到小数点后 3 位，超出部分会被舍弃）
			GammaMin             *float64 `protobuf:"fixed64,5,opt,name=gammaMin" json:"gammaMin,omitempty"`                         //希腊值 Gamma过滤起点（精确到小数点后 3 位，超出部分会被舍弃）
			GammaMax             *float64 `protobuf:"fixed64,6,opt,name=gammaMax" json:"gammaMax,omitempty"`                         //希腊值 Gamma过滤终点（精确到小数点后 3 位，超出部分会被舍弃）
			VegaMin              *float64 `protobuf:"fixed64,7,opt,name=vegaMin" json:"vegaMin,omitempty"`                           //希腊值 Vega过滤起点（精确到小数点后 3 位，超出部分会被舍弃）
			VegaMax              *float64 `protobuf:"fixed64,8,opt,name=vegaMax" json:"vegaMax,omitempty"`                           //希腊值 Vega过滤终点（精确到小数点后 3 位，超出部分会被舍弃）
			ThetaMin             *float64 `protobuf:"fixed64,9,opt,name=thetaMin" json:"thetaMin,omitempty"`                         //希腊值 Theta过滤起点（精确到小数点后 3 位，超出部分会被舍弃）
			ThetaMax             *float64 `protobuf:"fixed64,10,opt,name=thetaMax" json:"thetaMax,omitempty"`                        //希腊值 Theta过滤终点（精确到小数点后 3 位，超出部分会被舍弃）
			RhoMin               *float64 `protobuf:"fixed64,11,opt,name=rhoMin" json:"rhoMin,omitempty"`                            //希腊值 Rho过滤起点（精确到小数点后 3 位，超出部分会被舍弃）
			RhoMax               *float64 `protobuf:"fixed64,12,opt,name=rhoMax" json:"rhoMax,omitempty"`                            //希腊值 Rho过滤终点（精确到小数点后 3 位，超出部分会被舍弃）
			NetOpenInterestMin   *float64 `protobuf:"fixed64,13,opt,name=netOpenInterestMin" json:"netOpenInterestMin,omitempty"`    //净未平仓合约数过滤起点（精确到小数点后 0 位，超出部分会被舍弃）
			NetOpenInterestMax   *float64 `protobuf:"fixed64,14,opt,name=netOpenInterestMax" json:"netOpenInterestMax,omitempty"`    //净未平仓合约数过滤终点（精确到小数点后 0 位，超出部分会被舍弃）
			OpenInterestMin      *float64 `protobuf:"fixed64,15,opt,name=openInterestMin" json:"openInterestMin,omitempty"`          //未平仓合约数过滤起点（精确到小数点后 0 位，超出部分会被舍弃）
			OpenInterestMax      *float64 `protobuf:"fixed64,16,opt,name=openInterestMax" json:"openInterestMax,omitempty"`          //未平仓合约数过滤终点（精确到小数点后 0 位，超出部分会被舍弃）
			VolMin               *float64 `protobuf:"fixed64,17,opt,name=volMin" json:"volMin,omitempty"`                            //成交量过滤起点（精确到小数点后 0 位，超出部分会被舍弃）
			VolMax               *float64 `protobuf:"fixed64,18,opt,name=volMax" json:"volMax,omitempty"`                            //成交量过滤终点（精确到小数点后 0 位，超出部分会被舍弃）
	*/
	res := clt.Sync(adapt.ProtoID_Qot_GetOptionChain,
		adapt.With("code", "HK.00700"),
		adapt.With("begin", "2021-10-08"),
		adapt.With("end", "2021-10-30"),
		adapt.With("", adapt.Message{
			"type":      proto.Int32(1),
			"condition": proto.Int32(1),
		}),
		adapt.With("DataFilter", adapt.Message{
			"deltaMin": proto.Float64(0.1),
			"deltaMax": proto.Float64(0.9),
		}),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
