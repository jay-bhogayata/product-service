package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthResponse struct {
	Message string `json:"message" example:"ok"`
}

// @Summary					Health check endpoint
// @Description				Health check endpoint
// @Tags					health
// @Produce					json
// @Success					200 {object} healthResponse
// @Router					/health [get]
func Health(c echo.Context) error {
	res := healthResponse{Message: "ok"}
	return c.JSON(http.StatusOK, res)
}
