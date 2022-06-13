package events

import "context"

type Subscriber interface {
	Handle(ctx context.Context, payload interface{}) error
}

type AsyncSubscriber interface {
	Subscriber
	Async() bool
}
