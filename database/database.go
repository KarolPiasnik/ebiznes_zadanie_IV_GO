package database

import (
	"myapp/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	dsn := "host=localhost user=postgres password=postgres dbname=ebiznes port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Basket{})
	db.AutoMigrate(&model.Category{})
}

func DbManager() *gorm.DB {
	return db
}
