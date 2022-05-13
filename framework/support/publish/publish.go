package publish

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/zxdstyle/liey-admin/framework/support/publish/publisher"
)

var publishes = gmap.NewStrAnyMap(true)

func RegisterPublishes(group string, assets ...publisher.Publisher) error {
	for i, _ := range assets {
		if publishes.Contains(group) {
			return ErrRepeatedPublish
		}
		publishes.SetIfNotExist(group, assets[i])
	}
	return nil
}

func GetPublishes(groups ...string) map[string]publisher.Publisher {
	mis := make(map[string]publisher.Publisher)
	if len(groups) == 0 {
		publishes.Iterator(func(k string, v interface{}) bool {
			mis[k] = v.(publisher.Publisher)
			return true
		})
		return mis
	}

	for _, group := range groups {
		if val, ok := publishes.Get(group).(publisher.Publisher); ok {
			mis[group] = val
		}
	}
	return mis
}

func GetPublish(group string) publisher.Publisher {
	val, ok := publishes.Search(group)
	if !ok {
		return nil
	}
	return val.(publisher.Publisher)
}

func AllPublishKeys() []string {
	return publishes.Keys()
}
