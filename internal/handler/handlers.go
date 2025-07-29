package handler

import (
	"groupie-tracker/internal/api"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func ArtistHandler(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, _ := strconv.Atoi(idStr)

	data, err := api.FetchArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artist", 500)
		return
	}

	artists, err := api.ParseArtists(data)
	if err != nil {
		http.Error(w, "Failed to parse artists", 500)
		return
	}

	for _, artist := range artists {
		if artist.ID == id {
			tmpl := template.Must(template.ParseFiles("web/artist.html"))
			tmpl.Execute(w, artist)
			return
		}
	}

	http.NotFound(w, r)

}
