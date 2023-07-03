package api_server

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	api_server "go-scribblog/repo/api-server"
	"go-scribblog/repo/api-server/conf"
	"go-scribblog/repo/api-server/route"
	"go-scribblog/repo/log"
	"sync"
)

var configPath string
var serverCmd *cobra.Command
var once sync.Once

func GetServerCommand() *cobra.Command {
	// 创建 api server单例
	once.Do(func() {
		serverCmd = &cobra.Command{
			Use:   "api-server",
			Short: "api-server",
			Long:  `api-server`,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("run api-server with config = ", configPath)
				start()
			},
		}
		serverCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "../config/config.yaml",
			"config file path")
	})
	return serverCmd
}

func start() {
	config, err := conf.LoadConfigFromFile(configPath)
	if err != nil {
		panic(err)
	}
	if err = log.InitLog(&config.Log); err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := api_server.GetInstance(config, ctx)
	if err != nil {
		panic(err)
	}

	//注册路由
	route.RegisterRouter(server)

	if err = server.Start(); err != nil {
		panic(err)
	}
}
