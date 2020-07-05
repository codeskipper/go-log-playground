package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&TextFormatter{
		DisableColors: false,
		//ForceColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.999Z07:00",
		SpacePadding:    40,
	})

	log.Out = os.Stdout
	log.Level = logrus.DebugLevel
	log.ReportCaller = true
	log.Info("Logging from init")
}

func testFunc(l *logrus.Logger) {
	l.SetReportCaller(true)
	l.Printf("See if we get the report on the caller here")
}

func main() {
	/*
		logger := &logrus.Logger{
			Out:   os.Stderr,
			Level: logrus.DebugLevel,
			Formatter: &TextFormatter{
				//ForceColors:   true,
				DisableColors: false,
				//TimestampFormat : "2006-01-02 15:04:05",
				FullTimestamp: true,
				SpacePadding:  40,
			},
		}
	*/

	//logger.Printf("test")
	log.Debugf("There may be an issue here")
	log.Infof("Starting out in logrus territory")
	log.Warningf("Just so you know, the ice is thin")
	log.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	endgameLogger := log.WithFields(logrus.Fields{"place": "arctic", "situation": "hazardous"})
	endgameLogger.Errorf("As always, expect the unexpected")

	/* 	logger.SetReportCaller(true)
	   	logger.Printf("See if we get the report on the caller here")
	*/
	testFunc(log)
}
