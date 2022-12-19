package main

import (
	"github.com/jkarage/alinker/log"
	"github.com/jkarage/alinker/routers"
	"github.com/jkarage/alinker/utils"
)

func main() {
	log.Initialize()
	r := routers.InitRoute()
	utils.InitializeStore()
	r.Run(":9800")
}
