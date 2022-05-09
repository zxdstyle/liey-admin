package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zxdstyle/liey-admin/framework/exception"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
	validate "github.com/zxdstyle/liey-admin/framework/validator"
	"gorm.io/gorm"
	"net/http"
)

var exceptions = map[error]func(err error) *responses.Response{
	gorm.ErrRecordNotFound: func(err error) *responses.Response {
		return responses.Error(err, http.StatusNotFound)
	},
	exception.ErrUnauthorized: func(err error) *responses.Response {
		return responses.Error(err, http.StatusUnauthorized)
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
	case validator.ValidationErrors:
		return responses.Failed(validate.FirstErrorMsg(err.(validator.ValidationErrors)), http.StatusUnprocessableEntity)
	case *gerror.Error:
		return responses.Failed(err.Error(), gerror.Code(err).Code())
	case nil:
		return responses.Failed("Internal Server Error", http.StatusInternalServerError)
	default:
		if exce, ok := exceptions[err]; ok {
			return exce(err)
		}
		return responses.Error(err, http.StatusBadRequest)
	}
}
