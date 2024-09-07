package models

import (
	"fmt"
	"math"
	"math/rand/v2"
	"reflect"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var (
	BankMaxInterestRate uint8 = 20
	BankMaxTotalSum           = uint32(16)
	BankMaxRating       uint8 = 100
)

type OfficeStatus baseIntEnum

const (
	OfficeActive OfficeStatus = 1 << iota
	OfficeAbleToPlaceAtm
	OfficeCreditAvailable
)

func (of OfficeStatus) Name() string {
	return reflect.TypeOf(of).Name()
}

type AtmStatus baseIntEnum

const (
	AtmActive AtmStatus = 1 << iota
	AtmHaveMoney
	AtmWorkToDespenseMoney
	AtmAbleWithdraw
)

type EmployeeStatus baseIntEnum

const (
	EmployeeIsRemote EmployeeStatus = 1 << iota
	EmployeeCanGiveLoans
)

func (as AtmStatus) Name() string {
	return reflect.TypeOf(as).Name()
}

type Bank struct {
	gorm.Model
	Name         string `gorm:"size:256;unique"`
	BankOffices  []BankOffice
	BankAtms     []BankAtm
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

type BankOffice struct {
	gorm.Model
	Addres string       `gorm:"size:100"`
	Status OfficeStatus `gorm:"type_office_status"`
	Rental uint32

	Bank   Bank
	BankID uint
}

type BankAtm struct {
	gorm.Model
	Name         string    `gorm:"size:30"`
	Status       AtmStatus `gorm:"type:atm_status"`
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
	FirstName      string `gorm:"size:30"`
	SecondName     string `gorm:"size:30"`
	PatronymicName string `gorm:"size:30"`
	DateOfBirth    datatypes.Date

	Position string
	Status   EmployeeStatus
	Salary   uint

	Bank                 Bank
	BankOffice           BankOffice
	BankID, BankOfficeID uint
}

type CreditAccount struct {
	gorm.Model
}

type PaymentAccount struct{}

type User struct{}
