package controllers

import (
	"net/http"
	"strconv"

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
		"status":            "success",
		"productCategories": productsCategories,
	})
}

func GetProductCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get a product category, product category with ID " + c.Param("id") + " is not found",
		})
	}

	user, getErr := database.GetProductCategory(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func DeleteProductCategoryController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success create new productCategory",
		"productCategory": "test delete",
	})
}

func UpdateProductCategoryController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success create new productCategory",
		"productCategory": "test update",
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
		"message":         "success create new productCategory",
		"productCategory": productCategories,
	})
}
