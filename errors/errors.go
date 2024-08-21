package errors

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.ErrorLevel)
}

func NewError(msg string, hight bool) {
	if hight {
		err := errors.New(msg)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"msg":   "Error Hight level detected",
			}).Fatal("FIX THIS ERROR NOW!")
		}
	} else {
		err := errors.New(msg)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"msg":   "Error not hight priority",
			}).Error("Fix the error but not more fast or yes ?)")
		}
	}
}

func CheckErrors(err error) {
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"msg":   "Error detected read the error field for more information",
		}).Error("New error detected if this is a bug report in the repo")
	}
}
