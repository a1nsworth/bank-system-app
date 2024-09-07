package services

import "gorm.io/gorm"

type Repository[T any] interface {
	Create(model *T) (*T, error)
	GetByID(id uint) (*T, error)
	GetAll() ([]T, error)
	Update(model *T) (*T, error)
	Delete(id uint) error
}

type GormRepository[T any] struct {
	db    *gorm.DB
	model T
}

func NewGormRepository[T any](db *gorm.DB, model T) *GormRepository[T] {
	return &GormRepository[T]{db: db, model: model}
}

func (r *GormRepository[T]) Create(model *T) (*T, error) {
	result := r.db.Create(model)
	return model, result.Error
}

func (r *GormRepository[T]) GetByID(id uint) (*T, error) {
	result := r.db.First(&r.model, id)
	return &r.model, result.Error
}

func (r *GormRepository[T]) GetAll() ([]T, error) {
	var models []T
	result := r.db.Find(&models)
	return models, result.Error
}

func (r *GormRepository[T]) Update(model *T) (*T, error) {
	result := r.db.Save(model)
	return model, result.Error
}

func (r *GormRepository[T]) Delete(id uint) error {
	var model T
	result := r.db.Delete(&model, id)
	return result.Error
}
