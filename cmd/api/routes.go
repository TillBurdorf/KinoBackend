package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) route() http.Handler {
	r := mux.NewRouter()

	// Healthcheck
	r.HandleFunc("/v1/healthcheck", app.healthcheck).Methods("GET")

	// Movies
	r.HandleFunc("/v1/movies", app.getCreateMoviesHandler).Methods("GET", "POST")
	r.HandleFunc("/v1/movies/{id}", app.getUpdateDeleteMoviesHandler).Methods("GET", "DELETE", "PUT", "PATCH")

	// Users
	r.HandleFunc("/v1/users/register", app.registerUsersHandler).Methods("POST")
	r.HandleFunc("/v1/users/login", app.loginUsersHandler).Methods("POST")
	return r
}
