package validator

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/util/gvalid"
	"github.com/zxdstyle/liey-admin/framework/validator/rules"
	"reflect"
)

var (
	Instance Validator = &GValidator{}
	validate           = validator.New()
)

func init() {
	validate.SetTagName("v")
	if err := rules.RegisterCustomRules(func(rule rules.Rule) error {
		return validate.RegisterValidationCtx(rule.Name(), rule.Rule())
	}); err != nil {
		panic(err)
	}
}

type Validator interface {
	Validate(ctx context.Context, pointer interface{}) error
	Parse(ctx context.Context, pointer interface{}) error
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
		if field.Value.Kind() == reflect.Ptr && field.Value.IsNil() {
			continue
		}

		if rule := v.getFieldRule(field); len(rule) > 0 {
			fields = append(fields, fieldName)
		}
	}

	return validate.StructCtx(ctx, pointer)
}

func (GValidator) getFieldRule(field gstructs.Field) string {
	for _, tag := range gvalid.GetTags() {
		if rule := field.Tag(tag); len(rule) > 0 {
			return rule
		}
	}
	return ""
}
