package commands

import (
	"github.com/fatih/color"
	"github.com/gogf/gf/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/zxdstyle/liey-admin/framework/adm/instances"
	"os"
	"runtime"
)

var VersionCommand = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		dumpComponent()
	},
}

func dumpComponent() {
	color.Yellow(">>> Components")

	data := [][]string{
		{"GoLang", runtime.Version()},
		{"GoFrame", gf.VERSION},
		{"Liey Admin", instances.Version},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{color.HiGreenString("Component"), color.HiGreenString("Version")})
	table.SetAutoFormatHeaders(false)
	table.SetColMinWidth(0, 20)
	table.SetColMinWidth(1, 12)
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
