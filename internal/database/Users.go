package database

import (
	_ "fmt"

	"github.com/TillBurdorf/Kinoprojekt/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UsersStore struct {
	db *gorm.DB
}

func (s *UsersStore) Create(user models.User) (string, error) {
	user.Id = uuid.New()

	result := s.db.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.Id.String(), nil
}

func (s *UsersStore) GetByUsername(username string) (models.User, error) {
	var user models.User
	result := s.db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
