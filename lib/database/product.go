package database

import (
	"project/config"
	"project/models"
)

func GetProducts() (interface{}, error) {
	var products []models.Product

	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func CreateProduct(product *models.Product) (interface{}, error) {

	if err := config.DB.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
