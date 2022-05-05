package driver

import (
	"github.com/zxdstyle/liey-admin/framework/database/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDriver Driver = &MySql{}

type MySql struct {
}

func (MySql) Name() string {
	return "mysql"
}

func (MySql) Config() gorm.Option {
	return &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   nil,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	}
}

func (MySql) Dialector(conn config.Connection) gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN: dsn(conn),
	})
}
