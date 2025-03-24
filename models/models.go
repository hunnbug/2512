package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Username string
	Password string
}

type ErrorResponse struct {
	Err     error
	Message string
}
