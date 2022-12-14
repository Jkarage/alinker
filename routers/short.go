package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jkarage/alinker/controllers"
	"github.com/jkarage/alinker/middlewares"
)

func SetShortenerRouter(r *gin.Engine) {
	short := new(controllers.Shortener)
	r.GET("/", short.Home)
	r.GET("/docs", short.Docs)
	r.GET("/:shortUrl", short.HandleShortUrlRedirect)
	shortenerGroup := r.Group("/")
	shortenerGroup.Use(middlewares.Authentication())
	shortenerGroup.POST("/create-short-url", short.CreateShortUrl)
}
