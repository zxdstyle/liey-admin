package requests

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
	"gorm.io/gorm"
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

func (rest RestRequest) GetGRequest() *ghttp.Request {
	return rest.r
}

func (rest RestRequest) Validate(pointer interface{}) error {
	return rest.r.Parse(pointer)
}

func (rest RestRequest) ResourceID(key string) uint {
	return rest.r.GetRouter(key).Uint()
}

func (rest RestRequest) GetWithResources() Resources {
	if rest.with == nil {
		rest.with = parseWith(rest.r)
	}
	return rest.with
}

func (rest RestRequest) GetSelects() Selects {
	if rest.selects == nil {
		rest.selects = parseSelects(rest.r)
	}
	return rest.selects
}

func (rest RestRequest) GetOrder() (Orders, error) {
	if rest.orders == nil {
		orders, err := parseOrder(rest.r)
		if err != nil {
			return nil, err
		}
		rest.orders = orders
	}

	return rest.orders, nil
}

func (rest RestRequest) ToQuery(tx *gorm.DB) (*gorm.DB, error) {
	selects := rest.GetSelects()
	if selects != nil && len(selects) > 0 {
		tx = tx.Select(selects)
	}

	orders, err := rest.GetOrder()
	if err != nil {
		return tx, err
	}
	if orders != nil && len(orders) > 0 {
		tx = orders.Query(tx)
	}

	return tx, nil
}

func (rest RestRequest) Paginate(tx *gorm.DB) *gorm.DB {
	page := rest.GetPage()
	pageSize := rest.GetPageSize()
	offset := (page - 1) * pageSize
	return tx.Offset(offset).Limit(pageSize)
}

func (rest RestRequest) GetPage() int {
	return rest.r.GetQuery("page", 1).Int()
}

func (rest RestRequest) GetPageSize() int {
	return rest.r.GetQuery("pageSize", defaultPageSize).Int()
}

func (rest RestRequest) GetLimit() int {
	return allDataLimit
}

func (rest RestRequest) NeedPaginate() bool {
	return !rest.r.GetQuery("page").IsEmpty()
}

func (rest RestRequest) Paginator(mo interface{}) *responses.Paginator {
	return &responses.Paginator{
		Page:  rest.GetPage(),
		Data:  mo,
		Total: 0,
	}
}
