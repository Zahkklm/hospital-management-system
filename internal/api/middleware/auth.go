package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "hospital-management-system/pkg/utils"
)

// AuthMiddleware is a middleware function that checks for a valid JWT token in the request header.
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.Request.Header.Get("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
            c.Abort()
            return
        }

        // Remove "Bearer " prefix if present
        if strings.HasPrefix(tokenString, "Bearer ") {
            tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        }

        // Validate the token
        claims, err := utils.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Set the user information in the context
        c.Set("userID", claims.Username)
        c.Set("role", claims.Role)

        c.Next()
    }
}