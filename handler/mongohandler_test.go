package handler

import (
	"gourlshortener/structs"
	"os"
	"testing"
)

func init(){
	os.Setenv("MONGO_URI", "mongodb://localhost:27017")
}

func TestInsert(t *testing.T){
	url := structs.Url{}
	url.SetAddress("http://google.com")
	url.SetShortUrl("/go")
	url.SetUsername("alireza")
	_, err := Insert(url)
	if err != nil {
		panic(err)
	}
}

func TestFind(t *testing.T) {
	_, err := Find("/go")
	if err != nil {
		panic(err)
	}
}