package models

type Location struct {
	Name string `json:"name"`
	Category string `json:"category"`
}

type LocationResponse struct {
	Location []Location
}