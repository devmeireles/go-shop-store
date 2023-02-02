package services

import (
	database "github.com/devmeireles/go-shop-store/app/config"
	"github.com/devmeireles/go-shop-store/app/models"
)

func GetProducts() (*[]models.Product, error) {
	var err error
	products := []models.Product{}

	err = database.DB.Db.Model(&models.Product{}).
		Find(&products).Error

	if err != nil {
		return &[]models.Product{}, err
	}

	return &products, nil
}

func SaveProduct(product *models.Product) (*models.Product, error) {
	if err := database.DB.Db.Model(&models.Product{}).Create(&product).Error; err != nil {
		return &models.Product{}, err
	}

	return product, nil
}

func GetProductByID(id int) (*models.Product, error) {
	var err error
	product := models.Product{}

	err = database.DB.Db.Model(&models.Product{}).
		First(&product, id).Error

	if err != nil {
		return &models.Product{}, err
	}

	return &product, nil
}
