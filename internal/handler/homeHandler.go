package handler

import (
	"groupie-tracker/internal/api"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	rawData, err := api.FetchArtists()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Server Error", "Failed to load artists")
		return
	}

	artists, err := api.ParseArtists(rawData)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Server Error", "Failed to parse artists data")
		return
	}

	tmpl := template.Must(template.ParseFiles("web/index.html"))
	err = tmpl.Execute(w, artists)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Template Error", "Could not render the homepage")
		return
	}
}
