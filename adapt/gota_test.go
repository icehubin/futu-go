package adapt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/pb/keepalive"
	"google.golang.org/protobuf/proto"
)

func TestGoto(t *testing.T) {
	c2s := &keepalive.C2S{
		Time: proto.Int64(time.Now().Unix()),
	}
	m := adapt.PbParser().Map(c2s)
	fmt.Println(c2s, m)
}
