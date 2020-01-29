package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var logger *logrus.Logger

// Setup ...
func Setup() {
	logger = logrus.New()

	if viper.GetString("environment") == "production" {
		logger.SetLevel(logrus.ErrorLevel)
		return
	}

	switch viper.GetString("log_level") {
	case "trace":
		logger.SetLevel(logrus.TraceLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	default:
		logger.SetLevel(logrus.DebugLevel)
	}
}

// Log ...
func Log() *logrus.Logger {
	return logger
}
