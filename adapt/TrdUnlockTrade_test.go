package adapt_test

import (
	"crypto/md5"
	"fmt"
	"testing"

	"github.com/icehubin/futu-go/adapt"

	"github.com/icehubin/futu-go/client"
)

func TestTrdUnlockTrade(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	fmt.Println("请输入解锁密码：")
	var passwd string
	fmt.Scanf("%s", &passwd)
	pwdmd5 := fmt.Sprintf("%x", md5.Sum([]byte(passwd)))

	res := clt.Sync(adapt.ProtoID_Trd_UnlockTrade,
		adapt.With("unlock", true),
		adapt.With("SecurityFirm", int32(1)),
		adapt.With("pwd", pwdmd5),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
