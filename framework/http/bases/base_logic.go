package bases

import "context"

type BaseLogic struct {
	repo Repository
}

func NewBaseLogic(repo Repository) *BaseLogic {
	return &BaseLogic{repo}
}

func (b BaseLogic) Show(ctx context.Context, with []string, mo RepositoryModel) error {
	return b.repo.Show(ctx, with, mo)
}
