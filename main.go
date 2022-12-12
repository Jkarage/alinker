package main

import (
	"github.com/jkarage/alinker/routers"
	"github.com/jkarage/alinker/utils"
)

func main() {
	r := routers.InitRoute()
	utils.InitializeStore()
	r.Run(":9800")

}
