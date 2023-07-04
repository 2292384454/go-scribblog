package api_server

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"go-scribblog/repo/api-server/conf"
	"sync"
)

type Server struct {
	Name string
	Conf *conf.Config
	ctx  context.Context
	App  *iris.Application
}

var server *Server
var once sync.Once

// GetInstance 获取 server 实例
func GetInstance(cfg *conf.Config, ctx context.Context) *Server {
	// 生成 iris 实例
	once.Do(func() {
		app := iris.New()
		app.Validator = validator.New() //为 iris App 绑定一个 validator 实例
		server = &Server{
			Name: cfg.Server.Name,
			Conf: cfg,
			ctx:  ctx,
			App:  app,
		}
	})
	return server
}

// Start 启动server
func (server *Server) Start() error {
	return server.App.Run(
		iris.Addr(server.Conf.Server.WebAddr),
		iris.WithCharset("UTF-8"),
		iris.WithOptimizations,
	)
}
