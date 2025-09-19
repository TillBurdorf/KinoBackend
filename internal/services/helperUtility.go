package services

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetIdFromRequest(r *http.Request) (uuid.UUID, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok || idStr == "" {
		return uuid.Nil, errors.New("id parameter is missing")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}