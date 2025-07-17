package repositories

import (
	"app/internal/models"

	"gorm.io/gorm"
)

// 套用 base 版本
// type UserRepository struct {
// 	IRepository[models.User]
// }

// func NewUserRepository() *UserRepository {
// 	return &UserRepository{&repositories[models.User]{}}
// }

type IUserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	GetByUserName(username string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) GetAll() (*[]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return &users, err
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) Delete(id uint) error {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByUserName(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
