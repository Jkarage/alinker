package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jkarage/alinker/env"
	"github.com/jkarage/alinker/utils"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authentication")
		if len(tokenString) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Error": "Authentication is required",
			})
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			secretKey, err := env.Env("SECRET_KEY", "")
			if err != nil {
				return nil, err
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"Error": err.Error(),
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
			userservice := utils.Userservice{}
			user, err := userservice.FindByEmail(email)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"Error": "User not found",
				})
				return
			}
			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Error": "Token is not valid",
			})
			return
		}
	}
}

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Errors": c.Errors,
		})
	}
}
