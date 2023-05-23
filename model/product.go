package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id    uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Code  string `json:"code"`
	Price uint   `json:"price"`
	Name  string `json:"name"`
}
