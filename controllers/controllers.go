package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/RajVerma97/golang-url-shortner/models"
	"net/http"
	"strings"
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
	shortenUrl := generateShortCode(originalUrl)

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
func generateShortCode(originalUrl string) string {

	hasher := md5.New()
	hasher.Write([]byte(originalUrl)) //Converts the original url to a slice of byte
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8]
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
