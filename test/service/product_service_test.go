package service

import (
	"example.com/product-api/domain"
	"example.com/product-api/service"
	"example.com/product-api/service/dto"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var productService service.IProductService

func TestMain(m *testing.M) {
	initialProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Iron",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Washing Machine",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Floor Lamp",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "Decoration Palace",
		},
	}

	fakeProductRepository := NewFakeProductRepository(initialProducts)
	productService = service.NewProductService(fakeProductRepository)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_ShouldGetAllProducts(t *testing.T) {
	t.Run("ShouldGetAllProducts", func(t *testing.T) {
		actualProducts := productService.GetAll()
		assert.Equal(t, 4, len(actualProducts))
	})
}

func Test_WhenNoValidationErrorOccurred_ShouldAddProduct(t *testing.T) {
	t.Run("WhenNoValidationErrorOccurred_ShouldAddProduct", func(t *testing.T) {
		productService.Add(dto.ProductCreate{
			Name:     "Telephone",
			Price:    20000.0,
			Discount: 10.0,
			Store:    "Samsung",
		})
		actualProducts := productService.GetAll()
		assert.Equal(t, 5, len(actualProducts))
		assert.Equal(t, domain.Product{
			Id:       5,
			Name:     "Telephone",
			Price:    20000.0,
			Discount: 10.0,
			Store:    "Samsung",
		}, actualProducts[len(actualProducts)-1])
	})
}

func Test_WhenDiscountIsHigherThan70_ShouldNotAddProduct(t *testing.T) {
	t.Run("WhenDiscountIsHigherThan70_ShouldNotAddProduct", func(t *testing.T) {
		err := productService.Add(dto.ProductCreate{
			Name:     "Telephone",
			Price:    20000.0,
			Discount: 80.0,
			Store:    "Samsung",
		})
		actualProducts := productService.GetAll()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, "Discount can not be greater than 70", err.Error())
	})
}
