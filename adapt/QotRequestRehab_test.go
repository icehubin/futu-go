package adapt_test

import (
	"testing"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func TestQotRequestRehab(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	res := clt.Sync(adapt.ProtoID_Qot_RequestRehab,
		adapt.With("code", "US.TSLA"),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
