package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jay-bhogayata/product-service/config"
	_ "github.com/jay-bhogayata/product-service/docs"
	"github.com/jay-bhogayata/product-service/router"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title					product-service API
// @version					0.0.1
// @description				This is a sample server for product-service.
// @termsOfService			https://github.com/jay-bhogayata/product-service
// @contact.name			Jay Bhogayata
// @contact.email			jaybhogayata53@gmail.com
// @host					localhost:8080
// @BasePath				/api/v1
// @schemes					http
func Serve(cfg config.Config) error {

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	router.SetRoutes(e)

	s := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: e,
	}

	slog.Info("server is starting on", "port", cfg.Server.Port)

	err := s.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
