package jwtvalidation

import (
	"crypto/rsa"
	"net/http"
	"strings"

	"github.com/ak1m1tsu/tech-tinker/internal/lib/appcontext"
	"github.com/ak1m1tsu/tech-tinker/internal/lib/jwt"
	"github.com/ak1m1tsu/tech-tinker/internal/lib/response"
	"github.com/sirupsen/logrus"
)

const headerPrefix = "Bearer "

func New(key *rsa.PublicKey) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rawToken := r.Header.Get("Authorization")
			if rawToken == "" {
				response.Unauthorized(w)
				return
			}

			token := parseToken(rawToken)
			if token == "" {
				response.Unauthorized(w)
				return
			}

			claims, err := jwt.ValidateToken(token, key)
			if err != nil {
				logrus.WithError(err).Error("failed to validate token")

				response.Forbidden(w)
				return
			}

			ctx := appcontext.WithEmployeeID(r.Context(), claims.EmployeeID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func parseToken(token string) string {
	if strings.HasPrefix(token, headerPrefix) {
		return strings.TrimPrefix(token, headerPrefix)
	}
	return ""
}
