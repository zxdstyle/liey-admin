package database

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/zxdstyle/liey-admin/framework/exception"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

var migrations = gmap.NewStrAnyMap(true)

type (
	Migration interface {
		Models() []bases.RepositoryModel
	}
)

func RegisterMigration(group string, migration ...Migration) error {
	for i, _ := range migration {
		if migrations.Contains(group) {
			return exception.ErrRepeatedMigration
		}
		migrations.SetIfNotExist(group, migration[i])
	}

	return nil
}

func GetMigrations(groups ...string) map[string]Migration {
	mis := make(map[string]Migration)
	if len(groups) == 0 {
		migrations.Iterator(func(k string, v interface{}) bool {
			mis[k] = v.(Migration)
			return true
		})
		return mis
	}

	for _, group := range groups {
		mis[group] = migrations.Get(group).(Migration)
	}
	return mis
}

func GetMigration(group string) Migration {
	val, ok := migrations.Search(group)
	if !ok {
		return nil
	}
	return val.(Migration)
}

func AllMigrationKeys() []string {
	return migrations.Keys()
}
