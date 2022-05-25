package clauses

import (
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/model"
	"gorm.io/gorm"
	"strings"
)

// Preload
// key: with.resource
// value: id,name,created_at
type Preload struct {
	key      string
	value    interface{}
	resource string
	fields   string
}

func NewPreloadClause(key string, value interface{}) *Preload {
	return &Preload{key: key, value: value}
}

func (p Preload) Build(tx *gorm.DB, mo *model.QueryBuilderModel) (*gorm.DB, error) {
	// todo 嵌套资源加载及model允许加载的资源定制
	if err := p.resolve(mo); err != nil {
		return tx, err
	}
	//tx.Statement.Schema.
	//_, ok := tx.Statement.Schema.Relationships.Relations[p.resource]
	//if !ok {
	//	return tx, fmt.Errorf("not support with resource: %s", p.resource)
	//}
	//relations.Field.
	// todo 筛选with加载资源的字段
	return tx.Preload(p.resource), nil
}

func (p *Preload) resolve(mo *model.QueryBuilderModel) error {
	if len(p.resource) > 0 && len(p.fields) > 0 {
		return nil
	}

	keys := strings.Split(p.key, clauseKeySplit)
	if len(keys) != 2 {
		return fmt.Errorf("invalid preload synatx: %s", p.key)
	}
	p.resource = gstr.CaseCamel(keys[1])
	return nil
}
