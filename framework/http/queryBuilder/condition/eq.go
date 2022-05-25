package condition

import (
	"fmt"
	"gorm.io/gorm"
)

type Eq struct {
	tx *gorm.DB
}

func NewEqCondition(tx *gorm.DB) *Eq {
	return &Eq{tx}
}

func (e Eq) Build(field string, value interface{}) *gorm.DB {
	return e.tx.Where(fmt.Sprintf("`%s` = ?", field), value)
}
