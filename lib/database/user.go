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

func GetUser() (interface{}, error) {
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

func DeleteUser(user *models.User) (interface{}, error) {

	if err := config.DB.Delete(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *models.User) (interface{}, error) {
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
