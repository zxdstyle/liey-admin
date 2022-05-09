package adm

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/console"
	"github.com/zxdstyle/liey-admin/framework/adm/instances"
	"github.com/zxdstyle/liey-admin/framework/http/server"
	"gorm.io/gorm"
)

func Version() string {
	return instances.Version
}

func Debug() bool {
	ctx := context.Background()
	return g.Cfg("app").MustGet(ctx, "debug", false).Bool()
}

func DB(name ...string) *gorm.DB {
	return instances.DB(name...)
}

func Start() {
	ctx := context.Background()

	bootstrap(ctx)

	console.Execute()
}

func Server() *server.RestServer {
	return instances.RestServer()
}
