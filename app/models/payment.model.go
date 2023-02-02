package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Status      int     `json:"status" gorm:"int; not null; default: 1"`
}