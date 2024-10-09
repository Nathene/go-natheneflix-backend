package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *app) routes() http.Handler {
	mux := mux.NewRouter()

	mux.Use(recoveryMiddleware)
	mux.Use(a.enableCORS)

	mux.HandleFunc("/", a.handleHome)

	mux.HandleFunc("/movies", a.handleAllMovies)

	return mux
}

// recoveryMiddleware recovers from panics, logs the error, and returns a 500 status
// Chi has this feature, but for maintainence reasons i chose Mux
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %+v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
