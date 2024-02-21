package publicapi

import (
	"context"
	"errors"
	"fmt"
	"github.com/ak1m1tsu/go-libs/connector/postgresql"
	accountservice "github.com/insan1a/tech-tinker/internal/domain/services/account"
	customerrepo "github.com/insan1a/tech-tinker/internal/repository/customer"
	employeerepo "github.com/insan1a/tech-tinker/internal/repository/employee"
	orderrepo "github.com/insan1a/tech-tinker/internal/repository/order"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	accountcontroller "github.com/insan1a/tech-tinker/internal/delivery/http/controllers/account"
	"github.com/insan1a/tech-tinker/internal/delivery/http/middleware/jwtvalidation"
	"github.com/insan1a/tech-tinker/internal/delivery/http/router"
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

	orderRepo := orderrepo.New(conn)
	employeeRepo := employeerepo.New(conn)
	customerRepo := customerrepo.New(conn)

	service := accountservice.New(&accountservice.Config{}, employeeRepo, customerRepo, orderRepo)

	controller := accountcontroller.New(service)

	r := router.New()
	r.Use(jwtvalidation.New(cfg.JWT.PublicKey))
	r.Route("/api", func(r chi.Router) {
		r.Route("/account", func(r chi.Router) {
			r.Get("/", controller.HandleAccountInfo)
			r.Post("/stat", controller.HandleAccountStatistic)
			r.Route("/orders", func(r chi.Router) {
				r.Get("/", controller.HandleAccountOrders)
				r.Get("/{orderID}", controller.HandleAccountOrder)
			})
		})
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
