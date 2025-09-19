package services

import (
	"errors"

	"github.com/TillBurdorf/Kinoprojekt/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func CalculateUserFromRegistrationRequest(request models.RegistrationRequest) (models.User, error) {
	var u models.User
	u.Username = request.Username
	u.First_name = request.First_name
	u.Last_name = request.Last_name
	u.Email = request.Email
	password_hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, errors.New("error during hash process")
	}
	u.Password_hash = string(password_hash)
	return u, nil
}

func Login(user models.User, password string) error {
	err := checkPasswordHash([]byte(user.Password_hash), password)
	if err != nil {
		return errors.New("Login Unsuccessfull")
	}
	return nil
}

func checkPasswordHash(password_hash []byte, password string) error {
	err := bcrypt.CompareHashAndPassword(password_hash, []byte(password))
	return err
}
