package handlers

import (
	"MyResume/internal/models"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type ResumeHandler struct {
	tmpl *template.Template
}

func NewHandler(tmpl *template.Template) *ResumeHandler {
	return &ResumeHandler{
		tmpl: tmpl,
	}
}

func (h *ResumeHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {

	var page models.Page
	data, err := os.ReadFile("app/static/projects.txt")
	if err != nil {
		fmt.Printf("reading projects file fialed : %v", err)
		return
	}
	err = json.Unmarshal(data, &page)
	if err != nil {
		fmt.Printf("parsing projects failed : %v", err)
		return
	}
	/*
		page := models.Page{
			Name:     "Armin Ebrahimpour",
			Projects: projects,
		}
	*/
	h.tmpl.Execute(w, page)
}
