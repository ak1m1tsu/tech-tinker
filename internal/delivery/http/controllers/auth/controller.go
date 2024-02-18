package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/insan1a/tech-tinker/internal/domain/interfaces"
	"github.com/insan1a/tech-tinker/internal/lib/decoder"
	"github.com/insan1a/tech-tinker/internal/lib/response"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	service interfaces.AuthService
}

func New(service interfaces.AuthService) *Controller {
	return &Controller{service: service}
}

func (c *Controller) HandleAuthToken(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	log := logrus.WithFields(logrus.Fields{
		"request_id": middleware.GetReqID(r.Context()),
		"real_ip":    r.RemoteAddr,
		"uri":        r.RequestURI,
		"method":     r.Method,
	})

	if err := decoder.DecodeJSON(r.Body, &input); err != nil {
		log.WithError(err).Error("failed to decode request body")

		response.BadRequest(w, err.Error())

		return
	}

	ctx := r.Context()

	e, err := c.service.GetByEmail(ctx, input.Email)
	if err != nil {
		log.WithError(err).Error("failed to find user by email")

		response.Unauthorized(w)

		return
	}

	token, err := c.service.CreateToken(ctx, e)
	if err != nil {
		log.WithError(err).WithField("email", e.Email).Error("failed to create token for user")

		response.Unauthorized(w)

		return
	}

	response.JSON(w, http.StatusOK, response.M{
		"data": response.M{
			"token": token,
		},
		"success": true,
	})
}
