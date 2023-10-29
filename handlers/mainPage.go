package handlers

import (
	"net/http"
	"text/template"

	"groupieTracker/models"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Artists   []models.Artist
		Relations models.RelationsData
	}{
		Artists:   models.Artists,
		Relations: models.Relations,
	}

	tmpl, err := template.ParseFiles("templates/index.html", "templates/card.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
