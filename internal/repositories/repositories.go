package repositories

import (
	"bank-system-app/internal/models/constraints"

	"gorm.io/gorm"
)

type GormRepository[TModel any, TId constraints.ID] struct {
	db    *gorm.DB
	model TModel
}

func NewGormRepository[TModel any, TId constraints.ID](
	db *gorm.DB,
	model TModel,
) *GormRepository[TModel, TId] {
	return &GormRepository[TModel, TId]{db: db, model: model}
}

func (r *GormRepository[TModel, TId]) Create(model *TModel) (*TModel, error) {
	result := r.db.Create(model)
	return model, result.Error
}

func (r *GormRepository[TModel, TId]) GetByID(id TId) (*TModel, error) {
	result := r.db.First(&r.model, id)
	return &r.model, result.Error
}

func (r *GormRepository[TModel, TId]) GetAll() ([]TModel, error) {
	var models []TModel
	result := r.db.Find(&models)
	return models, result.Error
}

func (r *GormRepository[TModel, TId]) Update(model *TModel) (*TModel, error) {
	result := r.db.Save(model)
	return model, result.Error
}

func (r *GormRepository[TModel, TId]) DeleteById(id TId) error {
	var model TModel
	result := r.db.Delete(&model, id)
	return result.Error
}

func (r *GormRepository[TModel, TId]) DeleteByConditions(conditions ...any) error {
	var model TModel
	result := r.db.Delete(&model, conditions...)
	return result.Error
}
