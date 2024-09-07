package services_test

import (
	"bank-system-app/internal/database"
	"bank-system-app/internal/models"
	"bank-system-app/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB() (*gorm.DB, func()) {
	var db database.Database
	db, err := database.NewSQLiteDatabase("../migrations/test.db", &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	list_models := []any{
		&models.Bank{},
		&models.BankOffice{},
		&models.BankAtm{},
		&models.Employee{},
		&models.User{},
		&models.PaymentAccount{},
		&models.CreditAccount{},
	}
	if err := db.MigrateModels(list_models...); err != nil {
		panic("Failed to migrate models:")
	}
	conn := db.GetConnection()
	return conn, func() {
		if err := conn.Migrator().DropTable(list_models...); err != nil {
			panic("Error during table drop: %v")
		}
	}
}

func TestCreate(t *testing.T) {
	db, teardown := setupTestDB()
	defer teardown()

	bankRepo := services.NewGormRepository[models.Bank](db, models.Bank{})
	bank := &models.Bank{Name: "Test Bank"}
	createdBank, err := bankRepo.Create(bank)
	assert.NoError(t, err)
	assert.NotNil(t, createdBank)
	assert.Equal(t, "Test Bank", createdBank.Name)
}

func TestGetByID(t *testing.T) {
	db, teardown := setupTestDB()
	defer teardown()

	bankRepo := services.NewGormRepository[models.Bank](db, models.Bank{})

	bank := &models.Bank{Name: "Test Bank"}
	createdBank, _ := bankRepo.Create(bank)

	retrievedBank, err := bankRepo.GetByID(createdBank.ID)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedBank)
	assert.Equal(t, createdBank.ID, retrievedBank.ID)
	assert.Equal(t, "Test Bank", retrievedBank.Name)
}

func TestUpdate(t *testing.T) {
	db, teardown := setupTestDB()
	defer teardown()

	bankRepo := services.NewGormRepository[models.Bank](db, models.Bank{})

	bank := &models.Bank{Name: "Test Bank"}
	createdBank, _ := bankRepo.Create(bank)

	createdBank.Name = "Updated Bank"
	updatedBank, err := bankRepo.Update(createdBank)
	assert.NoError(t, err)
	assert.NotNil(t, updatedBank)
	assert.Equal(t, "Updated Bank", updatedBank.Name)
}

func TestDelete(t *testing.T) {
	db, teardown := setupTestDB()
	defer teardown()

	bankRepo := services.NewGormRepository[models.Bank](db, models.Bank{})

	bank := &models.Bank{Name: "Test Bank"}
	createdBank, err := bankRepo.Create(bank)
	if err != nil {
		t.Fatalf("failed to create bank: %v", err)
	}

	err = bankRepo.Delete(createdBank.ID)
	if err != nil {
		t.Fatalf("failed to delete bank: %v", err)
	}

	// Проверка удаления
	_, err = bankRepo.GetByID(createdBank.ID)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
}
