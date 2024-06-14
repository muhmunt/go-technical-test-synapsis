package repository

import (
	"go-technical-test-synapsis/src/entity"
	"strings"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(product entity.Product) (entity.Product, error)
	FindAll(offset, limit int) ([]entity.Product, error)
	FindByID(productID int) (entity.Product, error)
	FindByCategory(category string) ([]entity.Product, error)
	DeleteProductByID(product int) (entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) Save(product entity.Product) (entity.Product, error) {
	err := r.db.Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) FindAll(offset, limit int) ([]entity.Product, error) {
	var products []entity.Product

	err := r.db.Offset(offset).Limit(limit).Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil

}

func (r *productRepository) FindByCategory(category string) ([]entity.Product, error) {
	var products []entity.Product

	err := r.db.Where("category = ?", strings.ToUpper((category))).Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil

}

func (r *productRepository) FindByID(productID int) (entity.Product, error) {
	var getProduct entity.Product

	err := r.db.Where("id = ?", productID).First(&getProduct).Error

	if err != nil {
		return getProduct, err
	}

	return getProduct, nil

}

func (r *productRepository) DeleteProductByID(product int) (entity.Product, error) {
	var deleteProduct entity.Product
	err := r.db.Where("id = ?", product).Delete(&deleteProduct).Error

	if err != nil {
		return deleteProduct, err
	}

	return deleteProduct, nil

}
