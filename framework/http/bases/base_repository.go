package bases

import (
	"context"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormRepository struct {
	Orm *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		Orm: db,
	}
}

func (repo GormRepository) Paginate(ctx context.Context, req requests.Request, paginator *responses.Paginator) error {
	tx, err := req.ToQuery(repo.Orm.WithContext(ctx))
	if err != nil {
		return err
	}

	if er := tx.WithContext(ctx).Count(&paginator.Total).Error; er != nil {
		return er
	}

	tx = repo.loadResources(tx, req.GetWithResources(), paginator.Data.(RepositoryModels).GetModel(0))

	page := req.GetPage()
	pageSize := req.GetPageSize()
	offset := (page - 1) * pageSize
	return tx.Offset(offset).Limit(pageSize).Find(paginator.Data).Error
}

func (repo GormRepository) All(ctx context.Context, req requests.Request, mos RepositoryModels) error {
	tx, err := req.ToQuery(repo.Orm.WithContext(ctx))
	if err != nil {
		return err
	}

	tx = repo.loadResources(tx, req.GetWithResources(), mos.GetModel(0))

	return tx.Limit(req.GetLimit()).Find(mos).Error
}

func (repo GormRepository) Show(ctx context.Context, with requests.Resources, mo RepositoryModel) error {
	tx := repo.Orm.WithContext(ctx)
	tx = repo.loadResources(tx, with, mo)
	return tx.First(mo).Error
}

func (repo GormRepository) ExistsByKeys(ctx context.Context, keys []uint, exists *bool) error {
	var count int64
	if err := repo.CountByKeys(ctx, keys, &count); err != nil {
		return err
	}
	*exists = count == int64(len(keys))
	return nil
}

func (repo GormRepository) Exists(ctx context.Context, mo RepositoryModel, exists *bool) error {
	var count int64
	if err := repo.Count(ctx, mo, &count); err != nil {
		return err
	}
	*exists = count > 0
	return nil
}

func (repo GormRepository) Count(ctx context.Context, mo RepositoryModel, count *int64) error {
	return repo.Orm.WithContext(ctx).Where(mo).Count(count).Error
}

func (repo GormRepository) CountByKeys(ctx context.Context, keys []uint, count *int64) error {
	return repo.Orm.WithContext(ctx).Where("id IN ?", keys).Count(count).Error
}

func (repo GormRepository) Create(ctx context.Context, mo RepositoryModel) error {
	return repo.Orm.WithContext(ctx).Omit(clause.Associations).Create(mo).Error
}

func (repo GormRepository) BatchCreate(ctx context.Context, mos RepositoryModels) error {
	return repo.Orm.WithContext(ctx).Omit(clause.Associations).Create(mos).Error
}

func (repo GormRepository) Update(ctx context.Context, mo RepositoryModel) error {
	if err := repo.Orm.WithContext(ctx).Omit(clause.Associations).Model(mo).Updates(mo).Error; err != nil {
		return err
	}
	return repo.Show(ctx, nil, mo)
}

func (repo GormRepository) BatchUpdate(ctx context.Context, mos RepositoryModels) error {
	return repo.Orm.WithContext(ctx).Updates(mos).Error
}

func (repo GormRepository) Destroy(ctx context.Context, mo RepositoryModel) error {
	return repo.Orm.WithContext(ctx).Delete(mo).Error
}

func (repo GormRepository) DestroyById(ctx context.Context, ids ...uint) error {
	query := repo.Orm.WithContext(ctx)
	if len(ids) > 1 {
		query = query.Where("id IN ?", ids)
	} else {
		query = query.Where("id = ?", ids[0])
	}
	return query.Delete(repo.Orm.Statement.Model).Error
}

func (repo GormRepository) loadResources(tx *gorm.DB, with requests.Resources, mo RepositoryModel) *gorm.DB {
	if with == nil {
		return tx
	}

	// 自定义预加载
	if model, ok := mo.(HasPreload); ok {
		for _, resource := range with {
			preload := model.Preload(resource)
			if preload != nil {
				tx = preload(tx)
			}
		}
		return tx
	}

	for _, resource := range with {
		tx = tx.Preload(resource.String())
	}
	return tx
}
