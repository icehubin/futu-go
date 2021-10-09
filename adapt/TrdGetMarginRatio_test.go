package adapt_test

import (
	"testing"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
	"github.com/icehubin/futu-go/logger"
	"github.com/icehubin/futu-go/pb/trdgetacclist"
	"google.golang.org/protobuf/proto"
)

func TestTrdGetMarginRatio(t *testing.T) {
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
				res = clt.Sync(adapt.ProtoID_Trd_GetMarginRatio,
					adapt.With("Header", adapt.Message{
						"trdEnv":    proto.Int32(acc.GetTrdEnv()),
						"accID":     proto.Uint64(acc.GetAccID()),
						"trdMarket": proto.Int32(trdMarket),
					}),
					adapt.With("code_list", []string{"US.TSLA", "US.BIDU"}),
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
