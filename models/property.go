package models

type PropertyListRequest struct {
	Category  string
	Locations string
	Order     int
	Limit     int
	Items     int
	Device    string
	Page      int
}

type PropertyListResponse struct {
	Success bool            `json:"Success"`
	GeoInfo PropertyGeoInfo `json:"GeoInfo"`
	Result  PropertyResult  `json:"Result"`
}

type PropertyGeoInfo struct {
	CountryCode  string `json:"CountryCode"`
	LocationSlug string `json:"LocationSlug"`
}

type PropertyResult struct {
	Count   int      `json:"Count"`
	ItemIDs []string `json:"ItemIDs"`
}
