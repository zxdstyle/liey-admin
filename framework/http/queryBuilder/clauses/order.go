package clauses

import (
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/model"
	"gorm.io/gorm"
	"strings"
)

var (
	sortValues = garray.NewStrArrayFrom([]string{"desc", "asc", "DESC", "ASC"})
)

type Order struct {
	key   string
	value interface{}
	field string
	sort  string
}

// NewOrderClause
// key: order.field
// value: desc | asc | DESC | ASC
func NewOrderClause(key string, value interface{}) *Order {
	return &Order{key: key, value: value}
}

func (c Order) Build(tx *gorm.DB, mo *model.QueryBuilderModel) (*gorm.DB, error) {
	if err := c.resolve(); err != nil {
		return tx, err
	}
	return tx.Order(fmt.Sprintf("`%s` %s", c.field, c.sort)), nil
}

func (c *Order) resolve() error {
	if len(c.field) > 0 && len(c.sort) > 0 {
		return nil
	}

	keys := strings.Split(c.key, clauseKeySplit)
	if len(keys) != 2 {
		return fmt.Errorf("invalid query synatx: %s", c.key)
	}
	c.field = keys[1]

	val := gconv.String(c.value)
	if len(val) == 0 || !sortValues.Contains(val) {
		return fmt.Errorf("invalid order query value: %v", c.value)
	}
	c.sort = strings.ToUpper(val)

	return nil
}
