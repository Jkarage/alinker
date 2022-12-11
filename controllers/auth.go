package controllers

import (
	"net/http"

	"github.com/jkarage/alinker/models/user"
	"github.com/jkarage/alinker/utils"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct{}

func (auth Auth) Register(c *gin.Context) {
	type signupInfo struct {
		Email    string `idx:"{email},unique" bson:"email,omitempty" json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
		Username string `idx:"unique" json:"username" binding:"required,min=5,alphanum"`
		Apps     int    `bson:"apps" json:"apps"`
	}

	var info signupInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user user.User
	user.Email = info.Email
	user.Username = info.Username
	user.Apps = 0

	hash, err := bcrypt.GenerateFromPassword([]byte(info.Password), 15)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	user.Password = string(hash)
	var userservice utils.Userservice
	err = userservice.Create(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Success Insertion",
	})
}

func (auth Auth) Get(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var userservice utils.Userservice
	user, err := userservice.Get(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.String(http.StatusOK, user.String())
	}
}

func (auth Auth) Delete(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var userservice utils.Userservice

	err := userservice.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.String(200, id.String())
		return
	}
}

func (auth Auth) Login(c *gin.Context) {
	type logininfo struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var info logininfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var userservice utils.Userservice
	user, err := userservice.FindByEmail(info.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(info.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := user.GetToken()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Header("Authentication", token)
	// return
}
