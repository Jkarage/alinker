package logger

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jkarage/alinker/utils"
)

func SetLogOutput() {
	f, err := os.OpenFile("runtime.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	utils.CheckNilError(err)

	gin.DefaultWriter = f
}
