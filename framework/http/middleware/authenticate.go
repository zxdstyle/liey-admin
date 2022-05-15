package middleware

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zxdstyle/liey-admin/framework/exception"
	"github.com/zxdstyle/liey-admin/framework/support"
	"github.com/zxdstyle/liey-admin/framework/support/jwt"
)

func Authenticate(guardName string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		token := r.GetHeader("Authorization")
		if len(token) == 0 {
			r.SetError(exception.ErrMissingToken)
			return
		}

		claims, er := support.JWT().ParseToken(token)
		if er != nil {
			if er != jwt.TokenRefresh {
				r.SetError(gerror.NewCode(exception.CodeUnauthorized, er.Error()))
				return
			}

			newToken, err := support.JWT().RefreshToken(claims)
			if err != nil && err != jwt.TokenRefresh {
				r.SetError(err)
				return
			}
			r.Response.Header().Set("Authorization", newToken)

		}

		if claims.Guard != guardName {
			r.SetError(gerror.NewCode(exception.CodeUnauthorized, "Invalid token"))
			return
		}

		r.SetCtxVar("guard", guardName)
		r.SetCtxVar("AuthID", claims.AuthId)

		r.Middleware.Next()
	}
}
