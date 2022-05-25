package condition

import "gorm.io/gorm"

type Condition interface {
	Build(key string, value interface{}) *gorm.DB
}
