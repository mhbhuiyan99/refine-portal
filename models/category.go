package models

type CategoryResponse struct {
	GeoInfo GeoInfo `json:"GeoInfo"`
	Result  CategoryResult `json:"Result"`
}

type CategoryResult struct {
	Sections []CategorySection `json:"Sections"`
}

type CategorySection struct {
	Query    CategoryQuery `json:"Query"`
	Title    string        `json:"Title"`
	SubTitle string        `json:"SubTitle"`
	ID       string        `json:"ID"`

	Items []Item `json:"Items"`
}

type CategoryQuery struct {
	Order int `json:"order"`
}

type Item struct {
	ID       string   `json:"ID"`
	GeoInfo  GeoInfo  `json:"GeoInfo"`
	Partner  Partner  `json:"Partner"`
	Property Property `json:"Property"`
}

type GeoInfo struct {
    Name        string       `json:"Name"`
    ShortName   string       `json:"ShortName"`
    State       string       `json:"State"`
    Country     string       `json:"Country"`
    LocationID  string       `json:"LocationID"`
    Breadcrumbs []Breadcrumb `json:"Breadcrumbs"`
}

type Partner struct {
	URL string `json:"URL"`
}