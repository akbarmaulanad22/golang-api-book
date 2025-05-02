package app

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {

	logger := logrus.New()

	file, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}

	return logger
}