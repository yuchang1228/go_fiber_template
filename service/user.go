package service

import (
	"app/model"
	"app/repository"
)

type IUserService interface {
	CreateUser(user *model.User) error
	GetUserByID(id string) (*model.User, error)
	GetAllUsers() (*[]model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
	GetUserByUsername(username string) (*model.User, error)
}

type userService struct {
	userRepository *repository.UserRepository
}

func NewUserService(
	userRepository *repository.UserRepository,
) IUserService {
	return &userService{
		userRepository,
	}
}

func (s *userService) CreateUser(user *model.User) error {
	return s.userRepository.Create(user)
}

func (s *userService) GetUserByID(id string) (*model.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *userService) GetAllUsers() (*[]model.User, error) {
	return s.userRepository.GetAll()
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.userRepository.Update(user)
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepository.Delete(id)
}

func (s *userService) GetUserByUsername(username string) (*model.User, error) {
	user, err := s.userRepository.GetByUserName(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, nil
}
