package main

import (
	"cv-backend/internal/handlers"
	"cv-backend/internal/middleware"
	"cv-backend/internal/storage"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	if err := storage.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	log.Println("✅ Database initialized")

	// Initialize default user
	userStorage := storage.NewUserStorage()
	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminUsername == "" {
		adminUsername = "admin"
	}
	if adminPassword == "" {
		adminPassword = "cvadmin2024"
	}
	
	if err := userStorage.InitializeDefaultUser(adminUsername, adminPassword); err != nil {
		log.Printf("Warning: Failed to initialize default user: %v", err)
	}

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Gin router
	router := gin.Default()

	// Apply CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Initialize handlers
	authHandler := handlers.NewAuthHandler()
	cvHandler := handlers.NewCVHandler()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "cv-backend"})
	})

	// Public routes
	api := router.Group("/api")
	{
		// Authentication
		api.POST("/login", authHandler.Login)
		
		// CV download and viewing (public)
		api.GET("/download-cv", cvHandler.DownloadCV)
		api.GET("/view-cv", cvHandler.ViewCV)
	}

	// Protected routes (require authentication)
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Auth verification
		protected.GET("/verify", authHandler.VerifyToken)
		
		// Password management
		protected.PUT("/change-password", authHandler.ChangePassword)
		
		// CV management
		protected.POST("/upload-cv", cvHandler.UploadCV)
		protected.GET("/cv-info", cvHandler.GetCVInfo)
		protected.GET("/cv-stats", cvHandler.GetStats)
		protected.DELETE("/cv", cvHandler.DeleteCV)
	}

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("💾 File storage: data/ directory")
	log.Printf("🔐 Admin username: %s", adminUsername)
	
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}