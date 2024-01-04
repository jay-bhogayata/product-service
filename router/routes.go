package router

import (
	"github.com/jay-bhogayata/product-service/config"
	"github.com/jay-bhogayata/product-service/controllers"
	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) {

	api := e.Group("/api/v1")

	api.GET("/health", controllers.Health)

	cc := controllers.NewCategoryControllers(config.AppCfg)

	api.POST("/category", cc.CreateCategory)
}
