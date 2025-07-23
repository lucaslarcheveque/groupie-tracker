package api

import (
	"io"
	"net/http"
)

func fetchURL(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseData, err
}

func FetchArtists() ([]byte, error) {
	const artistURL = "https://groupietrackers.herokuapp.com/api/artists"
	return fetchURL(artistURL)
}

func FetchLocations() ([]byte, error) {
	const locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	return fetchURL(locationsURL)
}

func FetchDates() ([]byte, error) {
	const datesURL = "https://groupietrackers.herokuapp.com/api/dates"
	return fetchURL(datesURL)
}

func FetchRelation() ([]byte, error) {
	const relationURL = "https://groupietrackers.herokuapp.com/api/relation"
	return fetchURL(relationURL)
}
