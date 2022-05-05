package requests

import (
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/zxdstyle/liey-admin/framework/exception"
	"gorm.io/gorm"
)

const (
	OrderByDesc = "DESC"
	OrderByAsc  = "ASC"
)

type (
	Order struct {
		Order   string `json:"order"`
		OrderBy string `json:"order_by"`
	}

	Orders []Order
)

func (orders Orders) Query(tx *gorm.DB) *gorm.DB {
	for _, order := range orders {
		tx = tx.Order(fmt.Sprintf("`%s` %s", order.Order, order.OrderBy))
	}
	return tx
}

func parseOrder(r *ghttp.Request) (Orders, error) {
	query := r.GetQuery("_order")
	if query.IsEmpty() {
		return Orders{
			{Order: "id", OrderBy: OrderByDesc},
		}, nil
	}

	if !query.IsMap() {
		return nil, exception.ErrInvalidOrderArgument
	}

	var orders Orders
	for field, sort := range query.MapStrStr() {
		sort = gstr.ToUpper(sort)
		if sort != OrderByDesc && sort != OrderByAsc {
			return nil, exception.ErrInvalidOrderArgument
		}
		orders = append(orders, Order{OrderBy: field, Order: sort})
	}
	return orders, nil
}
