package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"refine-portal/models"
	"strings"
	"time"

	"github.com/beego/beego/v2/server/web"
)

const (
	locationAPIPath = "/api/location/v1"
)

func GetLocation(keyword string) (*models.LocationResponse, error) {
	// Get the base url
	baseURL, err := web.AppConfig.String("base_url")
	if err != nil {
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

	// Client HTTP request
	request, err := http.NewRequest(
		http.MethodGet,
		parsedURL.String(),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	request.Header.Set("Accept", "application/json")

	// Send request
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer response.Body.Close()

	// Check response status
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	// Decode response
	var location models.LocationResponse

	if err := json.NewDecoder(
		response.Body,
	).Decode(&location); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &location, nil
}
