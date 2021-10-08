package adapt_test

import (
	"testing"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func TestQotModifyUserSecurity(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	res := clt.Sync(adapt.ProtoID_Qot_ModifyUserSecurity,
		adapt.With("name", "指数"),
		adapt.With("op", "add"),
		adapt.With("code_list", []string{"SH.600519", "HK.00700"}),
	)

	res = clt.Sync(adapt.ProtoID_Qot_ModifyUserSecurity,
		adapt.With("name", "指数"),
		adapt.With("op", "remove"),
		adapt.With("code_list", []string{"HK.00700"}),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
