package main

import (
	"github.com/insan1a/tech-tinker/internal/app"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
}

func main() {
	if err := app.Start(); err != nil {
		logrus.WithError(err).Fatal("failed to start application")
	}
}
