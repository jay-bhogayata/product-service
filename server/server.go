package server

import (
	"log"
	"net/http"

	"github.com/jay-bhogayata/product-service/router"
	"github.com/labstack/echo/v4"
)

func Serve() error {

	e := echo.New()

	router.SetRoutes(e)

	s := http.Server{
		Addr:    ":8080",
		Handler: e,
	}

	log.Println("server is starting on port 8080")

	err := s.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
