package database

import (
	"github.com/TillBurdorf/Kinoprojekt/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Storage struct {
	Movies interface {
		Create(movie models.Movie) (error, string)
		Delete(id uuid.UUID) error
		Update(id uuid.UUID, movie models.Movie) error
		GetByID(id uuid.UUID) (models.Movie, error)
		GetAll() ([]models.Movie, error)
	}
	Users interface {
		Create(user models.User) (string, error)
		GetByUsername(username string) (models.User, error)
	}
}

func NewStorage(db *gorm.DB) Storage {
	return Storage{
		Movies: &MoviesStore{db},
		Users:  &UsersStore{db},
	}
}
