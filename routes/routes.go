package routes

import (
	"project/constants"
	"project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)

	// JWT Group
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.DELETE("/users/:id", controllers.DeleteUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)
	
	e.GET("/products", controllers.GetProductsController)
	e.POST("/products/store", controllers.CreateProductController)

	e.GET("/product_categories", controllers.GetProductCategoriesController)
	e.POST("/product_categories/store", controllers.CreateProductCategoryController)
	
	return e
}
