package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// Use the secret from config instead of hardcoded
var jwtSecret []byte

func SetJWTSecret(secret string) {
	jwtSecret = []byte(secret)
}

func GenerateToken(username, role string) (string, error) {
	if len(jwtSecret) == 0 {
		jwtSecret = []byte("asdj8123kdsavcilkdsamm129majksdIAnjdsaSM124") // fallback
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	if len(jwtSecret) == 0 {
		jwtSecret = []byte("asdj8123kdsavcilkdsamm129majksdIAnjdsaSM124") // fallback
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
