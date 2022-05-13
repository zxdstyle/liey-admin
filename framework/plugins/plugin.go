package plugins

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/framework/database/migrations"
	"github.com/zxdstyle/liey-admin/framework/exception"
	"github.com/zxdstyle/liey-admin/framework/support/publish"
	"github.com/zxdstyle/liey-admin/framework/support/publish/publisher"
)

var pluginsMap = gmap.NewStrAnyMap(true)

type (
	Plugin interface {
		Name() string
		Boot() error
		Migrations() []migrations.Migration
		Resources() []publisher.Publisher
	}

	defaultPlugin struct {
		name       string
		boot       func() error
		migrations []migrations.Migration
		resources  []publisher.Publisher
	}
)

func (p defaultPlugin) Name() string {
	return p.name
}

func (p defaultPlugin) Boot() error {
	return p.boot()
}

func (p defaultPlugin) Migrations() []migrations.Migration {
	return p.migrations
}

func (p defaultPlugin) Resources() []publisher.Publisher {
	return p.resources
}

func RegisterPlugin(ctx context.Context, plugins ...Plugin) error {
	for i, plugin := range plugins {
		name := plugin.Name()
		if pluginsMap.Contains(name) {
			return exception.ErrPluginExists
		}

		if err := migrations.RegisterMigration(name, plugin.Migrations()...); err != nil {
			return err
		}

		if err := publish.RegisterPublishes(name, plugin.Resources()...); err != nil {
			return err
		}

		pluginsMap.SetIfNotExist(name, plugins[i])

		g.Log().Noticef(ctx, "success to register plugin: %s", plugin.Name())
	}
	return nil
}

func Iterator(handler func(name string, plugin Plugin) error) {
	pluginsMap.Iterator(func(k string, v interface{}) bool {
		p := v.(Plugin)
		if err := handler(k, p); err != nil {
			return false
		}
		return true
	})
}

func WithRename(plugin Plugin, name string) Plugin {
	return defaultPlugin{
		name:       name,
		boot:       plugin.Boot,
		migrations: plugin.Migrations(),
		resources:  plugin.Resources(),
	}
}
