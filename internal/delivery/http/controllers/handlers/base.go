package handlers

import (
	"net/http"

	"github.com/ak1m1tsu/tech-tinker/internal/lib/response"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	response.NotFound(w)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	response.MethodNotAllowed(w)
}
