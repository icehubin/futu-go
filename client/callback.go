package client

import (
	"github.com/icehubin/futu-go/logger"
)

func DefaultCallback(p *ResPack) {

	h := *p.Header
	r := *p.Response

	logger.WithFields(logger.Fields{
		"head":     h,
		"response": r,
	}).Info("Host error")
}
