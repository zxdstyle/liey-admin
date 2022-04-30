package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gvalid"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
	"net/http"
)

func Exception(r *ghttp.Request) {
	r.Middleware.Next()

	var (
		err    = r.GetError()
		status = r.Response.Status
	)

	if status == http.StatusOK && err == nil {
		return
	}

	var resp *responses.Response
	switch status {
	case http.StatusNotFound:
		resp = responses.Failed("Not Found", http.StatusNotFound)
	default:
		switch err.(type) {
		case gvalid.Error:
			resp = responses.Error(err.(gvalid.Error).FirstError(), http.StatusUnprocessableEntity)
		default:
			resp = responses.Error(err, http.StatusInternalServerError)
		}
	}

	r.Response.ClearBuffer()
	r.Response.WriteHeader(resp.Code)
	r.Response.WriteJsonExit(resp)
}
