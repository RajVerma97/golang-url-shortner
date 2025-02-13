package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/RajVerma97/golang-url-shortner/models"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hello handle root`))
}

func HandleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var body map[string]string
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	originalUrl := body["url"]

	if originalUrl == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}
	shortenUrl := generateShortCode()

	err = models.CreateUrl(originalUrl, shortenUrl)
	if err != nil {
		http.Error(w, "unable to create url", http.StatusBadRequest)
		return

	}
	response := map[string]string{
		"message":      "Successfully Shortened Url",
		"shortenedUrl": shortenUrl,
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
func generateShortCode() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	var result string

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= length; i++ {
		randomIndex := random.Intn(len(chars))
		char := string(chars[randomIndex])
		result += char
	}

	return result

}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	segments := strings.Split(path, "/")
	shortenUrl := segments[len(segments)-1]

	originalUrl, err := models.GetOriginalUrl(shortenUrl)
	if err != nil {
		http.Error(w, "not found the Original Url", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalUrl, http.StatusMovedPermanently)
	w.Write([]byte(`hello handle redirect`))

}
