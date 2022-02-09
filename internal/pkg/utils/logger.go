package utils

import (
	"os"

	"github.com/sirupsen/logrus"

	"proxy_server/config"
)

func NewLogger(config *config.AppConfig) *logrus.Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(config.LoggingLevel)
	if err != nil {
		logrus.Fatalf("Could not parse logging level: %s", err)
	}
	logger.Writer()
	logger.SetLevel(level)
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{})

	return logger
}
