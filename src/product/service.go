package product

import (
	"go-technical-test-synapsis/src/entity"
	"go-technical-test-synapsis/src/helper"
	"go-technical-test-synapsis/src/repository"
)

type Service interface {
	FindProducts(offset, limit int) ([]entity.Product, error)
	StoreProduct(request ProductRequest) (entity.Product, error)
	FindProductByID(productID ProductIDRequest) (entity.Product, error)
	FindProductByCategory(category ProductCategoryRequest) ([]entity.Product, error)
	UpdateProduct(productID ProductIDRequest, request ProductRequest) (entity.Product, error)
	DeleteProductByID(productID ProductIDRequest) (entity.Product, error)
}

type service struct {
	repository repository.ProductRepository
}

func NewService(repository repository.ProductRepository) *service {
	return &service{repository}
}

func (s *service) FindProducts(offset, limit int) ([]entity.Product, error) {
	products, err := s.repository.FindAll(offset, limit)

	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) FindProductByCategory(category ProductCategoryRequest) ([]entity.Product, error) {
	products, err := s.repository.FindByCategory(category.Category)

	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) StoreProduct(request ProductRequest) (entity.Product, error) {
	product := entity.Product{}
	product.ProductName = request.Name
	product.Price = request.Price
	product.Description = request.Description
	product.Category = request.Category
	product.CreatedAt = helper.TimeNow()

	newProduct, err := s.repository.Save(product)

	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) FindProductByID(productID ProductIDRequest) (entity.Product, error) {
	product, err := s.repository.FindByID(productID.ID)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) UpdateProduct(productID ProductIDRequest, request ProductRequest) (entity.Product, error) {

	product, err := s.repository.FindByID(productID.ID)

	if err != nil {
		return product, err
	}

	product.ProductName = request.Name
	product.Price = request.Price
	product.Description = request.Description
	product.Category = request.Category
	product.UpdatedAt = helper.TimeNow()

	updateProduct, err := s.repository.Save(product)
	if err != nil {
		return updateProduct, err
	}

	return updateProduct, nil
}

func (s *service) DeleteProductByID(productID ProductIDRequest) (entity.Product, error) {
	product, err := s.repository.DeleteProductByID(productID.ID)

	if err != nil {
		return product, err
	}

	return product, nil
}
