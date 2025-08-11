package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	// Get database connection string from environment
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Fallback to individual components
		host := getEnvDefault("DB_HOST", "localhost")
		port := getEnvDefault("DB_PORT", "5432")
		user := getEnvDefault("DB_USER", "cvadmin")
		password := getEnvDefault("DB_PASSWORD", "cv2024secure")
		dbname := getEnvDefault("DB_NAME", "curriculum_vitae_dev")
		sslmode := getEnvDefault("DB_SSLMODE", "disable")
		
		dbURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslmode)
	}

	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}

// getEnvDefault gets environment variable or returns default value
func getEnvDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}