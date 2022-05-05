package validator

import (
	"context"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/util/gvalid"
	"github.com/zxdstyle/liey-admin/framework/validator/rules"
	"reflect"
)

var (
	Instance Validator = &GValidator{}
	validate           = validator.New()
	trans    ut.Translator
)

func init() {
	validate.SetTagName("v")
	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	err := en_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
	if err := rules.RegisterCustomRules(func(rule rules.Rule) error {
		if er := validate.RegisterTranslation(rule.Name(), trans, func(ut ut.Translator) error {
			return ut.Add(rule.Name(), rule.Message(), false)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(rule.Name(), rule.Translate(fe)...)
			return t
		}); er != nil {
			return er
		}
		return validate.RegisterValidationCtx(rule.Name(), rule.Rule())
	}); err != nil {
		panic(err)
	}
}

type Validator interface {
	Validate(ctx context.Context, pointer interface{}) error
	Parse(ctx context.Context, pointer interface{}) error
	Translate(err validator.FieldError) string
}

type GValidator struct {
}

func (GValidator) Validate(ctx context.Context, pointer interface{}) error {
	return validate.StructCtx(ctx, pointer)
}

func (v GValidator) Parse(ctx context.Context, pointer interface{}) error {
	fieldMap, er := gstructs.FieldMap(gstructs.FieldMapInput{
		Pointer:          pointer,
		PriorityTagArray: nil,
		RecursiveOption:  gstructs.RecursiveOptionEmbedded,
	})
	if er != nil {
		return er
	}

	fields := make([]string, 0)
	for fieldName, field := range fieldMap {
		//if rule := v.getFieldRule(field); len(rule) > 0 {
		//	fields = append(fields, v.resolveField(fieldName, field.Value)...)
		//}
		if field.Kind() == reflect.Ptr && field.Value.IsNil() {
			fields = append(fields, fieldName)
		}
	}
	return validate.StructExceptCtx(ctx, pointer, fields...)
}

func (v GValidator) Translate(err validator.FieldError) string {
	return err.Translate(trans)
}

func (GValidator) getFieldRule(field gstructs.Field) string {
	for _, tag := range gvalid.GetTags() {
		if rule := field.Tag(tag); len(rule) > 0 {
			return rule
		}
	}
	return ""
}

func (v GValidator) resolveField(fieldName string, value reflect.Value) []string {
	if value.Kind() == reflect.Ptr && value.IsNil() {
		return []string{}
	}

	if value.Kind() == reflect.Ptr {
		return v.resolveField(fieldName, value.Elem())
	}
	switch value.Kind() {
	case reflect.Struct:
		fields, _ := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         value.Interface(),
			RecursiveOption: gstructs.RecursiveOptionEmbedded,
		})
		res := make([]string, 0)
		for _, field := range fields {
			res = append(res, fmt.Sprintf("%s.%s", fieldName, field.Name()))
		}
		return res
	case reflect.Slice:

		res := make([]string, 0)
		fields, _ := gstructs.Fields(gstructs.FieldsInput{
			Pointer:         value.Interface(),
			RecursiveOption: gstructs.RecursiveOptionEmbedded,
		})
		for _, field := range fields {
			res = append(res, fmt.Sprintf("%s[0].%s", fieldName, field.Name()))
		}
		return res
	case reflect.Map:
		return []string{"Rules[0].PermissionRule.HttpPath"}
	default:
		return []string{fieldName, "Test.HttpPath"}
	}
}

func FirstErrorMsg(errs validator.ValidationErrors) string {
	if len(errs) == 0 {
		return ""
	}
	return translate(errs[0])
}

func translate(err validator.FieldError) string {
	return Instance.Translate(err)
}
