package main

import (
	"context"
	"example.com/product-api/common/app"
	"example.com/product-api/common/postgresql"
	"example.com/product-api/controller"
	"example.com/product-api/persistence"
	"example.com/product-api/service"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	e := echo.New()

	configurationManager := app.NewConfigurationManager()
	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	productRepository := persistence.NewProductRepository(dbPool)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	productController.RegisterRoutes(e)

	err := e.Start("localhost:8080")

	if err != nil {
		return
	}
}
