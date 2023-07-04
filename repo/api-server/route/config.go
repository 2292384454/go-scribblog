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
			newRouteHandler("POST", "/hello", controller.Hello),
		},
		auth: false,
	},
}
