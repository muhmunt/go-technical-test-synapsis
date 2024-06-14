package repository

import (
	"go-technical-test-synapsis/src/entity"

	"gorm.io/gorm"
)

type CartRepository interface {
	StoreToCart(cart entity.Cart) (entity.Cart, error)
	FindByID(userID int) (entity.Cart, error)
	DeleteCartByID(userID int) (entity.Cart, error)
}

type cartRepository struct {
	db *gorm.DB
}

func NewCart(db *gorm.DB) *cartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) StoreToCart(cart entity.Cart) (entity.Cart, error) {
	err := r.db.Save(&cart).Error

	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (r *cartRepository) FindByID(userID int) (entity.Cart, error) {
	var getCart entity.Cart

	err := r.db.Where("user_id = ?", userID).Preload("CartItems.Product").First(&getCart).Error

	if err != nil {
		return getCart, err
	}

	return getCart, nil

}

func (r *cartRepository) DeleteCartByID(userID int) (entity.Cart, error) {
	var deleteCart entity.Cart
	err := r.db.Where("id = ?", userID).Delete(&deleteCart).Error

	if err != nil {
		return deleteCart, err
	}

	return deleteCart, nil

}
