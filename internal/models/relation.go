package models

type Relation struct {
	ID           int                 `json:"id"`
	DateLocation map[string][]string `json:"datesLocations"`
}
