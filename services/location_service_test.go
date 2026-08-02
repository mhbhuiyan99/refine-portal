package services

import (
	"errors"
	"testing"

	"refine-portal/models"
	"refine-portal/requests"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetLocation_Success(t *testing.T) {

	expected := &models.LocationResponse{
		Success: true,
		Message: "success",
		GeoInfo: models.LocationGeoInfo{
			City: "Dhaka",
		},
	}

	patches := gomonkey.ApplyFunc(
		requests.GetLocationRequest,
		func(keyword string) (*models.LocationResponse, error) {

			assert.Equal(t, "Dhaka", keyword)

			return expected, nil
		},
	)

	defer patches.Reset()

	result, err := GetLocation("Dhaka")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetLocation_Error(t *testing.T) {

	expectedErr := errors.New("request failed")

	patches := gomonkey.ApplyFunc(
		requests.GetLocationRequest,
		func(keyword string) (*models.LocationResponse, error) {

			return nil, expectedErr
		},
	)

	defer patches.Reset()

	result, err := GetLocation("Dhaka")

	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
}


