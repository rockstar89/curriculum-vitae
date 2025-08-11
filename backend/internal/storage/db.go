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
		// Check if we're in development mode
		// Check if individual database components are provided
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		sslmode := os.Getenv("DB_SSLMODE")
		
		if host != "" && user != "" && password != "" && dbname != "" {
			// Use individual database components
			if port == "" {
				port = "5432"
			}
			if sslmode == "" {
				sslmode = "require" // Default to secure connection for production
			}
			dbURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
				host, port, user, password, dbname, sslmode)
			fmt.Printf("🔧 Using individual database environment variables\n")
		} else if os.Getenv("GIN_MODE") != "release" {
			// Development fallback
			host = getEnvDefault("DB_HOST", "localhost")
			port = getEnvDefault("DB_PORT", "5432")
			user = getEnvDefault("DB_USER", "cvadmin")
			password = getEnvDefault("DB_PASSWORD", "cv2024secure")
			dbname = getEnvDefault("DB_NAME", "curriculum_vitae_dev")
			sslmode = getEnvDefault("DB_SSLMODE", "disable")
			
			dbURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
				host, port, user, password, dbname, sslmode)
			fmt.Printf("🔧 Development mode: using local database connection\n")
		} else {
			return fmt.Errorf("DATABASE_URL environment variable or individual DB_* variables are required for production deployment")
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

// getEnvDefault gets environment variable or returns default value
func getEnvDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}