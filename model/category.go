package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID       uint64 `json:"ID" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name     string `json:"name"`
	Products []Product
}
