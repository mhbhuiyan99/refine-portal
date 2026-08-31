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
//
// go:noinline is required here: this function is small enough that
// the compiler would otherwise inline it, which breaks gomonkey's
// runtime patching in property_image_api_test.go (the mock silently
// never engages, and the real HTTP call runs instead).
//
//go:noinline
func GetPropertyImages(propertyID string) (*models.PropertyImagesResponse, error) {
	return requests.GetPropertyImagesRequest(propertyID)
}