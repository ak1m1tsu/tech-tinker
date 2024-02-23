package account

import (
	"errors"
	"net/http"

	"github.com/ak1m1tsu/tech-tinker/internal/domain/interfaces"
	"github.com/ak1m1tsu/tech-tinker/internal/domain/services/account"
	"github.com/ak1m1tsu/tech-tinker/internal/lib/appcontext"
	"github.com/ak1m1tsu/tech-tinker/internal/lib/decoder"
	"github.com/ak1m1tsu/tech-tinker/internal/lib/response"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

		if errors.Is(err, account.ErrOrderNotFound) {
			response.NotFound(w)

			return
		}

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

	var input struct {
		From string `json:"from"`
		To   string `json:"to"`
	}

	if err := decoder.DecodeJSON(r.Body, &input); err != nil {
		log.WithError(err).Errorf("failed to parse request body")

		response.BadRequest(w, "bad request body")

		return
	}

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
