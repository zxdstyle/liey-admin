package server

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zxdstyle/liey-admin/framework/http/middleware"
)

type RestServer struct {
	server *ghttp.Server
}

func NewRestServer(name string) *RestServer {
	ctx := context.Background()
	s := ghttp.GetServer(name)

	serverCfg := g.Cfg("server").MustData(ctx)
	if err := s.SetConfigWithMap(serverCfg); err != nil {
		panic(err)
	}
	s.SetDumpRouterMap(false)

	// error handler
	s.Use(middleware.Exception)

	if !g.Cfg("app").MustGet(ctx, "debug", false).Bool() {
		s.SetErrorStack(false)
	}
	return &RestServer{server: s}
}

func (rest RestServer) Group(prefix string, handle func(group *RouterGroup)) {
	rest.server.Group(prefix, func(group *ghttp.RouterGroup) {
		handle(newRouterGroup(group))
	})
}

func (rest RestServer) Run() {
	rest.server.Run()
}
