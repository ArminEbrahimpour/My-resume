package main

import (
	"net/http"

	"MyResume/internal/handlers"
)

func main() {

	http.HandleFunc("/", handlers.IndexHandler)

	http.ListenAndServe(":8080", nil)
}
