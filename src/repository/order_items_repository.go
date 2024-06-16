package repository

import (
	"go-technical-test-synapsis/src/entity"

	"gorm.io/gorm"
)

type OrderItemsRepository interface {
	Save(order entity.OrderItems) (entity.OrderItems, error)
}

type orderItemsRepository struct {
	db *gorm.DB
}

func NewOrderItems(db *gorm.DB) *orderItemsRepository {
	return &orderItemsRepository{db}
}

func (r *orderItemsRepository) Save(order entity.OrderItems) (entity.OrderItems, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
