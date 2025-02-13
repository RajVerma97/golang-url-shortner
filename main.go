package main

import (
	"fmt"
	"net/http"

	"github.com/RajVerma97/golang-url-shortner/config"
	"github.com/RajVerma97/golang-url-shortner/routes"
)

func main() {

	routes.Setup()
	config.ConnectDB()
	fmt.Println("Listening on PORT 3000")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}
}
