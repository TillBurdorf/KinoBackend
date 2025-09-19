package database

import (
	"fmt"

	"github.com/TillBurdorf/Kinoprojekt/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MoviesStore struct {
	db *gorm.DB
}

func (s *MoviesStore) Create(movie models.Movie) (error, string) {
	id := uuid.New()
	movie.ID = id

	result := s.db.Create(&movie)
	if result.Error != nil {
		return result.Error, ""
	}

	return nil, id.String()
}

func (s *MoviesStore) Delete(id uuid.UUID) error {
	result := s.db.Delete(&models.Movie{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no movie found with id %s", id)
	}

	return nil
}

func (s *MoviesStore) Update(id uuid.UUID, movie models.Movie) error {
	return nil
}

func (s *MoviesStore) GetByID(id uuid.UUID) (models.Movie, error) {
	var movie models.Movie

	result := s.db.First(&movie, "id = ?", id)

	if result.Error != nil {
		return models.Movie{}, result.Error
	}

	return movie, nil
}

func (s *MoviesStore) GetAll() ([]models.Movie, error) {
	var movies []models.Movie
	result := s.db.Find(&movies)

	if result.Error != nil {
		return nil, result.Error
	}

	return movies, nil
}
