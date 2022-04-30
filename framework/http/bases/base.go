package bases

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	RepositoryModel interface {
		GetKey() uint
		GetCreatedAt() *gtime.Time
		GetUpdatedAt() *gtime.Time
	}

	RepositoryModels interface {
		GetModel(i int) RepositoryModel
	}

	Logic interface {
		Show(ctx context.Context, with []string, mo RepositoryModel) error
	}

	Repository interface {
		Show(ctx context.Context, with []string, mo RepositoryModel) error
	}
)
