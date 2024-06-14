package cart

import (
	"go-technical-test-synapsis/src/entity"
	"go-technical-test-synapsis/src/helper"
	"go-technical-test-synapsis/src/repository"
)

type Service interface {
	StoreCart(request CartRequest) (entity.Cart, error)
	RemoveItemCart(request CartRequest) (entity.Cart, error)
	FindCartByID(userID CartUserIDRequest) (entity.Cart, error)
	UpdateCart(userID CartIDRequest, request CartRequest) (entity.Cart, error)
	DeleteCartByID(userID CartIDRequest) (entity.Cart, error)
}

type service struct {
	repository      repository.CartRepository
	itemsRepository repository.CartItemsRepository
}

func NewService(repository repository.CartRepository, itemsRepository repository.CartItemsRepository) *service {
	return &service{repository, itemsRepository}
}

func (s *service) StoreCart(request CartRequest) (entity.Cart, error) {
	cart := entity.Cart{}
	cart.UserID = request.UserID
	cart.CreatedAt = helper.TimeNow()

	getCart, err := s.repository.FindByID(request.UserID)

	if err != nil {
		getCart, err = s.repository.StoreToCart(cart)
	}

	checkItems, err := s.itemsRepository.FindByProductID(request.ProductID, request.UserID)

	checkItems.CartID = getCart.ID
	checkItems.ProductID = request.ProductID
	checkItems.Quantity = request.Quantity
	checkItems.CreatedAt = helper.TimeNow()

	if err != nil {
		getCart, err = s.repository.StoreToCart(cart)
	}

	_, err = s.itemsRepository.Save(checkItems)

	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (s *service) FindCartByID(userID CartUserIDRequest) (entity.Cart, error) {
	cart, err := s.repository.FindByID(userID.ID)

	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (s *service) UpdateCart(userID CartIDRequest, request CartRequest) (entity.Cart, error) {

	cart, err := s.repository.FindByID(userID.ID)

	if err != nil {
		return cart, err
	}

	return entity.Cart{}, nil
}

func (s *service) DeleteCartByID(userID CartIDRequest) (entity.Cart, error) {
	cart, err := s.repository.DeleteCartByID(userID.ID)

	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (s *service) RemoveItemCart(request CartRequest) (entity.Cart, error) {
	getCart, err := s.repository.FindByID(request.UserID)

	if err != nil {
		return getCart, err
	}

	_, err = s.itemsRepository.DeleteCartProductByID(request.ProductID, request.UserID)

	if err != nil {
		return getCart, err
	}

	return entity.Cart{}, nil
}
