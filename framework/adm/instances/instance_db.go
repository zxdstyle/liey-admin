package instances

import (
	"github.com/zxdstyle/liey-admin/framework/database"
	"gorm.io/gorm"
)

const (
	frameCoreDatabase = "liey.core.database"
)

func DB() *gorm.DB {
	return instances.GetOrSetFunc(frameCoreDatabase, func() interface{} {
		return database.ConnectGorm()
	}).(*gorm.DB)
}
