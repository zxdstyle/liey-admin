package requests

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder"
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/builder"
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/clauses"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
	"github.com/zxdstyle/liey-admin/framework/validator"
)

type (
	RestRequest struct {
		r       *ghttp.Request
		with    Resources
		selects Selects
		orders  Orders
	}
)

func ParseRequest(r *ghttp.Request) *RestRequest {
	return &RestRequest{
		r: r,
	}
}

func (rest RestRequest) GetRequest() *ghttp.Request {
	return rest.r
}

func (rest RestRequest) Validate(pointer interface{}) error {
	if err := rest.r.GetStruct(pointer); err != nil {
		return err
	}
	return validator.Instance.Validate(rest.r.Context(), pointer)
}

func (rest RestRequest) Parse(pointer interface{}) error {
	if err := rest.r.GetStruct(pointer); err != nil {
		return err
	}
	return validator.Instance.Parse(rest.r.Context(), pointer)
}

func (rest RestRequest) ResourceID(key string) uint {
	return rest.r.GetRouter(key).Uint()
}

func (rest RestRequest) ToQuery() *builder.Builder {
	return queryBuilder.NewBuilderFromRequest(rest.r)
}

func (rest RestRequest) GetClauses() []clauses.Clause {
	return queryBuilder.ParseClauses(rest.r)
}

func (rest RestRequest) GetPage() int {
	return rest.r.GetQuery("page", 1).Int()
}

func (rest RestRequest) GetPageSize() int {
	return rest.r.GetQuery("pageSize", defaultPageSize).Int()
}

func (rest RestRequest) GetLimit() int {
	return rest.r.GetQuery("limit", allDataLimit).Int()
}

func (rest RestRequest) NeedPaginate() bool {
	return !rest.r.GetQuery("page").IsEmpty()
}

func (rest RestRequest) Paginator(mo interface{}) *responses.Paginator {
	return &responses.Paginator{
		Data: mo,
		Meta: responses.Meta{
			Page:  rest.GetPage(),
			Total: 0,
		},
	}
}

func (rest RestRequest) ID() uint {
	return rest.r.GetCtxVar("AuthID").Uint()
}
