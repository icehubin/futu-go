package adapt_test

import (
	"testing"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/pb/trdgetacclist"

	"github.com/icehubin/futu-go/client"
)

func TestTrdSubAccPush(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	res := clt.Sync(adapt.ProtoID_Trd_GetAccList)

	if res.RetType == adapt.RetType_Succeed {
		accs := make([]uint64, 0)
		if _, ok := res.S2C.(*trdgetacclist.S2C); ok {
			for _, acc := range res.S2C.(*trdgetacclist.S2C).GetAccList() {
				accs = append(accs, acc.GetAccID())
			}
		}
		clt.Sync(adapt.ProtoID_Trd_SubAccPush,
			adapt.With("accids", accs),
		)
		if res.RetType != adapt.RetType_Succeed {
			t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
		} else {
			t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
		}
	}

	for {
		clt.Read()
	}

	// clt.Close()
}
