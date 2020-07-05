package main

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var l = logrus.New()

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		//ForceColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.999Z07:00",
		PadLevelText:    true,
		//DisableSorting: true,
	})

	log.SetLevel(log.DebugLevel)

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	log.SetReportCaller(true)

	log.Errorf("As always, expect the unexpected")

	l.SetReportCaller(true)
	l.Panic("not what you expected, eh?")
}
