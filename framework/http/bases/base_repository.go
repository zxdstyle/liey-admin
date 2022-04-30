package bases

import (
	"context"
	"gorm.io/gorm"
)

type GormRepository struct {
	Orm *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		Orm: db,
	}
}

func (repo GormRepository) Show(ctx context.Context, with []string, mo RepositoryModel) error {
	tx := repo.Orm.WithContext(ctx)
	return tx.First(mo).Error
}
