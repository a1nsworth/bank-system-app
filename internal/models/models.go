package models

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"

	"bank-system-app/internal/database"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	Name         string `gorm:"size:256;unique"`
	BankOffices  []BankOffice
	BankAtms     []BankAtm
	Users        []User `gorm:"many2many:banks_users"`
	Employees    []Employee
	Rating       uint8
	TotalSum     uint32
	InterestRate uint8
}

func (b *Bank) BeforeCreate(tx *gorm.DB) (err error) {
	r := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))
	b.Rating = uint8(r.Int32N(int32(BankMaxRating + 1)))

	fmt.Println("rating", b.Rating)
	b.TotalSum = uint32(r.Int32N(int32(BankMaxTotalSum + 1)))
	fmt.Println("totalSum", b.TotalSum)
	b.InterestRate = uint8(
		math.Round(
			(1 - float64(b.Rating)/float64(BankMaxRating)) * float64(
				BankMaxInterestRate,
			) * rand.Float64(),
		),
	)

	return
}

func (b *Bank) CountAtms() int {
	return len(b.BankAtms)
}

func (b *Bank) CountOffices() int {
	return len(b.BankOffices)
}

func (b *Bank) CountUsers() int {
	return len(b.Users)
}

func (b *Bank) CountEmployees() int {
	return len(b.Employees)
}

type BankOffice struct {
	gorm.Model
	Addres string        `gorm:"size:100"`
	Status *OfficeStatus `gorm:"type_office_status"`
	Rental uint32

	Bank   Bank
	BankID uint
}

type BankAtm struct {
	gorm.Model
	Name         string     `gorm:"size:30"`
	Status       *AtmStatus `gorm:"type:atm_status"`
	Amortization uint

	Bank                 Bank
	BankOffice           BankOffice
	BankOfficeID, BankID uint
}

func (atm BankAtm) Addres() string {
	return atm.BankOffice.Addres
}

func (atm BankAtm) Owner() string {
	return atm.Bank.Name
}

func (atm BankAtm) TotalSum() uint32 {
	return atm.Bank.TotalSum
}

type Employee struct {
	gorm.Model
	personModel

	Position string
	Status   *EmployeeStatus
	Salary   uint

	Bank                 Bank
	BankOffice           BankOffice
	BankID, BankOfficeID uint
}

type User struct {
	gorm.Model
	personModel

	PlaceOfWork     string `gorm:"size:30"`
	MonthlyIncome   uint
	BankCreditScore uint

	BanksUsed       []Bank `gorm:"many2many:banks_users"`
	CreditAccounts  []CreditAccount
	PaymentAccounts []PaymentAccount
}

type PaymentAccount struct {
	gorm.Model
	UserID, BankID uint
	Bank           Bank
	User           User

	Balance uint
}

type CreditAccount struct {
	gorm.Model
	UserID, BankID, EmployeeID, PaymentAccountID uint
	User                                         User
	Bank                                         Bank
	Employee                                     Employee
	PaymentAccount                               PaymentAccount

	LoanStartDate, LoanEndDate                datatypes.Date
	LoanDurationMounts                        uint8
	LoanAmount, MounthlyPayment, InterestRate uint
}

func MigrateTables(database database.Database) (err error) {
	println("MIGRATE")
	err = database.Migrate(
		&Bank{},
		&BankOffice{},
		&BankAtm{},
		&Employee{},
		&User{},
		&PaymentAccount{},
		&CreditAccount{},
	)
	if err != nil {
		return
	}
	return err
}

func DropTables(database database.Database) (err error) {
	err = database.Drop(
		&Bank{},
		&BankOffice{},
		&BankAtm{},
		&Employee{},
		&User{},
		&PaymentAccount{},
		&CreditAccount{},
	)
	if err != nil {
		return
	}
	return err
}
