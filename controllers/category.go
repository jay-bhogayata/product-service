package controllers

import (
	"net/http"
	"strconv"

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
// @Failure					500 {object} errorResponse
// @Security				BearerToken
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

// @Summary					Get available categories
// @Description  			Get available categories
// @Tags					category
// @Accept					n/a
// @Produce					json
// @Success					200 {array} models.Category
// @Failure					400 {object} errorResponse
// @Failure					500 {object} errorResponse
// @Security				BearerToken
// @Router					/category [get]
func (c *CategoryControllers) GetAvailableCategories(ctx echo.Context) error {

	categories, err := models.GetAvailableCategories(c.db)
	if err != nil {
		res := &errorResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, res)
	}

	return ctx.JSON(http.StatusOK, categories)
}

// @Summary					Get category by id
// @Description  			Get category by id
// @Tags					category
// @Accept					n/a
// @Produce					json
// @Param					id path int true "category id"
// @Success					200 {object} models.Category
// @Failure					400 {object} errorResponse
// @Failure					500 {object} errorResponse
// @Security				BearerToken
// @Router					/category/{id} [get]
func (c *CategoryControllers) GetCategoryById(ctx echo.Context) error {

	id := ctx.Param("id")
	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		res := &errorResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, res)
	}

	category, err := models.GetCategoryById(c.db, categoryID)
	if err != nil {
		res := &errorResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, res)
	}

	return ctx.JSON(http.StatusOK, category)
}

// @Summary					Edit category
// @Description  			Edit category
// @Tags					category
// @Accept					json
// @Produce					json
// @Param					id path int true "category id"
// @Param					category body category true "category"
// @Success					200 {object} successResponse
// @Failure					400 {object} errorResponse
// @Failure					500 {object} errorResponse
// @Security				BearerToken
// @Router					/category/{id} [put]
func (c *CategoryControllers) EditCategory(ctx echo.Context) error {

	id := ctx.Param("id")
	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		res := &errorResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, res)
	}

	var cat category

	err = ctx.Bind(&cat)
	if err != nil {
		res := &errorResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, res)
	}

	err = models.EditCategory(c.db, categoryID, cat.Name)
	if err != nil {
		res := &errorResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, res)
	}

	res := &successResponse{
		Success: true,
		Message: "category updated successfully",
	}
	return ctx.JSON(http.StatusOK, res)
}

// @Summary					Delete category
// @Description  			Delete category
// @Tags					category
// @Accept					n/a
// @Produce					json
// @Param					id path int true "category id"
// @Success					200 {object} successResponse
// @Failure					400 {object} errorResponse
// @Failure					500 {object} errorResponse
// @Security				BearerToken
// @Router					/category/{id} [delete]
func (c *CategoryControllers) DeleteCategoryById(ctx echo.Context) error {

	id := ctx.Param("id")
	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		res := &errorResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, res)
	}

	err = models.DeleteCategoryById(c.db, categoryID)
	if err != nil {
		res := &errorResponse{
			Success: false,
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, res)
	}

	res := &successResponse{
		Success: true,
		Message: "category deleted successfully",
	}
	return ctx.JSON(http.StatusOK, res)
}
