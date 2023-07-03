package route

import (
	"github.com/kataras/iris/v12"
	api_server "go-scribblog/repo/api-server"
)

type handler func(server *api_server.Server, ctx iris.Context)

type routeHandler struct {
	method  string
	url     string
	handler handler
}

func newRouteHandler(method string, url string, handleFunc handler) *routeHandler {
	return &routeHandler{
		method:  method,
		url:     url,
		handler: handleFunc,
	}
}

func routeWrapper(handleFunc handler, server *api_server.Server) func(iris.Context) {
	wrapperFunc := func(ctx iris.Context) {
		handleFunc(server, ctx)
	}
	return wrapperFunc
}

func RegisterRouter(server *api_server.Server) {
	for _, group := range routeGroupConfig {
		partyRouter := server.App.Party(group.partyUrl)
		for _, route := range group.routeHandles {
			partyRouter.Handle(route.method, route.url, routeWrapper(route.handler, server))
		}
	}
}
