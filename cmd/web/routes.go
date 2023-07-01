package main

import (
	"github.com/MelihEmreGuler/web-content-management/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// routes sets the routes for the application
func routes() http.Handler {
	mux := chi.NewRouter()

	// middlewares are executed in the order they are defined
	mux.Use(middleware.Recoverer) // recovering from panics
	mux.Use(WriteToConsole)       // write to console
	mux.Use(NoSurf)               // no csrf attacks
	mux.Use(SessionLoad)          // load and save session on every request

	mux.Get("/home", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux

	/*mux := pat.New()
	mux.Get("/home", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux*/
}
