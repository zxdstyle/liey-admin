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
		return responses.Failed("Unauthorized", http.StatusUnauthorized)
	},
	exception.ErrMissingToken: func(err error) *responses.Response {
		return responses.Failed("Unauthorized", http.StatusUnauthorized)
	},
	exception.ErrInvalidToken: func(err error) *responses.Response {
		return responses.Failed("Unauthorized", http.StatusUnauthorized)
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
	if ve, ok := err.(validator.ValidationErrors); ok {
		return responses.Failed(validate.FirstErrorMsg(ve), http.StatusUnprocessableEntity)
	}

	if exec, ok := exceptions[err]; ok {
		return exec(err)
	}

	switch err.(type) {
	case *gerror.Error:
		code := gerror.Code(err).Code()
		if code < http.StatusOK {
			code = http.StatusInternalServerError
		}
		return responses.Failed(err.Error(), code)
	case nil:
		return responses.Failed("Internal Server Error", http.StatusInternalServerError)
	default:
		return responses.Error(err, http.StatusBadRequest)
	}
}
