package rules

import (
	"github.com/go-playground/validator/v10"
)

type Rule interface {
	Name() string
	Rule() validator.FuncCtx
}

var customRules = []Rule{
	//&ExistsRule{},
	&UniqueRule{},
}

func RegisterCustomRules(register func(rule Rule) error) error {
	for _, rule := range customRules {
		if err := register(rule); err != nil {
			return err
		}
	}
	return nil
}
