package services

import (
	"fmt"

	"bank-system-app/internal/database"
	"bank-system-app/internal/models"
)

type BankService struct {
	crudService[models.Bank, uint]
}

func NewBankService(db database.Database) BankService {
	fmt.Println("new bank service")
	return BankService{
		crudService: newCRUDService[models.Bank, uint](
			db.GetConnection(),
			models.Bank{},
		),
	}
}

type BankAtmService struct {
	crudService[models.BankAtm, uint]
}

func NewBankAtmService(db database.Database) BankAtmService {
	return BankAtmService{
		crudService: newCRUDService[models.BankAtm, uint](
			db.GetConnection(),
			models.BankAtm{},
		),
	}
}

type BankOfficeService struct {
	crudService[models.BankOffice, uint]
}

func NewBankOfficeService(db database.Database) BankOfficeService {
	return BankOfficeService{
		crudService: newCRUDService[models.BankOffice, uint](
			db.GetConnection(),
			models.BankOffice{},
		),
	}
}
