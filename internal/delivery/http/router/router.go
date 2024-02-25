package router

import (
	"github.com/ak1m1tsu/tech-tinker/internal/delivery/http/controllers/handlers"
	"github.com/ak1m1tsu/tech-tinker/internal/delivery/http/middleware/jsonlogger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Router struct {
	chi.Router
}

func New() *Router {
	r := &Router{chi.NewMux()}

	r.NotFound(handlers.NotFoundHandler)
	r.MethodNotAllowed(handlers.MethodNotAllowedHandler)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(jsonlogger.New)
	r.Use(middleware.Recoverer)

	return r
}
