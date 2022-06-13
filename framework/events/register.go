package events

import (
	"fmt"
	"reflect"
	"sync"
)

var (
	events = make(map[string][]Subscriber, 32)
	locker = &sync.RWMutex{}
)

func eventName(event Event) string {
	t := reflect.TypeOf(event)
	return t.Name()
}

func Listen(event Event, subscribers ...Subscriber) error {
	locker.Lock()
	defer locker.Unlock()

	name := eventName(event)
	if _, ok := events[name]; ok {
		return fmt.Errorf("event '%s' has already been registered, please do not register again", name)
	}
	events[name] = subscribers
	return nil
}

func getSubscribers(event Event) ([]Subscriber, error) {
	locker.RLock()
	defer locker.RUnlock()

	name := eventName(event)
	if e, ok := events[name]; ok {
		return e, nil
	}
	return nil, fmt.Errorf("not found event subscribers")
}
