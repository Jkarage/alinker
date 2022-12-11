package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jkarage/alinker/controllers"
)

func SetAuthRouter(r *gin.Engine) {
	auth := new(controllers.Auth)
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)
	r.GET("user/:id", auth.Get)
	r.DELETE("user/:id", auth.Delete)
}
