package handler

import (
	"groupie-tracker/internal/api"
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	rawData, err := api.FetchArtists()
	if err != nil {
		http.Error(w, "Failed to load Artists", http.StatusInternalServerError)
		return
	}

	artists, err := api.ParseArtists(rawData)
	if err != nil {
		http.Error(w, "Failed to parse Artists data", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/index.html"))
	err = tmpl.Execute(w, artists)
	if err != nil {
		log.Println("Template execution error", err)
		return
	}
}
