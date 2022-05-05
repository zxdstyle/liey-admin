package requests

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
	"gorm.io/gorm"
)

type (
	Request interface {
		GetGRequest() *ghttp.Request
		Parse(pointer interface{}) error
		Validate(pointer interface{}) error
		ResourceID(key string) uint
		GetWithResources() Resources
		GetPage() int
		GetPageSize() int
		GetLimit() int
		Paginator(mo interface{}) *responses.Paginator
		NeedPaginate() bool
		GetSelects() Selects
		GetOrder() (Orders, error)
		ToQuery(tx *gorm.DB) (*gorm.DB, error)
		Paginate(tx *gorm.DB) *gorm.DB
	}
)

var (
	ctx             = context.Background()
	defaultPageSize = g.Cfg("app").MustGet(ctx, "defaultPageSize", 20).Int()
	allDataLimit    = g.Cfg("app").MustGet(ctx, "allDataLimit", 1000).Int()
)
