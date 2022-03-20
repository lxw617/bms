package utils

import (
	log "github.com/sirupsen/logrus"
)

func SetLog() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
