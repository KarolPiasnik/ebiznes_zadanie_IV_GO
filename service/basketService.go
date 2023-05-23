package service

import (
	"myapp/database"
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Baskets
func SaveBasket(c echo.Context) error {
	db := database.DbManager()
	basket := model.Basket{}
	result := db.Create(&basket)

	if result.Error == nil {
		return c.String(http.StatusOK, "Utworzono nowy koszyk")
	} else {
		return c.String(http.StatusInternalServerError, "Nie udało się utworzyć koszyka")
	}

}

func GetAllBaskets(c echo.Context) error {
	db := database.DbManager()
	baskets := []model.Basket{}
	db.Find(&baskets)
	return c.JSON(http.StatusOK, baskets)
}

func GetBasket(c echo.Context) error {
	db := database.DbManager()
	basket := model.Basket{}
	result := db.First(&basket, c.Param("id"))
	if result.Error == nil {
		return c.JSON(http.StatusOK, basket)
	} else {
		return c.String(http.StatusInternalServerError, "Nie udało się znaleźć koszyka")
	}
}

func UpdateBasket(c echo.Context) error {
	db := database.DbManager()
	basket := model.Basket{}
	result := db.First(&basket, c.Param("id"))
	if result.Error == nil {
		return c.JSON(http.StatusOK, basket)
	} else {
		return c.String(http.StatusInternalServerError, "Nie udało się edytować koszyka")
	}
}

func DeleteBasket(c echo.Context) error {
	db := database.DbManager()
	basket := model.Basket{}
	result := db.Delete(&basket, c.Param("id"))
	if result.Error == nil {
		return c.JSON(http.StatusOK, basket)
	} else {
		return c.String(http.StatusInternalServerError, "Nie udało się usunąć koszyka")
	}
}
