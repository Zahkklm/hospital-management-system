package utils

import (
    "regexp"
    "unicode"
)

type Validator struct{}

func NewValidator() *Validator {
    return &Validator{}
}

// IsPasswordStrong checks if password meets strength requirements
func (v *Validator) IsPasswordStrong(password string) bool {
    if len(password) < 8 {
        return false
    }

    var hasUpper, hasLower, hasNumber, hasSpecial bool

    for _, char := range password {
        switch {
        case unicode.IsUpper(char):
            hasUpper = true
        case unicode.IsLower(char):
            hasLower = true
        case unicode.IsDigit(char):
            hasNumber = true
        case unicode.IsPunct(char) || unicode.IsSymbol(char):
            hasSpecial = true
        }
    }

    return hasUpper && hasLower && hasNumber && hasSpecial
}

// IsValidEmail checks if email format is valid
func (v *Validator) IsValidEmail(email string) bool {
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return emailRegex.MatchString(email)
}