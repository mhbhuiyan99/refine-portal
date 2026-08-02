package services

import (
	"errors"
	"testing"

	"refine-portal/models"
	"refine-portal/requests"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetPropertyImages_Success(t *testing.T) {

	propertyID := "123"

	expected := &models.PropertyImagesResponse{
		Success: true,
	}

	patches := gomonkey.ApplyFunc(
		requests.GetPropertyImagesRequest,
		func(id string) (*models.PropertyImagesResponse, error) {

			assert.Equal(t, propertyID, id)

			return expected, nil
		},
	)

	defer patches.Reset()

	result, err := GetPropertyImages(propertyID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetPropertyImages_Error(t *testing.T) {

	propertyID := "123"

	expectedErr := errors.New("request failed")

	patches := gomonkey.ApplyFunc(
		requests.GetPropertyImagesRequest,
		func(id string) (*models.PropertyImagesResponse, error) {

			assert.Equal(t, propertyID, id)

			return nil, expectedErr
		},
	)

	defer patches.Reset()

	result, err := GetPropertyImages(propertyID)

	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
}