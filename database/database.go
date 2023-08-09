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
	errProduct := db.AutoMigrate(&model.Product{})
	errBasket := db.AutoMigrate(&model.Basket{})
	errCategory := db.AutoMigrate(&model.Category{})
	errPayment := db.AutoMigrate(&model.Payment{})

	if errProduct != nil {
		panic("failed to migrate product table")
	}

	if errBasket != nil {
		panic("failed to migrate category table")
	}

	if errCategory != nil {
		panic("failed to migrate basket table")
	}

	if errPayment != nil {
		panic("failed to migrate payment table")
	}
}

func DbManager() *gorm.DB {
	return db
}
