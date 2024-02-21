package jsonlogger

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func (w *wrappedResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *wrappedResponseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.size += size
	return size, err
}

func New(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := &wrappedResponseWriter{ResponseWriter: w}

		start := time.Now()
		defer func() {
			logrus.WithFields(logrus.Fields{
				"uri":        r.RequestURI,
				"method":     r.Method,
				"status":     ww.statusCode,
				"size":       ww.size,
				"duration":   time.Since(start).String(),
				"real_ip":    r.RemoteAddr,
				"request_id": middleware.GetReqID(r.Context()),
			}).Info("request completed")
		}()

		logrus.WithFields(logrus.Fields{
			"uri":        r.RequestURI,
			"method":     r.Method,
			"real_ip":    r.RemoteAddr,
			"request_id": middleware.GetReqID(r.Context()),
		}).Info("request started")
		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}
