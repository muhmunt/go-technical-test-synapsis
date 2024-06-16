package repository

import (
	"go-technical-test-synapsis/src/entity"

	"gorm.io/gorm"
)

type CartItemsRepository interface {
	Save(cart entity.CartItems) (entity.CartItems, error)
	FindByID(ID int) (entity.CartItems, error)
	FindByCartID(cartID int) (entity.CartItems, error)
	FindByProductID(productID, cartID int) (entity.CartItems, error)
	DeleteCartProductByID(productID, userID int) (entity.CartItems, error)
	DeleteCartByID(cartID int) (entity.CartItems, error)
	DeleteByID(itemID int) (entity.CartItems, error)
}

type cartItemsRepository struct {
	db *gorm.DB
}

func NewCartItems(db *gorm.DB) *cartItemsRepository {
	return &cartItemsRepository{db}
}

func (r *cartItemsRepository) Save(cart entity.CartItems) (entity.CartItems, error) {
	err := r.db.Save(&cart).Error

	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (r *cartItemsRepository) FindByID(ID int) (entity.CartItems, error) {
	var getCart entity.CartItems

	err := r.db.Where("id = ?", ID).Preload("Product").First(&getCart).Error

	if err != nil {
		return getCart, err
	}

	return getCart, nil

}

func (r *cartItemsRepository) FindByCartID(cartID int) (entity.CartItems, error) {
	var getCart entity.CartItems

	err := r.db.Where("cart_id = ?", cartID).Preload("Product").First(&getCart).Error

	if err != nil {
		return getCart, err
	}

	return getCart, nil

}

func (r *cartItemsRepository) FindByProductID(productID, cartID int) (entity.CartItems, error) {
	var getCart entity.CartItems

	err := r.db.Where("product_id = ?", productID).Where("cart_id = ? ", cartID).First(&getCart).Error

	if err != nil {
		return getCart, err
	}

	return getCart, nil

}

func (r *cartItemsRepository) DeleteCartProductByID(productID, userID int) (entity.CartItems, error) {
	var deleteCart entity.CartItems
	err := r.db.Where("product_id = ?", productID).Where("cart_id = ? ", userID).Delete(&deleteCart).Error

	if err != nil {
		return deleteCart, err
	}

	return deleteCart, nil

}

func (r *cartItemsRepository) DeleteCartByID(cart int) (entity.CartItems, error) {
	var deleteCart entity.CartItems
	err := r.db.Where("cart_id = ?", cart).Delete(&deleteCart).Error

	if err != nil {
		return deleteCart, err
	}

	return deleteCart, nil

}

func (r *cartItemsRepository) DeleteByID(itemID int) (entity.CartItems, error) {
	var deleteCart entity.CartItems
	err := r.db.Where("id = ?", itemID).Delete(&deleteCart).Error

	if err != nil {
		return deleteCart, err
	}

	return deleteCart, nil

}
