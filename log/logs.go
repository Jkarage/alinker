package log

import (
	"os"

	helpers "github.com/jkarage/alinker/env"
)

type Log struct{}

// CreateLogger checks if a runtime.log file exists
// Else creates it with permission 0666
func CreateLogger() error {
	f, err := helpers.Env("LOG_FILE", "runtime.log")
	if err != nil {
		return err
	}

	lf := "log/" + f
	if _, err = os.Stat(lf); os.IsNotExist(err) {
		_, err := os.Create(lf)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteLog() {

}
