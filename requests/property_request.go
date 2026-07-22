package requests

import (
	"fmt"
	"net/url"
	"refine-portal/models"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// Property Details API endpoint.
const (
	propertyDetailsAPIPath = "/api/property/bookmark/v1"
)

// GetPropertyDetailsRequest calls the Property Details API for
// a batch of property IDs and returns detailed property information.
func GetPropertyDetailsRequest(
	propertyIDs []string,
) (*models.PropertyDetailsResponse, error) {
	
	baseURL, err := web.AppConfig.String("base_url")
	if err != nil {
		logs.Error(
			"[PropertyDetailsRequest] Failed to read configuration | key=base_url | err=%v",
			err,
		)
		return nil, fmt.Errorf("failed to get 'base_url' from config: %w", err)
	}

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("parse base_url failed: %w", err)
	}

	parsedURL.Path = propertyDetailsAPIPath

	query := parsedURL.Query()

	query.Set(
		"propertyIdList",
		strings.Join(propertyIDs, ","),
	)

	parsedURL.RawQuery = query.Encode()

	logs.Debug(
		"[PropertyDetailsRequest] Calling Property Details API | propertyIdCount=%d | url=%s",
		len(propertyIDs),
		parsedURL.String(),
	)

	request, err := NewGETRequest(parsedURL.String())
	if err != nil {
		return nil, err
	}

	var result models.PropertyDetailsResponse

	err = DoRequest(
		request,
		&result,
	)

	if err != nil {
		return nil, err
	}

	logs.Debug(
		"[PropertyDetailsRequest] Property Details API success | propertyCount=%d",
		len(result.Items),
	)

	return &result, nil
}