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
	// Wrap the JSON in a struct that matches the API
	var wrapper struct {
		Index []models.Relation `json:"index"`
	}

	err := json.Unmarshal(data, &wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.Index, nil
}
