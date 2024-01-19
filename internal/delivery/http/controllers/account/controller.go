package account

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/insan1a/tech-tinker/internal/delivery/http/controllers"
	"github.com/insan1a/tech-tinker/internal/delivery/http/middleware/jwtvalidation"
	"github.com/insan1a/tech-tinker/internal/domain/model"
	"github.com/insan1a/tech-tinker/internal/lib/response"
	"net/http"
	"time"
)

type controller struct {
	config *Config
}

func newController(cfg *Config) (*controller, error) {
	if cfg == nil {
		return nil, controllers.ErrConfigMissing
	}

	return &controller{
		config: cfg,
	}, nil
}

func MountRoutes(cfg *Config, router chi.Router) error {
	c, err := newController(cfg)
	if err != nil {
		return err
	}

	router.With(
		jwtvalidation.New(c.config.rsaPubKey),
	).Route(
		"/account",
		func(r chi.Router) {
			r.Get("/", c.HandleAccountInfo)
			r.Route("/orders", func(r chi.Router) {
				r.Get("/", c.HandleAccountOrders)
				r.Get("/{orderID}", c.HandleAccountOrder)
			})
			r.Post("/statistics", c.HandleAccountStatistic)
		},
	)
	return nil
}

func (c *controller) HandleAccountInfo(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"account": response.M{
				"id":         uuid.NewString(),
				"first_name": "Ivan",
				"last_name":  "Ivanov",
				"email":      "ivan.ivanov@techtinker.com",
				"role":       model.EmployeeRoleTechnician.String(),
				"orders":     []response.M{},
				"created_at": "2020-01-01T00:00:00Z",
			},
		},
	})
}

func (c *controller) HandleAccountOrders(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"orders": []response.M{
				{
					"id":          uuid.NewString(),
					"number":      1,
					"price_limit": 15000000,
					"comment":     "this is very long comment",
					"address":     "Irkutsk, Donskaya Street, 4, 21",
					"status":      model.OrderStatusInProcess.String(),
					"customer": response.M{
						"id":           uuid.NewString(),
						"first_name":   "Ivan",
						"last_name":    "Ivanov",
						"email":        "ivan.ivanov@gmail.com",
						"phone_number": "78005553535",
					},
					"created_at": time.Now().Format(time.RFC3339),
				},
				{
					"id":          uuid.NewString(),
					"number":      2,
					"price_limit": 5000000,
					"comment":     "this is very long comment",
					"address":     "Moscow, Lenin Street, 14, 5",
					"status":      model.OrderStatusInProcess.String(),
					"customer": response.M{
						"id":           uuid.NewString(),
						"first_name":   "Roman",
						"last_name":    "Ivanov",
						"email":        "roman.ivanov@gmail.com",
						"phone_number": "78005553534",
					},
					"created_at": time.Now().Format(time.RFC3339),
				},
				{
					"id":          uuid.NewString(),
					"number":      3,
					"price_limit": 25000000,
					"comment":     "this is very long comment",
					"address":     "Novosibirsk, Pushkin Street, 6, 85",
					"status":      model.OrderStatusInProcess.String(),
					"customer": response.M{
						"id":           uuid.NewString(),
						"first_name":   "Vasiliy",
						"last_name":    "Ivanov",
						"email":        "vasiliy.ivanov@gmail.com",
						"phone_number": "78005553533",
					},
					"created_at": time.Now().Format(time.RFC3339),
				},
			},
		},
	})
}

func (c *controller) HandleAccountOrder(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"order": response.M{
				"id":          uuid.NewString(),
				"number":      1,
				"price_limit": 15000000,
				"comment":     "this is very long comment",
				"address":     "Irkutsk, Donskaya Street, 4, 21",
				"status":      model.OrderStatusInProcess.String(),
				"customer": response.M{
					"id":           uuid.NewString(),
					"first_name":   "Ivan",
					"last_name":    "Ivanov",
					"email":        "ivan.ivanov@gmail.com",
					"phone_number": "78005553535",
				},
				"created_at": time.Now().Format(time.RFC3339),
				"configurations": []response.M{
					{
						"id":         uuid.NewString(),
						"price":      15000000,
						"created_at": time.Now().Format(time.RFC3339),
					},
					{
						"id":         uuid.NewString(),
						"price":      14900000,
						"created_at": time.Now().Format(time.RFC3339),
					},
					{
						"id":         uuid.NewString(),
						"price":      14500000,
						"created_at": time.Now().Format(time.RFC3339),
					},
				},
			},
		},
	})
}

func (c *controller) HandleAccountStatistic(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"statistic": response.M{
				"from":  time.Now().Format(time.RFC3339),
				"to":    time.Now().Add(time.Hour * 24 * 30).Format(time.RFC3339),
				"total": 182000000,
				"budgets": []response.M{
					{
						"count": 6,
						"type":  model.BudgetTypeLowerThan50K.String(),
					},
					{
						"count": 2,
						"type":  model.BudgetTypeBetween50KAnd100K.String(),
					},
					{
						"count": 2,
						"type":  model.BudgetTypeBetween100KAnd500K.String(),
					},
					{
						"count": 1,
						"type":  model.BudgetTypeGreaterThan500K.String(),
					},
				},
			},
		},
	})
}
