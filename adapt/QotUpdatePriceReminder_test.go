package adapt_test

import (
	"testing"
	"time"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func TestQotUpdatePriceReminder(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	res := clt.Sync(adapt.ProtoID_Qot_Sub,
		adapt.With("code_list", []string{"SH.600519"}),
		adapt.With("subtype_list", []string{"QUOTE", "TICKER"}),
		adapt.With("IsFirstPush", true),
		adapt.With("Push", true),
	)
	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	// for {
	// 	clt.Read()
	// }
	time.Sleep(time.Second * 10)
	clt.Close()

}
