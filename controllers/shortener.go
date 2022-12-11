package controllers

import (
	"net/http"

	"github.com/jkarage/alinker/env"
	"github.com/jkarage/alinker/utils"

	"github.com/gin-gonic/gin"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

type Shortener struct{}

func (s Shortener) CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := utils.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	utils.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	host, err := env.Env("APP_HOST"+"/", "alinker.tk/")
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

func (s Shortener) HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := utils.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}

func (s Shortener) Home(c *gin.Context) {
	c.String(http.StatusOK, `Welcome to alinker API, A free url shortener api.
	Written using Golang.
	for more and endpoints found on this api, go to Documentation page
	at /docs`)
}
