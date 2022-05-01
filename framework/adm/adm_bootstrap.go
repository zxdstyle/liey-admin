package adm

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/framework/logger"
	"github.com/zxdstyle/liey-admin/framework/plugins"
	"github.com/zxdstyle/liey-admin/framework/validator"
)

func bootstrap(ctx context.Context) {
	g.Log().SetHandlers(logger.LoggingColorHandler)

	// 注册自定义规则
	validator.RegisterCustomRules()

	kernel, er := GetKernel()
	if er != nil {
		g.Log().Fatal(ctx, er.Error())
	}

	kernel.Boot()

	if err := plugins.RegisterPlugin(ctx, kernel.Plugins()...); err != nil {
		g.Log().Fatal(ctx, err)
	}
}
