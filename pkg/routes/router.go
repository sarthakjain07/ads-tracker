package routes

import (
	"ads-tracker/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/ads", func(r chi.Router) {
		r.Get("/", handlers.GetAds)
		r.Post("/click", handlers.PostClick)
		r.Get("/analytics", handlers.GetAnalytics)
	})

	return r
}
