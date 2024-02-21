package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ak1m1tsu/go-libs/connector/postgresql"
	"github.com/go-chi/chi/v5"
	authcontroller "github.com/insan1a/tech-tinker/internal/delivery/http/controllers/auth"
	"github.com/insan1a/tech-tinker/internal/delivery/http/router"
	authservice "github.com/insan1a/tech-tinker/internal/domain/services/auth"
	emprepo "github.com/insan1a/tech-tinker/internal/repository/employee"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func Run() error {
	cfg, err := newConfig()
	if err != nil {
		return err
	}

	conn, err := postgresql.Connect(
		cfg.DB.URL,
		postgresql.NewConfig().
			WithPoolSize(cfg.DB.PoolSize).
			WithConnectionAttempts(cfg.DB.ConnectionAttempts).
			WithConnectionTimeout(cfg.DB.ConnectionTimeout).
			WithRetryDelay(cfg.DB.RetryDelay),
	)
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := emprepo.New(conn)

	service := authservice.New(
		authservice.NewConfig().
			WithJWTPublicKey(cfg.JWT.PrivateKey).
			WithJWTTTL(cfg.JWT.TTL).
			WithCacheSize(cfg.Cache.Size).
			WithCacheTTL(cfg.Cache.TTL),
		repo,
	)

	controller := authcontroller.New(service)

	r := router.New()
	r.Route("/auth", func(r chi.Router) {
		r.Post("/token", controller.HandleAuthToken)
	})

	srv := http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
		WriteTimeout: cfg.HTTP.Timeout,
		ReadTimeout:  cfg.HTTP.Timeout,
		IdleTimeout:  cfg.HTTP.Timeout,
		ErrorLog:     slog.NewLogLogger(slog.NewTextHandler(logrus.StandardLogger().Out, nil), slog.LevelError),
	}

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
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
