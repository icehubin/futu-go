package adapt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/icehubin/futu-go/client"

	"github.com/icehubin/futu-go/adapt"
)

func TestGoto(t *testing.T) {
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
		mp := make([]map[string]interface{}, 0)
		if klList, ok := res.Data["klList"]; ok {
			klines := adapt.ResToArr(klList)
			for _, v := range klines {
				mp = append(mp, adapt.ResToMap(v,
					adapt.Field("changeRate", "change"),
					adapt.Field("closePrice", "close"),
					adapt.Field("highPrice", "high"),
					adapt.Field("lowPrice", "low"),
					adapt.Field("lastClosePrice", "close"),
					adapt.Field("openPrice", "open"),
					adapt.Field("time", ""),
					adapt.Field("turnover", ""),
					adapt.Field("turnoverRate", ""),
				))
			}
		}
		df := dataframe.LoadMaps(mp)
		fmt.Println(df)

		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

}
