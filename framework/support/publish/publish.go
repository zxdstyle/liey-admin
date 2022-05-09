package publish

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/zxdstyle/liey-admin/framework/exception"
)

var publishes = gmap.NewStrAnyMap(true)

func RegisterPublishes(group string, assets ...Publisher) error {
	for i, _ := range assets {
		if publishes.Contains(group) {
			return exception.ErrRepeatedMigration
		}
		publishes.SetIfNotExist(group, assets[i])
	}
	return nil
}

func GetPublishes(groups ...string) map[string]Publisher {
	mis := make(map[string]Publisher)
	if len(groups) == 0 {
		publishes.Iterator(func(k string, v interface{}) bool {
			mis[k] = v.(Publisher)
			return true
		})
		return mis
	}

	for _, group := range groups {
		if val, ok := publishes.Get(group).(Publisher); ok {
			mis[group] = val
		}
	}
	return mis
}

func GetPublish(group string) Publisher {
	val, ok := publishes.Search(group)
	if !ok {
		return nil
	}
	return val.(Publisher)
}

func AllPublishKeys() []string {
	return publishes.Keys()
}
