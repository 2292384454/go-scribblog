package controller

import (
	"github.com/kataras/iris/v12"
	apiserver "go-scribblog/repo/api-server"
	"go-scribblog/repo/log"
)

func Hello(server *apiserver.Server, ctx iris.Context) {
	log.Info("received the get request of /api/greet/hello")
	ctx.JSON("hello,scribblog")
}
