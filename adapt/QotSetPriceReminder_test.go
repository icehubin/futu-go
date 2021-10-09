package adapt_test

import (
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func TestQotSetPriceReminder(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	res := clt.Sync(adapt.ProtoID_Qot_SetPriceReminder,
		adapt.With("code", "HK.00700"),
		// adapt.With("op", int32(1)),
		// adapt.With("type", int32(1)),
		// adapt.With("value", float64(500)),
		// adapt.With("freq", int32(3)),
		// adapt.With("note", "addbyapi"),
		adapt.With("", adapt.Message{
			"op":    proto.Int32(1),
			"type":  proto.Int32(1),
			"value": proto.Float64(600),
			"freq":  proto.Int32(3),
			"note":  proto.String("addbyPb"),
		}),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
