package adm

import (
	"context"
	"github.com/zxdstyle/liey-admin/framework/events"
)

func Subscribe[T any](e events.Event[T], subscribers ...events.Subscriber[T]) {
	//events.DefaultBus.Subscribe(e, subscribers...)
}

func Dispatch(ctx context.Context, e events.Event[any]) {
	//events.DefaultBus.Dispatch(ctx, e)
}
