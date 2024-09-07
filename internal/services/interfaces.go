package services

type CreateService[T any] interface {
	Create(model *T) (*T, error)
}

type ReadService[T any] interface {
	GetByID(id uint) (*T, error)
	GetAll() ([]T, error)
}

type UpdateService[T any] interface {
	Update(model *T) (*T, error)
}

type DeleteService[T any] interface {
	Delete(id uint) error
}

type CRUDService[T any] interface {
	CreateService[T]
	ReadService[T]
	UpdateService[T]
	DeleteService[T]
}
