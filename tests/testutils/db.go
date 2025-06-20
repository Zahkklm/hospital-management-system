package testutils

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupMockDB creates a mock database connection for testing
func SetupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		sqlDB.Close()
		return nil, nil, err
	}

	return gormDB, mock, nil
}

// CleanupMockDB closes the mock database connection
func CleanupMockDB(db *gorm.DB, mock sqlmock.Sqlmock) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting SQL DB: %v", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		log.Printf("Unfulfilled expectations: %s", err)
	}

	sqlDB.Close()
}

// SetupTestDB creates a test database connection (for integration tests)
func SetupTestDB() (*gorm.DB, error) {
	// This would connect to a real test database
	// For now, return mock
	db, _, err := SetupMockDB()
	return db, err
}
