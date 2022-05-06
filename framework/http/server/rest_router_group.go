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
	router.GET(getResourceUriIndex(resource, base), handler.Index)
	router.GET(getResourceUriShow(resource, base), handler.Show)
	router.POST(getResourceUriStore(resource, base), handler.Create)
	router.PUT(getResourceUriUpdate(resource, base), handler.Update)
	router.DELETE(getResourceUriDestroy(resource, base), handler.Destroy)
}

func (router RouterGroup) GET(path string, handler RestHandlerFunc) {
	router.group.GET(path, wrapHandler(handler))
}

func (router RouterGroup) POST(path string, handler RestHandlerFunc) {
	router.group.POST(path, wrapHandler(handler))
}

func (router RouterGroup) PUT(path string, handler RestHandlerFunc) {
	router.group.PUT(path, wrapHandler(handler))
}

func (router RouterGroup) DELETE(path string, handler RestHandlerFunc) {
	router.group.DELETE(path, wrapHandler(handler))
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
