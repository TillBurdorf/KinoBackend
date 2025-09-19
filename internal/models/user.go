package models

import "github.com/google/uuid"

type User struct {
	Id     uuid.UUID `gorm:"type:uuid;primary_key;"`
	Username  string
	First_name string
	Last_name string
	Email string
	Password_hash string
}

type RegistrationRequest struct{
	Username string
	First_name string
	Last_name string
	Email string
	Password string
}

type LoginRequest struct{
	Username string
	Password string
}
