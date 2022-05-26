package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"gourlshortener/errors"
	"gourlshortener/structs"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getConnection() *mongo.Client{
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		fmt.Println("Please set your 'MONGO_URI' environment variable")
		uri = "mongodb://localhost:27017"
	}
	clientOption := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	return client
}

func Insert(url structs.Url) (*mongo.InsertOneResult, error){
	checkRecord, _ := Find(url.GetShortUrl())
	if checkRecord.GetUsername() != "" {
		urlException := errors.UrlException{
			Message: "This Short url exist",
		}
		return nil, urlException
	}
	client := getConnection()
	coll := client.Database("goshortener").Collection("url")
	result, err := coll.InsertOne(
		context.TODO(),
		bson.D{
			{"address", url.GetAddress()},
			{"username", url.GetUsername()},
			{"shortUrl", url.GetShortUrl()},
		},
	)
	if err != nil {
		panic(err)
	}
	return result, nil
}

func Find(shortAddress string) (structs.Url, error) {
	client := getConnection()
	coll := client.Database("goshortener").Collection("url")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"shortUrl", shortAddress}}).Decode(&result)
	urlException := errors.UrlException{}
	if err != nil {
		urlException.Message = "No document was found with the url " + shortAddress
		return structs.Url{}, urlException
	}
	jsonData, err := json.MarshalIndent(result, "", "	")
	var url structs.Url
	if err != nil {
		urlException.Message = err.Error()
		return structs.Url{}, urlException
	}
	err_jsonparse := json.Unmarshal(jsonData, &url)
	if err_jsonparse != nil {
		urlException.Message = err_jsonparse.Error()
		return structs.Url{}, urlException
	}
	return url, nil
}
