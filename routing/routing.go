package routing

import (
	"myapp/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/categories", service.SaveCategory)
	e.GET("/categories", service.GetAllCategories)
	e.GET("/categories/:id", service.GetCategory)
	e.PUT("/categories/:id", service.UpdateCategory)
	e.DELETE("/categories/:id", service.DeleteCategory)

	e.POST("/products", service.SaveProduct)
	e.GET("/products", service.GetAllProducts)
	e.GET("/products/:id", service.GetProduct)
	e.PUT("/products/:id", service.UpdateProduct)
	e.DELETE("/products/:id", service.DeleteProduct)

	e.POST("/baskets", service.SaveBasket)
	e.GET("/baskets", service.GetAllBaskets)
	e.GET("/baskets/:id", service.GetBasket)
	e.PUT("/baskets/:id", service.UpdateBasket)
	e.DELETE("/baskets/:id", service.DeleteBasket)
	return e
}
