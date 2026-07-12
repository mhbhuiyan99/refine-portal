package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"refine-portal/models"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

const (
	propertyDetailsAPIPath = "/api/property/bookmark/v1"
)

func GetPropertyDetails(
	req models.PropertyDetailsRequest,
) (*models.PropertyDetailsResponse, error) {
	// Get base url from config
	baseURL, err := web.AppConfig.String("base_url")
	if err != nil {
		logs.Error(
			"[PropertyDetailsService] Failed to read configuration | key=base_url | err=%v",
			err,
		)
		return nil, fmt.Errorf("failed to get 'base_url' from config: %w", err)
	}
	if strings.TrimSpace(baseURL) == "" {
		logs.Error(
			"[PropertyDetailsService] Configuration 'base_url' is empty",
		)
		return nil, fmt.Errorf("configuration 'base_url' is empty")
	}

	// Build URL
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("parse base_url failed: %w", err)
	}

	parsedURL.Path = propertyDetailsAPIPath

	query := parsedURL.Query()

	query.Set(
		"propertyIdList",
		strings.Join(req.PropertyIDList, ","),
	)

	parsedURL.RawQuery = query.Encode()

	logs.Debug(
		"[PropertyDetailsService] Calling Property Details API | propertyIdCount=%d | url=%s",
		len(req.PropertyIDList),
		parsedURL.String(),
	)

	// Create Request
	request, err := http.NewRequest(
		http.MethodGet,
		parsedURL.String(),
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	request.Header.Set("Accept", "application/json")

	// Send Request
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		logs.Error(
			"[PropertyDetailsService] HTTP request failed | url=%s | err=%v",
			parsedURL.String(),
			err,
    	)
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer response.Body.Close()

	// Validate Response
	if response.StatusCode != http.StatusOK {
		logs.Warn(
			"[PropertyDetailsService] Unexpected response | status=%d | url=%s",
			response.StatusCode,
			parsedURL.String(),
    	)
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	// Decode Response
	var propertyDetailsResponse models.PropertyDetailsResponse

	if err := json.NewDecoder(
		response.Body,
	).Decode(&propertyDetailsResponse); err != nil {
		logs.Error(
			"[PropertyDetailsService] Decode response failed | err=%v",
			err,
		)
		return nil, fmt.Errorf("decode response failed: %w", err)
	}

	/*logs.Debug(
		"[PropertyDetailsService] Property Details API success | returned=%d",
		len(propertyDetailsResponse.Properties),
	)*/
	return &propertyDetailsResponse, nil
}
