package events

import (
	"context"
	"fmt"
	"testing"
)

type TestEvent struct {
	Name string
}

func (t TestEvent) Payload() string {
	return t.Name
}

type TestSubscriber struct {
}

func (TestSubscriber) Handle(ctx context.Context, e Event[string]) error {
	fmt.Println(e.Payload())
	return nil
}

func TestName(t *testing.T) {
	Subscribe[string](defaultBus, TestEvent{}, TestSubscriber{})
	Dispatch[string](context.Background(), defaultBus, TestEvent{Name: "TestName"})
}
