package service

import (
	"groupie-tracker/internal/models"
	"strings"
)

// FilterArtists filtre les artistes dont le nom contient la chaîne recherchée (insensible à la casse)
func FilterArtists(artists []models.Artist, query string) []models.Artist {
	if query == "" {
		return artists // pas de recherche → renvoie tout
	}

	query = strings.ToLower(query)
	var result []models.Artist

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			result = append(result, artist)
		}
	}
	return result
}
