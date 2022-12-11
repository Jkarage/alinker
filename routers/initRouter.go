package routers

import "github.com/gin-gonic/gin"

func InitRoute() *gin.Engine {
	r := gin.Default()
	SetAuthRouter(r)
	SetShortenerRouter(r)
	return r
}
