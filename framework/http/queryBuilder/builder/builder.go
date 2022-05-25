package builder

import (
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/clauses"
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/model"
	"gorm.io/gorm"
)

type Builder struct {
	clauses []clauses.Clause
}

func NewBuilder(clauses []clauses.Clause) *Builder {
	return &Builder{clauses: clauses}
}

func (b *Builder) WithContext(tx *gorm.DB, mo interface{}) *Statement {
	return &Statement{tx: tx.Model(mo), builder: b}
}

func (b *Builder) query(tx *gorm.DB) (*gorm.DB, error) {
	mo := model.New(tx.Statement.Model)
	var err error
	for _, cla := range b.clauses {
		tx, err = cla.Build(tx, mo)
		if err != nil {
			return tx, err
		}
	}
	return tx, err
}
