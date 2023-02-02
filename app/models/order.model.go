package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Cart    Cart    `json:"cart" gorm:"foreignKey:ID"`
	Payment Payment `json:"payment" gorm:"foreignKey:ID"`
	Total   float32 `json:"total"`
	Status  int     `json:"status" gorm:"int; not null; default: 1"`
}
