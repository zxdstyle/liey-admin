package container

import (
	"fmt"
	"testing"
)

type Cache struct {
	key string
}

func (c Cache) Test() string {
	return c.key
}

func TestSingleton(t *testing.T) {
	Singleton(func() (Cache, error) {
		return Cache{key: "Singleton"}, nil
	})

	test := func(cache Cache) {
		fmt.Println("key:", cache.Test())
	}

	Call(test)

}
