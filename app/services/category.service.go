package services

import (
	database "github.com/devmeireles/go-shop-store/app/config"
	"github.com/devmeireles/go-shop-store/app/models"
)

func SaveCategory(category *models.Category) (*models.Category, error) {
	if err := database.DB.Db.Model(&models.Category{}).Create(&category).Error; err != nil {
		return &models.Category{}, err
	}

	return category, nil
}

func GetCategories() (*[]models.Category, error) {
	var err error
	categories := []models.Category{}

	err = database.DB.Db.Model(&models.Category{}).
		Find(&categories).Error

	if err != nil {
		return &[]models.Category{}, err
	}

	return &categories, nil
}


func GetCategoryByID(id int) (*models.Category, error) {
	var err error
	category := models.Category{}

	err = database.DB.Db.Model(&models.Category{}).
		First(&category, id).Error

	if err != nil {
		return &models.Category{}, err
	}

	return &category, nil
}

func UpdateCategory(category models.Category, id int) error {
	if err := database.DB.Db.Model(&models.Category{}).
		Where("id = ?", id).
		Updates(&category).Error; err != nil {
		return err
	}

	return nil
}

func DeleteCategory(category *models.Category, id int) error {
	if err := database.DB.Db.Delete(category, id).Error; err != nil {
		return err
	}

	return nil
}
