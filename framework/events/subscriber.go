package events

import "context"

type (
	Subscriber[V any] interface {
		Handle(ctx context.Context, e Event[V]) error
	}

	SubscriberFailed interface {
		OnFailed(ctx context.Context, e Event[any], err error)
	}

	AsyncSubscriber interface {
		Async() bool
	}

	AsyncSubscriberQueue struct {
		event      any
		subscriber any
	}

	SubscriberFunc[V any] func(ctx context.Context, e Event[V]) error

	SubscriberSorted struct {
		subscribers []any
	}
)

func (fn SubscriberFunc[V]) Handle(ctx context.Context, e Event[V]) error {
	return fn(ctx, e)
}

func (s *SubscriberSorted) Push(subscribers ...any) {
	s.subscribers = append(s.subscribers, subscribers...)
}

func (s *SubscriberSorted) Iterator(fn func(subscriber any) bool) {
	for _, subscriber := range s.subscribers {
		if !fn(subscriber) {
			break
		}
	}
}

func (s SubscriberSorted) IsEmpty() bool {
	return s.subscribers == nil || len(s.subscribers) == 0
}
