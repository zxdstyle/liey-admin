package queryBuilder

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/builder"
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/clauses"
	"strings"
)

const (
	whereClausePrefix   = "where"
	selectClausePrefix  = "select"
	preloadClausePrefix = "with"
	orderClausePrefix   = "order"
)

func ParseClauses(r *ghttp.Request) []clauses.Clause {
	rawQuery := strings.Replace(r.URL.RawQuery, ".", "*", -1)
	queries, _ := gstr.Parse(rawQuery)
	cls := make([]clauses.Clause, 0)
	for key, value := range queries {
		key = strings.Replace(key, "*", ".", -1)
		if gstr.HasPrefix(key, whereClausePrefix) {
			cls = append(cls, clauses.NewWhereClause(key, value))
		}
		if gstr.HasPrefix(key, orderClausePrefix) {
			cls = append(cls, clauses.NewOrderClause(key, value))
		}
		if key == selectClausePrefix {
			cls = append(cls, clauses.NewSelectClause(key, value))
		}
		if gstr.HasPrefix(key, preloadClausePrefix) {
			cls = append(cls, clauses.NewPreloadClause(key, value))
		}
	}
	return cls
}

func NewBuilderFromRequest(cls []clauses.Clause) *builder.Builder {
	return builder.NewBuilder(cls)
}
