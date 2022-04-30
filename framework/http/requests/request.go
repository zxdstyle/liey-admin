package requests

import "github.com/gogf/gf/v2/net/ghttp"

type (
	Request interface {
		GetGRequest() *ghttp.Request
	}
)
