package events

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var eventsBus = make(chan subscriberWithPayload, 1024)

// todo waitDone 优雅退出
func init() {
	go doInit()
}

type subscriberWithPayload struct {
	payload    interface{}
	subscriber Subscriber
}

func (s subscriberWithPayload) Consume(ctx context.Context) {
	if err := s.subscriber.Handle(ctx, s.payload); err != nil {
		g.Log().Error(ctx, err)
	}
}

func wrapSubscriberWithPayload(subscriber Subscriber, payload interface{}) subscriberWithPayload {
	return subscriberWithPayload{payload: payload, subscriber: subscriber}
}

func Dispatch(ctx context.Context, event Event) error {
	subscribers, err := getSubscribers(event)
	if err != nil {
		return err
	}

	for _, subscriber := range subscribers {
		if asyncSubscriber, ok := subscriber.(AsyncSubscriber); ok && asyncSubscriber.Async() {
			eventsBus <- wrapSubscriberWithPayload(asyncSubscriber, event.Payload())
			continue
		}

		if e := subscriber.Handle(ctx, event.Payload()); e != nil {
			return e
		}
	}
	return nil
}

func doInit() {
	ctx := context.Background()
	for {
		select {
		case <-ctx.Done():
			break
		case subscriber := <-eventsBus:
			go subscriber.Consume(ctx)
		}
	}
}
