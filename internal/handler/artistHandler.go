package handler

import (
	"groupie-tracker/internal/api"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, _ := strconv.Atoi(idStr)

	data, err := api.FetchArtists()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Server Error", "Failed to fetch artist")
		return
	}

	artists, err := api.ParseArtists(data)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Server Error", "Failed to parse artists data")
		return
	}

	for _, artist := range artists {
		if artist.ID == id {
			tmpl := template.Must(template.ParseFiles("web/artist.html"))
			if err := tmpl.Execute(w, artist); err != nil {
				ErrorHandler(w, http.StatusInternalServerError, "Template Error", "Could not render artist page")
			}
			return
		}
	}

	ErrorHandler(w, http.StatusNotFound, "Not Found", "The artist you are looking for does not exist.")

}
