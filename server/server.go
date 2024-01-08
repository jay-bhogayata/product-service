package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/jay-bhogayata/product-service/config"
	_ "github.com/jay-bhogayata/product-service/docs"
	"github.com/jay-bhogayata/product-service/router"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
// @schemes					http https
// @securityDefinitions.apiKey	BearerToken
// @in							header
// @name						Authorization
func Serve(cfg config.Config) error {

	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		Skipper:     pathSkipper,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				slog.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				slog.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))
	e.Use(middleware.CORS())
	e.Use(echojwt.WithConfig(
		echojwt.Config{
			Skipper:    pathSkipper,
			SigningKey: []byte(cfg.Server.JWTSecret),
		},
	))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/docs/*", echoSwagger.EchoWrapHandler(echoSwagger.InstanceName("private")))

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

func pathSkipper(c echo.Context) bool {
	paths := []string{"/swagger", "/docs", "/api/v1/health"}
	for _, path := range paths {
		if strings.Contains(c.Path(), path) {
			return true
		}
	}
	return false
}
