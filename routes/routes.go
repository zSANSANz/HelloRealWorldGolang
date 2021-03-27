package routes

import (
	"project/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users/store", controllers.CreateUserController)
	e.DELETE("/users/:id/delete", controllers.DeleteUserController)
	e.PUT("/users/update", controllers.UpdateUserController)

	e.GET("/products", controllers.GetProductsController)
	e.POST("/products/store", controllers.CreateProductController)

	e.GET("/product_categories", controllers.GetProductCategoriesController)
	// e.POST("/product_categories/store", controllers.CreateProductCategoryController)
	// e.DELETE("/product_categories/delete", controllers.GetProductCategoriesController)
	
	return e
}
