package logger

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// SetupLog ...
func SetupLog() {
	logger = logrus.New()
}

// Log ...
func Log() *logrus.Logger {
	return logger
}
