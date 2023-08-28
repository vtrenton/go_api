package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	// "github.com/vtrenton/go_api/internal/middleware"
)

// Since the function started with an upper case H is is now public
// functions that start with a lower case are private
func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)
}
