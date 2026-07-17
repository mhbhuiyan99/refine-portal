package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"refine-portal/models"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

const (
	propertyListAPIPath = "/api/properties/category/v1"
)

func GetProperties(
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
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("parse base_url failed: %w", err)
	}

	parsedURL.Path = propertyListAPIPath

	query := parsedURL.Query()

	query.Set("category", req.Category)
	query.Set("locations", req.Locations)
	query.Set("order", fmt.Sprintf("%d", req.Order))
	query.Set("limit", fmt.Sprintf("%d", req.Limit))
	query.Set("items", fmt.Sprintf("%d", req.Items))
	query.Set("device", req.Device)
	query.Set("page", fmt.Sprintf("%d", req.Page))

	parsedURL.RawQuery = query.Encode()

	logs.Debug(
		"[PropertyService] Calling Property List API | category=%s | locations=%s | page=%d | limit=%d | order=%d | url=%s",
		req.Category,
		req.Locations,
		req.Page,
		req.Limit,
		req.Order,
		parsedURL.String(),
	)

	// Create Request
	request, err := NewGETRequest(parsedURL.String())
	if err != nil {
		return nil, err
	}

	// Send Request
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Validate Response
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	// Decode Response
	var propertyListResponse models.PropertyListResponse

	if err := json.NewDecoder(
		response.Body,
	).Decode(&propertyListResponse); err != nil {
		return nil, fmt.Errorf("decode response failed: %w", err)
	}

	logs.Debug(
		"[PropertyService] Property List API success | propertyCount=%d",
		len(propertyListResponse.Result.ItemIDs),
	)
	return &propertyListResponse, nil
}
