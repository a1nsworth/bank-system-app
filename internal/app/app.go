package app

import (
	"bank-system-app/internal/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Run() {
	db, err := gorm.Open(sqlite.Open("migrations/main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&models.Bank{},
		&models.BankOffice{},
		&models.BankAtm{},
		// &models.Employee{},
		// &models.CreditAccount{},
		// &models.User{},
		// &models.PaymentAccount{},
	)
	if err != nil {
		panic(err)
	}

	bank := models.Bank{
		Name:        "SBERID",
		BankAtms:    []models.BankAtm{{Status: models.AtmHaveMoney | models.AtmAbleWithdraw}},
		BankOffices: []models.BankOffice{{}},
	}
	result := db.Create(&bank)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	var banks []models.Bank

	result = db.Find(&banks)
	fmt.Println(banks)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
}
