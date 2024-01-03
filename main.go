package main

import (
	"log"

	"github.com/jay-bhogayata/product-service/server"
)

func main() {

	err := server.Serve()
	if err != nil {
		log.Panicln(err)
	}
}
