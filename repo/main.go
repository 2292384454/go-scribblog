package main

import (
	"fmt"
	apiserver "go-scribblog/repo/cmd/api-server"
	"os"
)

func main() {
	server := apiserver.GetServerCommand()
	if err := server.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
