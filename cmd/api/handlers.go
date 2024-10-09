package main

import (
	"backend/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"time"
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

	out, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (a *app) handleAllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "2003-01-02")

	matrix := models.Movie{
		ID:          1,
		Title:       "Matrix",
		ReleaseDate: rd,
		MPAARating:  "M",
		RunTime:     154,
		Description: "Matrix Movie!",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	rd, _ = time.Parse("2006-01-02", "2006-02-05")
	matrix2 := models.Movie{
		ID:          2,
		Title:       "Matrix 2",
		ReleaseDate: rd,
		MPAARating:  "M",
		RunTime:     132,
		Description: "Matrix Movie! (Number two!)",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, matrix, matrix2)

	out, err := json.Marshal(movies)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
