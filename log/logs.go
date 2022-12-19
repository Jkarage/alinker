package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = &logrus.Logger{
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.DebugLevel,
}

// CreateLogger checks if a runtime.log file exists
// Else creates it with permission 0666
func CreateLogger() error {
	file, err := os.OpenFile("log/runtime.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)

	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	return nil
}

func Initialize() {
	CreateLogger()
}

func Write(err error) {
	log.Error(err)
}
