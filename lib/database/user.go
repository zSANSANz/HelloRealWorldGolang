package database

import (
	"project/config"
	"project/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user *models.User) (interface{}, error) {

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
