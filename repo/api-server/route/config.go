package route

import "go-scribblog/repo/api-server/controller"

type routeGroup struct {
	partyUrl     string
	routeHandles []*routeHandler
	auth         bool
}

var routeGroupConfig = []*routeGroup{
	{
		partyUrl: "api/greet",
		routeHandles: []*routeHandler{
			newRouteHandler("GET", "/hello", controller.Hello),
		},
		auth: false,
	},
}
