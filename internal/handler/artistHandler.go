package handler

import (
	"groupie-tracker/internal/api"
	"groupie-tracker/internal/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Extract artist ID from URL
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest, "Invalid ID", "The artist ID is invalid.")
		return
	}

	// Fetch and parse artists
	artistData, err := api.FetchArtists()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Server Error", "Failed to fetch artists")
		return
	}
	artists, err := api.ParseArtists(artistData)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Server Error", "Failed to parse artists data")
		return
	}

	// Fetch and parse relations
	relationData, err := api.FetchRelation()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Server Error", "Failed to fetch relations")
		return
	}
	relations, err := api.ParseRelation(relationData)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Server Error", "Failed to parse relations")
		return
	}

	// Find the selected artist
	var selectedArtist models.Artist
	found := false
	for _, a := range artists {
		if a.ID == id {
			selectedArtist = a
			found = true
			break
		}
	}
	if !found {
		ErrorHandler(w, http.StatusNotFound, "Not Found", "The artist you are looking for does not exist.")
		return
	}

	// Find the corresponding relation
	var artistRelation models.Relation
	for _, r := range relations {
		if r.ID == id {
			artistRelation = r
			break
		}
	}

	// Pre-format members as a single string
	membersStr := strings.Join(selectedArtist.Members, ", ")

	// Combine all data into one struct for the template
	dataToRender := struct {
		Artist        models.Artist
		Relation      models.Relation
		MembersString string
	}{
		Artist:        selectedArtist,
		Relation:      artistRelation,
		MembersString: membersStr,
	}

	// Parse and execute template
	tmpl, err := template.ParseFiles("web/artist.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Template Error", "Could not parse template")
		return
	}

	if err := tmpl.Execute(w, dataToRender); err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Template Error", "Could not render artist page")
	}
}
