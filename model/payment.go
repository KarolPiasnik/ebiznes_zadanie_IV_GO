package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID     uint64 `json:"ID" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Amount uint64 `json:"amount"`
}
