package requests

import (
	"fmt"
	"net/url"
	"refine-portal/models"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
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

	// Get the base url
	baseURL, err := web.AppConfig.String("base_url")
	if err != nil {
		logs.Error(
			"[LocationRequest] Failed to read configuration | key=base_url | err=%v",
			err,
		)
		return nil, fmt.Errorf("failed to read 'base_url' from configuration: %w", err)
	}

	if strings.TrimSpace(baseURL) == "" {
		return nil, fmt.Errorf("configuration 'base_url' is empty")
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
