package storage

import (
	"cv-backend/internal/models"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection with GORM and retry logic
func InitDB() error {
	// Get database connection string from environment
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// For development mode, use default local database
		if os.Getenv("GIN_MODE") != "release" {
			dsn = "postgresql://cvadmin:cv2024secure@localhost:5432/curriculum_vitae_dev?sslmode=disable"
			fmt.Printf("🔧 Development mode: using default local database\n")
		} else {
			return fmt.Errorf("DATABASE_URL environment variable is required for production deployment")
		}
	}

	fmt.Printf("Attempting to connect to database with GORM...\n")

	// Configure GORM logger
	var gormLogger logger.Interface
	if os.Getenv("GIN_MODE") == "release" {
		gormLogger = logger.Default.LogMode(logger.Silent)
	} else {
		gormLogger = logger.Default.LogMode(logger.Info)
	}

	// Retry connection with backoff
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})
		
		if err == nil {
			// Test the connection
			sqlDB, err := DB.DB()
			if err == nil && sqlDB.Ping() == nil {
				fmt.Printf("✅ Database connected successfully with GORM\n")
				
				// Auto-migrate the schema
				if err := autoMigrate(); err != nil {
					return fmt.Errorf("failed to migrate database schema: %w", err)
				}
				fmt.Printf("✅ Database schema migrated\n")
				
				return nil
			}
		}
		
		fmt.Printf("⏳ Database connection attempt %d/%d failed, retrying in %d seconds...\n", 
			i+1, maxRetries, (i+1)*2)
		time.Sleep(time.Duration((i+1)*2) * time.Second)
	}

	return fmt.Errorf("failed to connect to database after %d attempts", maxRetries)
}

// GetDB returns the GORM database instance
func GetDB() *gorm.DB {
	return DB
}

// autoMigrate runs GORM auto-migration for all models
func autoMigrate() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.CVFile{},
	)
}