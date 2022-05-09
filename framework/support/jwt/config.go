package jwt

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Config struct {
	Secret               []byte
	Issuer               string
	TTL                  int
	RefreshTTL           int
	BlacklistGracePeriod uint
}

var defaultConfig = Config{
	Secret:               []byte("liey-admin-secret-key"),
	Issuer:               "liey-admin",
	TTL:                  3600,
	RefreshTTL:           60 * 60 * 24 * 7,
	BlacklistGracePeriod: 10,
}

func GetConfig() Config {
	ctx := context.Background()
	val, _ := g.Cfg("auth").Get(ctx, "jwt")
	var cfg Config
	if err := val.Struct(&cfg); err != nil {
		g.Log().Warningf(ctx, "failed to get jwt configuration, using default configuration: %s", err.Error())
		return defaultConfig
	}
	return cfg
}
