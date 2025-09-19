package main

import "context"

// GetMovie holt einen einzelnen Film anhand seiner ID.
//encore:api public method=GET path=/v1/movies/:id
func GetMovie(ctx context.Context, id int) (string, error) {
    // Logik, um einen Film mit der gegebenen ID zu finden...
    return "test", nil
}

// DeleteMovie löscht einen Film.
//encore:api public method=DELETE path=/v1/movies/:id
func DeleteMovie(ctx context.Context, id int) error {
    // Logik zum Löschen des Films mit der ID...
    return nil
}