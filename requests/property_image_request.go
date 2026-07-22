package requests

import (
	"net/url"
	"refine-portal/models"

	"github.com/beego/beego/v2/core/logs"
)

const propertyImagesAPIPath = "/api/property/images/v1"

// GetPropertyImagesRequest retrieves all images for a property.
//
// Responsibilities:
//   - Build the Property Images API URL.
//   - Create the HTTP request.
//   - Execute the HTTP request.
//   - Decode the JSON response.
//   - Return the property images.
func GetPropertyImagesRequest(
	propertyID string,
) (*models.PropertyImagesResponse, error) {

	baseURL, err := GetBaseURL()
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Set("propertyId", propertyID)

	requestURL, err := BuildURL(
		baseURL,
		propertyImagesAPIPath,
		query,
	)
	if err != nil {
		return nil, err
	}

	logs.Debug(
		"[PropertyImagesRequest] Calling Images API | property=%s",
		propertyID,
	)

	request, err := NewGETRequest(requestURL)
	if err != nil {
		return nil, err
	}

	var result models.PropertyImagesResponse

	if err := DoRequest(
		request,
		&result,
	); err != nil {
		return nil, err
	}

	return &result, nil
}