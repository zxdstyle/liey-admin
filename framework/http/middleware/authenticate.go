package middleware

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zxdstyle/liey-admin/framework/auth"
	"github.com/zxdstyle/liey-admin/framework/exception"
	"github.com/zxdstyle/liey-admin/framework/support/jwt"
)

func Authenticate(guardName string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		token := r.GetHeader("Authorization")
		if len(token) == 0 {
			r.SetError(exception.ErrMissingToken)
			return
		}

		guard, er := auth.Guard(guardName)
		if er != nil {
			r.SetError(gerror.NewCode(exception.CodeInternalError, er.Error()))
			return
		}

		ctx, err := guard.Check(r.Context(), token)
		if err != nil {
			if err == jwt.TokenExpired {

			}

			r.SetError(err)
			return
		}

		r.SetCtx(ctx)

		r.SetCtxVar("guard", guard)

		r.Middleware.Next()
	}
}
