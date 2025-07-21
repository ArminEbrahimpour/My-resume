package main

import (
	"html/template"
	"net/http"

	"MyResume/internal/handlers"
)

func main() {
	tmpl := template.Must(template.ParseFiles("/app/internal/templates/resume.gohtml"))

	resumeHandler := handlers.NewHandler(tmpl)
	http.HandleFunc("/", resumeHandler.IndexHandler)

	http.ListenAndServe(":8080", nil)
}
