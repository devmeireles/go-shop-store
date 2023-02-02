package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `json:"name" gorm:"text; not null; default:null"`
	Email    string  `json:"email" gorm:"text; not null; default:null"`
	Password string  `json:"password" gorm:"text; not null; default:null"`
	Phone    string  `json:"phone" gorm:"text; not null; default:null"`
	Status   int     `json:"status" gorm:"int; not null; default: 1"`
	Address  Address `json:"address" gorm:"foreignKey:ID"`
}
