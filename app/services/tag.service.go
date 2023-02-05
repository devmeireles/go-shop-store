package services

import (
	database "github.com/devmeireles/go-shop-store/app/config"
	"github.com/devmeireles/go-shop-store/app/models"
)

func SaveTag(tag *models.Tag) (*models.Tag, error) {
	if err := database.DB.Db.Model(&models.Tag{}).Create(&tag).Error; err != nil {
		return &models.Tag{}, err
	}

	return tag, nil
}

func GetTags() (*[]models.Tag, error) {
	var err error
	tags := []models.Tag{}

	err = database.DB.Db.Model(&models.Tag{}).
		Find(&tags).Error

	if err != nil {
		return &[]models.Tag{}, err
	}

	return &tags, nil
}

func GetTagByID(id int) (*models.Tag, error) {
	var err error
	tag := models.Tag{}

	err = database.DB.Db.Model(&models.Tag{}).
		First(&tag, id).Error

	if err != nil {
		return &models.Tag{}, err
	}

	return &tag, nil
}

func UpdateTag(tag models.Tag, id int) error {
	if err := database.DB.Db.Model(&models.Tag{}).
		Where("id = ?", id).
		Updates(&tag).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTag(tag *models.Tag, id int) error {
	if err := database.DB.Db.Delete(tag, id).Error; err != nil {
		return err
	}

	return nil
}
