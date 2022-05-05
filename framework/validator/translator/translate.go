package translator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/zxdstyle/liey-admin/framework/translation"
)

func FirstErrorMsg(errs validator.ValidationErrors) string {
	if len(errs) == 0 {
		return ""
	}
	return translate(errs[0])
}

func translate(err validator.FieldError) string {
	key := fmt.Sprintf("validation.%s", err.Tag())
	return translation.Translator().Translate(key, err.Field())
}
