package events

import "context"

type (
	Listener[V any] interface {
		Handle(ctx context.Context, e Event[V]) error
	}

	ListenerFunc[V any] func(ctx context.Context, e Event[V]) error

	ListenerItem[V any] struct {
		Priority int
		Listener Listener[V]
	}
	ListenerSorted[V any] struct {
		listeners []*Listener[V]
	}
)

func (fn ListenerFunc[V]) Handle(ctx context.Context, e Event[V]) error {
	return fn(ctx, e)
}
