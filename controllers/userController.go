package controllers

import (
	"net/http"
	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	users, e := database.CreateUser(&user)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    users,
	})
}
