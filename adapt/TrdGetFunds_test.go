package adapt_test

import (
	"testing"

	"github.com/icehubin/futu-go/adapt"
	"google.golang.org/protobuf/proto"

	"github.com/icehubin/futu-go/client"
)

func TestTrdGetFunds(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	res := clt.Sync(adapt.ProtoID_Trd_GetFunds,
		adapt.With("Header", adapt.Message{
			"trdEnv":    proto.Int32(1),
			"accID":     proto.Uint64(281756460285261810),
			"trdMarket": proto.Int32(2),
		}),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
