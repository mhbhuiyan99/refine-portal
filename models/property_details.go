package models

type PropertyDetailsRequest struct {
	PropertyIDList []string
}

type PropertyDetailsResponse struct {
	Success bool              `json:"Success"`
	Message string            `json:"Message"`
	Error   any               `json:"Error"`
	Items   []PropertyDetails `json:"Items"`
}

type PropertyDetails struct {
	Property Property `json:"Property"`
}

type Property struct {
	PropertyName      string  `json:"PropertyName"`
	PropertyType      string  `json:"PropertyType"`
	PropertySlug      string  `json:"PropertySlug"`
	FeatureImage      string  `json:"FeatureImage"`
	Price             float64 `json:"Price"`
	ReviewScore       float64 `json:"ReviewScore"`
	PropertyAttribute string  `json:"PropertyAttribute"`

	Amenities    	  map[string]string `json:"Amenities"`
	TopListedAmenities []Amenity `json:"TopListedAmenities"`

	Counts Counts `json:"Counts"`
}

type Counts struct {
	Bedroom   int `json:"Bedroom"`
	Bathroom  int `json:"Bathroom"`
	Occupancy int `json:"Occupancy"`
}

type Amenity struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}