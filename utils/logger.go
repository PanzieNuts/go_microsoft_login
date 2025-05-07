package utils

import (
	"log"
	"os"
)

var logger *log.Logger

// InitLogger initializes the logger
func InitLogger() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
}

// GetLogger returns the logger instance
func GetLogger() *log.Logger {
	return logger
}
