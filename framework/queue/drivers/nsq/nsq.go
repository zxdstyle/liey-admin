package nsq

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/framework/queue/drivers/contract"
)

var driverName = "nsq"

type (
	Nsq struct {
	}
)

func NewNsqDriver() *Nsq {
	return &Nsq{}
}

func (n Nsq) Name() string {
	return driverName
}

func (n Nsq) Connect(connName string) (contract.Connection, error) {
	val, err := g.Cfg("queue").Get(context.Background(), fmt.Sprintf("connections.%s", connName))
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, fmt.Errorf("Missing configuration for connection `%s`", connName)
	}

	var cfg Config
	if err := val.Scan(&cfg); err != nil {
		return nil, err
	}
	return NewConnection(cfg)
}
