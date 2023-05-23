package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id   uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name uint   `json:"name"`
}
