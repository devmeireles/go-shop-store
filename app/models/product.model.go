package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string     `json:"title" gorm:"text; not null;" validate:"required"`
	Description string     `json:"description" gorm:"text; not null" validate:"required"`
	Price       float64    `json:"price" gorm:"float; not null" validate:"required"`
	Status      int        `json:"status" gorm:"int; not null; default: 1"`
	Tags        []Tag      `json:"tags" gorm:"many2many:product_tags"`
	Categories  []Category `json:"categories" gorm:"many2many:product_categories"`
	Inventory   Inventory  `json:"inventory" gorm:"foreignKey:ID"`
}
