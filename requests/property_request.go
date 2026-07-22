package requests

import (
	"net/url"
	"refine-portal/models"
	"strings"

	"github.com/beego/beego/v2/core/logs"
)

// Property Details API endpoint.
const (
	propertyDetailsAPIPath = "/api/property/bookmark/v1"
)

// GetPropertyDetailsRequest retrieves detailed property information
// for a batch of property IDs from the Property Details API.
//
// Responsibilities:
//   - Build the Property Details API URL.
//   - Create the HTTP request.
//   - Execute the HTTP request.
//   - Decode the JSON response.
//   - Return the property details response.
func GetPropertyDetailsRequest(
	propertyIDs []string,
) (*models.PropertyDetailsResponse, error) {

	baseURL, err := GetURLFromConfig("base_url")
	if err != nil {
		return nil, err
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
