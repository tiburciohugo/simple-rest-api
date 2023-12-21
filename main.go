package main

import (
	"net/http"

	"github.com/tiburciohugo/simple-rest-api/router"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	router.Initialize()
}

func main() {
	http.HandleFunc("/", HandleRequest)
	http.ListenAndServe(":8080", nil)
}
