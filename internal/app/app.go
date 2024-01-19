package app

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/insan1a/tech-tinker/internal/delivery/http/controllers/account"
	"github.com/insan1a/tech-tinker/internal/delivery/http/controllers/auth"
	"github.com/insan1a/tech-tinker/internal/delivery/http/controllers/router"
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
	cfg, err := newConfig()
	if err != nil {
		return err
	}
	mux := chi.NewMux()
	mux.NotFound(router.NotFoundHandler)
	mux.MethodNotAllowed(router.MethodNotAllowedHandler)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	mux.Use(jsonlogger.New)
	mux.Use(middleware.Recoverer)

	err = auth.MountRoutes(
		auth.NewConfig().
			WithRSAPrivateKey(cfg.RSA.privateKey),
		mux,
	)
	if err != nil {
		return err
	}

	err = account.MountRoutes(
		account.NewConfig().
			WithRSAPublicKey(cfg.RSA.publicKey),
		mux,
	)
	if err != nil {
		return err
	}

	srv := http.Server{
		Handler:      mux,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Minute,
		ErrorLog:     slog.NewLogLogger(slog.NewTextHandler(os.Stderr, nil), slog.LevelError),
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
