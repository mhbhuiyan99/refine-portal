package services

import (
	"encoding/json"
	"fmt"
	"net/url"
	"refine-portal/models"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

const (
	categoryAPIPath = "/api/v1/category/details"
)

func GetCategory(
	slug string,
	countryCode string,
) (*models.CategoryResponse, error) {

	// Read base URL
	baseURL, err := web.AppConfig.String("base_url")
	if err != nil {
		return nil, fmt.Errorf("failed to read base_url: %w", err)
	}

	if strings.TrimSpace(baseURL) == "" {
		return nil, fmt.Errorf("base_url is empty")
	}

	// Parse URL
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("parse base_url failed: %w", err)
	}

	// Build path
	parsedURL.Path = categoryAPIPath + "/" + slug

	// Build query
	query := parsedURL.Query()

	query.Set("aggsAvgPrice", "1")
	query.Set("aggsAvgRating", "1")
	query.Set("aggsAvgRoomSize", "1")
	query.Set("aggsCategory", "1")
	query.Set("device", "desktop")
	query.Set("items", "1")
	query.Set("locations", countryCode)
	query.Set("sections", "1")

	parsedURL.RawQuery = query.Encode()

	logs.Debug(
		"[CategoryService] Calling Category API | slug=%s | country=%s | url=%s",
		slug,
		countryCode,
		parsedURL.String(),
	)

	// Create request
	request, err := NewGETRequest(parsedURL.String())
	if err != nil {
		return nil, err
	}

	// Execute request
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Decode response
	var category models.CategoryResponse

	if err := json.NewDecoder(response.Body).Decode(&category); err != nil {
		logs.Error(
			"[CategoryService] Decode failed | url=%s | err=%v",
			parsedURL.String(),
			err,
		)
		return nil, err
	}

	logs.Debug(
		"[CategoryService] Category API success | location=%s",
		category.GeoInfo.Name,
	)

	return &category, nil
}