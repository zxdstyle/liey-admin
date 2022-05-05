package config

import (
	"github.com/gogf/gf/v2/container/gmap"
	"gorm.io/plugin/dbresolver"
)

var (
	policies = gmap.NewStrAnyMapFrom(map[string]interface{}{
		"random": dbresolver.RandomPolicy{},
	}, true)
)

func RegisterPolicy(name string, policy dbresolver.Policy) {
	policies.Set(name, policy)
}

func GetPolicy(name string) dbresolver.Policy {
	key := "random"
	if len(name) > 0 {
		key = name
	}
	val, ok := policies.Search(key)
	if !ok {
		return nil
	}
	return val.(dbresolver.Policy)
}
