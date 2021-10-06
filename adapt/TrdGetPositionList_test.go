package adapt_test

import (
	"testing"

	"github.com/icehubin/futu-go/logger"
	"github.com/icehubin/futu-go/pb/trdgetacclist"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func TestTrdGetPositionList(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	/*
		Header       *trdcommon.TrdHeader  `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`             //交易公共参数头
		PositionList []*trdcommon.Position `protobuf:"bytes,2,rep,name=positionList" json:"positionList,omitempty"` //持仓列表
	*/

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
			res = clt.Sync(adapt.ProtoID_Trd_GetPositionList,
				adapt.With("Header", adapt.TrdHeader{
					TrdEnv:    acc.GetTrdEnv(),
					AccID:     acc.GetAccID(),
					TrdMarket: acc.GetTrdMarketAuthList()[0],
				}),
			)
		}
	}

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
