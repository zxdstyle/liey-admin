package commands

import (
	"context"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/spf13/cobra"
	"github.com/zxdstyle/liey-admin/framework/adm/instances"
	"github.com/zxdstyle/liey-admin/framework/database"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

var (
	MigrateCommand = &cobra.Command{
		Use:   "migrate",
		Short: "Run the database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			if len(args) == 0 {
				multiMigrateData(ctx)
			}

			instances.DB().AutoMigrate()
		},
	}
)

func multiMigrateData(ctx context.Context) {
	options := database.AllMigrationKeys()

	sv := &survey.Select{
		Message: "Please select the data to migrate",
		Options: options,
	}

	var migrateKey string
	if err := survey.AskOne(sv, &migrateKey); err != nil {
		g.Log().Fatal(ctx, err)
	}

	mi := database.GetMigration(migrateKey)
	if mi == nil {
		g.Log().Fatal(ctx, "Incorrect migration data selected")
	}
	if err := migrate(ctx, mi.Models()...); err != nil {
		g.Log().Fatal(ctx, err)
	}

	g.Log().Noticef(ctx, "success to migrate: %s", migrateKey)
}

func migrate(ctx context.Context, mos ...bases.RepositoryModel) error {
	var models []interface{}
	for i, _ := range mos {
		models = append(models, mos[i])
	}
	return instances.DB().Debug().AutoMigrate(models...)
}
