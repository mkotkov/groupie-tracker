package handlers

import (
	"net/http"
	"strings"
	"text/template"

	"groupieTracker/models"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.NotFound(w, r)
		return
	}
	name := parts[2]

	if name == "" {
		http.NotFound(w, r)
		return
	}

	var artistInfo models.Artist
	found := false
	for _, a := range models.Artists {
		if a.Name == name {
			artistInfo = a
			found = true
			break
		}
	}

	if !found {
		http.NotFound(w, r)
		return
	}

	var artistRelations models.RelationsData
	for _, relation := range models.Relations.Index {
		if relation.Id == artistInfo.Id {
			artistRelations.Index = append(artistRelations.Index, relation)
		}
	}
	tmplFuncs := template.FuncMap{
		"formatRelations": formatRelations,
	}

	tmpl, err := template.New("artist.html").Funcs(tmplFuncs).ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		ArtistInfo      models.Artist
		ArtistRelations models.RelationsData
	}{
		ArtistInfo:      artistInfo,
		ArtistRelations: artistRelations,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formatRelations(relations []models.Relation) map[string][]string {
	formattedRelations := make(map[string][]string)
	for _, relation := range relations {
		for location, dates := range relation.DatesLocation {
			formattedRelations[location] = dates
		}
	}
	return formattedRelations
}
