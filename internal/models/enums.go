package models

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

type namer interface {
	Name() string
}

type baseIntEnum int

func (enum baseIntEnum) Name() string {
	return reflect.TypeOf(enum).Name()
}

func (enum baseIntEnum) Value() (driver.Value, error) {
	return int64(enum), nil
}

func (enum *baseIntEnum) Scan(src any) error {
	v, ok := src.(int64)
	if !ok {
		panic(fmt.Errorf("cannot scan type %s into int64", enum.Name()))
	}

	*enum = baseIntEnum(v)
	return nil
}

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
