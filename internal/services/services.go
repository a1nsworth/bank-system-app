package services

type CrudService[T any] struct {
	repo Repository[T]
}

func NewCrudService[T any](repo Repository[T]) *CrudService[T] {
	return &CrudService[T]{repo: repo}
}

func (s *CrudService[T]) Create(model *T) (*T, error) {
	return s.repo.Create(model)
}

func (s *CrudService[T]) GetByID(id uint) (*T, error) {
	return s.repo.GetByID(id)
}

func (s *CrudService[T]) GetAll() ([]T, error) {
	return s.repo.GetAll()
}

func (s *CrudService[T]) Update(model *T) (*T, error) {
	return s.repo.Update(model)
}

func (s *CrudService[T]) Delete(id uint) error {
	return s.repo.Delete(id)
}
