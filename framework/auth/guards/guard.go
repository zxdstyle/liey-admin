package guards

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/zxdstyle/liey-admin/framework/auth/authenticate"
)

type (
	Guard interface {
		Login(auth authenticate.Authenticate) (interface{}, error)
		Check(ctx context.Context, param interface{}) (context.Context, error)
		SetUser(ctx context.Context, auth authenticate.Authenticate) context.Context
		User(ctx context.Context) authenticate.Authenticate
		GetUser(ctx context.Context) (auth authenticate.Authenticate, err error)
	}
)

var (
	guards      = gmap.NewStrAnyMap(true)
	authority   = "Authority"
	authorityID = "AuthorityID"
)

func GetGuard(name string) Guard {
	if val, ok := guards.Search(name); ok {
		return val.(Guard)
	}
	return nil
}

func SetGuard(name string, guard Guard) {
	guards.Set(name, guard)
}
