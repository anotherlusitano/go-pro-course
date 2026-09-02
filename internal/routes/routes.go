package routes

import (
	"github.com/anotherlusitano/goProject/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	return r
}
