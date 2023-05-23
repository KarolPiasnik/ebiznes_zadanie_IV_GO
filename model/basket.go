package model

import (
	"gorm.io/gorm"
)

type Basket struct {
	gorm.Model
	Id uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
}
