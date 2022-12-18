package log

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jkarage/alinker/env"
)

type Log struct{}

// CreateLogger checks if a runtime.log file exists
// Else creates it with permission 0666
func (l Log) CreateLogger() error {
	f, err := env.Env("LOG_FILE", "runtime.log")
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

func (l Log) Log(c *gin.Context) {
	// Logging  Implementation to come
}

func (l Log) Initialize() {
	l.CreateLogger()
}
