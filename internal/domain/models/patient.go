package models

import "time"

type Patient struct {
    ID        int       `json:"id" db:"id"`
    FirstName string    `json:"first_name" db:"first_name"`
    LastName  string    `json:"last_name" db:"last_name"`
    DOB       time.Time `json:"dob" db:"date_of_birth"`
    Gender    string    `json:"gender" db:"gender"`
    Phone     string    `json:"phone" db:"phone_number"`
    Email     string    `json:"email" db:"email"`
    Address   string    `json:"address" db:"address"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}