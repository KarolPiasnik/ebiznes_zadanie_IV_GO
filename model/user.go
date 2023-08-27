package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              string `json:"ID" sql:"unique" gorm:"primary_key"`
	Username        string `json:"username"`
	Token           string `json:"token"`
	TokenExpiration string `json:"tokenExpiration"`
}
