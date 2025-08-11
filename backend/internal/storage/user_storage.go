package storage

import (
	"cv-backend/internal/models"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserStorage handles GORM-based user operations
type UserStorage struct {
	db *gorm.DB
}

// NewUserStorage creates a new UserStorage instance
func NewUserStorage() *UserStorage {
	return &UserStorage{
		db: GetDB(),
	}
}

// InitializeDefaultUser creates the default admin user if it doesn't exist
func (us *UserStorage) InitializeDefaultUser(username, password string) error {
	var user models.User
	
	// Check if user already exists
	result := us.db.Where("username = ?", username).First(&user)
	if result.Error == nil {
		return nil // User already exists, nothing to do
	}
	
	if result.Error != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to check if user exists: %w", result.Error)
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Create the user
	newUser := models.User{
		Username:           username,
		PasswordHash:       string(hashedPassword),
		FirstLogin:         true,
		LoginCount:         0,
		LastPasswordChange: time.Now(),
		LastLoginAt:        time.Now(),
	}

	if err := us.db.Create(&newUser).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// ValidatePassword validates a user's password and updates login stats
func (us *UserStorage) ValidatePassword(username, password string) (*models.User, error) {
	var user models.User
	
	// Find user by username
	if err := us.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	// Update login stats
	updates := map[string]interface{}{
		"login_count":   gorm.Expr("login_count + ?", 1),
		"last_login_at": time.Now(),
		"updated_at":    time.Now(),
	}
	
	if err := us.db.Model(&user).Updates(updates).Error; err != nil {
		// Log the error but don't fail the login
		fmt.Printf("Warning: failed to update login stats: %v\n", err)
	}

	// Refresh user data to get updated login count
	us.db.Where("username = ?", username).First(&user)

	return &user, nil
}

// GetUser retrieves a user by username
func (us *UserStorage) GetUser(username string) (*models.User, error) {
	var user models.User
	
	if err := us.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// ChangePassword changes a user's password
func (us *UserStorage) ChangePassword(username, newPassword string) error {
	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Update password and mark as not first login
	result := us.db.Model(&models.User{}).
		Where("username = ?", username).
		Updates(map[string]interface{}{
			"password_hash":         string(hashedPassword),
			"first_login":          false,
			"last_password_change": time.Now(),
			"updated_at":           time.Now(),
		})

	if result.Error != nil {
		return fmt.Errorf("failed to update password: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}