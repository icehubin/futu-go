package adapt

import (
	"fmt"
	"testing"

	"github.com/icehubin/futu-go/pb/trdgethistoryorderlist"

	"github.com/icehubin/futu-go/pb/qotcommon"
	"github.com/icehubin/futu-go/pb/qotgetmarketstate"

	"github.com/icehubin/futu-go/pb/qotsetpricereminder"
	"google.golang.org/protobuf/proto"
)

func TestFillProto(t *testing.T) {
	pm := &qotsetpricereminder.C2S{
		Value:    proto.Float64(1.1),
		Security: &qotcommon.Security{},
	}
	fmt.Println("pm", pm)
	protoFill(pm, Message{
		"op":   proto.Int32(1234),
		"key":  proto.Int64(134242343),
		"type": proto.Int64(111),
	})
	protoFill(pm.Security, Message{
		"market": proto.Int32(1),
		"code":   proto.String("haha"),
	})
	fmt.Println("pm", pm)
}

func TestArrInt(t *testing.T) {
	pm := &trdgethistoryorderlist.C2S{}
	fmt.Println("pm", pm)
	protoFill(pm, Message{
		"filterStatusList": []int32{1, 2},
	})
	fmt.Println("pm", pm)
}

func TestAppendProto(t *testing.T) {
	seclist := make([]*qotcommon.Security, 0)
	pm := &qotgetmarketstate.C2S{
		SecurityList: seclist,
	}
	fmt.Println("pm", pm)
	protoAppend(&pm.SecurityList, Message{
		"market": proto.Int32(1),
		"code":   proto.String("haha"),
	})
	fmt.Println("pm", pm)
	protoAppend(&pm.SecurityList, Message{
		"market": proto.Int32(2),
		"code":   proto.String("hahaha"),
	})
	fmt.Println("pm", pm)
}
