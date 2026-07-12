package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"refine-portal/models"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

const (
	locationAPIPath = "/api/location/v1"
)

func GetLocation(keyword string) (*models.LocationResponse, error) {
	// Get the base url
	baseURL, err := web.AppConfig.String("base_url")
	if err != nil {
		logs.Error(
			"[LocationService] Failed to read configuration | key=base_url | err=%v",
			err,
		)
		return nil, fmt.Errorf("failed to read 'base_url' from configuration: %w", err)
	}

	if strings.TrimSpace(baseURL) == "" {
		return nil, fmt.Errorf("configuraion 'base_url' is empty")
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
		"[LocationService] Calling Location API | keyword=%s | url=%s",
		keyword,
		parsedURL.String(),
	)
	
	// Client HTTP request
	request, err := NewGETRequest(parsedURL.String())
	if err != nil {
		return nil, err
	}

	// Send request
	response, err := httpClient.Do(request)
	if err != nil {
		logs.Error(
			"[LocationService] HTTP request failed | url=%s | err=%v",
			parsedURL.String(),
			err,
    	)
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer response.Body.Close()

	// Check response status
	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)

		logs.Warn(
			"[LocationService] Unexpected response | status=%d | body=%s",
			response.StatusCode,
			string(body),
		)

		return nil, fmt.Errorf(
			"unexpected status code: %d",
			response.StatusCode,
		)
	}

	// Decode response
	var location models.LocationResponse

	if err := json.NewDecoder(
		response.Body,
	).Decode(&location); err != nil {
		logs.Error(
			"[LocationService] Decode response failed | err=%v",
			err,
		)
		return nil, fmt.Errorf("decode response: %w", err)
	}

	logs.Debug(
		"[LocationService] Location API success | locationID=%s",
		location.GeoInfo.LocationID,
	)

	return &location, nil
}
