package main

import (
	"math/rand"
	"gourlshortener/handler"
	"time"

	"github.com/gin-gonic/gin"
)

func main()  {
	
	urlShortenerServer := gin.Default()
	rand.Seed(time.Now().UnixNano())

	// define new address
	urlShortenerServer.POST("/urls", func(ctx *gin.Context) {
		handler.DefineNewUrl(ctx)
	})

	// redirect to short url
	urlShortenerServer.GET("/u/*shortUrl", func(ctx *gin.Context) {
		handler.RedirectToUrl(ctx)
	})

	// run server on port 8080
	urlShortenerServer.Run(":8080")

}