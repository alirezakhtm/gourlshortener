package handler

import (
	"fmt"
	"gourlshortener/structs"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefineNewUrl(ctx *gin.Context){
	var url structs.Url
	err_bind := ctx.BindJSON(&url)
	if err_bind != nil {
		ctx.String(http.StatusOK, "Input is not true")
	}
	shorUrl := GenerateSafeRandomAddress()
	url.SetShortUrl(shorUrl)
	_, err := Insert(url)
	if err != nil {
		ctx.String(http.StatusOK, err.Error())
	}
	ctx.JSON(http.StatusOK, url)
}

func RedirectToUrl(ctx *gin.Context){
	shortUrl := ctx.Param("shortUrl")
	url, err := Find(shortUrl)
	if err != nil {
		ctx.String(http.StatusOK, "this url doesn't exist")
	}
	originalAddress := url.GetAddress()
	ctx.Redirect(http.StatusMovedPermanently, originalAddress)
}

func GenerateSafeRandomAddress() string {
	shortUrl := ""
	for i := 0; i < 100; i++ {	
		shortUrl = "/" + GenerateRandomAddress()
		url, _ := Find(shortUrl)
		if url.GetShortUrl() == "" {
			break
		}
	}
	return shortUrl
}

func GenerateRandomAddress() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
    b := make([]rune, 3)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
	ans := string(b)
	fmt.Println(ans)
    return string(b)
}