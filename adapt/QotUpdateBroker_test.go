package adapt_test

import (
	"testing"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func TestQotUpdateBroker(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}
	defer clt.Close()

	res := clt.Sync(adapt.ProtoID_Qot_Sub,
		adapt.With("code_list", []string{"HK.00700"}),
		adapt.With("subtype_list", []string{"QUOTE", "BROKER"}),
		adapt.With("IsFirstPush", true),
		adapt.With("Push", true),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	for {
		clt.Read()
	}
	// time.Sleep(time.Second * 10)
}
