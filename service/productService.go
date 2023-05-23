package service

import (
	"myapp/database"
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Products
func SaveProduct(c echo.Context) error {
	db := database.DbManager()
	product := model.Product{}
	err := c.Bind(&product)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	result := db.Create(&product)

	if result.Error == nil {
		return c.String(http.StatusOK, "Utworzono nowy produkt")
	} else {
		return c.String(http.StatusInternalServerError, "Nie udało się utworzyć produktu")
	}
}

func GetAllProducts(c echo.Context) error {
	db := database.DbManager()
	products := []model.Product{}
	db.Find(&products)
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	db := database.DbManager()
	product := model.Product{}
	result := db.First(&product, c.Param("id"))
	if result.Error == nil {
		return c.JSON(http.StatusOK, product)
	} else {
		return c.String(http.StatusInternalServerError, "Nie udało się znaleźć produktu")
	}
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	b := new(model.Product)
	db := database.DbManager()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	product := new(model.Product)

	if err := db.First(&product, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	product.Name = b.Name
	product.Code = b.Code
	product.Price = b.Price

	if err := db.Save(&product).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": product,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteProduct(c echo.Context) error {
	db := database.DbManager()
	product := model.Product{}
	result := db.Delete(&product, c.Param("id"))
	if result.Error == nil {
		return c.JSON(http.StatusOK, product)
	} else {
		return c.String(http.StatusInternalServerError, "Nie udało się usunąć produktu")
	}
}
