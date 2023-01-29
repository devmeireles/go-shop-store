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
