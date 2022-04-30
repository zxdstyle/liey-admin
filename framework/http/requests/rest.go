package requests

import "github.com/gogf/gf/v2/net/ghttp"

type (
	RestRequest struct {
		r *ghttp.Request
	}
)

func ParseRequest(r *ghttp.Request) *RestRequest {
	return &RestRequest{
		r: r,
	}
}

func (rest RestRequest) GetGRequest() *ghttp.Request {
	return rest.r
}
