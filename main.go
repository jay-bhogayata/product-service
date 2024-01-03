package main

import (
	"log/slog"

	"github.com/jay-bhogayata/product-service/config"
	"github.com/jay-bhogayata/product-service/server"
)

func main() {

	err := config.Init()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	err = server.Serve(config.AppCfg)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
