package clauses

import (
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/model"
	"gorm.io/gorm"
	"strings"
)

var (
	eq    = "eq"    // 等于
	lt    = "lt"    // 小于
	lte   = "lte"   // 小于等于
	gt    = "gt"    // 大于
	gte   = "gte"   // 大于等于
	in    = "in"    // in
	match = "match" // 匹配

	operates = garray.NewStrArrayFrom([]string{eq, lt, lte, gt, gte, in, match})
)

type Where struct {
	key     string
	value   interface{}
	field   string
	operate string
}

// NewWhereClause
// key: where.field.eq
func NewWhereClause(key string, value interface{}) *Where {
	return &Where{
		key:   key,
		value: value,
	}
}

func (w Where) Build(tx *gorm.DB, mo *model.QueryBuilderModel) (*gorm.DB, error) {
	if err := w.resolve(mo); err != nil {
		return tx, err
	}

	// todo n
	switch w.operate {
	case eq:
		return tx.Where(fmt.Sprintf("`%s` = ?", w.field), w.value), nil
	case lt:
		return tx.Where(fmt.Sprintf("`%s` < ?", w.field), w.value), nil
	case lte:
		return tx.Where(fmt.Sprintf("`%s` <= ?", w.field), w.value), nil
	case gt:
		return tx.Where(fmt.Sprintf("`%s` > ?", w.field), w.value), nil
	case gte:
		return tx.Where(fmt.Sprintf("`%s` >= ?", w.field), w.value), nil
	case in:
		val := w.value
		if !gvar.New(w.value).IsSlice() {
			v := gconv.String(w.value)
			val = strings.Split(v, ",")
		}
		return tx.Where(fmt.Sprintf("`%s` IN ?", w.field), val), nil
	case match:
		return tx.Where(fmt.Sprintf("`%s` LIKE ?", w.field), gconv.String(w.value)+"%"), nil
	default:
		return tx, nil
	}
}

func (w *Where) resolve(mo *model.QueryBuilderModel) error {
	keys := strings.Split(w.key, clauseKeySplit)
	if len(keys) != 3 {
		return fmt.Errorf("invalid query synatx: %s", w.key)
	}
	if !operates.ContainsI(keys[2]) {
		return fmt.Errorf("not support query operate: %s", w.key)
	}
	w.operate = keys[2]

	if !mo.Fields.Contains(keys[1]) {
		return fmt.Errorf("not support query field: %s", w.key)
	}
	w.field = keys[1]
	return nil
}
