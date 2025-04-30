package repository

import "app/model"

type UserRepository struct {
	IRepository[model.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{&Repository[model.User]{}}
}
