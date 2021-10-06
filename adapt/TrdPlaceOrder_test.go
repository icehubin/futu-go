package adapt_test

import (
	"crypto/md5"
	"fmt"
	"testing"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
	"github.com/icehubin/futu-go/logger"
	"github.com/icehubin/futu-go/pb/trdgetacclist"
)

func TestTrdPlaceOrder(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	//解锁交易
	fmt.Println("请输入解锁密码：")
	var passwd string
	fmt.Scanf("%s", &passwd)
	// passwd = "111111"
	pwdmd5 := fmt.Sprintf("%x", md5.Sum([]byte(passwd)))

	res := clt.Sync(adapt.ProtoID_Trd_UnlockTrade,
		adapt.With("unlock", true),
		adapt.With("SecurityFirm", int32(1)),
		adapt.With("pwd", pwdmd5),
	)

	//获取账户列表
	res = clt.Sync(adapt.ProtoID_Trd_GetAccList)
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
				//
				res = clt.Sync(adapt.ProtoID_Trd_PlaceOrder,
					adapt.With("Header", adapt.TrdHeader{
						TrdEnv:    acc.GetTrdEnv(),
						AccID:     acc.GetAccID(),
						TrdMarket: acc.GetTrdMarketAuthList()[0],
					}),
					adapt.With("code", "US.TSLA"),
					adapt.With("trdside", "buy"),
					adapt.With("qty", float64(1)),
					adapt.With("price", float64(1.1)),
				)
			}
			if trdMarket == int32(1) { //港股
				res = clt.Sync(adapt.ProtoID_Trd_PlaceOrder,
					adapt.With("Header", adapt.TrdHeader{
						TrdEnv:    acc.GetTrdEnv(),
						AccID:     acc.GetAccID(),
						TrdMarket: acc.GetTrdMarketAuthList()[0],
					}),
					adapt.With("code", "HK.00700"),
					adapt.With("trdside", "buy"),
					adapt.With("qty", float64(100)),
					adapt.With("price", float64(400)),
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
