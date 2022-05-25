package clauses

import (
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/model"
	"gorm.io/gorm"
)

const (
	clauseKeySplit = "."
)

type Clause interface {
	Build(tx *gorm.DB, mo *model.QueryBuilderModel) (*gorm.DB, error)
}
