package main

import (
	"net/http"
)

func (a *app) handleHome(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Natheneflix up and running",
		Version: "1.0.0",
	}

	a.writeJSON(w, http.StatusOK, payload)
}

func (a *app) handleAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := a.DB.AllMovies()
	if err != nil {
		a.errorJSON(w, err)
	}
	a.writeJSON(w, http.StatusOK, movies)
}
