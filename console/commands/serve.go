package commands

import (
	"github.com/spf13/cobra"
	"github.com/zxdstyle/liey-admin/framework/adm/instances"
)

var ServerCommand = &cobra.Command{
	Use:   "serve",
	Short: "Start http server",
	Run: func(cmd *cobra.Command, args []string) {
		Serve()
	},
}

func Serve() {
	instances.RestServer().Run()
}
