package bases

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/zxdstyle/liey-admin/framework/exception"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

type (
	Logic interface {
	}

	BaseLogic struct {
		repo Repository
	}
)

func NewBaseLogic(repo Repository) *BaseLogic {
	return &BaseLogic{repo}
}

func (b BaseLogic) Paginate(ctx context.Context, req requests.Request, paginator *responses.Paginator) error {
	return b.repo.Paginate(ctx, req, paginator)
}

func (b BaseLogic) All(ctx context.Context, req requests.Request, mos RepositoryModels) error {
	return b.repo.All(ctx, req, mos)
}

func (b BaseLogic) Show(ctx context.Context, req requests.Request, mo RepositoryModel) error {
	if mo.GetKey() == 0 {
		return gerror.NewCode(exception.CodeNotFound, "not found record")
	}

	return b.repo.Show(ctx, req.GetClauses(), mo)
}

func (b BaseLogic) Create(ctx context.Context, mo RepositoryModel) error {
	return b.repo.Create(ctx, mo)
}

func (b BaseLogic) Update(ctx context.Context, mo RepositoryModel) error {
	return b.repo.Update(ctx, mo)
}

func (b BaseLogic) Destroy(ctx context.Context, mo RepositoryModel) error {
	return b.repo.Destroy(ctx, mo)
}

func (b BaseLogic) DestroyById(ctx context.Context, ids ...uint) error {
	return b.repo.DestroyById(ctx, ids...)
}
