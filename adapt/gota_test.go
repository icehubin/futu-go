package adapt_test

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"

	"github.com/icehubin/futu-go/pb/getuserinfo"

	"github.com/icehubin/futu-go/adapt"
)

func TestGoto(t *testing.T) {
	c2s := &getuserinfo.S2C{
		NickName:  proto.String("icehu"),
		AvatarUrl: proto.String("http://wwww.baidu.com/logo.png"),
		ApiLevel:  proto.String("一级"),
		UserID:    proto.Int64(1342342323434),
	}
	m := adapt.PbParser(
		adapt.Field("userID", "user"),
		adapt.Field("nickName", ""),
	).Map(c2s)
	fmt.Println(c2s)
	fmt.Println(m)
}
