package config

import (
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DatabaseURL string
    JWTSecret   string
    Port        string
    Environment string
}

func LoadConfig() *Config {
    // Load .env file if it exists
    godotenv.Load()

    return &Config{
        DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:password@localhost:5432/hospital_management?sslmode=disable"),
        JWTSecret:   getEnv("JWT_SECRET", "asdj8123kdsavcilkdsamm129majksdIAnjdsaSM124"),
        Port:        getEnv("PORT", "8080"),
        Environment: getEnv("ENV", "development"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}