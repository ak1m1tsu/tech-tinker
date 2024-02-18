package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/insan1a/tech-tinker/internal/delivery/http/controllers/account"
	"github.com/insan1a/tech-tinker/internal/delivery/http/controllers/auth"
	"github.com/insan1a/tech-tinker/internal/delivery/http/controllers/handlers"
	"github.com/insan1a/tech-tinker/internal/delivery/http/middleware/jsonlogger"
	"github.com/insan1a/tech-tinker/internal/domain/interfaces"
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

func (r *Router) MountAccountRoutes() {
	controller := account.New()

	r.Route("/account", func(r chi.Router) {
		r.Get("/", controller.HandleAccountInfo)
		r.Post("/stat", controller.HandleAccountStatistic)
		r.Route("/orders", func(r chi.Router) {
			r.Get("/", controller.HandleAccountOrders)
			r.Get("/{orderID}", controller.HandleAccountOrder)
		})
	})
}

func (r *Router) MountAuthRoutes(service interfaces.AuthService) {
	controller := auth.New(service)

	r.Post("/token", controller.HandleAuthToken)
}
