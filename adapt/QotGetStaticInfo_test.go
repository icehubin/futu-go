package adapt_test

import (
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
)

func TestQotGetStaticInfo(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	res := clt.Sync(adapt.ProtoID_Qot_GetStaticInfo,
		adapt.With("code_list", []string{"SH.600519", "HK.03333"}),
		adapt.With("", adapt.Message{
			"market":  proto.Int32(1),
			"secType": proto.Int32(3),
		}),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
