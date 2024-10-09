package main

import (
	"log"
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

func (a *app) handleAuthenticate(w http.ResponseWriter, r *http.Request) {
	// read json payload

	// validate user against database

	// check password

	// create a jwt user
	u := jwtUser{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
	}

	// generate tokens
	tokens, err := a.auth.GenerateTokenPair(&u)
	if err != nil {
		a.errorJSON(w, err)
		return
	}
	log.Println(tokens.Token)
	a.writeJSON(w, 200, tokens)
}
