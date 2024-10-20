package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Get("/natheneflix/", app.Home)

	mux.Post("/natheneflix/authenticate", app.authenticate)
	mux.Get("/natheneflix/refresh", app.refreshToken)
	mux.Get("/natheneflix/logout", app.logout)

	mux.Get("/natheneflix/movies", app.AllMovies)
	mux.Get("/natheneflix/movies/{id}", app.GetMovie)

	mux.Get("/natheneflix/genres", app.AllGenres)
	mux.Get("/natheneflix/movies/genres/{id}", app.AllMoviesByGenre)

	mux.Post("/natheneflix/graphql", app.moviesGraphql)

	mux.Route("/natheneflix/admin", func(mux chi.Router) {
		mux.Use(app.authRequired)

		mux.Get("/movies", app.MovieCatalogue)
		mux.Get("/movies/{id}", app.MovieForEdit)
		mux.Put("/movies/0", app.InsertMovie)
		mux.Patch("/movies/{id}", app.UpdateMovie)
		mux.Delete("/movies/{id}", app.DeleteMovie)
	})

	return mux
}
