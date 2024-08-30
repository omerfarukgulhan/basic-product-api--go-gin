package service

import (
	"errors"
	"example.com/product-api/domain"
	"example.com/product-api/persistence"
	"fmt"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

func (fakeProductRepository *FakeProductRepository) GetAll() []domain.Product {
	return fakeProductRepository.products
}

func (fakeProductRepository *FakeProductRepository) GetById(productId int64) (domain.Product, error) {
	for _, product := range fakeProductRepository.products {
		if product.Id == productId {
			return product, nil
		}
	}

	return domain.Product{}, errors.New(fmt.Sprintf("Product not found with id %d", productId))
}

func (fakeProductRepository *FakeProductRepository) GetAllByStore(storeName string) []domain.Product {
	products := make([]domain.Product, 0)

	for _, product := range fakeProductRepository.products {
		if product.Store == storeName {
			products = append(products, product)
		}
	}

	return products
}

func (fakeProductRepository *FakeProductRepository) Add(product domain.Product) error {
	fakeProductRepository.products = append(fakeProductRepository.products, domain.Product{
		Id:       int64(len(fakeProductRepository.products)) + 1,
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})

	return nil
}

func (fakeProductRepository *FakeProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	for i, product := range fakeProductRepository.products {
		if product.Id == productId {
			fakeProductRepository.products[i].Price = newPrice
			return nil
		}
	}

	return errors.New("product not found")
}

func (fakeProductRepository *FakeProductRepository) DeleteById(productId int64) error {
	for i, product := range fakeProductRepository.products {
		if product.Id == productId {
			// Remove product by index
			fakeProductRepository.products = append(fakeProductRepository.products[:i], fakeProductRepository.products[i+1:]...)
			return nil
		}
	}

	return errors.New("product not found")
}
