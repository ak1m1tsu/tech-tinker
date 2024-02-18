package publicapi

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/insan1a/tech-tinker/internal/delivery/http/middleware/jwtvalidation"
	"github.com/insan1a/tech-tinker/internal/delivery/http/router"
)

func Run() error {
	cfg, err := newConfig()
	if err != nil {
		return err
	}

	mux := router.New()
	mux.Use(jwtvalidation.New(cfg.RSA.PublicKey))
	mux.MountAccountRoutes()

	srv := http.Server{
		Handler:      mux,
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
