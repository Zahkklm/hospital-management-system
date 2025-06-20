package routes

import (
	"hospital-management-system/internal/api/handlers"
	"hospital-management-system/internal/api/middleware"
	"hospital-management-system/internal/config"
	"hospital-management-system/internal/infrastructure/database"
	"hospital-management-system/internal/infrastructure/repository"
	"hospital-management-system/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// Initialize repositories
	db := database.GetDB()
	userRepo := repository.NewUserRepository(db)
	patientRepo := repository.NewPatientRepository(db)

	// Initialize services
	cfg := config.LoadConfig()
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	userService := services.NewUserService(userRepo)
	patientService := services.NewPatientService(patientRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService, userService)
	userHandler := handlers.NewUserHandler(userService)
	patientHandler := handlers.NewPatientHandler(patientService)

	// Public routes
	router.GET("/", authHandler.ShowLoginPage)
	router.GET("/login", authHandler.ShowLoginPage)
	router.GET("/register", authHandler.ShowRegisterPage)
	router.POST("/api/auth/login", authHandler.Login)
	router.POST("/api/auth/register", authHandler.Register)
	router.GET("/dashboard", authHandler.ShowDashboard)
	router.GET("/api/dashboard", authHandler.ShowDashboard)

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/logout", authHandler.Logout)

		// Patient routes
		api.GET("/patients", patientHandler.GetAllPatients)
		api.POST("/patients", patientHandler.CreatePatient)
		api.GET("/patients/:id", patientHandler.GetPatient)
		api.PUT("/patients/:id", patientHandler.UpdatePatient)
		api.DELETE("/patients/:id", patientHandler.DeletePatient)

		// User routes
		api.GET("/users/:id", userHandler.GetUser)
		api.PUT("/users/:id", userHandler.UpdateUser)
	}
}
