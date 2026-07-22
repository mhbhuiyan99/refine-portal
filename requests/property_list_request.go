package requests

import (
	"fmt"
	"net/url"
	"refine-portal/models"

	"github.com/beego/beego/v2/core/logs"
)

// Property API endpoint.
const (
	propertyListAPIPath = "/api/properties/category/v1"
)

// GetPropertyListRequest retrieves a list of properties
// based on the provided search filters.
//
// Responsibilities:
//   - Build the Property List API URL.
//   - Create the HTTP request.
//   - Execute the HTTP request.
//   - Decode the JSON response.
//   - Return the property list response.
func GetPropertyListRequest(
	req models.PropertyListRequest,
) (*models.PropertyListResponse, error) {

	baseURL, err := GetURLFromConfig("base_url")
	if err != nil {
		return nil, err
	}

	// Build URL
	query := url.Values{}
	query.Set("category", req.Category)
	query.Set("locations", req.Locations)
	query.Set("order", fmt.Sprintf("%d", req.Order))
	query.Set("limit", fmt.Sprintf("%d", req.Limit))
	query.Set("items", fmt.Sprintf("%d", req.Items))
	query.Set("device", req.Device)
	query.Set("page", fmt.Sprintf("%d", req.Page))

	requestURL, err := BuildURL(
		baseURL, 
		propertyListAPIPath, 
		query,
	)
	if err != nil {
		return nil, err
	}

	logs.Debug(
		"[PropertyListRequest] Calling Property List API | category=%s | locations=%s | page=%d | limit=%d | order=%d | url=%s",
		req.Category,
		req.Locations,
		req.Page,
		req.Limit,
		req.Order,
		requestURL,
	)

	request, err := NewGETRequest(requestURL)
	if err != nil {
		return nil, err
	}

	// Execute Request
	var propertyListResponse models.PropertyListResponse

	if err := DoRequest(request, &propertyListResponse); err != nil {
		return nil, fmt.Errorf("property list request failed: %w", err)
	}

	logs.Debug(
		"[PropertyListRequest] Property List API success | propertyCount=%d",
		len(propertyListResponse.Result.ItemIDs),
	)
	return &propertyListResponse, nil
}
