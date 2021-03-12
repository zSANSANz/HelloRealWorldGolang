package database

import (
	"project/config"
	"project/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func CreateUser() (interface{}, error) {
	user := models.User{}

	if e := config.DB.Save(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}
