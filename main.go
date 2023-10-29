package main

import (
	"fmt"
	"log"
	"net/http"

	"groupieTracker/handlers"
	"groupieTracker/middleware"
	"groupieTracker/models"
)

func main() {
	artistsData, err := middleware.FetchData(models.ArtistsURL)
	if err != nil {
		log.Fatal("Error fetching artist data:", err)
	}

	relationData, err := middleware.FetchData(models.RelationURL)
	if err != nil {
		log.Fatal("Error fetching relation data:", err)
	}

	// Parsing data into separate slices
	err = middleware.ParseData(artistsData, &models.Artists)
	if err != nil {
		log.Fatal("Error parsing artist data:", err)
	}

	err = middleware.ParseData(relationData, &models.Relations)
	if err != nil {
		log.Fatal("Error parsing relation data:", err)
	}

	// Setting up routes and starting the server
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

	http.HandleFunc("/", handlers.MainPage)
	http.HandleFunc("/artist/", handlers.Artist)
	
	fmt.Println("Server is running!\nhttp://localhost:8080\nTo shut down the server ctrl+c")
	http.ListenAndServe(":8080", nil)
	
}
