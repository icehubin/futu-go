package adapt_test

import (
	"testing"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
	"github.com/icehubin/futu-go/logger"
	"github.com/icehubin/futu-go/pb/trdgetacclist"
)

func TestTrdGetHistoryOrderFillList(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	res := clt.Sync(adapt.ProtoID_Trd_GetAccList)
	if res.RetType != adapt.RetType_Succeed {
		logger.WithField("error", "GetAccListFailed").Warn("Failed")
		return
	}
	acclist := res.S2C.(*trdgetacclist.S2C).GetAccList()
	if nil != acclist {
		for _, acc := range acclist {
			if acc.GetTrdEnv() != int32(1) {
				//跳过测试环境
				continue
			}
			trdMarket := acc.GetTrdMarketAuthList()[0]
			if trdMarket == int32(2) { //美股
				res = clt.Sync(adapt.ProtoID_Trd_GetHistoryOrderFillList,
					adapt.With("Header", adapt.TrdHeader{
						TrdEnv:    acc.GetTrdEnv(),
						AccID:     acc.GetAccID(),
						TrdMarket: acc.GetTrdMarketAuthList()[0],
					}),
					adapt.With("Conditions", adapt.TrdFilterConditions{
						CodeList:  []string{"TSLA"},
						BeginTime: "2021-01-01",
						EndTime:   "2022-01-01",
					}),
				)
			}
		}
	}

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
