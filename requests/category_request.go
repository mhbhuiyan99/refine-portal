package requests

import (
	"fmt"
	"net/url"
	"refine-portal/models"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// Category API endpoint.
const (
	categoryAPIPath = "/api/v1/category/details"
)

// GetCategoryRequest retrieves category page data from the Category API.
//
// Responsibilities:
//   - Build the Category API URL.
//   - Create the HTTP request.
//   - Execute the HTTP request.
//   - Decode the JSON response.
//   - Return the category response.
func GetCategoryRequest(
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
	query := url.Values{}
	query.Set("aggsAvgPrice", "1")
	query.Set("aggsAvgRating", "1")
	query.Set("aggsAvgRoomSize", "1")
	query.Set("aggsCategory", "1")
	query.Set("device", "desktop")
	query.Set("items", "1")
	query.Set("locations", countryCode)
	query.Set("sections", "1")

	requestURL, err := BuildURL(baseURL, categoryAPIPath+"/"+slug, query)
	if err != nil {
		return nil, err
	}

	logs.Debug(
		"[CategoryRequest] Calling Category API | slug=%s | country=%s | url=%s",
		slug,
		countryCode,
		requestURL,
	)

	request, err := NewGETRequest(requestURL)
	if err != nil {
		logs.Error(
			"[CategoryRequest] Create request failed | err=%v",
			err,
		)
		return nil, err
	}

	var category models.CategoryResponse

	err = DoRequest(
		request,
		&category,
	)

	if err != nil {
		return nil, err
	}

	logs.Debug(
		"[CategoryRequest] Category API success | location=%s",
		category.GeoInfo.Name,
	)

	return &category, nil
}
