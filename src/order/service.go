package order

import (
	"errors"
	"go-technical-test-synapsis/src/entity"
	"go-technical-test-synapsis/src/helper"
	"go-technical-test-synapsis/src/repository"
)

type service struct {
	repository           repository.OrderRepository
	cartItemsRepository  repository.CartItemsRepository
	orderItemsrepository repository.OrderItemsRepository
}

type Service interface {
	CreateOrder(request OrderRequest) (entity.Order, error)
}

func NewService(repository repository.OrderRepository, cartItemsRepository repository.CartItemsRepository, orderItemsrepository repository.OrderItemsRepository) *service {
	return &service{repository, cartItemsRepository, orderItemsrepository}
}

func (s *service) CreateOrder(request OrderRequest) (entity.Order, error) {

	items := request.CartItemsID

	totalAmount := 0
	for _, item := range items {
		price, err := s.cartItemsRepository.FindByID(item)

		if err != nil {
			return entity.Order{}, errors.New("Product Not Found")
		}

		totalAmount += price.Product.Price
	}

	transaction := entity.Order{}
	transaction.UserID = request.UserID
	transaction.TotalAmount = totalAmount

	newOrder, err := s.repository.Save(transaction)
	if err != nil {
		return newOrder, err
	}

	for _, item := range items {
		getItem, _ := s.cartItemsRepository.FindByID(item)

		orderItem := entity.OrderItems{}
		orderItem.OrderID = newOrder.ID
		orderItem.ProductID = getItem.ProductID
		orderItem.Quantity = getItem.Quantity
		orderItem.UnitPrice = getItem.Product.Price
		orderItem.CreatedAt = helper.TimeNow()

		_, err := s.orderItemsrepository.Save(orderItem)
		if err == nil {
			_, _ = s.cartItemsRepository.DeleteByID(item)
		}
	}

	return newOrder, nil
}
