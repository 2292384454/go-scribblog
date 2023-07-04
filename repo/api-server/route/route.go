package route

import (
	"github.com/kataras/iris/v12"
	apiserver "go-scribblog/repo/api-server"
)

type handler func(server *apiserver.Server, ctx iris.Context)

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

func routeWrapper(handleFunc handler, server *apiserver.Server) func(iris.Context) {
	wrapperFunc := func(ctx iris.Context) {
		handleFunc(server, ctx)
	}
	return wrapperFunc
}

func RegisterRouter(server *apiserver.Server) {
	for _, group := range routeGroupConfig {
		partyRouter := server.App.Party(group.partyUrl)
		for _, route := range group.routeHandles {
			partyRouter.Handle(route.method, route.url, routeWrapper(route.handler, server))
		}
	}
}
