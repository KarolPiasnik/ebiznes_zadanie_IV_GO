package service

import (
	"myapp/database"
	"myapp/model"
)

// Users
func SaveUser(user model.User) {
	db := database.DbManager()
	db.Create(&user)
}

func UpdateUser(user model.User) {
	db := database.DbManager()
	db.Save(&user)
}

func GetUser(id string) model.User {
	db := database.DbManager()
	user := model.User{}
	user.ID = id
	result := db.First(&user)
	if result.Error == nil {
		return user
	} else {
		return model.User{}
	}
}

func GetUserByToken(token string) model.User {
	db := database.DbManager()
	user := model.User{}
	user.Token = token
	result := db.First(&user)
	if result.Error == nil {
		return user
	} else {
		return model.User{}
	}
}
