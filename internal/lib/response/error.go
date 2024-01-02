package response

import "net/http"

type M map[string]any

func BadRequest(w http.ResponseWriter, message string) {
	JSON(w, http.StatusBadRequest, M{"error": message})
}

func NotFound(w http.ResponseWriter) {
	message := "the resource not found"
	JSON(w, http.StatusNotFound, M{"error": message})
}

func InternalServerError(w http.ResponseWriter) {
	message := "internal server error"
	JSON(w, http.StatusInternalServerError, M{"error": message})
}

func Unauthorized(w http.ResponseWriter) {
	message := "unauthorized"
	JSON(w, http.StatusUnauthorized, M{"error": message})
}

func Forbidden(w http.ResponseWriter) {
	message := "forbidden"
	JSON(w, http.StatusForbidden, M{"error": message})
}

func MethodNotAllowed(w http.ResponseWriter) {
	message := "method not allowed"
	JSON(w, http.StatusMethodNotAllowed, M{"error": message})
}
