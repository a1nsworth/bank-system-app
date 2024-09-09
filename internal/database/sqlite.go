package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteDatabase struct {
	db *gorm.DB
}

func NewSQLiteDatabase(path string, config *gorm.Config) (*SQLiteDatabase, error) {
	db, err := gorm.Open(sqlite.Open(path), config)
	if err != nil {
		return nil, err
	}
	return &SQLiteDatabase{db: db}, nil
}

func (db *SQLiteDatabase) GetConnection() *gorm.DB {
	return db.db
}

func (db *SQLiteDatabase) Migrate(models ...any) error {
	return db.db.AutoMigrate(models...)
}

func (db *SQLiteDatabase) Drop(models ...any) error {
	return db.db.Migrator().DropTable(models...)
}
