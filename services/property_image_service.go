package services

import (
	"refine-portal/models"
	"refine-portal/requests"
)

// GetPropertyImages returns all images for one property.
//
// Responsibilities:
//   - Call the request layer.
//   - Return the image response.
func GetPropertyImages(
	propertyID string,
) (*models.PropertyImagesResponse, error) {

	return requests.GetPropertyImagesRequest(propertyID)
}