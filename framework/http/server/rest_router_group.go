package server

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
)

type RouterGroup struct {
	group *ghttp.RouterGroup
}

func newRouterGroup(group *ghttp.RouterGroup) *RouterGroup {
	return &RouterGroup{group}
}

func (router RouterGroup) Resource(resource string, handler RestHandler) {
	base := getBaseName(resource)

	router.group.GET(getResourceUriIndex(resource, base), wrapHandler(handler.Index))
	router.group.GET(getResourceUriShow(resource, base), wrapHandler(handler.Show))
	router.group.POST(getResourceUriCreate(resource, base), wrapHandler(handler.Create))
	router.group.PUT(getResourceUriUpdate(resource, base), wrapHandler(handler.Update))
	router.group.DELETE(getResourceUriDestroy(resource, base), wrapHandler(handler.Destroy))
}

func (router RouterGroup) Group(prefix string, handle func(group *RouterGroup)) {
	router.group.Group(prefix, func(group *ghttp.RouterGroup) {
		handle(newRouterGroup(group))
	})
}

func wrapHandler(handler RestHandlerFunc) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		resp, err := handler(r.Context(), requests.ParseRequest(r))
		r.SetError(err)
		if resp != nil {
			r.Response.ClearBuffer()
			r.Response.WriteHeader(resp.Code)
			r.Response.WriteJson(resp)
		}
	}
}
