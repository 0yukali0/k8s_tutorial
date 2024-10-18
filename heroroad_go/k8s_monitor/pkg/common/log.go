package common

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

func GetLogger() *logrus.Logger {
	return logger
}

func init() {
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
}
