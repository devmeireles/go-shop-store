package models

import (
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	Quantity    int    `json:"quantity" gorm:"number; not null;"`
	Description string `json:"description" gorm:"text; not null; default:null"`
	Status      int    `json:"status" gorm:"int; not null; default: 1"`
}
