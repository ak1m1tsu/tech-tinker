package account

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/insan1a/tech-tinker/internal/domain/model"
	"github.com/insan1a/tech-tinker/internal/lib/response"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) HandleAccountInfo(w http.ResponseWriter, r *http.Request) {
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

func (c *Controller) HandleAccountOrders(w http.ResponseWriter, r *http.Request) {
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

func (c *Controller) HandleAccountOrder(w http.ResponseWriter, r *http.Request) {
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

func (c *Controller) HandleAccountStatistic(w http.ResponseWriter, r *http.Request) {
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
