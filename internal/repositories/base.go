package repositories

import "app/config"

type IRepository[T any] interface {
	Create(models *T) error
	GetByID(id string) (*T, error)
	GetAll() (*[]T, error)
	Update(models *T) error
	Delete(id string) error
}

type Repository[T any] struct{}

func (r *Repository[T]) Create(models *T) error {
	db := config.GORM_DB
	if err := db.Create(&models).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository[T]) GetByID(id string) (*T, error) {
	db := config.GORM_DB
	var models T
	if err := db.First(&models, id).Error; err != nil {
		return nil, err
	}
	return &models, nil
}

func (r *Repository[T]) GetAll() (*[]T, error) {
	db := config.GORM_DB
	var models []T
	if err := db.Find(&models).Error; err != nil {
		return nil, err
	}
	return &models, nil
}

func (r *Repository[T]) Update(models *T) error {
	db := config.GORM_DB
	if err := db.Save(&models).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository[T]) Delete(id string) error {
	db := config.GORM_DB
	var models T
	if err := db.First(&models, id).Error; err != nil {
		return err
	}
	if err := db.Delete(&models).Error; err != nil {
		return err
	}
	return nil
}
