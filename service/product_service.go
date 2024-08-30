package service

import (
	"errors"
	"example.com/product-api/domain"
	"example.com/product-api/persistence"
	"example.com/product-api/service/dto"
)

type IProductService interface {
	GetAll() []domain.Product
	GetById(productId int64) (domain.Product, error)
	GetAllByStore(storeName string) []domain.Product
	Add(productCreate dto.ProductCreate) error
	UpdatePrice(productId int64, newPrice float32) error
	DeleteById(productId int64) error
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (productService *ProductService) GetAll() []domain.Product {
	return productService.productRepository.GetAll()
}

func (productService *ProductService) GetById(productId int64) (domain.Product, error) {
	return productService.productRepository.GetById(productId)
}

func (productService *ProductService) GetAllByStore(storeName string) []domain.Product {
	return productService.productRepository.GetAllByStore(storeName)
}

func (productService *ProductService) Add(productCreate dto.ProductCreate) error {
	validateErr := validateProductCreate(productCreate)

	if validateErr != nil {
		return validateErr
	}

	return productService.productRepository.Add(domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	})
}

func (productService *ProductService) UpdatePrice(productId int64, newPrice float32) error {
	return productService.productRepository.UpdatePrice(productId, newPrice)
}

func (productService *ProductService) DeleteById(productId int64) error {
	return productService.productRepository.DeleteById(productId)
}

func validateProductCreate(productCreate dto.ProductCreate) error {
	if productCreate.Discount > 70.0 {
		return errors.New("Discount can not be greater than 70")
	}
	return nil
}
