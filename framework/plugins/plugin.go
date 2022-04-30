package plugins

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/framework/database"
	"github.com/zxdstyle/liey-admin/framework/exception"
)

type (
	Plugin interface {
		Name() string
		Boot() error
		Migrations() []database.Migration
	}
)

var pluginsMap = gmap.NewStrAnyMap(true)

func RegisterPlugin(ctx context.Context, plugins ...Plugin) error {
	for i, plugin := range plugins {
		name := plugin.Name()
		if pluginsMap.Contains(name) {
			return exception.ErrPluginExists
		}

		if err := plugin.Boot(); err != nil {
			return err
		}

		if err := database.RegisterMigration(name, plugin.Migrations()...); err != nil {
			return err
		}

		pluginsMap.SetIfNotExist(name, plugins[i])

		g.Log().Noticef(ctx, "success to register plugin: %s", plugin.Name())
	}
	return nil
}
