package drivers

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/framework/exception"
	"github.com/zxdstyle/liey-admin/framework/queue/drivers/contract"
	"github.com/zxdstyle/liey-admin/framework/queue/drivers/nsq"
)

var (
	drivers = gmap.NewStrAnyMap(true)
)

type (
	Handler func(ctx context.Context, payload []byte) error

	Driver interface {
		Name() string
		Connect(name string) (contract.Connection, error)
	}
)

func init() {
	ctx := context.Background()
	if err := Register(nsq.NewNsqDriver()); err != nil {
		g.Log().Fatal(ctx, err)
	}
}

func Register(drs ...Driver) error {
	for _, driver := range drs {
		name := driver.Name()
		if drivers.Contains(name) {
			return gerror.NewCode(exception.CodeInternalError, fmt.Sprintf("driver `%s` already exists", name))
		}
		drivers.Set(name, driver)
	}
	return nil
}

func GetDriver(name string) (Driver, error) {
	val, ok := drivers.Search(name)
	if !ok {
		return nil, gerror.NewCode(exception.CodeInternalError, fmt.Sprintf("driver `%s` already exists", name))
	}
	return val.(Driver), nil
}
