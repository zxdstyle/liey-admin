package validator

import (
	"github.com/gogf/gf/v2/util/gvalid"
	"github.com/zxdstyle/liey-admin/framework/validator/rules"
)

type Rule interface {
	Name() string
	Rule() gvalid.RuleFunc
}

var customRules = []Rule{
	&rules.ExistsRule{},
	&rules.UniqueRule{},
}

func RegisterCustomRules() {
	for _, rule := range customRules {
		gvalid.RegisterRule(rule.Name(), rule.Rule())
	}
}
