package requests

import (
	"net/url"
	"refine-portal/models"

	"github.com/beego/beego/v2/core/logs"
)

// Location API endpoint.
const (
	locationAPIPath = "/api/location/v1"
)

// GetLocationRequest retrieves location suggestions from the Location API.
//
// Responsibilities:
//   - Build the Location API URL.
//   - Create the HTTP request.
//   - Execute the HTTP request.
//   - Decode the JSON response.
//   - Return the location response.
func GetLocationRequest(
	keyword string,
) (*models.LocationResponse, error) {

	baseURL, err := GetBaseURL()
	if err != nil {
		return nil, err
	}

	// Build URL
	query := url.Values{}
	query.Set("keyword", keyword)
	query.Set("isLocationEntity", "true")

	requestURL, err := BuildURL(baseURL, locationAPIPath, query)
	if err != nil {
		return nil, err
	}

	logs.Debug(
		"[LocationRequest] Calling Location API | keyword=%s | url=%s",
		keyword,
		requestURL,
	)

	request, err := NewGETRequest(requestURL)
	if err != nil {
		logs.Error(
			"[LocationRequest] Create request failed | err=%v",
			err,
		)
		return nil, err
	}

	var location models.LocationResponse

	err = DoRequest(
		request,
		&location,
	)

	if err != nil {
		return nil, err
	}

	logs.Debug(
		"[LocationRequest] Location API success | location=%s",
		location.GeoInfo.Name,
	)

	return &location, nil
}
