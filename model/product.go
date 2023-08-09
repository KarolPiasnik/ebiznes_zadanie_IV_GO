package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID         uint64   `json:"ID" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Code       string   `json:"code" gorm:"unique"`
	Price      uint     `json:"price"`
	Name       string   `json:"name"`
	CategoryID uint64   `json:"categoryID"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
