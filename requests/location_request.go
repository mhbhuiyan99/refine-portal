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
	parsedURL, err := url.Parse(baseURL) // Validates and converts the raw string into a structured *url.URL object.
	if err != nil {
		return nil, fmt.Errorf("parse base_url failed: %w", err)
	}

	parsedURL.Path = locationAPIPath // append the endpoint path

	query := parsedURL.Query() // Extract existing query parameters into a url.Values map
	query.Set("keyword", keyword)
	query.Set("isLocationEntity", "true")
	parsedURL.RawQuery = query.Encode() // Marshals the map back into a raw string

	logs.Debug(
		"[LocationRequest] Calling Location API | keyword=%s | url=%s",
		keyword,
		parsedURL.String(),
	)
	
	// Create HTTP request
	request, err := NewGETRequest(parsedURL.String())
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