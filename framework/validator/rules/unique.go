package rules

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/zxdstyle/liey-admin/framework/database"
)

// UniqueRule unique=db.table field ignoreFieldName
type UniqueRule struct {
}

func (UniqueRule) Name() string {
	return "unique-db"
}

func (UniqueRule) Message() string {
	return "The {0} already exists."
}

func (UniqueRule) Translate(err validator.FieldError) []string {
	return []string{
		err.Field(),
	}
}

func (UniqueRule) Rule() validator.FuncCtx {
	return func(ctx context.Context, fl validator.FieldLevel) bool {
		var (
			db    string
			table string
			field = gstr.CaseSnake(fl.FieldName())
		)
		rs := gstr.Split(fl.Param(), " ")
		if len(rs) == 0 {
			return false
		}

		db, table = resolveTable(rs[0])
		l := len(rs)
		if l >= 2 {
			field = rs[1]
		}

		p := fl.Parent()
		var count int64
		query := database.GetDB(db).WithContext(ctx).Table(table).Where(fmt.Sprintf("`%s` = ?", field), fl.Field().Interface())

		//data := in.Data.MapStrVar()
		key := p.FieldByName("ID")
		if !key.IsZero() {
			query = query.Where("`id` <> ?", key.Interface())
		}

		if err := query.Count(&count).Error; err != nil {
			return false
		}

		return count == 0
	}
}

func resolveTable(p string) (db string, table string) {
	database := gstr.Split(p, ".")
	switch len(database) {
	case 2:
		db = database[0]
		table = database[1]
	case 1:
		table = database[0]
	}
	if len(db) == 0 {
		db = "default"
	}
	return
}
