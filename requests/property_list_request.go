package requests

import (
	"fmt"
	"net/url"
	"refine-portal/models"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// Property API endpoint.
const (
	propertyListAPIPath = "/api/properties/category/v1"
)

// GetPropertyListRequest calls the Property List API and returns
// the list of properties for the requested location and filters.
func GetPropertyListRequest(
	req models.PropertyListRequest,
) (*models.PropertyListResponse, error) {

	// Get base url from config
	baseURL, err := web.AppConfig.String("base_url")
	if err != nil {
		return nil, fmt.Errorf("failed to read base_url: %w", err)
	}
	if strings.TrimSpace(baseURL) == "" {
		return nil, fmt.Errorf("base_url is empty")
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

	requestURL, err := BuildURL(baseURL, propertyListAPIPath, query)
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
