package models

type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type Page struct {
	Name     string    `json:"name"`
	Projects []Project `json:"projects"`
}
