package services

import (
	"testing"
	"errors"

	"refine-portal/models"
	"refine-portal/requests"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetProperties_Success(t *testing.T) {

	req := models.PropertyListRequest{
		Category:  "dhaka",
		Locations: "BD",
		Order:     1,
		Limit:     10,
	}

	expected := &models.PropertyListResponse{
		Success: true,
		Result: models.PropertyResult{
			Count: 2,
			ItemIDs: []string{
				"101",
				"102",
			},
		},
	}

	patches := gomonkey.ApplyFunc(
		requests.GetPropertyListRequest,
		func(r models.PropertyListRequest) (*models.PropertyListResponse, error) {

			assert.Equal(t, req, r)

			return expected, nil
		},
	)

	defer patches.Reset()

	result, err := GetProperties(req)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetProperties_Error(t *testing.T) {

	req := models.PropertyListRequest{}

	expectedErr := errors.New("request failed")

	patches := gomonkey.ApplyFunc(
		requests.GetPropertyListRequest,
		func(r models.PropertyListRequest) (*models.PropertyListResponse, error) {

			return nil, expectedErr
		},
	)

	defer patches.Reset()

	result, err := GetProperties(req)

	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
}