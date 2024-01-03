package main

import (
	"log"

	"github.com/jay-bhogayata/product-service/config"
	"github.com/jay-bhogayata/product-service/server"
)

func main() {

	err := config.Init()
	if err != nil {
		log.Panicln(err)
	}

	err = server.Serve(config.AppCfg)
	if err != nil {
		log.Panicln(err)
	}
}
