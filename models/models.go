package models

import (
	"context"
	"errors"
	"github.com/RajVerma97/golang-url-shortner/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Url struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"` //Maps to MongoDB's _id
	OriginalUrl string             `json:"originalUrl" bson:"originalUrl"`
	ShortenUrl  string             `json:"shortenUrl" bson:"shortenUrl"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func SaveUrlToDb(originalUrl, shortenUrl string) error {
	if originalUrl == "" {
		return errors.New("original url cannot be empty")
	}
	if shortenUrl == "" {
		return errors.New("shorten url cannot be empty")
	}
	_, err := GetURlFromDb(shortenUrl)

	if err == nil {
		return errors.New("url with the shorten url already exists")
	}

	newUrl := Url{
		OriginalUrl: originalUrl,
		ShortenUrl:  shortenUrl,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = config.Collection.InsertOne(ctx, newUrl)
	if err != nil {
		return errors.New("unable to insert url in the db")
	}

	return nil
}

func GetURlFromDb(shortenUrl string) (Url, error) {

	var url Url
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := config.Collection.FindOne(ctx, bson.M{"shortenUrl": shortenUrl}).Decode(&url)
	if err != nil {
		return Url{}, errors.New("error finding the url in the db")
	}
	return url, nil
}
func GetOriginalUrl(shortenUrl string) (string, error) {
	if shortenUrl == "" {
		return "", errors.New("empty shorten url ")
	}
	url, err := GetURlFromDb(shortenUrl)

	if err != nil {
		return "", errors.New("error finding the url in the db")
	}
	return url.OriginalUrl, nil
}
