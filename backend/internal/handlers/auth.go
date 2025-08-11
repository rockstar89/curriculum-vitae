package handlers

import (
	"cv-backend/internal/auth"
	"cv-backend/internal/models"
	"cv-backend/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{
	userStorage *storage.UserStorage
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userStorage: storage.NewUserStorage(),
	}
}

// Login handles admin login
func (h *AuthHandler) Login(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Validate credentials using file-based storage
	user, err := h.userStorage.ValidatePassword(loginReq.Username, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(loginReq.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, models.LoginResponse{
		Token:      token,
		Message:    "Login successful",
		FirstLogin: user.FirstLogin,
	})
}

// VerifyToken checks if the provided token is valid
func (h *AuthHandler) VerifyToken(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get user to check if first login
	user, err := h.userStorage.GetUser(username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"valid":       true,
		"username":    username,
		"message":     "Token is valid",
		"first_login": user.FirstLogin,
	})
}

// ChangePassword handles password change requests
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var changeReq models.ChangePasswordRequest
	if err := c.ShouldBindJSON(&changeReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Verify current password
	_, err := h.userStorage.ValidatePassword(username.(string), changeReq.CurrentPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Current password is incorrect"})
		return
	}

	// Validate new password strength (basic validation)
	if len(changeReq.NewPassword) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password must be at least 8 characters long"})
		return
	}

	// Change password
	if err := h.userStorage.ChangePassword(username.(string), changeReq.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change password"})
		return
	}

	c.JSON(http.StatusOK, models.ChangePasswordResponse{
		Success: true,
		Message: "Password changed successfully",
	})
}