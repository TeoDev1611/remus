package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true, QuoteEmptyFields: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func Info(msg string) {
	log.WithFields(log.Fields{
		"info":       msg,
		"msgOfRemus": "If you need help open a discussion in github.com/TeoDev1611/remus",
	}).Info(msg)
}
