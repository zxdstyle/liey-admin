package contract

import (
	"context"
	"github.com/zxdstyle/liey-admin/framework/queue/job"
)

type (
	Handler func(ctx context.Context, payload []byte) error

	Connection interface {
		Produce(name string, payload []byte) error
		Consume(j job.Job) error
		Close()
	}
)
