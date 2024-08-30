package persistence

import (
	"context"
	"errors"
	"example.com/product-api/domain"
	"example.com/product-api/persistence/common"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAll() []domain.Product
	GetById(productId int64) (domain.Product, error)
	GetAllByStore(storeName string) []domain.Product
	Add(product domain.Product) error
	UpdatePrice(productId int64, newPrice float32) error
	DeleteById(productId int64) error
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{dbPool: dbPool}
}

func (productRepository *ProductRepository) GetAll() []domain.Product {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "Select * from products")

	if err != nil {
		log.Error("Couldn't get products")
		return []domain.Product{}
	}

	return extractProductsFromRows(productRows)
}

func (productRepository *ProductRepository) GetById(productId int64) (domain.Product, error) {
	ctx := context.Background()
	getByIdSql := `Select * from products where id = $1`
	queryRow := productRepository.dbPool.QueryRow(ctx, getByIdSql, productId)

	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	scanErr := queryRow.Scan(&id, &name, &price, &discount, &store)

	if scanErr != nil && scanErr.Error() == common.NOT_FOUND {
		return domain.Product{}, errors.New(fmt.Sprintf("Product not found with id %d", productId))
	}

	if scanErr != nil {
		return domain.Product{}, errors.New(fmt.Sprintf("Error while getting product with id %d", productId))
	}

	return domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}, nil
}

func (productRepository *ProductRepository) GetAllByStore(storeName string) []domain.Product {
	ctx := context.Background()

	getProductsByStoreNameSql := `Select * from products where store = $1`

	productRows, err := productRepository.dbPool.Query(ctx, getProductsByStoreNameSql, storeName)

	if err != nil {
		log.Error("Error while getting all products %v", err)
		return []domain.Product{}
	}

	return extractProductsFromRows(productRows)
}

func (productRepository *ProductRepository) Add(product domain.Product) error {
	ctx := context.Background()
	insert_sql := `INSERT INTO products(name, price, discount, store) VALUES($1, $2, $3, $4)`
	newProduct, err := productRepository.dbPool.Exec(ctx, insert_sql, product.Name, product.Price, product.Discount, product.Store)

	if err != nil {
		log.Error("Error while inserting product %v", err)
		return err
	}

	log.Info(fmt.Printf("Product added with %v", newProduct))

	return nil
}

func (productRepository *ProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	ctx := context.Background()
	updateSql := `Update products set price = $1 where id = $2`
	_, err := productRepository.dbPool.Exec(ctx, updateSql, newPrice, productId)

	if err != nil {
		return errors.New(fmt.Sprintf("Error while updating product with id: %d", productId))
	}

	log.Info("Product %d price updated with new price %v", productId, newPrice)

	return nil
}

func (productRepository *ProductRepository) DeleteById(productId int64) error {
	ctx := context.Background()
	_, getErr := productRepository.GetById(productId)

	if getErr != nil {
		return errors.New("product not found")
	}

	deleteSql := `Delete from products where id = $1`
	_, err := productRepository.dbPool.Exec(ctx, deleteSql, productId)

	if err != nil {
		return errors.New(fmt.Sprintf("Error while deleting product with id %d", productId))
	}

	log.Info("Product deleted")

	return nil
}

func extractProductsFromRows(productRows pgx.Rows) []domain.Product {
	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for productRows.Next() {
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})
	}

	return products
}
