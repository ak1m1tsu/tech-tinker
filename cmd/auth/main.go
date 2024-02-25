package main

import (
	"os"

	app "github.com/ak1m1tsu/tech-tinker/internal/app/auth"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(new(logrus.TextFormatter))
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
}

func main() {
	if err := app.Run(); err != nil {
		logrus.WithError(err).Fatal("Failed to start application")
	}
}
