package service

import (
	"fmt"
	"myapp/database"
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Payments
func SavePayment(c echo.Context) error {
	db := database.DbManager()
	payment := model.Payment{}
	err := c.Bind(&payment)
	if err != nil {
		fmt.Print(err)
		return c.String(http.StatusBadRequest, "bad request")
	}
	result := db.Create(&payment)

	if result.Error == nil {
		return c.String(http.StatusOK, "Utworzono nową płatność")
	} else {
		return c.String(http.StatusInternalServerError, "Nie udało się utworzyć płatności")
	}

}

func GetAllPayments(c echo.Context) error {
	db := database.DbManager()
	payments := []model.Payment{}
	db.Find(&payments)
	return c.JSON(http.StatusOK, payments)
}

func GetPayment(c echo.Context) error {
	db := database.DbManager()
	payment := model.Payment{}
	result := db.First(&payment, c.Param("id"))
	if result.Error == nil {
		return c.JSON(http.StatusOK, payment)
	} else {
		return c.String(http.StatusInternalServerError, "Nie udało się znaleźć płatności")
	}
}

func UpdatePayment(c echo.Context) error {

	return c.String(http.StatusInternalServerError, "Payments can't be updated")

}

func DeletePayment(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "Payments can't be deleted")
}
