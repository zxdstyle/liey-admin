package adm

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/console"
	"github.com/zxdstyle/liey-admin/framework/adm/instances"
	"github.com/zxdstyle/liey-admin/framework/http/server"
	"github.com/zxdstyle/liey-admin/framework/logger"
	"github.com/zxdstyle/liey-admin/framework/plugins"
	"github.com/zxdstyle/liey-admin/framework/queue"
	"github.com/zxdstyle/liey-admin/framework/queue/job"
	"gorm.io/gorm"
)

func Version() string {
	return instances.Version
}

func Debug() bool {
	ctx := context.Background()
	res, _ := g.Cfg("app").Get(ctx, "debug", false)
	return res.Bool()
}

func DB(name ...string) *gorm.DB {
	return instances.DB(name...)
}

func Start(kernel instances.Kernel) {
	ctx := context.Background()

	g.Log().SetHandlers(logger.LoggingColorHandler)

	kernel.Boot()

	if hk, ok := kernel.(instances.HttpKernel); ok {
		if err := plugins.RegisterPlugin(ctx, hk.Plugins()...); err != nil {
			g.Log().Fatal(ctx, err)
		}

		if err := job.RegisterJob(ctx, hk.Queues()...); err != nil {
			g.Log().Fatal(ctx, err)
		}

		queue.InitQueueWithConfig()
	}

	console.Execute()
}

func Server() *server.RestServer {
	return instances.RestServer()
}
