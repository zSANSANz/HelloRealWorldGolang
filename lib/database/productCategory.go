package database

import (
	"project/config"
	"project/models"
)

func GetProductCategories() (interface{}, error) {
	var productCategories []models.ProductCategory

	if err := config.DB.Find(&productCategories).Error; err != nil {
		return nil, err
	}
	return productCategories, nil
}

func GetProductCategory(id int) (interface{}, error) {
	var productCategory []models.ProductCategory

	if err := config.DB.First(&productCategory, id).Error; err != nil {
		return nil, err
	}
	return productCategory, nil
}

func CreateProductCategory(productCategory *models.ProductCategory) (interface{}, error) {
	if err := config.DB.Save(productCategory).Error; err != nil {
		return nil, err
	}
	return productCategory, nil
}
