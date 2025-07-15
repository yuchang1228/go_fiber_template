package test

import (
	"app/internal/model"
	"app/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(id uint) (*model.User, error) {
	args := m.Called(id)
	if user, ok := args.Get(0).(*model.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) GetAll() (*[]model.User, error) {
	args := m.Called()
	if users, ok := args.Get(0).(*[]model.User); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Update(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) GetByUserName(username string) (*model.User, error) {
	args := m.Called(username)
	if user, ok := args.Get(0).(*model.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestUserService_Create(t *testing.T) {
	mockUserRepository := new(MockUserRepository)
	userService := service.NewUserService(mockUserRepository)

	user := &model.User{
		Model:    gorm.Model{ID: 1},
		Username: "testuser",
	}
	mockUserRepository.On("Create", user).Return(nil)

	err := userService.Create(user)
	assert.NoError(t, err, "Expected no error when creating user")

	mockUserRepository.AssertExpectations(t)
}

func TestUserService_GetByID(t *testing.T) {
	mockUserRepository := new(MockUserRepository)
	userService := service.NewUserService(mockUserRepository)

	user := &model.User{
		Model:    gorm.Model{ID: 1},
		Username: "testuser",
	}
	mockUserRepository.On("GetByID", uint(1)).Return(user, nil)

	result, err := userService.GetByID(1)
	assert.NoError(t, err, "Expected no error when getting user by ID")
	assert.Equal(t, user, result, "Expected returned user to match")

	mockUserRepository.AssertExpectations(t)
}

func TestUserService_GetAll(t *testing.T) {
	mockUserRepository := new(MockUserRepository)
	userService := service.NewUserService(mockUserRepository)

	users := []model.User{
		{Model: gorm.Model{ID: 1}, Username: "user1"},
		{Model: gorm.Model{ID: 2}, Username: "user2"},
	}
	mockUserRepository.On("GetAll").Return(&users, nil)

	result, err := userService.GetAll()
	assert.NoError(t, err, "Expected no error when getting all users")
	assert.Equal(t, &users, result, "Expected returned users to match")

	mockUserRepository.AssertExpectations(t)
}

func TestUserService_Update(t *testing.T) {
	mockUserRepository := new(MockUserRepository)
	userService := service.NewUserService(mockUserRepository)

	user := &model.User{
		Model:    gorm.Model{ID: 1},
		Username: "testuser",
	}
	mockUserRepository.On("Update", user).Return(nil)

	err := userService.Update(user)
	assert.NoError(t, err, "Expected no error when updating user")

	mockUserRepository.AssertExpectations(t)
}

func TestUserService_Delete(t *testing.T) {
	mockUserRepository := new(MockUserRepository)
	userService := service.NewUserService(mockUserRepository)

	mockUserRepository.On("Delete", uint(1)).Return(nil)

	err := userService.Delete(1)
	assert.NoError(t, err, "Expected no error when deleting user")

	mockUserRepository.AssertExpectations(t)
}
