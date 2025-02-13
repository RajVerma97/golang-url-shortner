package routes

import (
	"github.com/RajVerma97/golang-url-shortner/controllers"
	"net/http"
)

func Setup() {
	http.HandleFunc("/", controllers.HandleRoot)
	http.HandleFunc("/shorten", controllers.HandleShorten)
	http.HandleFunc("/redirect/", controllers.HandleRedirect)
}
