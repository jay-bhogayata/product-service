package controllers

import (
	"net/http"

	"github.com/jay-bhogayata/product-service/config"
	"github.com/jay-bhogayata/product-service/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryControllers struct {
	db *gorm.DB
}

func NewCategoryControllers(cfg config.Config) *CategoryControllers {
	return &CategoryControllers{
		db: cfg.Db,
	}
}

type category struct {
	Name string `json:"name" example:"electronics"`
}

type successResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"category created successfully"`
}

type errorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"category already exists"`
}

// @Summary					Create category
// @Description  			Create category
// @Tags					category
// @Accept					json
// @Produce					json
// @Param					category body category true "category"
// @Success					201 {object} successResponse
// @Failure					400 {object} errorResponse
// @Router					/category [post]
func (c *CategoryControllers) CreateCategory(ctx echo.Context) error {

	var cat category

	err := ctx.Bind(&cat)
	if err != nil {
		res := &errorResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, res)
	}

	err = models.CreateCategory(c.db, cat.Name)
	if err != nil {
		res := &errorResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, res)
	}

	res := &successResponse{
		Success: true,
		Message: "category created successfully",
	}
	return ctx.JSON(http.StatusCreated, res)
}
