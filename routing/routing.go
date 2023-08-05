package routing

import (
	"myapp/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
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

	e.POST("/payments", service.SavePayment)
	e.GET("/payments", service.GetAllPayments)
	e.GET("/payments/:id", service.GetPayment)
	e.PUT("/payments/:id", service.UpdatePayment)
	e.DELETE("/payments/:id", service.DeletePayment)
	return e
}
