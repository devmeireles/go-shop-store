package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	AddressLine1 string `json:"addres_line_1" gorm:"text; not null;"`
	AddressLine2 string `json:"addres_line_2" gorm:"text; not null; default:null"`
	Postcode     string `json:"postcode" gorm:"text; not null; default:null"`
	City         string `json:"city" gorm:"text; not null; default:null"`
	Country      string `json:"country" gorm:"text; not null; default:null"`
	Note         string `json:"note" gorm:"text; default:null"`
	Status       int    `json:"status" gorm:"int; not null; default: 1"`
}
