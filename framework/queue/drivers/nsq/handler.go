package nsq

import (
	"context"
	"github.com/nsqio/go-nsq"
	"github.com/zxdstyle/liey-admin/framework/queue/drivers/contract"
)

type (
	defaultHandler struct {
		handler contract.Handler
	}
)

func newDefaultHandler(handler contract.Handler) *defaultHandler {
	return &defaultHandler{handler}
}

func (h defaultHandler) HandleMessage(message *nsq.Message) error {
	return h.handler(context.Background(), message.Body)
}
