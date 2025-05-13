// utils/logger.go
package utils

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitLogger() {
	logger = logrus.New()
	logger.SetLevel(logrus.InfoLevel)
}

func GetLogger() *logrus.Logger {
	return logger
}
