package storage

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB initializes the database connection with retry logic
func InitDB() error {
	// Get database connection string from environment
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// For development mode, use default local database
		if os.Getenv("GIN_MODE") != "release" {
			dbURL = "postgresql://cvadmin:cv2024secure@localhost:5432/curriculum_vitae_dev?sslmode=disable"
			fmt.Printf("🔧 Development mode: using default local database\n")
		} else {
			return fmt.Errorf("DATABASE_URL environment variable is required for production deployment")
		}
	}

	fmt.Printf("Attempting to connect to database...\n")

	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Retry connection with backoff
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		if err = db.Ping(); err == nil {
			fmt.Printf("✅ Database connected successfully\n")
			
			// Create tables if they don't exist
			if err := createTables(); err != nil {
				return fmt.Errorf("failed to create database tables: %w", err)
			}
			fmt.Printf("✅ Database tables initialized\n")
			
			return nil
		}
		
		fmt.Printf("⏳ Database connection attempt %d/%d failed, retrying in %d seconds...\n", 
			i+1, maxRetries, (i+1)*2)
		time.Sleep(time.Duration((i+1)*2) * time.Second)
	}

	return fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}

// createTables creates the required database tables if they don't exist
func createTables() error {
	// Create users table
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		first_login BOOLEAN DEFAULT true,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		last_password_change TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		login_count INTEGER DEFAULT 0,
		last_login_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
	`

	// Create cv_files table
	cvFilesTable := `
	CREATE TABLE IF NOT EXISTS cv_files (
		id SERIAL PRIMARY KEY,
		filename VARCHAR(255) NOT NULL,
		original_name VARCHAR(255) NOT NULL,
		content_type VARCHAR(100) NOT NULL DEFAULT 'application/pdf',
		file_size BIGINT NOT NULL,
		file_data BYTEA NOT NULL,
		is_current BOOLEAN DEFAULT true,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE INDEX IF NOT EXISTS idx_cv_files_current ON cv_files(is_current) WHERE is_current = true;
	CREATE INDEX IF NOT EXISTS idx_cv_files_created_at ON cv_files(created_at);
	`

	// Execute table creation
	if _, err := db.Exec(usersTable); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	if _, err := db.Exec(cvFilesTable); err != nil {
		return fmt.Errorf("failed to create cv_files table: %w", err)
	}

	return nil
}

