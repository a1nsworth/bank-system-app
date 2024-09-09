package services

import (
	"bank-system-app/internal/models/constraints"
	"bank-system-app/internal/repositories"

	"gorm.io/gorm"
)

type baseService[TModel any, TId constraints.ID] struct {
	repository repositories.Repository[TModel, TId]
}

func newBaseService[TModel any, TId constraints.ID](
	db *gorm.DB,
	model TModel,
) baseService[TModel, TId] {
	return baseService[TModel, TId]{
		repository: repositories.NewGormRepository[TModel, TId](db, model),
	}
}

type crudService[TModel any, TId constraints.ID] struct {
	baseService[TModel, TId]
}

func newCRUDService[TModel any, TId constraints.ID](
	db *gorm.DB,
	model TModel,
) crudService[TModel, TId] {
	return crudService[TModel, TId]{baseService: newBaseService[TModel, TId](db, model)}
}

func (s *crudService[TModel, TId]) Create(model *TModel) (*TModel, error) {
	return s.repository.Create(model)
}

func (s *crudService[TModel, TId]) GetByID(id TId) (*TModel, error) {
	return s.repository.GetByID(id)
}

func (s *crudService[TModel, TId]) GetAll() ([]TModel, error) {
	return s.repository.GetAll()
}

func (s *crudService[TModel, TId]) Update(model *TModel) (*TModel, error) {
	return s.repository.Update(model)
}

func (s *crudService[TModel, TId]) DeleteById(id TId) error {
	return s.repository.DeleteById(id)
}

func (s *crudService[TModel, TId]) DeleteByConditions(conditions ...any) error {
	return s.repository.DeleteByConditions(conditions)
}
