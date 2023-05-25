package model

import (
	"gorm.io/gorm"
)

type Basket struct {
	gorm.Model
	ID uint64 `json:"ID" sql:"AUTO_INCREMENT" gorm:"primary_key"`
}
