package instances

import (
	"github.com/zxdstyle/liey-admin/framework/database"
	"gorm.io/gorm"
)

func DB(name ...string) *gorm.DB {
	key := "default"
	if len(name) > 0 {
		key = name[0]
	}
	return database.GetDB(key)
}
