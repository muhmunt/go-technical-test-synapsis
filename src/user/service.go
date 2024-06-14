package user

import (
	"errors"
	"go-technical-test-synapsis/src/entity"
	"go-technical-test-synapsis/src/helper"
	"go-technical-test-synapsis/src/repository"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(request RegisterRequest) (entity.User, error)
	CheckUsernameAvailable(username string) (bool, error)
	Login(request LoginRequest) (entity.User, error)
	GetUserByID(ID int) (entity.User, error)
}

type service struct {
	repository repository.UserRepository
}

func NewService(repository repository.UserRepository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(request RegisterRequest) (entity.User, error) {
	user := entity.User{}
	user.Name = request.Name
	user.Username = request.Username
	user.Role = "USER"
	user.CreatedAt = helper.TimeNow()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) CheckUsernameAvailable(username string) (bool, error) {
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) Login(request LoginRequest) (entity.User, error) {
	username := request.Username
	password := request.Password

	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Username not found.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserByID(ID int) (entity.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found.")
	}

	return user, nil
}
