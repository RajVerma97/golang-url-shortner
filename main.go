package main

import (
	"fmt"
	"github.com/RajVerma97/golang-url-shortner/routes"
	"net/http"
)

func main() {

	routes.Setup()
	fmt.Println("Listening on PORT 3000")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}
}
