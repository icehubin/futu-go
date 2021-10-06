package adapt_test

import (
	"testing"

	"github.com/icehubin/futu-go/adapt"

	"github.com/icehubin/futu-go/client"
)

func TestTrdGetFunds(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	res := clt.Sync(adapt.ProtoID_Trd_GetFunds,
		adapt.With("Header", adapt.TrdHeader{
			TrdEnv:    1,
			AccID:     281756460285261810,
			TrdMarket: 2,
		}),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
