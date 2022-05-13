package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/zxdstyle/liey-admin/framework/adm/instances"
	"github.com/zxdstyle/liey-admin/framework/plugins"
)

var ServerCommand = &cobra.Command{
	Use:   "serve",
	Short: "Start http server",
	Run: func(cmd *cobra.Command, args []string) {
		Serve()
	},
}

func Serve() {
	ctx := context.Background()
	installPlugins(ctx)

	instances.RestServer().Run()
}

func installPlugins(ctx context.Context) {
	plugins.Iterator(func(name string, plugin plugins.Plugin) error {
		return plugin.Boot()
	})
}
