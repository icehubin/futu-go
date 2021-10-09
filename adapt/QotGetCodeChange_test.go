package adapt_test

import (
	"testing"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
	"google.golang.org/protobuf/proto"
)

func TestQotGetCodeChange(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	res := clt.Sync(adapt.ProtoID_Qot_GetCodeChange,
		adapt.With("code_list", []string{"HK.00700", "HK.09888"}),
		/*
			Type      *int32  `protobuf:"varint,1,req,name=type" json:"type,omitempty"`          //TimeFilterType, 过滤类型
			BeginTime *string `protobuf:"bytes,2,opt,name=beginTime" json:"beginTime,omitempty"` //开始时间点
			EndTime   *string `protobuf:"bytes,3,opt,name=endTime" json:"endTime,omitempty"`     //结束时间点
		*/
		adapt.With("TimeFilter", adapt.Message{
			"type":      proto.Int32(1),
			"beginTime": proto.String("2020-08-01"),
			"endTime":   proto.String("2021-10-10"),
		}),
		adapt.With("TimeFilter", adapt.Message{
			"type":      proto.Int32(2),
			"beginTime": proto.String("2021-09-01"),
			"endTime":   proto.String("2021-10-10"),
		}),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
