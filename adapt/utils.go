package adapt

import (
	"strings"

	"github.com/icehubin/futu-go/pb/qotcommon"
)

//quote
func Stock2Security(stock string) *qotcommon.Security {
	arr := strings.Split(stock, ".")
	if len(arr) != 2 {
		return nil
	}
	market_code, ok := Market_Value[strings.ToUpper(arr[0])]
	if ok {
		return &qotcommon.Security{
			Market: &market_code,
			Code:   &arr[1],
		}
	}
	return nil
}

func StocksToSecurity(stocks []string) []*qotcommon.Security {
	sa := make([]*qotcommon.Security, 0)
	for _, v := range stocks {
		s := Stock2Security(v)
		if s != nil {
			sa = append(sa, s)
		}
	}
	return sa
}

func MapSecurityToStock(sm map[string]interface{}) string {
	market_name := MARKET_NONE
	if market, ok := sm["market"]; ok {
		if market, ok := market.(int32); ok {
			market_name, ok = Market_Name[market]
		}
	}
	code, _ := sm["code"].(string)
	return market_name + "." + code
}

func SecurityToStock(s *qotcommon.Security) string {
	market_name, ok := Market_Name[*s.Market]
	if ok {
		return market_name + "." + *s.Code
	} else {
		return MARKET_NONE + string(*s.Market) + "." + *s.Code
	}
}

//trd
type TrdStock struct {
	Code      string
	SecMarket int32
}

func StockToTrd(stock string) *TrdStock {
	arr := strings.Split(stock, ".")
	if len(arr) != 2 {
		return nil
	}
	market_code, ok := SecMarket_Value[strings.ToUpper(arr[0])]
	if ok {
		return &TrdStock{
			Code:      arr[1],
			SecMarket: market_code,
		}
	}
	return nil
}

// func TrdStockToString()
