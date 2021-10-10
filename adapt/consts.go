package adapt

import (
	"github.com/icehubin/futu-go/pb/qotcommon"
	"github.com/icehubin/futu-go/pb/trdcommon"
)

const (
	//默认的ClientID, 用于区分不同的api
	DEFULAT_CLIENT_ID = "GoByIcehu"
	CLIENT_VERSION    = 300
)

const (
	RetType_Succeed    = 0    //成功
	RetType_Failed     = -1   //失败
	RetType_TimeOut    = -100 //超时
	RetType_DisConnect = -200 //连接断开
	RetType_Unknown    = -400 //未知结果
	RetType_Invalid    = -500 //包内容非法
)

//交易所前缀对应关系表
const (
	MARKET_NONE      = "N/A"       //未知市场
	MARKET_HK        = "HK"        //香港市场
	MARKET_US        = "US"        //美国市场
	MARKET_SH        = "SH"        //沪股市场
	MARKET_SZ        = "SZ"        //深股市场
	MARKET_HK_FUTURE = "HK_FUTURE" //港期货(已废弃，使用QotMarket_HK_Security即可)
	MARKET_SG        = "SG"        //新加坡市场
	MARKET_JP        = "JP"        //日本市场
)

var (
	Market_Value = map[string]int32{
		MARKET_NONE: int32(qotcommon.QotMarket_QotMarket_Unknown),
		MARKET_HK:   int32(qotcommon.QotMarket_QotMarket_HK_Security),
		MARKET_US:   int32(qotcommon.QotMarket_QotMarket_US_Security),
		MARKET_SH:   int32(qotcommon.QotMarket_QotMarket_CNSH_Security),
		MARKET_SZ:   int32(qotcommon.QotMarket_QotMarket_CNSZ_Security),
		MARKET_SG:   int32(qotcommon.QotMarket_QotMarket_SG_Security),
		MARKET_JP:   int32(qotcommon.QotMarket_QotMarket_JP_Security),
	}

	Market_Name = map[int32]string{
		int32(qotcommon.QotMarket_QotMarket_Unknown):       MARKET_NONE,
		int32(qotcommon.QotMarket_QotMarket_HK_Security):   MARKET_HK,
		int32(qotcommon.QotMarket_QotMarket_US_Security):   MARKET_US,
		int32(qotcommon.QotMarket_QotMarket_CNSH_Security): MARKET_SH,
		int32(qotcommon.QotMarket_QotMarket_CNSZ_Security): MARKET_SZ,
		int32(qotcommon.QotMarket_QotMarket_SG_Security):   MARKET_SG,
		int32(qotcommon.QotMarket_QotMarket_JP_Security):   MARKET_JP,
	}

	SecMarket_Value = map[string]int32{
		MARKET_NONE: int32(trdcommon.TrdSecMarket_TrdSecMarket_Unknown),
		MARKET_HK:   int32(trdcommon.TrdSecMarket_TrdSecMarket_HK),
		MARKET_US:   int32(trdcommon.TrdSecMarket_TrdSecMarket_US),
		MARKET_SH:   int32(trdcommon.TrdSecMarket_TrdSecMarket_CN_SH),
		MARKET_SZ:   int32(trdcommon.TrdSecMarket_TrdSecMarket_CN_SZ),
		MARKET_SG:   int32(trdcommon.TrdSecMarket_TrdSecMarket_SG),
		MARKET_JP:   int32(trdcommon.TrdSecMarket_TrdSecMarket_JP),
	}

	SecMarket_Name = map[int32]string{
		int32(trdcommon.TrdSecMarket_TrdSecMarket_Unknown): MARKET_NONE,
		int32(trdcommon.TrdSecMarket_TrdSecMarket_HK):      MARKET_HK,
		int32(trdcommon.TrdSecMarket_TrdSecMarket_US):      MARKET_US,
		int32(trdcommon.TrdSecMarket_TrdSecMarket_CN_SH):   MARKET_SH,
		int32(trdcommon.TrdSecMarket_TrdSecMarket_CN_SZ):   MARKET_SZ,
		int32(trdcommon.TrdSecMarket_TrdSecMarket_SG):      MARKET_SG,
		int32(trdcommon.TrdSecMarket_TrdSecMarket_JP):      MARKET_JP,
	}
)
