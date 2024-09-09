package repositories

import (
	"bank-system-app/internal/models/constraints"
)

type Repository[TModel any, TId constraints.ID] interface {
	Create(model *TModel) (*TModel, error)
	GetByID(id TId) (*TModel, error)
	GetAll() ([]TModel, error)
	Update(model *TModel) (*TModel, error)
	DeleteById(id TId) error
	DeleteByConditions(conditions ...any) error
}
