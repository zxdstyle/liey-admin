package database

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectGorm() *gorm.DB {
	ctx := context.Background()

	cfg := g.Cfg("database")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MustGet(ctx, "username", "root").String(),
		cfg.MustGet(ctx, "password").String(),
		cfg.MustGet(ctx, "host", "127.0.0.1").String(),
		cfg.MustGet(ctx, "port", 3306).Int(),
		cfg.MustGet(ctx, "database", "liey-admin").String(),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if g.Cfg("app").MustGet(ctx, "debug", false).Bool() {
		db = db.Debug()
	}

	return db
}
