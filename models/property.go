package models

type Property struct {
    ID          int
    Name        string
    Price       float64
    Thumbnail   string

    Rating      float64
    Images       []string
    Amenities    []string
    Description  string
}

type PropertyListResponse struct {
    Properties []Property
}

type PropertyDetailsResponse struct {
    Properties []Property
}