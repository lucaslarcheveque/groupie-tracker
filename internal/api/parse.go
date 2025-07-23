package api

import (
	"encoding/json"
	"groupie-tracker/internal/models"
)

func ParseArtists(data []byte) ([]models.Artist, error) {
	var artists []models.Artist
	err := json.Unmarshal(data, &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func ParseLocation(data []byte) ([]models.Location, error) {
	var locations []models.Location
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func ParseDate(data []byte) ([]models.Date, error) {
	var dates []models.Date
	err := json.Unmarshal(data, &dates)
	if err != nil {
		return nil, err
	}
	return dates, nil
}

func ParseRelation(data []byte) ([]models.Relation, error) {
	var relations []models.Relation
	err := json.Unmarshal(data, &relations)
	if err != nil {
		return nil, err
	}
	return relations, nil
}
