package adapt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func TestQotRequestHistoryKL(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	fmt.Println(clt.Sync(adapt.ProtoID_Qot_Sub,
		adapt.With("code_list", []string{"SH.600519"}),
		adapt.With("subtype_list", []string{"QUOTE", "K_DAY"}),
		adapt.With("IsFirstPush", true),
	))
	time.Sleep(time.Microsecond * 500)
	res := clt.Sync(adapt.ProtoID_Qot_RequestHistoryKL,
		adapt.With("code", "SH.600519"),
		adapt.With("begin", "2021-09-01"),
		adapt.With("end", "2021-10-01"),
		adapt.With("k_type", "K_DAY"),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()

}
