package main

import (
	"fmt"
	api_server "go-scribblog/repo/cmd/api-server"
	"os"
)

func main() {
	server := api_server.GetServerCommand()
	if err := server.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
