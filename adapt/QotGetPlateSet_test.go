package adapt_test

import (
	"testing"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
	"google.golang.org/protobuf/proto"
)

func TestQotGetPlateSet(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	/*

		QotMarket_QotMarket_Unknown       QotMarket = 0  //未知市场
		QotMarket_QotMarket_HK_Security   QotMarket = 1  //香港市场
		QotMarket_QotMarket_HK_Future     QotMarket = 2  //港期货(已废弃，使用QotMarket_HK_Security即可)
		QotMarket_QotMarket_US_Security   QotMarket = 11 //美国市场
		QotMarket_QotMarket_CNSH_Security QotMarket = 21 //沪股市场
		QotMarket_QotMarket_CNSZ_Security QotMarket = 22 //深股市场
		QotMarket_QotMarket_SG_Security   QotMarket = 31 //新加坡市场
		QotMarket_QotMarket_JP_Security   QotMarket = 41 //日本市场

		PlateSetType_PlateSetType_All      PlateSetType = 0 //所有板块
		PlateSetType_PlateSetType_Industry PlateSetType = 1 //行业板块
		PlateSetType_PlateSetType_Region   PlateSetType = 2 //地域板块,港美股市场的地域分类数据暂为空
		PlateSetType_PlateSetType_Concept  PlateSetType = 3 //概念板块
		PlateSetType_PlateSetType_Other    PlateSetType = 4 //其他板块, 仅用于3207（获取股票所属板块）协议返回,不可作为其他协议的请求参数

	*/
	res := clt.Sync(adapt.ProtoID_Qot_GetPlateSet,
		// adapt.With("market", int32(21)),
		// adapt.With("plateSetType", int32(1)),
		adapt.With("", adapt.Message{
			"market":       proto.Int32(21),
			"plateSetType": proto.Int32(2),
		}),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
