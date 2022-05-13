package adm

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/framework/adm/instances"
	"github.com/zxdstyle/liey-admin/framework/logger"
	"github.com/zxdstyle/liey-admin/framework/plugins"
	"github.com/zxdstyle/liey-admin/framework/queue"
	"github.com/zxdstyle/liey-admin/framework/queue/job"
)

func bootstrap(ctx context.Context) {
	g.Log().SetHandlers(logger.LoggingColorHandler)

	kernel, er := instances.GetKernel()
	if er != nil {
		g.Log().Fatal(ctx, er.Error())
	}

	kernel.Boot()

	if err := plugins.RegisterPlugin(ctx, kernel.Plugins()...); err != nil {
		g.Log().Fatal(ctx, err)
	}

	if err := job.RegisterJob(ctx, kernel.Queues()...); err != nil {
		g.Log().Fatal(ctx, err)
	}

	queue.InitQueueWithConfig()
}
