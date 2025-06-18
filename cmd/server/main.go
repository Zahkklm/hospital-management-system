package main

import (
	"log"

	"hospital-management-system/internal/api/middleware"
	"hospital-management-system/internal/api/routes"
	"hospital-management-system/internal/config"
	"hospital-management-system/internal/infrastructure/database"
	"hospital-management-system/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Set JWT secret
	utils.SetJWTSecret(cfg.JWTSecret)

	// Initialize database connection
	database.Connect()
	db := database.GetDB()
	defer db.Close()

	// Set up Gin router
	router := gin.Default()

	// Add middleware
	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Serve static files
	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("web/templates/*")

	// Set up routes
	routes.SetupRoutes(router)

	// Start the server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
