package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zxdstyle/liey-admin/framework/auth"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
	"net/http"
)

func Authenticate(guardName string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		token := r.GetHeader("Authorization")
		if len(token) == 0 {
			failed(r, "Unauthorized", http.StatusUnauthorized)
		}

		guard, er := auth.Guard(guardName)
		if er != nil {
			failed(r, er.Error(), http.StatusInternalServerError)
		}

		ctx, err := guard.Check(r.Context(), token)
		if err != nil {
			failed(r, err.Error(), http.StatusUnauthorized)
		}

		r.SetCtx(ctx)

		r.SetCtxVar("guard", guard)

		r.Middleware.Next()
	}
}

func failed(r *ghttp.Request, message string, code int) {
	r.Response.ClearBuffer()
	r.Response.WriteHeader(code)
	r.Response.WriteJsonExit(responses.Failed(message, code))
}
