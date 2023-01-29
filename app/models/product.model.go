package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string  `json:"title" gorm:"text; not null; default:null"`
	Description string  `json:"description" gorm:"text; not null; default:null"`
	Price       float32 `json:"price" gorm:"float; not null"`
	Status      int     `json:"status" gorm:"int; not null; default: 1"`
}
