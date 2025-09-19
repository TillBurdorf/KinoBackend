package main

import (
	//public modules
	"encoding/json"
	"fmt"
	_ "fmt"
	"net/http"

	//internal modules
	"github.com/TillBurdorf/Kinoprojekt/internal/models"
	"github.com/TillBurdorf/Kinoprojekt/internal/services"

)

// Users-Handler

//encore:api public method=GET path=/v1/movies
func (app *application) registerUsersHandler(w http.ResponseWriter, r *http.Request) {
	var request models.RegistrationRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.CalculateUserFromRegistrationRequest(request)

	id, err := app.storage.Users.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User with ID: %s added.", id)
}

func (app *application) loginUsersHandler(w http.ResponseWriter, r *http.Request) {
	var request models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := app.storage.Users.GetByUsername(request.Username)
	if(err != nil){
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	//login user
	err = services.Login(user, request.Password)
	if(err != nil){
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "found user %s", user.Username)
}
