package database

import "gorm.io/gorm"

type Database interface {
	GetConnection() *gorm.DB
	Migrate(models ...any) error
	Drop(models ...any) error
}
