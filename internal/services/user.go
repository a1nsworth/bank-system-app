package services

import (
	"bank-system-app/internal/database"
	"bank-system-app/internal/models"
)

type UserService struct {
	crudService[models.User, uint]
}

func NewUserService(db database.Database) UserService {
	return UserService{
		crudService: newCRUDService[models.User, uint](
			db.GetConnection(),
			models.User{},
		),
	}
}

type EmployeeService struct {
	crudService[models.Employee, uint]
}

func NewEmployeeService(db database.Database) EmployeeService {
	return EmployeeService{
		crudService: newCRUDService[models.Employee, uint](
			db.GetConnection(),
			models.Employee{},
		),
	}
}
