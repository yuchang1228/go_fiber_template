package service

import (
	"app/model"
	"app/repository"
)

type IUserService interface {
	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetAll() (*[]model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
}

type userService struct {
	userRepository repository.IUserRepository
}

func NewUserService(
	userRepository repository.IUserRepository,
) IUserService {
	return &userService{userRepository}
}

func (s *userService) Create(user *model.User) error {
	return s.userRepository.Create(user)
}

func (s *userService) GetByID(id uint) (*model.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *userService) GetAll() (*[]model.User, error) {
	return s.userRepository.GetAll()
}

func (s *userService) Update(user *model.User) error {
	return s.userRepository.Update(user)
}

func (s *userService) Delete(id uint) error {
	return s.userRepository.Delete(id)
}
