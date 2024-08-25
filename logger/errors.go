package logger

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true, QuoteEmptyFields: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.ErrorLevel)
}

func NewError(msg string, high bool) {
	if high {
		err := errors.New(msg)
		if err != nil {
			log.WithFields(log.Fields{
				"error":      err,
				"msgOfRemus": "Report this error if is necessary in github.com/TeoDev1611/remus",
			}).Fatal(err)
		}
	} else {
		err := errors.New(msg)
		if err != nil {
			log.WithFields(log.Fields{
				"errorInfo":  err,
				"msgOfRemus": "Report this error if is necessary in github.com/TeoDev1611/remus",
			}).Error(err)
		}
	}
}

func CheckErrors(err error) {
	if err != nil {
		log.WithFields(log.Fields{
			"errorInfo":  err,
			"msgOfRemus": "Report this error if is necessary in github.com/TeoDev1611/remus",
		}).Error(err)
	}
}
