package utils_test

import (
	"testing"
	"time"

	"hospital-management-system/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	// Setup
	utils.SetJWTSecret("test-secret-key")
	username := "testuser"
	role := "receptionist"

	// Execute
	token, err := utils.GenerateToken(username, role)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken(t *testing.T) {
	// Setup
	utils.SetJWTSecret("test-secret-key")
	username := "testuser"
	role := "receptionist"

	// Generate token
	token, err := utils.GenerateToken(username, role)
	assert.NoError(t, err)

	// Execute
	claims, err := utils.ValidateToken(token)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, username, claims.Username)
	assert.Equal(t, role, claims.Role)
}

func TestValidateInvalidToken(t *testing.T) {
	// Setup
	utils.SetJWTSecret("test-secret-key")
	invalidToken := "invalid.token.here"

	// Execute
	claims, err := utils.ValidateToken(invalidToken)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestGenerateTokenWithUserID(t *testing.T) {
	// Setup
	utils.SetJWTSecret("test-secret-key")
	userID := 123
	username := "testuser"
	role := "doctor"

	// Execute
	token, err := utils.GenerateTokenWithUserID(userID, username, role)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateJWTWithUserID(t *testing.T) {
	// Setup
	utils.SetJWTSecret("test-secret-key")
	userID := 123
	username := "testuser"
	role := "doctor"

	// Generate token
	token, err := utils.GenerateTokenWithUserID(userID, username, role)
	assert.NoError(t, err)

	// Execute
	extractedUserID, extractedUsername, extractedRole, err := utils.ValidateJWTWithUserID(token)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, userID, extractedUserID)
	assert.Equal(t, username, extractedUsername)
	assert.Equal(t, role, extractedRole)
}

func TestValidateJWT(t *testing.T) {
	// Setup
	utils.SetJWTSecret("test-secret-key")
	username := "testuser"
	role := "receptionist"

	// Generate token
	token, err := utils.GenerateToken(username, role)
	assert.NoError(t, err)

	// Execute
	extractedUsername, extractedRole, err := utils.ValidateJWT(token)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, username, extractedUsername)
	assert.Equal(t, role, extractedRole)
}

func TestTokenExpiration(t *testing.T) {
	// This would test token expiration logic
	// For now, we test that tokens have proper expiration claims
	utils.SetJWTSecret("test-secret-key")
	token, err := utils.GenerateToken("testuser", "receptionist")
	assert.NoError(t, err)

	claims, err := utils.ValidateToken(token)
	assert.NoError(t, err)

	// Token should expire in 24 hours (86400 seconds)
	expectedExpiry := time.Now().Add(24 * time.Hour).Unix()
	assert.InDelta(t, expectedExpiry, claims.ExpiresAt, 60) // Allow 60 second difference
}
