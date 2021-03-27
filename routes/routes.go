package routes

import (
	"project/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	e.POST("/users/store", controllers.CreateUserController)

	e.GET("/products", controllers.GetProductsController)
	e.POST("/products/store", controllers.CreateProductController)

	e.GET("/product_categories", controllers.GetProductCategoriesController)
	e.POST("/product_categories/store", controllers.CreateProductCategoryController)

	return e
}
