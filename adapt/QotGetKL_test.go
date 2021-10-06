package adapt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
	"github.com/icehubin/futu-go/pb/qotgetkl"
)

func TestQotGetKL(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	fmt.Println(clt.Sync(adapt.ProtoID_Qot_Sub,
		adapt.With("code_list", []string{"SH.600519", "SZ.300957"}),
		adapt.With("subtype_list", []string{"QUOTE", "TICKER", "K_DAY"}),
		adapt.With("IsFirstPush", true),
	))
	time.Sleep(time.Microsecond * 500)
	res := clt.Sync(adapt.ProtoID_Qot_GetKL,
		adapt.With("code", "SZ.300957"),
		adapt.With("ktype", "K_DAY"),
		adapt.With("reqNum", int32(10)),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		Klist := res.S2C.(*qotgetkl.S2C).GetKlList()
		mp := make([]map[string]interface{}, 0)
		for _, v := range Klist {
			mp = append(mp, adapt.PbParser().Map(v))
		}
		df := dataframe.LoadMaps(mp)
		fmt.Println(df)
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
