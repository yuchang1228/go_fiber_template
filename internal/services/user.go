package services

import (
	"app/internal/models"
	"app/internal/repositories"
)

type IUserService interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(
	userRepository repositories.IUserRepository,
) IUserService {
	return &userService{userRepository}
}

func (s *userService) Create(user *models.User) error {
	return s.userRepository.Create(user)
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *userService) GetAll() (*[]models.User, error) {
	return s.userRepository.GetAll()
}

func (s *userService) Update(user *models.User) error {
	return s.userRepository.Update(user)
}

func (s *userService) Delete(id uint) error {
	return s.userRepository.Delete(id)
}
