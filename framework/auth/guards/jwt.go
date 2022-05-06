package guards

import (
	"context"
	"github.com/zxdstyle/liey-admin/framework/auth/authenticate"
	"github.com/zxdstyle/liey-admin/framework/database"
	"github.com/zxdstyle/liey-admin/framework/exception"
	JWT "github.com/zxdstyle/liey-admin/framework/support/jwt"
)

type JWTGuard struct {
	jwt  *JWT.JWT
	auth authenticate.Authenticate
}

func NewJWTGuard(auth authenticate.Authenticate) *JWTGuard {
	return &JWTGuard{
		jwt:  JWT.NewJWT(),
		auth: auth,
	}
}

func (g JWTGuard) Login(auth authenticate.Authenticate) (interface{}, error) {
	return g.jwt.CreateToken(auth.Guard(), auth.GetKey())
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

	if claims.Guard != g.auth.Guard() {
		return ctx, exception.ErrInvalidToken
	}

	return context.WithValue(ctx, "AuthorityId", claims.AuthId), nil
}

func (g JWTGuard) SetUser(ctx context.Context, auth authenticate.Authenticate) context.Context {
	return context.WithValue(ctx, authority, auth)
}

func (g JWTGuard) User(ctx context.Context) authenticate.Authenticate {
	val := ctx.Value(authority)
	if val == nil {
		return nil
	}
	return val.(authenticate.Authenticate)
}

func (g JWTGuard) GetUser(ctx context.Context) (auth authenticate.Authenticate, err error) {
	db := "default"
	if val, ok := g.auth.(database.Connector); ok {
		db = val.Connection()
	}
	err = database.GetDB(db).WithContext(ctx).First(&auth, ctx.Value(authorityID)).Error

	return
}
