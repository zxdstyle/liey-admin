package adm

import (
	"github.com/gogf/gf/v2/container/gmap"
)

var providers = gmap.NewAnyAnyMap(true)

type SingletonKey string

func Singleton(key SingletonKey, initializer func() interface{}) interface{} {
	return providers.GetOrSetFuncLock(key, func() interface{} {
		return initializer()
	})
}
