package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	apiserver "go-scribblog/repo/api-server"
	"go-scribblog/repo/api-server/view"
	"go-scribblog/repo/log"
)

func Hello(server *apiserver.Server, ctx iris.Context) {
	req := view.HelloReq{}
	if err := ParseRequest(ctx, &req); err != nil {
		ReturnRequestParseError(ctx, err.Error())
		return
	}
	log.Info("received the get request of /api/greet/hello")
	ctx.JSON(view.Success(fmt.Sprintf("hello,%s", req.Name)))
}
