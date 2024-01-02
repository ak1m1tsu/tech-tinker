package app

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/insan1a/tech-tinker/internal/delivery/http/controller/account"
	"github.com/insan1a/tech-tinker/internal/delivery/http/controller/router"
	"github.com/insan1a/tech-tinker/internal/delivery/http/middleware/jsonlogger"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() error {
	mux := chi.NewMux()
	mux.NotFound(router.NotFoundHandler)
	mux.MethodNotAllowed(router.MethodNotAllowedHandler)

	mux.Use(jsonlogger.New)
	account.MountRoutes(mux)

	srv := http.Server{
		Handler:      mux,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Minute,
		ErrorLog:     slog.NewLogLogger(slog.NewJSONHandler(os.Stderr, nil), slog.LevelError),
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		logrus.WithFields(logrus.Fields{
			"address": srv.Addr,
		}).Info("starting http server")
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})
	g.Go(func() error {
		<-gCtx.Done()
		logrus.Info("shutting down http server")
		if err := srv.Shutdown(context.Background()); err != nil {
			return err
		}
		return nil
	})

	return g.Wait()
}
