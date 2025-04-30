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
}

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(
	r repository.UserRepository,
) IUserService {
	return &UserService{
		repository: r,
	}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.repository.Create(user)
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	return s.repository.GetByID(id)
}

func (s *UserService) GetAllUsers() (*[]model.User, error) {
	return s.repository.GetAll()
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.repository.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repository.Delete(id)
}
