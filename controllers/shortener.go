package controllers

import (
	"net/http"
	"os"

	"github.com/jkarage/alinker/env"
	"github.com/jkarage/alinker/utils"

	"github.com/gin-gonic/gin"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

type Shortener struct{}

func (s Shortener) CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl, err := utils.GenerateShortLink(creationRequest.LongUrl, c)
	utils.CheckNilError(err)
	utils.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	host, err := env.Env("APP_HOST"+"/", "alinker.tk/")
	utils.CheckNilError(err)
	c.JSON(http.StatusOK, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

func (s Shortener) HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := utils.RetrieveInitialUrl(shortUrl)
	c.Redirect(http.StatusPermanentRedirect, initialUrl)
}

func (s Shortener) Home(c *gin.Context) {
	f, err := os.ReadFile("assets/index.html")
	utils.CheckNilError(err)
	c.Status(http.StatusOK)
	c.Writer.Write(f)
}

func (s Shortener) Docs(c *gin.Context) {
	f, err := os.ReadFile("assets/docs.html")
	utils.CheckNilError(err)
	c.Status(http.StatusOK)
	c.Writer.Write(f)
}
