package console

import (
	"github.com/spf13/cobra"
	"github.com/zxdstyle/liey-admin/console/commands"
	"log"
)

var rootCommand = &cobra.Command{
	Use:   "adm",
	Short: "Liey-Admin is an web framework",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Serve()
	},
}

func init() {
	RegisterCmd(commands.ServerCmd, commands.VersionCmd, commands.MigrateCmd, commands.PublishCmd, commands.QueueCmd)
}

func RegisterCmd(commands ...*cobra.Command) {
	rootCommand.AddCommand(commands...)
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
