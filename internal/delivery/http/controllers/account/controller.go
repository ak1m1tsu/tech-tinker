package account

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/insan1a/tech-tinker/internal/domain/interfaces"
	"github.com/insan1a/tech-tinker/internal/lib/appcontext"
	"github.com/insan1a/tech-tinker/internal/lib/response"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Controller struct {
	service interfaces.AccountService
}

func New(service interfaces.AccountService) *Controller {
	return &Controller{
		service: service,
	}
}

type AccountInfoResponse struct {
	ID        string                `json:"id"`
	FirstName string                `json:"first_name"`
	LastName  string                `json:"last_name"`
	Email     string                `json:"email"`
	Role      string                `json:"role"`
	CreatedAt string                `json:"created_at"`
	Orders    AccountOrderResponses `json:"orders,omitempty"`
}

type AccountOrderResponses []AccountOrderResponse

type AccountOrderResponse struct {
	ID         string `json:"id"`
	Number     int    `json:"number"`
	PriceLimit int    `json:"price_limit"`
	Comment    string `json:"comment"`
	Address    string `json:"address"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
}

func (c *Controller) HandleAccountInfo(w http.ResponseWriter, r *http.Request) {
	log := logrus.WithFields(logrus.Fields{
		"request_id": middleware.GetReqID(r.Context()),
		"real_ip":    r.RemoteAddr,
		"uri":        r.RequestURI,
		"method":     r.Method,
	})

	employeeID := appcontext.GetEmployeeID(r.Context())

	account, err := c.service.GetAccount(r.Context(), employeeID)
	if err != nil {
		log.WithError(err).Error("failed to find user by id")

		response.InternalServerError(w)

		return
	}

	var aor AccountOrderResponses
	if account.Orders != nil {
		aor = make(AccountOrderResponses, 0, len(account.Orders))
		for _, order := range account.Orders {
			aor = append(aor, AccountOrderResponse{
				ID:         order.ID,
				Number:     order.Number,
				PriceLimit: order.PriceLimit,
				Comment:    order.Comment,
				Address:    order.Address,
				Status:     order.Status.String(),
				CreatedAt:  order.CreatedAt.Format(time.RFC3339),
			})
		}
	}

	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"account": AccountInfoResponse{
				ID:        account.ID,
				FirstName: account.FirstName,
				LastName:  account.LastName,
				Email:     account.Email,
				Role:      account.Role.String(),
				CreatedAt: account.CreatedAt.Format(time.RFC3339),
				Orders:    aor,
			},
		},
	})
}

func (c *Controller) HandleAccountOrders(w http.ResponseWriter, r *http.Request) {
	log := logrus.WithFields(logrus.Fields{
		"request_id": middleware.GetReqID(r.Context()),
		"real_ip":    r.RemoteAddr,
		"uri":        r.RequestURI,
		"method":     r.Method,
	})

	employeeID := appcontext.GetEmployeeID(r.Context())

	orders, err := c.service.GetOrders(r.Context(), employeeID)
	if err != nil {
		log.WithError(err).Error("failed to find user by id")

		response.InternalServerError(w)

		return
	}

	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"orders": orders,
		},
	})
}

func (c *Controller) HandleAccountOrder(w http.ResponseWriter, r *http.Request) {
	log := logrus.WithFields(logrus.Fields{
		"request_id": middleware.GetReqID(r.Context()),
		"real_ip":    r.RemoteAddr,
		"uri":        r.RequestURI,
		"method":     r.Method,
	})

	orderID := chi.URLParam(r, "orderID")
	order, err := c.service.GetOrder(r.Context(), orderID)
	if err != nil {
		log.WithError(err).Error("failed to find user by id")

		response.InternalServerError(w)

		return
	}

	if order == nil {
		response.NotFound(w)

		return
	}

	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"order": order,
		},
	})
}

func (c *Controller) HandleAccountStatistic(w http.ResponseWriter, r *http.Request) {
	log := logrus.WithFields(logrus.Fields{
		"request_id": middleware.GetReqID(r.Context()),
		"real_ip":    r.RemoteAddr,
		"uri":        r.RequestURI,
		"method":     r.Method,
	})

	employeeID := appcontext.GetEmployeeID(r.Context())

	statistic, err := c.service.GetStatistic(r.Context(), employeeID)
	if err != nil {
		log.WithError(err).Error("failed to find user by id")

		response.InternalServerError(w)

		return
	}

	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"statistic": statistic,
		},
	})
}
