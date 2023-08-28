package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/vtrenton/go_api/internal/middleware"
)

// Since the function started with an upper case H is is now public
// functions that start with a lower case are private
func Handler(r *chi.Mux) {
	// Global middleware
	// This will strip the trailing slash at the end of the URL
	// This way localhost:8000/myapp/ -> localhost:8000/myapp
	r.Use(chimiddle.StripSlashes)
	// Set up a Route
	r.Route("/account", func(router chi.Router) {
		// Middleware to check Authorization
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
