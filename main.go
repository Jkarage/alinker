package main

import (
	logger "github.com/jkarage/alinker/log"
	"github.com/jkarage/alinker/routers"
	"github.com/jkarage/alinker/utils"
)

func main() {
	logger.SetLogOutput()
	r := routers.InitRoute()
	utils.InitializeStore()
	r.Run(":9800")
}
