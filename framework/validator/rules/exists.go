package rules

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gvalid"
	"github.com/zxdstyle/liey-admin/framework/adm/instances"
)

// ExistsRule exists:table,field,ignore
type ExistsRule struct {
}

func (ExistsRule) Name() string {
	return "exists"
}

func (ExistsRule) Rule() gvalid.RuleFunc {
	return func(ctx context.Context, in gvalid.RuleFuncInput) error {
		rs := gstr.Split(in.Rule, ":")
		if len(rs) == 0 {
			return nil
		}

		var table string
		if len(rs) == 1 {
			return fmt.Errorf("validation rule exists requires at least 1 parameters")
		}

		args := gstr.Split(rs[1], ",")
		if len(args) == 0 {
			return fmt.Errorf("validation rule exists requires at least 1 parameters")
		}

		field := "id"
		table = args[0]
		if len(args) >= 2 {
			field = args[1]
		}
		if len(args) >= 3 {
			ignore := args[2]
			if ignore == in.Value.String() {
				return nil
			}
		}

		var count int64
		if err := instances.DB().WithContext(ctx).Table(table).Where(fmt.Sprintf("%s = ?", field), in.Value.Val()).Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			if gstr.HasPrefix(in.Message, "The") {
				return fmt.Errorf("the %s in the %s does not exist", field, table)
			}
			return fmt.Errorf(in.Message)
		}

		return nil
	}
}
