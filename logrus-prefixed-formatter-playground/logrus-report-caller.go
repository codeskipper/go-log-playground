package main

import (
	"strings"

	"github.com/sirupsen/logrus"
	//prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"github.com/codeskipper/go-log-playground/internal/projectpath"
	prefixed "github.com/codeskipper/logrus-prefixed-formatter"
)

var log = logrus.New()

func init() {
	//formatter := new(prefixed.TextFormatter)
	formatter := &prefixed.TextFormatter{
		// CallerPrettyfier function is provided to let you introduce a custom format.
		//CallerPrettyfier: func(f *runtime.Frame) (string, string) {
		// scooped up for this example from https://trierra.dev/how-to-configure-a-golang-logger-logrus-for-production/
		// In my case I wanted file and line to look like this `file="engine.go:141`
		// but f.File provides a full path along with the file name.
		// So in `formatFilePath()` function I just trimmed everything before the file name
		// and added the line number at the end
		//	return f.Function, fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		//},

		// CallerFormatter function lets you to introduce a custom format.
		//CallerFormatter: func(fileVal string, funcVal string) (caller string) {

		//	return fmt.Sprintf("(%s: %s)", fileVal, funcVal)
		//},
		DisableColors: false,
		//FieldSpacing:  8,
	}
	log.SetFormatter(formatter)
	log.Level = logrus.DebugLevel
	log.SetReportCaller(true)
	//var projectPath = prefixed.GetProjectPath()
	log.Debug("Discovered project path:", projectpath.Root)
	log.Debug("logging ready")
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			// Fatal message
			log.WithFields(logrus.Fields{
				"omg":    true,
				"number": 100,
			}).Fatal("The ice breaks!")
		}
	}()

	// You could either provide a map key called `prefix` to add prefix
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	// Or you can simply add prefix in square brackets within message itself
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Debug("A group of walrus emerges from the ocean")

	// Warning message
	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	// Information message
	log.WithFields(logrus.Fields{
		"prefix":      "sensor",
		"temperature": -4,
	}).Info("Temperature changes")

	// Panic message
	log.WithFields(logrus.Fields{
		"prefix": "sensor",
		"animal": "orca",
		"size":   9009,
	}).Panic("It's over 9000!")
}
