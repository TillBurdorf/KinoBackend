package main

import (
	//public modules
	"encoding/json"
	"net/http"

	//internal modules
	"github.com/TillBurdorf/Kinoprojekt/internal/models"
)

// Healthcheck-Handler
func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	//Example how to convert data into JSON and return it as a JSON object
	data := models.HealthcheckResponse{
		Status:      "available",
		Environment: app.config.env,
		Version:     version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}