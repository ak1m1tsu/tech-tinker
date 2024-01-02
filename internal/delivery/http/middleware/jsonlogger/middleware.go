package jsonlogger

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
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
				"uri":      r.RequestURI,
				"method":   r.Method,
				"status":   ww.statusCode,
				"size":     ww.size,
				"duration": time.Since(start).String(),
			}).Info("request completed")
		}()

		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}
