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

// Movies-Handler
func (app *application) getCreateMoviesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		movies, err := app.storage.Movies.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movies)

	case http.MethodPost:
		var m models.Movie
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err, id := app.storage.Movies.Create(m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Movie with ID: %s added.", id)
	}
}

func (app *application) getUpdateDeleteMoviesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		id, err := services.GetIdFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		m, err := app.storage.Movies.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(m)

	case http.MethodDelete:
		id, err := services.GetIdFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = app.storage.Movies.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Movie Deleted Successfully")
	}
}