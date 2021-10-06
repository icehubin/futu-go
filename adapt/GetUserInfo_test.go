package adapt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/icehubin/futu-go/adapt"

	"github.com/icehubin/futu-go/client"
)

func TestGetUserInfo(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	time.Sleep(time.Microsecond * 500)
	res := clt.Sync(adapt.ProtoID_GetUserInfo)

	fmt.Println(adapt.PbParser().Map(res.S2C))

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
