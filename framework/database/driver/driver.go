package driver

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/framework/database/config"
	"gorm.io/gorm"
)

type Driver interface {
	Name() string
	Config() gorm.Option
	Dialector(conn config.Connection) gorm.Dialector
}

var (
	drivers = gmap.NewStrAnyMapFrom(map[string]interface{}{
		mysqlDriver.Name(): mysqlDriver,
	}, true)
)

func RegisterDriver(driver Driver) {
	ctx := context.TODO()
	if driver == nil {
		g.Log().Fatal(ctx, "the driver is invalid")
	}
	if _, ok := drivers.Search(driver.Name()); ok {
		g.Log().Fatalf(ctx, "the driver %s already exists, if you want to replace it, please use `ReplaceDriver`", driver.Name())
	}

	drivers.Set(driver.Name(), driver)
}

func ReplaceDriver(driver Driver) {
	drivers.Set(driver.Name(), driver)
}

func GetDriver(name string) Driver {
	dr, ok := drivers.Search(name)
	if !ok {
		return nil
	}
	return dr.(Driver)
}

func dsn(conn config.Connection) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conn.Username,
		conn.Password,
		conn.Host,
		conn.Port,
		conn.Database,
	)
}
