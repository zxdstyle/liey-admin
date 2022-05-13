package database

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/zxdstyle/liey-admin/framework/database/config"
	"github.com/zxdstyle/liey-admin/framework/database/driver"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var (
	cfg            config.Config
	connectionPool = gmap.NewStrAnyMap(true)
)

type (
	Connector interface {
		Connection() string
	}
)

func init() {
	ctx := context.Background()
	val := g.Cfg("database").MustData(ctx)
	if err := gconv.Scan(val, &cfg); err != nil {
		panic(err)
	}

	debug, _ := g.Cfg("app").Get(ctx, "debug", false)

	initDialectors(cfg.Connections)

	connect(cfg.Connections, debug.Bool())
}

func connect(connections map[string]config.Connection, debug bool) {

	for name, connection := range connections {
		dial := dialectors.Get(name)
		db, er := gorm.Open(dial)
		if er != nil {
			panic(er)
		}

		policy, e := connection.GetPolicy()
		if e != nil {
			panic(e)
		}

		if err := db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  dialectors.Gets(connection.Sources...),
			Replicas: dialectors.Gets(connection.Replicas...),
			Policy:   policy,
		})); err != nil {
			panic(err)
		}

		if debug {
			db = db.Debug()
		}

		connectionPool.Set(name, db)
	}
}

func initDialectors(options map[string]config.Connection) {
	for name, option := range options {
		dr := driver.GetDriver(option.Driver)
		dialectors.Set(name, dr.Dialector(option))
	}
}

func GetDB(name string) *gorm.DB {
	val, ok := connectionPool.Search(name)
	if !ok {
		return nil
	}
	return val.(*gorm.DB)
}
