package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jkarage/alinker/models/user"
)

// TokenDecoder decodes the Authentication Key from the header
// Of the request and lookup for the corresponding user in the database
// Returns the User found, Nil if No user found
// TODO: Making sure the Authentication key corresponds to the user signed in
// Add Expiration time in the key as well
func TokenDecoder(c *gin.Context) *user.User {
	tokenString := c.Request.Header.Get("Authentication")
	if len(tokenString) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": "Authentication is required for this endpoint",
		})
		return nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secretKey := "secretkey"
		return []byte(secretKey), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"Error": err.Error(),
		})
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		userservice := Userservice{}
		user, err := userservice.FindByEmail(email)
		if err != nil {
			c.AbortWithStatusJSON(402, gin.H{
				"Error": "User not found",
			})
			return nil
		} else {
			return user
		}
	} else {
		return nil
	}

}
