package controller

import (
	"example.com/product-api/controller/request"
	"example.com/product-api/controller/response"
	"example.com/product-api/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
		return c.JSON(http.StatusOK, response.ToProductResponseList(productController.productService.GetAll()))
	}

	return c.JSON(http.StatusOK, response.ToProductResponseList(productController.productService.GetAllByStore(store)))
}

func (productController *ProductController) GetById(c echo.Context) error {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: "enter valid id"})
	}

	product, err := productController.productService.GetById(int64(productId))

	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return c.JSON(http.StatusOK, response.ToProductResponse(product))
}

func (productController *ProductController) Add(c echo.Context) error {
	var addProductRequest request.AddProductRequest
	err := c.Bind(&addProductRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	err = productController.productService.Add(addProductRequest.ToModel())

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return c.JSON(http.StatusCreated, addProductRequest.ToModel())

}

func (productController *ProductController) UpdatePrice(c echo.Context) error {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: "enter valid id"})
	}

	newPrice := c.QueryParam("newPrice")
	if len(newPrice) == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: "Parameter newPrice is required!",
		})
	}

	convertedPrice, err := strconv.ParseFloat(newPrice, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: "NewPrice Format Disrupted!",
		})
	}

	err = productController.productService.UpdatePrice(int64(productId), float32(convertedPrice))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return c.NoContent(http.StatusOK)
}

func (productController *ProductController) Delete(c echo.Context) error {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: "enter valid id"})
	}

	err = productController.productService.DeleteById(int64(productId))

	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return c.NoContent(http.StatusOK)
}
