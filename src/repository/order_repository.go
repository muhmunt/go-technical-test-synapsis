package repository

import (
	"go-technical-test-synapsis/src/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Save(order entity.Order) (entity.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrder(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Save(order entity.Order) (entity.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
