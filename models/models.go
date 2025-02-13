package models

import (
	"errors"
	"time"
)

type Url struct {
	ID          int       `json:"id"`
	OriginalUrl string    `json:"originalUrl"`
	ShortenUrl  string    `json:"shortenUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var urls = make(map[string]Url)

func CreateUrl(originalUrl, shortenUrl string) error {
	if originalUrl == "" {
		return errors.New("original url cannot be empty")
	}
	if shortenUrl == "" {
		return errors.New("shorten url cannot be empty")
	}

	if _, exists := urls[shortenUrl]; exists {

		return errors.New("shorten url already exist")
	}
	newUrl := Url{
		ID:          1,
		OriginalUrl: originalUrl,
		ShortenUrl:  shortenUrl,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	urls[shortenUrl] = newUrl
	return nil
}

func GetOriginalUrl(shortenUrl string) (string, error) {
	if shortenUrl == "" {
		return "", errors.New("empty shorten url ")
	}

	url, exists := urls[shortenUrl]
	if !exists {
		return "", errors.New("url not found")
	}
	return url.OriginalUrl, nil
}
