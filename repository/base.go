package repository

import (
	"app/database"
)

type IRepository[T any] interface {
	Create(model *T) error
	GetByID(id string) (*T, error)
	GetAll() (*[]T, error)
	Update(model *T) error
	Delete(id string) error
}

type Repository[T any] struct{}

func (r *Repository[T]) Create(model *T) error {
	db := database.GORM_DB
	if err := db.Create(&model).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository[T]) GetByID(id string) (*T, error) {
	db := database.GORM_DB
	var model T
	if err := db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *Repository[T]) GetAll() (*[]T, error) {
	db := database.GORM_DB
	var models []T
	if err := db.Find(&models).Error; err != nil {
		return nil, err
	}
	return &models, nil
}

func (r *Repository[T]) Update(model *T) error {
	db := database.GORM_DB
	if err := db.Save(&model).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository[T]) Delete(id string) error {
	db := database.GORM_DB
	var model T
	if err := db.First(&model, id).Error; err != nil {
		return err
	}
	if err := db.Delete(&model).Error; err != nil {
		return err
	}
	return nil
}
