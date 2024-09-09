package services

import (
	"bank-system-app/internal/models"

	"gorm.io/gorm"
)

type UserService struct {
	crudService[models.User, uint]
}

func NewUserService(db *gorm.DB) UserService {
	return UserService{crudService: newCRUDService[models.User, uint](db, models.User{})}
}

type EmployeeService struct {
	crudService[models.Employee, uint]
}

func NewEmployeeService(db *gorm.DB) EmployeeService {
	return EmployeeService{
		crudService: newCRUDService[models.Employee, uint](
			db,
			models.Employee{},
		),
	}
}
