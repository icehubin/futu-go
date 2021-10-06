package client_test

import (
	"testing"

	"github.com/icehubin/futu-go/client"
)

func TestWorker(t *testing.T) {
	var p = func() *client.Client {
		clt, err := client.New("127.0.0.1:11111")
		if err != nil {
			panic("Client init error")
		}
		return clt
	}

	work := client.NewWorker()
	work.PrepareClient(p)
	work.SetDefaultHandle(func(r *client.ResPack) {
		//do noth.
	})
	work.Work()
}
