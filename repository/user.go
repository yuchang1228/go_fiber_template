package repository

import (
	"app/database"
	"app/model"
)

type UserRepository struct {
	IRepository[model.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{&Repository[model.User]{}}
}

func (r *UserRepository) GetByUserName(username string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
