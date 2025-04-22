// Logrus logger setup
// Removed duplicate package declaration
package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger(logLevel string) {
	Log = logrus.New()

	// Set output to both file and stdout
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(file)
	} else {
		Log.SetOutput(os.Stdout)
		Log.Warn("Failed to log to file, using default stdout")
	}

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	Log.SetLevel(logrus.DebugLevel)
}
