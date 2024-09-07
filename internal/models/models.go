package models

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"

	"gorm.io/gorm"
)

var (
	BankMaxInterestRate uint8 = 20
	BankMaxTotalSum           = uint32(16)
	BankMaxRating       uint8 = 100
)

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
	BankID uint
}

type BankAtm struct {
	gorm.Model
	BankID uint
}

type CreditAccount struct{}

type Employee struct{}

type PaymentAccount struct{}

type User struct{}
