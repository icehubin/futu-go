package client_test

import (
	"testing"
	"time"

	"github.com/icehubin/futu-go/client"
)

func TestClient(t *testing.T) {
	clt, err := client.New("127.0.0.1:11111")
	if err != nil {
		return
	}

	time.Sleep(time.Second * 1)

	clt.KeepAlive()

	time.Sleep(time.Second * 1)

	clt.Close()
}
