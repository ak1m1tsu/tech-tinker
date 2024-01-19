package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/insan1a/tech-tinker/internal/delivery/http/controllers"
	"github.com/insan1a/tech-tinker/internal/domain/model"
	"github.com/insan1a/tech-tinker/internal/lib/decoder"
	"github.com/insan1a/tech-tinker/internal/lib/jwt"
	"github.com/insan1a/tech-tinker/internal/lib/response"
	"github.com/sirupsen/logrus"
	"net/http"
)

type controller struct {
	config *Config
}

func newController(cfg *Config) (*controller, error) {
	if cfg == nil {
		return nil, controllers.ErrConfigMissing
	}

	return &controller{config: cfg}, nil
}

func MountRoutes(cfg *Config, router chi.Router) error {
	c, err := newController(cfg)
	if err != nil {
		return err
	}

	router.Route("/auth", func(r chi.Router) {
		r.Post("/token", c.HandleAuthToken)
	})

	return nil
}

func (c *controller) HandleAuthToken(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := decoder.DecodeJSON(r.Body, &input); err != nil {
		logrus.WithError(err).Error("failed to decode request body")

		response.BadRequest(w, err.Error())

		return
	}

	token, err := jwt.GenerateToken(
		&jwt.Employee{
			ID:   input.Email,
			Role: model.EmployeeRoleTechnician.String(),
		},
		c.config.tokenTTL,
		c.config.key,
	)
	if err != nil {
		logrus.WithError(err).Error("failed to generate token")

		response.InternalServerError(w)

		return
	}

	response.JSON(w, http.StatusOK, response.M{
		"data": response.M{
			"token": token,
		},
		"success": true,
	})
}
