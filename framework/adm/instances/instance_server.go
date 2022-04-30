package instances

import (
	"github.com/zxdstyle/liey-admin/framework/http/server"
)

const (
	frameCoreRestServer = "liey.core.rest.server"
)

func RestServer() *server.RestServer {
	return instances.GetOrSetFunc(frameCoreRestServer, func() interface{} {
		return server.NewRestServer("default")
	}).(*server.RestServer)
}
