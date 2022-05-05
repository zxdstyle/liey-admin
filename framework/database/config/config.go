package config

import (
	"fmt"
	"gorm.io/plugin/dbresolver"
)

type Config struct {
	Default     string
	Connections map[string]Connection
}

type Connection struct {
	Driver   string
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Sources  []string // 主数据库
	Replicas []string // 从数据库
	Policy   string   // 负载均衡策略
}

func (conn Connection) GetPolicy() (dbresolver.Policy, error) {
	policy := GetPolicy(conn.Policy)
	if policy == nil {
		return nil, fmt.Errorf("%s policy is not supported", conn.Policy)
	}
	return policy, nil
}
