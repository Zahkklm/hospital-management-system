package utils_test

import (
	"testing"

	"hospital-management-system/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword123"

	// Execute
	hash, err := utils.HashPassword(password)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, password, hash)
	assert.Greater(t, len(hash), 10) // Bcrypt hashes are longer than 10 chars
}

func TestCheckPasswordHash_ValidPassword(t *testing.T) {
	password := "testpassword123"
	hash, err := utils.HashPassword(password)
	assert.NoError(t, err)

	// Execute
	isValid := utils.CheckPasswordHash(password, hash)

	// Assert
	assert.True(t, isValid)
}

func TestCheckPasswordHash_InvalidPassword(t *testing.T) {
	password := "testpassword123"
	wrongPassword := "wrongpassword"
	hash, err := utils.HashPassword(password)
	assert.NoError(t, err)

	// Execute
	isValid := utils.CheckPasswordHash(wrongPassword, hash)

	// Assert
	assert.False(t, isValid)
}

func TestHashPassword_EmptyPassword(t *testing.T) {
	password := ""

	// Execute
	hash, err := utils.HashPassword(password)

	// Assert - Should handle empty password gracefully
	if err != nil {
		assert.Error(t, err)
		assert.Empty(t, hash)
	} else {
		assert.NotEmpty(t, hash)
	}
}

func TestCheckPasswordHash_EmptyInputs(t *testing.T) {
	// Test with empty password
	isValid := utils.CheckPasswordHash("", "somehash")
	assert.False(t, isValid)

	// Test with empty hash
	isValid = utils.CheckPasswordHash("password", "")
	assert.False(t, isValid)

	// Test with both empty
	isValid = utils.CheckPasswordHash("", "")
	assert.False(t, isValid)
}

func TestPasswordHashConsistency(t *testing.T) {
	password := "consistencytest123"

	// Generate multiple hashes of the same password
	hash1, err1 := utils.HashPassword(password)
	hash2, err2 := utils.HashPassword(password)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotEqual(t, hash1, hash2) // Should be different due to salt

	// But both should validate the same password
	assert.True(t, utils.CheckPasswordHash(password, hash1))
	assert.True(t, utils.CheckPasswordHash(password, hash2))
}
