package bases

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
	"gorm.io/gorm"
)

type (
	RepositoryModel interface {
		GetKey() uint
		SetKey(id uint)
		GetCreatedAt() *gtime.Time
		GetUpdatedAt() *gtime.Time
	}

	RepositoryModels interface {
		GetModel(i int) RepositoryModel
	}

	HasPreload interface {
		Preload(resource requests.Resource) Filter
	}

	Logic interface {
		Show(ctx context.Context, with []string, mo RepositoryModel) error
	}

	Repository interface {
		Paginate(ctx context.Context, req requests.Request, paginator *responses.Paginator) error
		All(ctx context.Context, req requests.Request, mos RepositoryModels) error
		Show(ctx context.Context, with requests.Resources, mo RepositoryModel) error
		ExistsByKeys(ctx context.Context, keys []uint, exists *bool) error
		Exists(ctx context.Context, mo RepositoryModel, exists *bool) error
		Count(ctx context.Context, mo RepositoryModel, count *int64) error
		CountByKeys(ctx context.Context, keys []uint, count *int64) error
		Create(ctx context.Context, mo RepositoryModel) error
		BatchCreate(ctx context.Context, mos RepositoryModels) error
		Update(ctx context.Context, mo RepositoryModel) error
		BatchUpdate(ctx context.Context, mos RepositoryModels) error
		Destroy(ctx context.Context, mo RepositoryModel) error
		DestroyById(ctx context.Context, ids ...uint) error
	}

	Filter func(tx *gorm.DB) *gorm.DB
)
