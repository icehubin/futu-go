package adapt_test

import (
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func TestQotGetWarrant(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	res := clt.Sync(adapt.ProtoID_Qot_GetWarrant,
		adapt.With("code", "HK.00700"),
		adapt.With("", adapt.Message{
			"num":   proto.Int32(10),
			"begin": proto.Int32(0),
		}),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
