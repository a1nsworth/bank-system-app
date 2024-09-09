package services

import (
	"fmt"

	"bank-system-app/internal/models"

	"gorm.io/gorm"
)

type BankService struct {
	crudService[models.Bank, uint]
}

func NewBankService(db *gorm.DB) BankService {
	fmt.Println("new bank service")
	return BankService{crudService: newCRUDService[models.Bank, uint](db, models.Bank{})}
}

type BankAtmService struct {
	crudService[models.BankAtm, uint]
}

func NewBankAtmService(db *gorm.DB) BankAtmService {
	return BankAtmService{crudService: newCRUDService[models.BankAtm, uint](db, models.BankAtm{})}
}

type BankOfficeService struct {
	crudService[models.BankOffice, uint]
}

func NewBankOfficeService(db *gorm.DB) BankOfficeService {
	return BankOfficeService{
		crudService: newCRUDService[models.BankOffice, uint](
			db,
			models.BankOffice{},
		),
	}
}
