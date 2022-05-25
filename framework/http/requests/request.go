package requests

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/builder"
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/clauses"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

type (
	Request interface {
		GetRequest() *ghttp.Request
		Parse(pointer interface{}) error
		Validate(pointer interface{}) error
		ResourceID(key string) uint
		GetPage() int
		GetPageSize() int
		GetLimit() int
		Paginator(mo interface{}) *responses.Paginator
		NeedPaginate() bool
		ToQuery() *builder.Builder
		GetClauses() []clauses.Clause
		// ID 当前授权用户标识
		ID() uint
	}
)

var (
	ctx             = context.Background()
	defaultPageSize = g.Cfg("app").MustGet(ctx, "defaultPageSize", 20).Int()
	allDataLimit    = g.Cfg("app").MustGet(ctx, "allDataLimit", 1000).Int()
)
