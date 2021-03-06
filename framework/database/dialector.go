package database

import (
	"gorm.io/gorm"
	"sync"
)

var dialectors = &Dialector{
	data: make(map[string]gorm.Dialector),
	mu:   &sync.RWMutex{},
}

type Dialector struct {
	data map[string]gorm.Dialector
	mu   *sync.RWMutex
}

func (d Dialector) Get(name string) gorm.Dialector {
	d.mu.RLock()
	defer d.mu.RUnlock()

	val, ok := d.data[name]
	if !ok {
		return nil
	}
	return val
}

func (d Dialector) Gets(name ...string) (dials []gorm.Dialector) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	for _, key := range name {
		val, ok := d.data[key]
		if ok {
			dials = append(dials, val)
		}
	}
	return
}

func (d Dialector) Set(name string, dial gorm.Dialector) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[name] = dial
}
