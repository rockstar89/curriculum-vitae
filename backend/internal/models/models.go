package models

// LoginRequest represents the login request payload
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token      string `json:"token"`
	Message    string `json:"message"`
	FirstLogin bool   `json:"first_login,omitempty"`
}

// ChangePasswordRequest represents the password change request payload
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
}

// ChangePasswordResponse represents the password change response
type ChangePasswordResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}