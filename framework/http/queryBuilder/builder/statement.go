package builder

import (
	"gorm.io/gorm"
)

type Statement struct {
	tx      *gorm.DB
	builder *Builder
}

func (s *Statement) RawQuery() (*gorm.DB, error) {
	return s.builder.query(s.tx)
}

func (s *Statement) Filter() (*gorm.DB, error) {
	return s.builder.filter(s.tx)
}
