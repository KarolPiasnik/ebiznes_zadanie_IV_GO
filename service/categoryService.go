package service

import (
	"myapp/database"
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Categories
func SaveCategory(c echo.Context) error {
	db := database.DbManager()
	category := model.Category{}
	err := c.Bind(&category)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	result := db.Create(&category)

	if result.Error == nil {
		return c.String(http.StatusOK, "Utworzono nową kategorię")
	} else {
		return c.String(http.StatusInternalServerError, "Nie udało się utworzyć kategorii")
	}

}

func GetAllCategories(c echo.Context) error {
	db := database.DbManager()
	categories := []model.Category{}
	db.Preload("Products").Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func GetCategory(c echo.Context) error {
	db := database.DbManager()
	category := model.Category{}
	result := db.Preload("products").First(&category, c.Param("id"))
	if result.Error == nil {
		return c.JSON(http.StatusOK, category)
	} else {
		return c.String(http.StatusInternalServerError, "Nie udało się znaleźć kategorii")
	}
}

func UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	b := new(model.Category)
	db := database.DbManager()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	category := new(model.Category)

	if err := db.First(&category, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	category.Name = b.Name
	if err := db.Save(&category).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": category,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteCategory(c echo.Context) error {
	db := database.DbManager()
	category := model.Category{}
	result := db.Delete(&category, c.Param("id"))
	if result.Error == nil {
		return c.JSON(http.StatusOK, category)
	} else {
		return c.String(http.StatusInternalServerError, "Nie udało się usunąć kategorii")
	}
}
