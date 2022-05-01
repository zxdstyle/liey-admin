package responses

import (
	"net/http"
)

type (
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func (e Response) Error() string {
	return e.Message
}

func Error(err error, code int) *Response {
	return Failed(err.Error(), code)
}

func Failed(message string, code int) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func Custom(data interface{}, code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func Success(data interface{}) *Response {
	return Custom(data, http.StatusOK, "Success")
}

func Created(data interface{}) *Response {
	return Custom(data, http.StatusCreated, "Created")
}

func NoContent() *Response {
	return Custom(nil, http.StatusNoContent, "Deleted")
}
