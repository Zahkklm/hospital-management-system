package utils

import (
	"regexp"
	"strings"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) IsPasswordStrong(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)

	return hasUpper && hasLower && hasNumber && hasSpecial
}

func (v *Validator) IsValidEmail(email string) bool {
	if email == "" {
		return false
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func (v *Validator) IsValidPhone(phone string) bool {
	if phone == "" {
		return false
	}

	// Remove common formatting characters
	cleaned := regexp.MustCompile(`[\s\-\(\)\+]`).ReplaceAllString(phone, "")

	// Check if it's all digits and reasonable length
	phoneRegex := regexp.MustCompile(`^\d{10,15}$`)
	return phoneRegex.MatchString(cleaned)
}

func (v *Validator) IsValidRole(role string) bool {
	validRoles := []string{"receptionist", "doctor"}
	role = strings.ToLower(strings.TrimSpace(role))

	for _, validRole := range validRoles {
		if role == validRole {
			return true
		}
	}
	return false
}

func (v *Validator) IsValidGender(gender string) bool {
	validGenders := []string{"male", "female", "other"}
	gender = strings.ToLower(strings.TrimSpace(gender))

	for _, validGender := range validGenders {
		if gender == validGender {
			return true
		}
	}
	return false
}

func (v *Validator) IsValidName(name string) bool {
	name = strings.TrimSpace(name)
	if len(name) == 0 || len(name) > 100 {
		return false
	}

	// Allow letters, spaces, hyphens, and apostrophes
	nameRegex := regexp.MustCompile(`^[a-zA-Z\s\-']+$`)
	return nameRegex.MatchString(name)
}
