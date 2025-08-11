package storage

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the database
type User struct {
	ID                 int       `json:"id"`
	Username           string    `json:"username"`
	PasswordHash       string    `json:"password_hash"`
	FirstLogin         bool      `json:"first_login"`
	CreatedAt          time.Time `json:"created_at"`
	LastPasswordChange time.Time `json:"last_password_change"`
	LoginCount         int       `json:"login_count"`
	LastLoginAt        time.Time `json:"last_login_at"`
}

// UserStorage handles database-based user operations
type UserStorage struct {
	db *sql.DB
}

// NewUserStorage creates a new UserStorage instance
func NewUserStorage() *UserStorage {
	return &UserStorage{
		db: GetDB(),
	}
}

// InitializeDefaultUser creates the default admin user if it doesn't exist
func (us *UserStorage) InitializeDefaultUser(username, password string) error {
	// Check if user already exists
	var exists bool
	err := us.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}

	if exists {
		return nil // User already exists, nothing to do
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Create the user
	_, err = us.db.Exec(`
		INSERT INTO users (username, password_hash, first_login, created_at, last_password_change, login_count, last_login_at)
		VALUES ($1, $2, true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0, CURRENT_TIMESTAMP)
	`, username, string(hashedPassword))

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// ValidatePassword validates a user's password and updates login stats
func (us *UserStorage) ValidatePassword(username, password string) (*User, error) {
	var user User
	err := us.db.QueryRow(`
		SELECT id, username, password_hash, first_login, created_at, last_password_change, login_count, last_login_at
		FROM users 
		WHERE username = $1
	`, username).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.FirstLogin,
		&user.CreatedAt, &user.LastPasswordChange, &user.LoginCount, &user.LastLoginAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	// Update login stats
	_, err = us.db.Exec(`
		UPDATE users 
		SET login_count = login_count + 1, last_login_at = CURRENT_TIMESTAMP
		WHERE username = $1
	`, username)
	if err != nil {
		// Log the error but don't fail the login
		fmt.Printf("Warning: failed to update login stats: %v\n", err)
	}

	return &user, nil
}

// GetUser retrieves a user by username
func (us *UserStorage) GetUser(username string) (*User, error) {
	var user User
	err := us.db.QueryRow(`
		SELECT id, username, password_hash, first_login, created_at, last_password_change, login_count, last_login_at
		FROM users 
		WHERE username = $1
	`, username).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.FirstLogin,
		&user.CreatedAt, &user.LastPasswordChange, &user.LoginCount, &user.LastLoginAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
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
	result, err := us.db.Exec(`
		UPDATE users 
		SET password_hash = $1, first_login = false, last_password_change = CURRENT_TIMESTAMP
		WHERE username = $2
	`, string(hashedPassword), username)

	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}