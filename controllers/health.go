package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthResponse struct {
	Message string `json:"message"`
}

func Health(c echo.Context) error {
	res := healthResponse{Message: "ok"}
	return c.JSON(http.StatusOK, res)
}
