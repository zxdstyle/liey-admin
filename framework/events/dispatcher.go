package events

import (
	"fmt"
	"sync"
)

type Bus struct {
	sync.RWMutex
	// storage user custom Event instance.
	events map[string]Event[any]
	// storage all event name and ListenerQueue map
	listeners map[string]ListenerSorted[any]
}

func (b *Bus) Subscribe(event Event[any], listeners ...Listener[any]) {
	b.Lock()
	defer b.Unlock()

	name := b.resolveEventName(event)
	for _, listener := range listeners {
		b.events[name] = event
		b.addListenerItem(name, &ListenerItem[any]{Listener: listener})
	}

}

func (b *Bus) addListenerItem(name string, item *ListenerItem[any]) {

}

func (b *Bus) resolveEventName(e Event[any]) string {
	event, ok := e.(NamedEvent)
	if ok {
		name := event.Name()
		if len(name) > 0 {
			return name
		}
	}

	// struct
	return fmt.Sprintf("%T", e)
}
