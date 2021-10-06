package adapt_test

import (
	"testing"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func TestExampleAdapt(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	// fmt.Println(clt.Sync(adapt.ProtoID_Qot_Sub,
	// 	adapt.With("code_list", []string{"SH.600519"}),
	// 	adapt.With("subtype_list", []string{"QUOTE", "TICKER"}),
	// 	adapt.With("IsFirstPush", true),
	// ))

	res := clt.Sync(adapt.ProtoID_Example_Adapt)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
