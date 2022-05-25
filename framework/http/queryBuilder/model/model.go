package model

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

const (
	columnTag = "column"
)

type (
	SpecifiedDB interface {
		Connection() string
	}

	QueryBuilderModel struct {
		Fields    *garray.StrArray
		Resources *garray.StrArray
	}
)

func New(mo interface{}) *QueryBuilderModel {
	fieldMap, _ := gstructs.FieldMap(gstructs.FieldMapInput{
		Pointer:          mo,
		PriorityTagArray: []string{"gorm"},
		RecursiveOption:  1,
	})
	fields := garray.NewStrArray()
	resources := garray.NewStrArray()

	for name, field := range fieldMap {
		var column string
		if strings.Contains(name, columnTag) {
			column = resolveColumnTag(name)
		} else if name == "-" {
			continue
		} else if tag, ok := field.TagLookup("json"); ok {
			column = gstr.ReplaceByMap(tag, map[string]string{",omitempty": ""})
		} else {
			column = gstr.CaseSnake(name)
		}

		fields.PushRight(column)
	}
	return &QueryBuilderModel{
		fields,
		resources,
	}
}

func resolveColumnTag(tag string) string {
	tags := strings.Split(tag, ";")
	for _, t := range tags {
		if strings.HasPrefix(t, columnTag) {
			column := strings.TrimLeft(t, columnTag)
			return gstr.ReplaceByMap(column, map[string]string{":": "", ";": ""})
		}
	}
	return ""
}
