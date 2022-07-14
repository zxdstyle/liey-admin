package events

import (
	"context"
	"sync"
	"time"
)

var defaultBus = NewBus()

type Bus struct {
	sync.RWMutex
	// storage all event name and ListenerQueue map
	subscribers map[string]*SubscriberSorted
	asyncQueue  chan AsyncSubscriberQueue
}

func NewBus() *Bus {
	return &Bus{
		RWMutex:     sync.RWMutex{},
		subscribers: make(map[string]*SubscriberSorted, 64),
		asyncQueue:  make(chan AsyncSubscriberQueue, 1024),
	}
}

func (b *Bus) exists(name string) bool {
	b.RLock()
	defer b.RUnlock()

	_, ok := b.subscribers[name]
	return ok
}

func (b *Bus) get(name string) (any, bool) {
	b.RLock()
	defer b.RUnlock()

	s, ok := b.subscribers[name]
	return s, ok
}

func (b *Bus) set(name string, subscribers ...any) {
	b.Lock()
	defer b.Unlock()

	sub, ok := b.subscribers[name]
	if !ok || sub.IsEmpty() {
		b.subscribers[name] = &SubscriberSorted{subscribers}
		return
	}
	sub.Push(subscribers...)
}

func (b *Bus) GetSubscribers(name string) *SubscriberSorted {
	b.RLock()
	defer b.RUnlock()

	item, ok := b.subscribers[name]
	if !ok || item == nil {
		return nil
	}
	return item
}

func (b *Bus) Shutdown() {
	for len(b.asyncQueue) > 0 {
		time.Sleep(time.Second)
	}
}

func (b *Bus) doDispatch(ctx context.Context, subscriber any, e any) {
	if async, ok := subscriber.(AsyncSubscriber); ok && async.Async() {
		b.asyncQueue <- AsyncSubscriberQueue{
			event:      e,
			subscriber: subscriber,
		}
		return
	}
	b.doHandle(ctx, subscriber, e)
}

func (b *Bus) doHandle(ctx context.Context, subscriber any, e any) {
	sub, ok := subscriber.(Subscriber[any])
	if !ok {
		return
	}

	if err := sub.Handle(ctx, e); err != nil {
		onFailed, ok := subscriber.(SubscriberFailed)
		if !ok {
			return
		}
		onFailed.OnFailed(ctx, e, err)
	}
}

func (b *Bus) addSubscriberItem(name string, subscribers ...any) {
	b.Lock()
	defer b.Unlock()

	sub, ok := b.subscribers[name]
	if !ok || sub.IsEmpty() {
		b.subscribers[name] = &SubscriberSorted{subscribers}
		return
	}
	sub.Push(subscribers...)
}
