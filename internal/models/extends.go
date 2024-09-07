package models

import "gorm.io/datatypes"

type fioModel struct {
	FirstName      string `gorm:"size:30"`
	SecondName     string `gorm:"size:30"`
	PatronymicName string `gorm:"size:30"`
}

type dateOfBirthModel struct {
	DateOfBirth datatypes.Date
}

type personModel struct {
	fioModel
	dateOfBirthModel
}
