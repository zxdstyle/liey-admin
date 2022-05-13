package support

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/zxdstyle/liey-admin/framework/support/jwt"
)

var instances = gmap.NewStrAnyMap(true)

func JWT() *jwt.JWT {
	return instances.GetOrSetFuncLock("support.jwt", func() interface{} {
		return jwt.NewJWT()
	}).(*jwt.JWT)
}
