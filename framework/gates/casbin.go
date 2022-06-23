package gates

import (
	"context"
	"github.com/casbin/casbin/v2"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/framework/exception"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	CasbinGate struct {
		enforcer *casbin.SyncedEnforcer
	}
)

var (
	defaultCasbinModelPath = "config/rbac_model.conf"
	ErrForbidden           = gerror.NewCode(exception.CodeForbidden, "无权限访问")

	casbinInstance *CasbinGate
)

func newCasbinGate(db *gorm.DB) (*CasbinGate, error) {
	if casbinInstance != nil {
		return casbinInstance, nil
	}
	ctx := context.Background()
	enforcer, err := initAdapter(ctx, db)
	if err != nil {
		return nil, err
	}
	g := &CasbinGate{
		enforcer: enforcer,
	}
	return g, nil
}

func initAdapter(ctx context.Context, db *gorm.DB) (*casbin.SyncedEnforcer, error) {
	db = db.Session(&gorm.Session{
		Logger: db.Logger.LogMode(logger.Error),
	})
	a, err := adapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	val, _ := g.Cfg("scaffold").Get(ctx, "model_path", defaultCasbinModelPath)
	e, er := casbin.NewSyncedEnforcer(val.String(), a)
	if er != nil {
		return nil, er
	}
	if err := e.LoadPolicy(); err != nil {
		return nil, err
	}
	return e, nil
}

func (g CasbinGate) Can(authID uint, path, action string) error {
	ok, err := g.enforcer.Enforce(authID, path, action)
	if err != nil {
		return err
	}
	if !ok {
		return ErrForbidden
	}
	return nil
}
