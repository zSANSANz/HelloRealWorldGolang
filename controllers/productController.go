package controllers

import (
	"net/http"
	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetProductsController(c echo.Context) error {
	products, e := database.GetProducts()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"products":  products,
	})
}

func CreateProductController(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)
	products, e := database.CreateProduct(&product)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new product",
		"product":    products,
	})
}
