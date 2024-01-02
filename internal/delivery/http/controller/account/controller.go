package account

import (
	"github.com/go-chi/chi/v5"
	"github.com/insan1a/tech-tinker/internal/lib/response"
	"net/http"
)

type controller struct{}

func MountRoutes(router chi.Router) {
	c := &controller{}

	router.Route("/account", func(r chi.Router) {
		r.Get("/", c.HandleAccountInfo)
		r.Route("/orders", func(r chi.Router) {
			r.Get("/", c.HandleAccountOrders)
			r.Get("/{orderID}", c.HandleAccountOrder)
		})
		r.Post("/statistics", c.HandleAccountStatistic)
	})
}

func (c *controller) HandleAccountInfo(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"account": response.M{},
		},
	})
}

func (c *controller) HandleAccountOrders(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"orders": []response.M{},
		},
	})
}

func (c *controller) HandleAccountOrder(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"order": response.M{},
		},
	})
}

func (c *controller) HandleAccountStatistic(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"statistic": response.M{},
		},
	})
}
