package account

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/insan1a/tech-tinker/internal/domain/interfaces"
	"github.com/insan1a/tech-tinker/internal/lib/appcontext"
	"github.com/insan1a/tech-tinker/internal/lib/response"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	service interfaces.AccountService
}

func New(service interfaces.AccountService) *Controller {
	return &Controller{service: service}
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
		log.WithError(err).Errorf("failed to find user by id [%s]", employeeID)

		response.InternalServerError(w)

		return
	}

	var aor OrderResponses
	if account.Orders != nil {
		aor = make(OrderResponses, 0, len(account.Orders))
		for _, order := range account.Orders {
			aor = append(aor, OrderResponse{
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
			"account": FillInfo(account),
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
		log.WithError(err).Errorf("failed to find orders for user [%s]", employeeID)

		response.InternalServerError(w)

		return
	}

	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"orders": FillOrders(orders),
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

	employeeID := appcontext.GetEmployeeID(r.Context())
	orderID := chi.URLParam(r, "orderID")

	if _, err := uuid.Parse(orderID); err != nil {
		log.WithError(err).Errorf("invalid order id [%s]", orderID)

		response.NotFound(w)

		return
	}

	order, err := c.service.GetOrder(r.Context(), orderID)
	if err != nil {
		log.WithError(err).Errorf("failed to find order [%s] for user [%s]", orderID, employeeID)

		response.InternalServerError(w)

		return
	}

	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"order": FillOrder(order),
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
		log.WithError(err).Errorf("failed to generate statistic for user %s", employeeID)

		response.InternalServerError(w)

		return
	}

	response.JSON(w, http.StatusOK, response.M{
		"success": true,
		"data": response.M{
			"statistic": FillStatistic(statistic),
		},
	})
}
