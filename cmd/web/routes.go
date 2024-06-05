package main

import (
	"github.com/RazikaBengana/Go-BnB/pkg/config"
	"github.com/RazikaBengana/Go-BnB/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

// routes sets up the application's routing using the chi router and handlers
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) // Add middleware to recover from panics and return a 500 error
	mux.Use(NoSurf)               // Add CSRF protection middleware to all routes
	mux.Use(SessionLoad)          // Load and save the session on every request

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
