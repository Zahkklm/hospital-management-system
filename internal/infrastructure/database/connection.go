package database

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

// Connect establishes a connection to the PostgreSQL database.
func Connect() {
    var err error
    connStr := os.Getenv("DATABASE_URL")
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    if err = db.Ping(); err != nil {
        log.Fatalf("Failed to ping the database: %v", err)
    }

    log.Println("Database connection established")
}

// GetDB returns the database connection.
func GetDB() *sql.DB {
    return db
}