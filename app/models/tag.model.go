package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name       string  `json:"name" gorm:"text; not null; default:null"`
	Description string  `json:"description" gorm:"text; not null; default:null"`
	Status      int     `json:"status" gorm:"int; not null; default: 1"`
}