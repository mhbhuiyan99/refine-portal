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

	query := url.Values{}
	query.Set(
		"propertyIdList",
		strings.Join(propertyIDs, ","),
	)

	requestURL, err := BuildURL(baseURL, propertyDetailsAPIPath, query)
	if err != nil {
		return nil, err
	}

	logs.Debug(
		"[PropertyDetailsRequest] Calling Property Details API | propertyIdCount=%d | url=%s",
		len(propertyIDs),
		requestURL,
	)

	request, err := NewGETRequest(requestURL)
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
