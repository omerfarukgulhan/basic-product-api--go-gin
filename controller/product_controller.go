package controller

import (
	"example.com/product-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/products", productController.GetAll)
	e.GET("/api/products/:id", productController.GetById)
	e.POST("/api/products", productController.Add)
	e.PUT("/api/products/:id", productController.UpdatePrice)
	e.DELETE("/api/products/:id", productController.Delete)
}

func (productController *ProductController) GetAll(c echo.Context) error {
	store := c.QueryParam("store")

	if len(store) == 0 {
		return c.JSON(http.StatusOK, productController.productService.GetAll())

	}

	return c.JSON(http.StatusOK, productController.productService.GetAllByStore(store))
}

func (productController *ProductController) GetById(c echo.Context) error {
	//return c.JSON(http.StatusOK, productController.productService.GetById(c.Request().Header.Get(i)))
	return nil
}

func (productController *ProductController) UpdatePrice(c echo.Context) error {
	return nil
}

func (productController *ProductController) Add(c echo.Context) error {
	return nil
}

func (productController *ProductController) Delete(c echo.Context) error {
	return nil
}
