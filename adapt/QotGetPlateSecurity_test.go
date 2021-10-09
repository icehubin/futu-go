package adapt_test

import (
	"testing"

	"github.com/icehubin/futu-go/adapt"
	"github.com/icehubin/futu-go/client"
	"google.golang.org/protobuf/proto"
)

func TestQotGetPlateSecurity(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	res := clt.Sync(adapt.ProtoID_Qot_GetPlateSecurity,
		adapt.With("plate", "HK.BK1003"),
		/*
			SortField *int32              `protobuf:"varint,2,opt,name=sortField" json:"sortField,omitempty"` //Qot_Common.SortField,根据哪个字段排序,不填默认Code排序
			Ascend    *bool               `protobuf:"varint,3,opt,name=ascend" json:"ascend,omitempty"`       //升序ture, 降序false, 不填默认升序
		*/
		adapt.With("", adapt.Message{
			"sortField": proto.Int32(2),
			"ascend":    proto.Bool(true),
		}),
	)

	if res.RetType != adapt.RetType_Succeed {
		t.Errorf("Error,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	} else {
		t.Logf("PASS,excepted:%v, got:%v", adapt.RetType_Succeed, res.RetType)
	}

	clt.Close()
}
