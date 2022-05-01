package rules

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gvalid"
	"github.com/zxdstyle/liey-admin/framework/adm/instances"
)

// UniqueRule unique:table,field,ignore
type UniqueRule struct {
}

func (UniqueRule) Name() string {
	return "unique"
}

func (UniqueRule) Rule() gvalid.RuleFunc {
	return func(ctx context.Context, in gvalid.RuleFuncInput) error {
		rs := gstr.Split(in.Rule, ":")
		if len(rs) == 0 {
			return nil
		}

		if len(rs) == 1 {
			return fmt.Errorf("validation rule exists requires at least 1 parameters")
		}

		args := gstr.Split(rs[1], ",")
		if len(args) == 0 {
			return fmt.Errorf("validation rule exists requires at least 1 parameters")
		}

		field := "id"
		table := args[0]
		if len(args) >= 2 {
			field = args[1]
		}

		if in.Value.IsEmpty() {
			return nil
		}

		var count int64
		query := instances.DB().WithContext(ctx).Table(table).Where(fmt.Sprintf("%s = ?", field), in.Value.Val())

		data := in.Data.MapStrVar()
		if val, ok := data["id"]; ok && !val.IsEmpty() {
			query = query.Where("id <> ?", val.Int())
		}

		if err := query.Count(&count).Error; err != nil {
			return err
		}

		if count > 0 {
			if gstr.HasPrefix(in.Message, "The") {
				return fmt.Errorf("the %s has already been taken", field)
			}
			return fmt.Errorf(in.Message)
		}

		return nil
	}
}
