package controllers

import (
	"net/http"
	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetProductCategoriesController(c echo.Context) error {
	productsCategories, e := database.GetProductCategories()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"productCategories": productsCategories,
	})
}

func CreateProductCategoryController(c echo.Context) error {
	productCategory := models.ProductCategory{}
	c.Bind(&productCategory)
	productCategories, e := database.CreateProductCategory(&productCategory)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new productCategory",
		"productCategory": productCategories,
	})
}