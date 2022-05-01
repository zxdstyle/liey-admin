package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gvalid"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
	"gorm.io/gorm"
	"net/http"
)

var exceptions = map[error]func(err error) *responses.Response{
	gorm.ErrRecordNotFound: func(err error) *responses.Response {
		return responses.Error(err, http.StatusNotFound)
	},
}

func Exception(r *ghttp.Request) {
	r.Middleware.Next()

	var (
		err    = r.GetError()
		status = r.Response.Status
	)

	if isSuccess(status) && err == nil {
		return
	}

	var resp *responses.Response
	switch status {
	case http.StatusNotFound:
		resp = responses.Failed("Not Found", http.StatusNotFound)
	default:
		resp = rejectError(err)
	}

	r.Response.ClearBuffer()
	r.Response.WriteHeader(resp.Code)
	r.Response.WriteJsonExit(resp)
}

func isSuccess(status int) bool {
	return status == http.StatusOK || status == http.StatusCreated || status == http.StatusAccepted || status == http.StatusNoContent
}

func rejectError(err error) (resp *responses.Response) {
	switch err.(type) {
	case gvalid.Error:
		return responses.Error(err.(gvalid.Error).FirstError(), http.StatusUnprocessableEntity)
	default:
		if exce, ok := exceptions[err]; ok {
			return exce(err)
		}
		return responses.Error(err, http.StatusInternalServerError)
	}
}
