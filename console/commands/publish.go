package commands

import (
	"context"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/zxdstyle/liey-admin/framework/support/publish"
)

var (
	PublishCommand = &cobra.Command{
		Use:   "publish",
		Short: "Publish any publishable assets from plugin packages",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			publishData(ctx, args...)
		},
	}
)

func publishData(ctx context.Context, keys ...string) {
	if len(keys) == 0 {
		options := publish.AllPublishKeys()
		if len(options) == 0 {
			logger.Fatal(ctx, "no resources to publish")
		}
		sv := &survey.Select{
			Message: "Please select the assets to publish",
			Options: options,
		}

		var key string
		if err := survey.AskOne(sv, &key); err != nil {
			logger.Fatal(ctx, err)
		}
		keys = append(keys, key)
	}

	data := publish.GetPublishes(keys...)
	if len(data) == 0 {
		logger.Fatal(ctx, "no resources to publish")
	}
	for name, datum := range data {
		if err := datum.Publish(name); err != nil {
			logger.Fatal(ctx, err)
		}
	}
	logger.Notice(ctx, "assets published successfully.")
}
