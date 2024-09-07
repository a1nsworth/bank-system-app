package database

import "gorm.io/gorm"

type Database interface {
	GetConnection() *gorm.DB
	MigrateModels(models ...any) error
}
