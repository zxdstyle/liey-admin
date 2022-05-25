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
