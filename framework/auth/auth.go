package auth

import (
	"fmt"
	"github.com/zxdstyle/liey-admin/framework/auth/guards"
)

func Guard(name string) (guards.Guard, error) {
	guard := guards.GetGuard(name)
	if guard == nil {
		return guard, fmt.Errorf("auth guard %s is not defined", name)
	}
	return guard, nil
}
