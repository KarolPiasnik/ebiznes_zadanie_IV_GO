package database

import (
	"myapp/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	dsn := "host=" + os.Getenv("POSTGRES_HOST") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=" + os.Getenv("POSTGRES_PORT") + " sslmode=disable TimeZone=" + os.Getenv("POSTGRES_TIMEZONE")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	println(dsn)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	errProduct := db.AutoMigrate(&model.Product{})
	errBasket := db.AutoMigrate(&model.Basket{})
	errCategory := db.AutoMigrate(&model.Category{})
	errPayment := db.AutoMigrate(&model.Payment{})
	errUser := db.AutoMigrate(&model.User{})

	if errProduct != nil {
		panic("failed to migrate product table")
	}

	if errBasket != nil {
		panic("failed to migrate basket table")
	}

	if errCategory != nil {
		panic("failed to migrate category table")
	}

	if errPayment != nil {
		panic("failed to migrate payment table")
	}

	if errUser != nil {
		panic("failed to migrate user table")
	}
}

func DbManager() *gorm.DB {
	return db
}
