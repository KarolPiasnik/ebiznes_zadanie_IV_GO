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

	categoryBasePath := "/categories"
	categoryIdParametrizedPath := categoryBasePath + "/:id"

	e.POST(categoryBasePath, service.SaveCategory)
	e.GET(categoryBasePath, service.GetAllCategories)
	e.GET(categoryIdParametrizedPath, service.GetCategory)
	e.PUT(categoryIdParametrizedPath, service.UpdateCategory)
	e.DELETE(categoryIdParametrizedPath, service.DeleteCategory)

	productBasePath := "/products"
	productIdParametrizedPath := categoryBasePath + "/:id"

	e.POST(productBasePath, service.SaveProduct)
	e.GET(productBasePath, service.GetAllProducts)
	e.GET(productIdParametrizedPath, service.GetProduct)
	e.PUT(productIdParametrizedPath, service.UpdateProduct)
	e.DELETE(productIdParametrizedPath, service.DeleteProduct)

	basketBasePath := "/baskets"
	basketIdParametrizedPath := categoryBasePath + "/:id"

	e.POST(basketBasePath, service.SaveBasket)
	e.GET(basketBasePath, service.GetAllBaskets)
	e.GET(basketIdParametrizedPath, service.GetBasket)
	e.PUT(basketIdParametrizedPath, service.UpdateBasket)
	e.DELETE(basketIdParametrizedPath, service.DeleteBasket)

	paymentBasePath := "/payments"
	paymentIdParametrizedPath := categoryBasePath + "/:id"

	e.POST(paymentBasePath, service.SavePayment)
	e.GET(paymentBasePath, service.GetAllPayments)
	e.GET(paymentIdParametrizedPath, service.GetPayment)
	e.PUT(paymentIdParametrizedPath, service.UpdatePayment)
	e.DELETE(paymentIdParametrizedPath, service.DeletePayment)
	return e
}
