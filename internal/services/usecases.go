package services

import (
	"fmt"
	"log"
	"time"

	"bank-system-app/internal/database"
	"bank-system-app/internal/models"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const migrationsPath = "migrations/main.db"

func exampleUsage() {
	var db database.Database
	db, err := database.NewSQLiteDatabase(migrationsPath, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err = models.MigrateTables(db); err != nil {
		panic("Failed to migrate models:")
	}

	bankService := NewBankService(db)
	bankOfficeService := NewBankOfficeService(db)
	atmService := NewBankAtmService(db)
	employeeService := NewEmployeeService(db)
	userService := NewUserService(db)
	paymentService := NewPaymentAccountService(db)
	creditService := NewCreditAccountService(db)

	bank := &models.Bank{Name: "Test Bank"}
	createdBank, err := bankService.Create(bank)
	if err != nil {
		log.Fatal("failed to create bank:", err)
	}

	bankOffice := &models.BankOffice{
		Addres: "123 Test St",
		Status: func() *models.OfficeStatus {
			status := models.OfficeActive
			return &status
		}(),
		Rental: 5000,
		BankID: createdBank.ID,
	}
	createdBankOffice, err := bankOfficeService.Create(bankOffice)
	if err != nil {
		log.Fatal("failed to create bank office:", err)
	}

	atm := &models.BankAtm{
		Name: "ATM 1",
		Status: func() *models.AtmStatus {
			status := models.AtmActive | models.AtmHaveMoney
			return &status
		}(),
		Amortization: 10000,
		BankID:       createdBank.ID,
		BankOfficeID: createdBankOffice.ID,
	}
	createdAtm, err := atmService.Create(atm)
	if err != nil {
		log.Fatal("failed to create ATM:", err)
	}

	employee := &models.Employee{
		Position: "Bank Manager",
		Status: func() *models.EmployeeStatus {
			status := models.EmployeeCanGiveLoans
			return &status
		}(),
		Salary:       60000,
		BankID:       createdBank.ID,
		BankOfficeID: createdBankOffice.ID,
	}
	createdEmployee, err := employeeService.Create(employee)
	if err != nil {
		log.Fatal("failed to create employee:", err)
	}

	user := &models.User{
		PlaceOfWork:     "Test Company",
		MonthlyIncome:   3000,
		BankCreditScore: 150,
		BanksUsed:       []models.Bank{*createdBank},
	}
	createdUser, err := userService.Create(user)
	if err != nil {
		log.Fatal("failed to create user:", err)
	}

	paymentAccount := &models.PaymentAccount{
		UserID:  createdUser.ID,
		BankID:  createdBank.ID,
		Balance: 1000,
	}
	createdPaymentAccount, err := paymentService.Create(paymentAccount)
	if err != nil {
		log.Fatal("failed to create payment account:", err)
	}

	creditAccount := &models.CreditAccount{
		UserID:             createdUser.ID,
		BankID:             createdBank.ID,
		EmployeeID:         createdEmployee.ID,
		PaymentAccountID:   createdPaymentAccount.ID,
		LoanStartDate:      datatypes.Date(time.Now()),
		LoanEndDate:        datatypes.Date(time.Now().AddDate(1, 0, 0)),
		LoanDurationMounts: 12,
		LoanAmount:         5000,
		MounthlyPayment:    450,
		InterestRate:       5,
	}
	createdCreditAccount, err := creditService.Create(creditAccount)
	if err != nil {
		log.Fatal("failed to create credit account:", err)
	}

	// Проверка и вывод созданных данных
	createdBank, err = bankService.GetByID(createdBank.ID)
	if err != nil {
		log.Fatal("failed to get bank by ID:", err)
	}
	fmt.Printf(
		"Bank: ID=%d, Name=%s, Rating=%d\n",
		createdBank.ID,
		createdBank.Name,
		createdBank.Rating,
	)

	createdBankOffice, err = bankOfficeService.GetByID(createdBankOffice.ID)
	if err != nil {
		log.Fatal("failed to get bank office by ID:", err)
	}
	fmt.Printf(
		"BankOffice: ID=%d, Address=%s, Status=%v\n",
		createdBankOffice.ID,
		createdBankOffice.Addres,
		createdBankOffice.Status,
	)

	createdAtm, err = atmService.GetByID(createdAtm.ID)
	if err != nil {
		log.Fatal("failed to get ATM by ID:", err)
	}
	fmt.Printf(
		"BankAtm: ID=%d, Name=%s, Status=%v\n",
		createdAtm.ID,
		createdAtm.Name,
		createdAtm.Status,
	)

	createdEmployee, err = employeeService.GetByID(createdEmployee.ID)
	if err != nil {
		log.Fatal("failed to get employee by ID:", err)
	}
	fmt.Printf(
		"Employee: ID=%d, Position=%s, Status=%v\n",
		createdEmployee.ID,
		createdEmployee.Position,
		createdEmployee.Status,
	)

	createdUser, err = userService.GetByID(createdUser.ID)
	if err != nil {
		log.Fatal("failed to get user by ID:", err)
	}
	fmt.Printf(
		"User: ID=%d, PlaceOfWork=%s, MonthlyIncome=%d\n",
		createdUser.ID,
		createdUser.PlaceOfWork,
		createdUser.MonthlyIncome,
	)

	createdPaymentAccount, err = paymentService.GetByID(createdPaymentAccount.ID)
	if err != nil {
		log.Fatal("failed to get payment account by ID:", err)
	}
	fmt.Printf(
		"PaymentAccount: ID=%d, UserID=%d, BankID=%d, Balance=%d\n",
		createdPaymentAccount.ID,
		createdPaymentAccount.UserID,
		createdPaymentAccount.BankID,
		createdPaymentAccount.Balance,
	)

	createdCreditAccount, err = creditService.GetByID(createdCreditAccount.ID)
	if err != nil {
		log.Fatal("failed to get credit account by ID:", err)
	}
	fmt.Printf(
		"CreditAccount: ID=%d, UserID=%d, BankID=%d, LoanAmount=%d, InterestRate=%d\n",
		createdCreditAccount.ID,
		createdCreditAccount.UserID,
		createdCreditAccount.BankID,
		createdCreditAccount.LoanAmount,
		createdCreditAccount.InterestRate,
	)
}
