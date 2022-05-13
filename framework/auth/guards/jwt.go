package guards

import (
	"context"
	"github.com/zxdstyle/liey-admin/framework/auth/authenticate"
	"github.com/zxdstyle/liey-admin/framework/exception"
	JWT "github.com/zxdstyle/liey-admin/framework/support/jwt"
)

type (
	JWTGuard struct {
		jwt  *JWT.JWT
		auth authenticate.Authenticate
	}

	contextKey string
)

func NewJWTGuard(auth authenticate.Authenticate) *JWTGuard {
	return &JWTGuard{
		jwt:  JWT.NewJWT(),
		auth: auth,
	}
}

func (g JWTGuard) Login(auth authenticate.Authenticate) (interface{}, error) {
	return g.jwt.CreateToken(auth)
}

func (g JWTGuard) ParseToken() {

}
func (g JWTGuard) Check(ctx context.Context, param interface{}) (context.Context, error) {
	token, ok := param.(string)
	if !ok {
		return ctx, exception.ErrInvalidToken
	}

	claims, err := g.jwt.ParseToken(token)
	if err != nil {
		return ctx, err
	}

	if claims.Guard != g.auth.GuardName() {
		return ctx, exception.ErrInvalidToken
	}

	return context.WithValue(ctx, contextKey(g.auth.GuardName()), claims.AuthId), nil
}

func (g JWTGuard) ID(ctx context.Context) uint {
	val := ctx.Value(contextKey(authorityID))
	if val == nil {
		return 0
	}
	return val.(uint)
}
