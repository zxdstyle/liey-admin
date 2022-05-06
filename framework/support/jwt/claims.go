package jwt

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	AuthId    uint
	RefreshAt *jwt.NumericDate
	Guard     string
}

func NewClaims(guard string, authId uint, cfg Config) jwt.Claims {
	now := gtime.Now()
	return Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.Issuer,
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(cfg.TTL) * time.Second).Time),
			NotBefore: jwt.NewNumericDate(now.Time),
			IssuedAt:  jwt.NewNumericDate(now.Time),
		},
		AuthId:    authId,
		RefreshAt: jwt.NewNumericDate(now.Add(time.Duration(cfg.RefreshTTL) * time.Second).Time),
		Guard:     guard,
	}
}
