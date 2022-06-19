package gates

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/framework/database"
)

func Casbin() (*CasbinGate, error) {
	ctx := context.Background()
	cfg, err := g.Cfg().Get(ctx, "auth.casbin.dbConnect")
	if err != nil {
		return nil, err
	}
	db := database.GetDB(cfg.String())

	return newCasbinGate(db)
}
