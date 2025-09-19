package models

import "github.com/google/uuid"

type Movie struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title  string
	Omdbid string
}
