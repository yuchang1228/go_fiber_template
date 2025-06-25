package repository

import (
	"app/model"

	"gorm.io/gorm"
)

// 套用 base 版本
// type UserRepository struct {
// 	IRepository[model.User]
// }

// func NewUserRepository() *UserRepository {
// 	return &UserRepository{&Repository[model.User]{}}
// }

type IUserRepository interface {
	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetAll() (*[]model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
	GetByUserName(username string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) GetAll() (*[]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return &users, err
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) Delete(id uint) error {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByUserName(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
