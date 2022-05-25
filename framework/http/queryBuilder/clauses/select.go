package clauses

import (
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/model"
	"gorm.io/gorm"
)

type Select struct {
	key   string
	value interface{}
}

func NewSelectClause(key string, value interface{}) *Select {
	return &Select{
		key:   key,
		value: value,
	}
}

func (s Select) Build(tx *gorm.DB, mo *model.QueryBuilderModel) (*gorm.DB, error) {
	//TODO implement me
	panic("implement me")
}
