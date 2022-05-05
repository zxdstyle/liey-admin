package rules

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
)

// UniqueRule unique=db.table field ignoreFieldName
type UniqueRule struct {
}

func (UniqueRule) Name() string {
	return "unique"
}

func (UniqueRule) Rule() validator.FuncCtx {
	return func(ctx context.Context, fl validator.FieldLevel) bool {
		fmt.Println(fl.Param())
		//rs := gstr.Split(in.Rule, ":")
		//if len(rs) == 0 {
		//	return nil
		//}
		//
		//if len(rs) == 1 {
		//	return fmt.Errorf("validation rule exists requires at least 1 parameters")
		//}
		//
		//args := gstr.Split(rs[1], ",")
		//if len(args) == 0 {
		//	return fmt.Errorf("validation rule exists requires at least 1 parameters")
		//}
		//
		//field := "id"
		//table := args[0]
		//if len(args) >= 2 {
		//	field = args[1]
		//}
		//
		//if in.Value.IsEmpty() {
		//	return nil
		//}
		//
		//var count int64
		//query := instances.DB().WithContext(ctx).Table(table).Where(fmt.Sprintf("%s = ?", field), in.Value.Val())
		//
		//data := in.Data.MapStrVar()
		//if val, ok := data["id"]; ok && !val.IsEmpty() {
		//	query = query.Where("id <> ?", val.Int())
		//}
		//
		//if err := query.Count(&count).Error; err != nil {
		//	return err
		//}
		//
		//if count > 0 {
		//	msg := gstr.ReplaceByMap(in.Message, map[string]string{
		//		"{attribute}": field,
		//		"{value}":     in.Value.String(),
		//	})
		//	return fmt.Errorf(msg)
		//}

		return false
	}
}
