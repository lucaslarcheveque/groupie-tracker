package tests

import (
	"groupie-tracker/internal/api"
	"testing"
)

func TestParseArtists(t *testing.T) {
	sampleJSON := []byte(`[
        {
            "id": 1,
            "name": "The Beatles",
            "members": ["John Lennon", "Paul McCartney"],
            "start_year": 1960,
            "first_album_date": "1963-03-22",
            "image": "http://example.com/beatles.jpg"
        }
    ]`)

	artists, err := api.ParseArtists(sampleJSON)
	if err != nil {
		t.Fatalf("ParseArtists failed: %v", err)
	}

	if len(artists) != 1 {
		t.Fatalf("Expected 1 artist, got %d", len(artists))
	}

	if artists[0].Name != "The Beatles" {
		t.Errorf("Expected artist name 'The Beatles', got %q", artists[0].Name)
	}
}
