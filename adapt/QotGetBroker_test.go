package adapt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func TestQotGetBroker(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	fmt.Println(clt.Sync(adapt.ProtoID_Qot_Sub,
		adapt.With("code_list", []string{"HK.00700"}),
		adapt.With("subtype_list", []string{"BROKER"}),
		adapt.With("IsFirstPush", true),
	))
	time.Sleep(time.Microsecond * 500)
	res := clt.Sync(adapt.ProtoID_Qot_GetBroker,
		adapt.With("code", "HK.00700"),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
