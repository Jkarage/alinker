package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jkarage/alinker/routers"
	"github.com/jkarage/alinker/utils"
)

func main() {
	r := gin.Default()
	routers.InitRoute()
	utils.InitializeStore()
	r.Run(":9808")

}
